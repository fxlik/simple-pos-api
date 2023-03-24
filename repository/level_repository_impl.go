package repository

import (
	"context"
	"database/sql"
	"errors"
	"fxlik/simple-post-api/helper"
	"fxlik/simple-post-api/model/domain"
)

type LevelRepositoryImpl struct {
}

func NewLevelRepositoryImpl() *LevelRepositoryImpl {
	return &LevelRepositoryImpl{}
}

func (repository *LevelRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, level domain.Level) domain.Level {
	SQL := "INSERT INTO level(name) VALUES (?)"
	result, err := tx.ExecContext(ctx, SQL, level.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	level.Id = int32(id)
	return level
}

func (repository *LevelRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, level domain.Level) domain.Level {
	SQL := "UPDATE level SET name = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, level.Name, level.Id)
	helper.PanicIfError(err)
	return level
}

func (repository *LevelRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, level domain.Level) {
	SQL := "DELETE FROM level WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, level.Id)
	helper.PanicIfError(err)
}

func (repository *LevelRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, levelId int32) (domain.Level, error) {
	SQL := "SELECT id, name FROM level WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, levelId)
	helper.PanicIfError(err)
	defer rows.Close()
	level := domain.Level{}
	if rows.Next() {
		err := rows.Scan(&level.Id, &level.Name)
		helper.PanicIfError(err)
		return level, nil
	} else {
		return level, errors.New("level is not found")
	}
}

func (repository *LevelRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Level {
	SQL := "SELECT * FROM level"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var levels []domain.Level
	for rows.Next() {
		level := domain.Level{}
		err := rows.Scan(&level.Id, &level.Name)
		helper.PanicIfError(err)
		levels = append(levels, level)
	}
	return levels
}
