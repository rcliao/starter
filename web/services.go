package web

import (
	"github.com/rcliao/starter"
	"gopkg.in/mgo.v2/bson"
)

// ensure TodoService implements starter TodoServiceInterface
var _ starter.TodoService = &TodoService{}

// TodoService represents service layer to handle todos
type TodoService struct {
	client *Client
}

// List lists all todos
func (s *TodoService) List() ([]*starter.Todo, error) {
	var result []*starter.Todo
	c := s.client.db.session.DB("go-starter").C("todo")
	err := c.Find(bson.M{}).All(&result)
	if err != nil {
		return nil, err
	}

	return result, err
}

// Get returns a single Todo item based on todo id
func (s *TodoService) Get(id starter.TodoID) (*starter.Todo, error) {
	result := &starter.Todo{}
	c := s.client.db.session.DB("go-starter").C("todo")
	err := c.Find(id).One(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Add a todo to the database
func (s *TodoService) Add(name, description string) error {
	c := s.client.db.session.DB("go-starter").C("todo")
	return c.Insert(&starter.Todo{
		Name:        name,
		Description: description,
		Done:        false,
	})
}

// Done marks a single todo to be done
func (s *TodoService) Done(id starter.TodoID) error {
	return nil
}
