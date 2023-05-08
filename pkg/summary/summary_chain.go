package summary

import (
	"github.com/tmc/langchaingo/exp/chains"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/prompts"
)

func NewSummaryChain() (chains.LLMChain, error) {
	llm, err := openai.New()
	if err != nil {
		return chains.LLMChain{}, err
	}

	prompt, err := prompts.NewPromptTemplate(PromptTemplate, []string{
		"function",
	})

	if err != nil {
		return chains.LLMChain{}, err
	}
	return chains.NewLLMChain(llm, prompt), nil
}

func NewRefineSummaryChain() (chains.LLMChain, error) {
	llm, err := openai.New()
	if err != nil {
		return chains.LLMChain{}, err
	}

	prompt, err := prompts.NewPromptTemplate(RefinePromptTemplate, []string{
		"function_summary",
		"called_functions_summaries",
	})

	if err != nil {
		return chains.LLMChain{}, err
	}
	return chains.NewLLMChain(llm, prompt), nil
}
