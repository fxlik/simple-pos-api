package web

import "fxlik/simple-post-api/model/domain"

type LevelResponse struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
}

func ToLevelResponse(level domain.Level) LevelResponse {
	return LevelResponse{
		Id:   level.Id,
		Name: level.Name,
	}
}

func ToLevelResponses(levels []domain.Level) []LevelResponse {
	var levelResponses []LevelResponse
	for _, level := range levels {
		levelResponses = append(levelResponses, ToLevelResponse(level))
	}
	return levelResponses
}
