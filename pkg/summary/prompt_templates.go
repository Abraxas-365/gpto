package summary

const RefinePromptTemplate = `
	Your job is to produce a final summary for a Go Lang function.
	A function summary has the folowing format:

	----------
	Function Name
	Input Parameters
	Output Parameters
	Description of the function
	----------

	We have provided you an existing summary up to a certain point:

	{function_summary}

	This summary may reference the call to some functions.
	We have the opportunity to refine the existing summary (only if needed) with the summaries of the called functions below.

	{called_functions_summaries}

	Given the new context refine the original summary.
	If no refinement is needed return the original function.
	`

const PromptTemplate = `
	Your job is to produce a summary for a Go Lang function.
	A function summary has the folowing format:

	----------
	Function Name
	Input Parameters
	Output Parameters
	Description of the function
	----------

	Summarize the following function:

	{function}
	`
