package service

import (
	"context"
	"fxlik/simple-post-api/model/web"
)

type LevelService interface {
	Save(ctx context.Context, request web.LevelCreateRequest) web.LevelResponse
	Update(ctx context.Context, request web.LevelUpdateRequest) web.LevelResponse
	Delete(ctx context.Context, levelId int32)
	FindById(ctx context.Context, levelId int32) web.LevelResponse
	FindAll(ctx context.Context) []web.LevelResponse
}
