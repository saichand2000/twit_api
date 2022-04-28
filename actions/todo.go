package actions

import (
	"fmt"
	"net/http"
	"twit_api/models"

	"github.com/gobuffalo/buffalo"
)

// TodoIndex default implementation.
func TodoIndex(c buffalo.Context) error {
	// Create an array to receive todos
	todos := []models.Todo{}
	//get all the todos from database
	err := models.DB.All(&todos)
	// handle any error
	if err != nil {
		return c.Render(http.StatusOK, r.JSON(err))
	}
	//return list of todos as json
	return c.Render(http.StatusOK, r.JSON(todos))
}

// TodoShow default implementation.
func TodoShow(c buffalo.Context) error {
	// grab the id url parameter defined in app.go
	id := c.Param("id")
	// create a variable to receive the todo
	todo := models.Todo{}
	// grab the todo from the database
	err := models.DB.Find(&todo, id)
	// handle possible error
	if err != nil {
		fmt.Println(err)
		return c.Render(http.StatusOK, r.JSON(err))
	}
	fmt.Println(todo)
	//return the data as json
	return c.Render(http.StatusOK, r.JSON(todo))
}

// TodoAdd default implementation.
func TodoAdd(c buffalo.Context) error {

	//get item from url query
	item := c.Param("item")

	//create new instance of todo
	todo := &models.Todo{Item: item}

	// Create a fruit without running validations
	err := models.DB.Create(todo)

	fmt.Println(todo)

	// handle error
	if err != nil {
		return c.Render(http.StatusOK, r.JSON(err))
	}

	//return new todo as json
	return c.Render(http.StatusOK, r.JSON(todo))
}
