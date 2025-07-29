package errors

// TaskNotFoundError occurs when task is not found
type TaskNotFoundError struct{}

func (e TaskNotFoundError) Error() string {
	return "task not found"
}

// InvalidTaskIDError occurs when task ID is invalid
type InvalidTaskIDError struct{}

func (e InvalidTaskIDError) Error() string {
	return "invalid task ID"
}

// TaskCreationError occurs when task creation fails
type TaskCreationError struct {
	Message string
}

func (e TaskCreationError) Error() string {
	return e.Message
}

// TaskUpdateError occurs when task update fails
type TaskUpdateError struct {
	Message string
}

func (e TaskUpdateError) Error() string {
	return e.Message
} 