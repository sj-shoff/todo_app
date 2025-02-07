package handler

import (
	"net/http"
	"strconv"
	"todo"

	"github.com/gin-gonic/gin"
)

// @Summary Create a new todo item
// @Description Creates a new todo item in a specific todo list.
// @ID create-todo-item
// @Accept  json
// @Produce  json
// @Param listId path int true "Todo list ID"
// @Param input body todo.TodoItem true "Todo item information"
// @Success 200 {object} map[string]interface{} "Returns the ID of the created todo item"
// @Failure 400 {object} errorResponse "Invalid input data or list ID"
// @Failure 500 {object} errorResponse "Internal server error"
// @Router /api/lists/{listId}/items [post]
func (h *Handler) createItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	var input todo.TodoItem
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.TodoItem.Create(userId, listId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Get all todo items in a list
// @Description Retrieves all todo items in a specific todo list.
// @ID get-all-todo-items
// @Accept  json
// @Produce  json
// @Param listId path int true "Todo list ID"
// @Success 200 {array} todo.TodoItem "Returns an array of todo items"
// @Failure 400 {object} errorResponse "Invalid list ID"
// @Failure 500 {object} errorResponse "Internal server error"
// @Router /api/lists/{listId}/items [get]
func (h *Handler) getAllItems(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	items, err := h.services.TodoItem.GetAll(userId, listId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, items)
}

// @Summary Get a todo item by ID
// @Description Retrieves a specific todo item by its ID.
// @ID get-todo-item-by-id
// @Accept  json
// @Produce  json
// @Param itemId path int true "Todo item ID"
// @Success 200 {object} todo.TodoItem "Returns the todo item"
// @Failure 400 {object} errorResponse "Invalid item ID"
// @Failure 404 {object} errorResponse "Item not found"
// @Failure 500 {object} errorResponse "Internal server error"
// @Router /api/items/{itemId} [get]
func (h *Handler) getItemById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	item, err := h.services.TodoItem.GetById(userId, itemId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}

// @Summary Update a todo item
// @Description Updates a todo item's title, description, and/or done status.
// @ID update-todo-item
// @Accept  json
// @Produce  json
// @Param itemId path int true "Todo item ID"
// @Param input body todo.UpdateItemInput true "Fields to update"
// @Success 200 {object} statusResponse "Returns success status"
// @Failure 400 {object} errorResponse "Invalid input data or item ID"
// @Failure 404 {object} errorResponse "Item not found"
// @Failure 500 {object} errorResponse "Internal server error"
// @Router /api/items/{itemId} [put]
func (h *Handler) updateItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input todo.UpdateItemInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.TodoItem.Update(userId, id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

// @Summary Delete a todo item
// @Description Deletes a todo item by its ID.
// @ID delete-todo-item
// @Accept  json
// @Produce  json
// @Param itemId path int true "Todo item ID"
// @Success 200 {object} statusResponse "Returns success status"
// @Failure 400 {object} errorResponse "Invalid item ID"
// @Failure 404 {object} errorResponse "Item not found"
// @Failure 500 {object} errorResponse "Internal server error"
// @Router /api/items/{itemId} [delete]
func (h *Handler) deleteItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	err = h.services.TodoItem.Delete(userId, itemId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}
