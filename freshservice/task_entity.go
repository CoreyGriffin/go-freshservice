package freshservice

import "time"

// Tasks holds a list of Freshservice task details
type Tasks struct {
	List []TaskDetails `json:"tasks"`
}

// Task holds the details of a specific Freshservice task
type Task struct {
	Details TaskDetails `json:"task"`
}

// TaskDetails are the details related to a specific task in Freshservice
type TaskDetails struct {
	ID           int       `json:"id"`
	AgentID      int       `json:"agent_id"`
	Status       int       `json:"status"`
	DueDate      time.Time `json:"due_date"`
	NotifyBefore int       `json:"notify_before"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	ClosedAt     int       `json:"closed_at"`
	GroupID      int       `json:"group_id"`
}
