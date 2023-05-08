package application

import (
	"context"
	"strings"
	"sync"

	"github.com/Abraxas-365/gpto/internal/domain/ports"
	"github.com/Abraxas-365/gpto/internal/domain/ports/models"
	"github.com/Abraxas-365/gpto/pkg/funcnode"
	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/llms/openai"
)

type App struct {
	repo     ports.RepoInterface
	embedder embeddings.Embedder
	llm      *openai.LLM
}

func NewApp(repo ports.RepoInterface) (*App, error) {
	embedder, err := embeddings.NewOpenAI()
	if err != nil {
		return nil, err
	}

	llm, err := openai.New()
	if err != nil {
		return nil, err
	}

	return &App{
		repo:     repo,
		embedder: embedder,
		llm:      llm,
	}, nil
}

//TODO: check if function exist and if it exit check if function is equal, if its equal skip the embedding and save process
func (a *App) SaveMetaData(node funcnode.FuncNode) error {
	ctx := context.Background()

	embedding, err := a.embedder.EmbedQuery(ctx, node.Summary)
	if err != nil {
		return err
	}

	metadata := models.NewMetaData(node.Body, node.Summary, float64SliceToFloat32Slice(embedding))
	if err := a.repo.SaveMetaData(ctx, metadata); err != nil {
		return err
	}

	return nil
}

func (a *App) SaveMetadataConcurrently(nodes []funcnode.FuncNode) error {
	var wg sync.WaitGroup
	errorCh := make(chan error, len(nodes))

	for _, node := range nodes {
		wg.Add(1)
		go func(node funcnode.FuncNode) {
			defer wg.Done()
			if err := a.SaveMetaData(node); err != nil {
				errorCh <- err
			}
		}(node)
	}

	wg.Wait()
	close(errorCh)

	for err := range errorCh {
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) GetMostSimilarVectors(query string, limit int) ([]models.MetaData, error) {
	ctx := context.Background()
	embeddendQuery, err := a.embedder.EmbedQuery(ctx, query)
	if err != nil {
		return nil, err
	}

	similar, err := a.repo.GetMostSimilarVectors(ctx, float64SliceToFloat32Slice(embeddendQuery), limit)
	if err != nil {
		return nil, err
	}

	return similar, nil
}

func (a *App) Chat(query string) (string, error) {
	prompt := `As an AI language model, You are able to provide assistance with understanding and gathering information about codebases. 
			Yopu will resive a specific codebase or programming language to help with,
			along with any questions or details I would like you to address.
			You will always retuer your answers in markdown notation 
	
	`
	similar, err := a.GetMostSimilarVectors(query, 3)
	if err != nil {
		return "", err
	}
	var sb strings.Builder

	// Iterate through the metaDataSlice and concatenate the Summary and FunctionBody fields
	for _, metaData := range similar {
		sb.WriteString(metaData.Summary)
		sb.WriteString("\n")
		sb.WriteString(metaData.FunctionBody)
		sb.WriteString("\n")
	}

	prompt += sb.String()
	prompt += "Question: " + query + "\n"

	return a.llm.Chat(prompt)
}

func float64SliceToFloat32Slice(float64Slice []float64) []float32 {
	float32Slice := make([]float32, len(float64Slice))
	for i, value := range float64Slice {
		float32Slice[i] = float32(value)
	}
	return float32Slice
}
