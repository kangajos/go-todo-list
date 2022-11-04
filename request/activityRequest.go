package request

type ActivityRequest struct {
	ID    int
	Title string `binding:"required"`
	Email string `binding:"required,email"`
}
