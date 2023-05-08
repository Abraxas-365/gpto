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
	fmt.Println("SUMARY CHAIN: ", body)
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
		return ""
	}

	if len(calledFunctionsSummaries) != 0 {
		refinedSummary, err := chains.Call(ctx, refineSummaryChain, map[string]any{
			"function_summary":           summary["text"].(string),
			"called_functions_summaries": calledFunctionSummariesString,
		})

		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("REFINED SUMMARY: ", refinedSummary["text"].(string))
		return refinedSummary["text"].(string)
	}
	fmt.Println("SUMMARY: ", summary["text"].(string))
	return summary["text"].(string)
}
