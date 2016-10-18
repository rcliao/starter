package web

import (
	"github.com/rcliao/starter"
	mgo "gopkg.in/mgo.v2"
)

// Client wrap the service connection
type Client struct {
	todoService TodoService

	// Path is to define db path to connect to
	Path string
	db   *DB
}

// DB is a wrapper to wrap mongo session
type DB struct {
	session *mgo.Session
}

// NewClient is the constructor pattern to create new client
func NewClient() *Client {
	c := &Client{}
	c.todoService.client = c
	return c
}

// Connect opens up connection to mongodb and keep session for future use
func (c *Client) Connect() error {
	session, err := mgo.Dial(c.Path)
	c.db = &DB{session}
	if err != nil {
		return err
	}
	return nil
}

// Disconnect close the current connection (if any) to mongodb
func (c *Client) Disconnect() {
	if c.db != nil && c.db.session != nil {
		c.db.session.Close()
	}
}

// TodoService is to get the service from the client
func (c *Client) TodoService() starter.TodoService {
	return &c.todoService
}
