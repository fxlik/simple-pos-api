package web

type LevelCreateRequest struct {
	Name string `validate:"required,min=1,max=100" json:"name"`
}

type LevelUpdateRequest struct {
	Id   int32  `validate:"required" json:"id"`
	Name string `validate:"required,min=1,max=100" json:"name"`
}
