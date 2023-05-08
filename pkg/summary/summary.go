package summary

import (
	"context"
	"fmt"
	"strings"

	"github.com/tmc/langchaingo/exp/chains"
)

func Create(body string, calledFunctionsSummaries []string) string {
	ctx := context.Background()
	summaryChain, err := NewSummaryChain()
	if err != nil {
		fmt.Println(err)
	}
	refineSummaryChain, err := NewRefineSummaryChain()
	if err != nil {
		fmt.Println(err)
	}

	calledFunctionSummariesString := strings.Join(calledFunctionsSummaries, "\n")

	summary, err := chains.Call(ctx, summaryChain, map[string]any{
		"function": body,
	})
	if err != nil {
		fmt.Println(err)
	}

	refinedSummary, err := chains.Call(ctx, refineSummaryChain, map[string]any{
		"function_summary":           summary,
		"called_functions_summaries": calledFunctionSummariesString,
	})

	return refinedSummary["text"].(string)
}
