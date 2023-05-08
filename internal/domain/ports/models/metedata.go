package models

type MetaData struct {
	ID           int64     `json:"id"`
	FunctionBody string    `json:"function_body"`
	Summary      string    `json:"summary"`
	Embedding    []float32 `json:"embedding"`
}

func NewMetaData(
	fnBody string,
	summary string,
	embedding []float32,
) *MetaData {
	return &MetaData{
		FunctionBody: fnBody,
		Summary:      summary,
		Embedding:    embedding,
	}
}
