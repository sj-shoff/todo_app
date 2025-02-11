package todo

import "errors"

// TodoList represents a todo list
type TodoList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

// UsersList represents a relation between users and lists
type UsersList struct {
	Id     int
	UserId int
	ListId int
}

// TodoItem represents a todo item
type TodoItem struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	Done        bool   `json:"done" db:"done"`
}

// ListsItem represents a relation between lists and items
type ListsItem struct {
	Id     int
	ListId int
	ItemId int
}

// UpdateListInput represents an input for updating a list
type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

// Validate validates the update list input
func (i *UpdateListInput) Validate() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("update struscture has no values")
	}
	return nil
}

// UpdateItemInput represents an input for updating an item
type UpdateItemInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Done        *bool   `json:"done"`
}

// Validate validates the update item input
func (i UpdateItemInput) Validate() error {
	if i.Title == nil && i.Description == nil && i.Done == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
