package models

type TodoStatus string

const (
	TodoStatusPending    TodoStatus = "Pending"
	TodoStatusInProgress TodoStatus = "In Progress"
	TodoStatusCompleted  TodoStatus = "Completed"
)
