package handler

import (
	"net/http"
	"strconv"
	"todo"

	"github.com/gin-gonic/gin"
)

// @Summary Create a new todo list
// @Tags lists
// @Description Creates a new todo list for the authenticated user.
// @ID create-todo-list
// @Accept  json
// @Produce  json
// @Param input body todo.TodoList true "Todo list information"
// @Success 200 {object} map[string]interface{} "Returns the ID of the created todo list"
// @Failure 400 {object} errorResponse "Invalid input data"
// @Failure 500 {object} errorResponse "Internal server error"
// @Router /api/lists [post]
func (h *Handler) createList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input todo.TodoList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.TodoList.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

type getAllListsResponse struct {
	Data []todo.TodoList `json:"data"`
}

// @Summary Get all todo lists
// @Tags lists
// @Description Retrieves all todo lists for the authenticated user.
// @ID get-all-todo-lists
// @Accept  json
// @Produce  json
// @Success 200 {array} todo.TodoList "Returns an array of todo lists"
// @Failure 500 {object} errorResponse "Internal server error"
// @Router /api/lists [get]
func (h *Handler) getAllLists(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	lists, err := h.services.TodoList.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})
}

// @Summary Get a todo list by ID
// @Tags lists
// @Description Retrieves a specific todo list by its ID.
// @ID get-todo-list-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Todo list ID"
// @Success 200 {object} todo.TodoList "Returns the todo list"
// @Failure 400 {object} errorResponse "Invalid ID"
// @Failure 404 {object} errorResponse "List not found"
// @Failure 500 {object} errorResponse "Internal server error"
// @Router /api/lists/{id} [get]
func (h *Handler) getListById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	list, err := h.services.TodoList.GetById(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}

// @Summary Update a todo list
// @Tags lists
// @Description Updates a todo list's title and/or description.
// @ID update-todo-list
// @Accept  json
// @Produce  json
// @Param id path int true "Todo list ID"
// @Param input body todo.UpdateListInput true "Fields to update"
// @Success 200 {object} statusResponse "Returns success status"
// @Failure 400 {object} errorResponse "Invalid input data or ID"
// @Failure 404 {object} errorResponse "List not found"
// @Failure 500 {object} errorResponse "Internal server error"
// @Router /api/lists/{id} [put]
func (h *Handler) updateList(c *gin.Context) {
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

	var input todo.UpdateListInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.TodoList.Update(userId, id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

// @Summary Delete a todo list
// @Tags lists
// @Description Deletes a todo list by its ID.
// @ID delete-todo-list
// @Accept  json
// @Produce  json
// @Param id path int true "Todo list ID"
// @Success 200 {object} statusResponse "Returns success status"
// @Failure 400 {object} errorResponse "Invalid ID"
// @Failure 404 {object} errorResponse "List not found"
// @Failure 500 {object} errorResponse "Internal server error"
// @Router /api/lists/{id} [delete]
func (h *Handler) deleteList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.TodoList.Delete(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
