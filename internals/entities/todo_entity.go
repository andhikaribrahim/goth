package entities

type TodoStatus string

const (
	Pending TodoStatus = "pending"
	Completed TodoStatus = "completed"
)

type Todo struct {
	Title string
	Status TodoStatus
	Order int
}
