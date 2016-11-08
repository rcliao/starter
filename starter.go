package starter

// TodoID represents a todo identifier
type TodoID string

// Todo represents a single todo item with name, description and whether its done
type Todo struct {
	ID          TodoID `json:"id" bson:"_id,omitempty"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	Done        bool   `json:"done" bson:"done"`
}

// Client creates a connection to the services
type Client interface {
	TodoService() TodoService
}

// TodoService represents a service to manage todos
type TodoService interface {
	List() ([]*Todo, error)
	Get(id TodoID) (*Todo, error)
	Add(name, description string) error
	Done(id TodoID) error
}
