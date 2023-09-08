package request

type RequestID struct {
	ID uint `json:"id" validate:"required"`
}
