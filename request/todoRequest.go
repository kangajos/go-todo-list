package request

type TodoRequest struct {
	ID              int
	ActivityGroupID int    `binding:"required" json:"activity_group_id"`
	Title           string `binding:"required"`
	IsActive        bool   `binding:"required" json:"is_active"`
	Priority        string `binding:"required"`
}
