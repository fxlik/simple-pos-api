package repository

import (
	"context"
	"database/sql"
	"fxlik/simple-post-api/model/domain"
)

type LevelRepository interface {
	Save(ctx context.Context, tx *sql.Tx, level domain.Level) domain.Level
	Update(ctx context.Context, tx *sql.Tx, level domain.Level) domain.Level
	Delete(ctx context.Context, tx *sql.Tx, level domain.Level)
	FindById(ctx context.Context, tx *sql.Tx, levelId int32) (domain.Level, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Level
}
