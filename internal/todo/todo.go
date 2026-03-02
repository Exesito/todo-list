package todo

import "time"

type Status string

const (
	StatusPending   Status = "pending"
	StatusCompleted Status = "completed"
)

type Todo struct {
	ID          int        `json:"id"`
	Description string     `json:"description"`
	Status      Status     `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
}

type List []Todo

func NewTodo(id int, description string) Todo {
	return Todo{
		ID:          id,
		Description: description,
		Status:      StatusPending,
		CreatedAt:   time.Now(),
	}
}

func (l *List) Add(description string) {
	id := 1
	if len(*l) > 0 {
		id = (*l)[len(*l)-1].ID + 1
	}
	*l = append(*l, NewTodo(id, description))
}

func (l *List) Complete(id int) bool {
	for i, t := range *l {
		if t.ID == id {
			now := time.Now()
			(*l)[i].Status = StatusCompleted
			(*l)[i].CompletedAt = &now
			return true
		}
	}
	return false
}

func (l *List) Delete(id int) bool {
	for i, t := range *l {
		if t.ID == id {
			*l = append((*l)[:i], (*l)[i+1:]...)
			return true
		}
	}
	return false
}
