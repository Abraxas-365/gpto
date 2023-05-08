package ports

import (
	"context"

	"github.com/Abraxas-365/gpto/internal/domain/ports/models"
)

type RepoInterface interface {
	SaveMetaData(ctx context.Context, paragraph *models.MetaData) error
	GetMostSimilarVectors(ctx context.Context, embedding []float32, limit int) ([]models.MetaData, error)
	ContentExists(ctx context.Context, content string) (bool, error)
}
