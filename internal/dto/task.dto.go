package dto

type CreateTaskDTO struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status      string `json:"status" binding:"required,oneof=low medium high critical"` // Status should be one of the given (case sensitive)
}

type UpdateTaskDTO struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status" binding:"oneof=low medium high critical"` // Status should be one of the given (case sensitive)
}
