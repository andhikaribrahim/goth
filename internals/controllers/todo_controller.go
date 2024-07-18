package controllers

import (
	"goth/internals/entities"
	"goth/internals/utils"
	sharedView "goth/views/shared"
	view "goth/views/todo"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TodoController struct {
	todos []entities.Todo
}

func NewTodoController() *TodoController {
	return &TodoController{
		todos: []entities.Todo{},
	}
}

func (tc *TodoController) GetIndexView(c echo.Context) error {
	indexView := view.Index(tc.todos)

	return utils.RenderView(c, indexView)
}

func (tc *TodoController) PostTodo(c echo.Context) error {
	body, err := c.FormParams()
	if err != nil {
		return err
	}

	newTodo := new(entities.Todo)
	if err := c.Bind(newTodo); err != nil {
		return err
	}

	newTodo.Order = len(tc.todos) + 1
	newTodo.Status = entities.Pending
	newTodo.Title = body.Get("title")
	tc.todos = append(tc.todos, *newTodo)

	return utils.RenderView(c, view.TodoItem(tc.todos))
}

func (tc *TodoController) PutTodoStatus(c echo.Context) error {
	params := c.QueryParams()
	println(params.Get)
	id := params.Get("Order")

	for i, todo := range tc.todos {
		order, err := strconv.Atoi(id)
		if err != nil {
			return err
		}

		if todo.Order == order && todo.Status == entities.TodoStatus(entities.Pending) {
			tc.todos[i].Status = entities.TodoStatus(entities.Completed)
		} else if todo.Order == order && todo.Status == entities.TodoStatus(entities.Completed) {
			/* Throw error here */
			utils.RenderView(c, sharedView.FlashErrorMessage("Cannot change status of completed item"))
			return utils.RenderView(c, view.TodoItem(tc.todos))
		}
	}

	return utils.RenderView(c, view.TodoItem(tc.todos))
}

func (tc *TodoController) DeleteTodo(c echo.Context) error {
	body, err := c.FormParams()
	if err != nil {
		return err
	}

	id := body.Get("Order")

	for i, todo := range tc.todos {
		order, err := strconv.Atoi(id)
		if err != nil {
			return err
		}

		if todo.Order == order {
			tc.todos = append(tc.todos[:i], tc.todos[i+1:]...)
		}
	}

	return utils.RenderView(c, view.TodoItem(tc.todos))
}
