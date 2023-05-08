package repository

import (
	"context"
	"fmt"

	"github.com/Abraxas-365/gpto/internal/domain/ports"
	"github.com/Abraxas-365/gpto/internal/domain/ports/models"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pgvector/pgvector-go"
)

type metadataRepository struct {
	pool *pgxpool.Pool
}

func NewMetadataRepository(pool *pgxpool.Pool) ports.RepoInterface {
	return &metadataRepository{pool: pool}
}

func (r *metadataRepository) SaveMetaData(ctx context.Context, metadata *models.MetaData) error {
	fmt.Println("Saving metadata", metadata)
	query := `INSERT INTO "public"."metadata" (function_body, summary, embedding) VALUES ($1, $2, $3)`
	_, err := r.pool.Exec(ctx, query, metadata.FunctionBody, metadata.Summary, pgvector.NewVector(metadata.Embedding))
	return err
}

func (r *metadataRepository) GetMostSimilarVectors(ctx context.Context, embedding []float32, limit int) ([]models.MetaData, error) {
	query := ` 
	SELECT id, function_body, summary
	FROM "public"."metadata"
	ORDER BY embedding <-> $1
	LIMIT $2;
	`

	rows, err := r.pool.Query(ctx, query, pgvector.NewVector(embedding), limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var similarmetadatas []models.MetaData

	for rows.Next() {
		var sp models.MetaData

		err := rows.Scan(&sp.ID, &sp.FunctionBody, &sp.Summary)
		if err != nil {
			return nil, err
		}

		similarmetadatas = append(similarmetadatas, sp)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return similarmetadatas, nil
}

func (r *metadataRepository) ContentExists(ctx context.Context, functionBody string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM "public"."metadata" WHERE function_body = $1)`
	err := r.pool.QueryRow(ctx, query, functionBody).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
