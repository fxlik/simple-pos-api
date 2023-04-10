package web

type LevelCreateRequest struct {
	Name string `validate:"required,min=1,max=100" json:"name" form:"name"`
}

type LevelUpdateRequest struct {
	Id   int32  `validate:"required" json:"id" form:"id"`
	Name string `validate:"required,min=1,max=100" json:"name" form:"name"`
}
