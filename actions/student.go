package actions

import (
	"fmt"
	"net/http"
	"strconv"
	"twit_api/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
	//"github.com/gobuffalo/buffalo/Worker"
)

func Signin(c buffalo.Context) error {

	student := &models.Student{Username: c.Param("student[username]"), Password: c.Param("student[password]")}

	if err := c.Bind(student); err != nil {
		return nil
	}

	err := student.Authorize()

	if err != nil {
		fmt.Println("Error: %s", err)
	}
	return c.Render(http.StatusOK, r.JSON(student))
}

func (u Student) Authorize() error {
	err := DB.Where("username = ?", string.ToLower(u.Username)).First(u)

	if err != nil {
		fmt.Println("Error: %s", err)
	}

	err = DB.Where("password = ?", string.ToLower(u.Password)).First(u)

	if err != nil {
		fmt.Println("Error: %s", err)
	}
	return nil

}

// StudentIndex default implementation.
func StudentIndex(c buffalo.Context) error {
	// Create an array to receive students
	students := []models.Student{}
	//get all the todos from database
	err := models.DB.All(&students)
	// handle any error
	if err != nil {
		return c.Render(http.StatusOK, r.JSON(err))
	}
	//return list of todos as json
	return c.Render(http.StatusOK, r.JSON(students))
}

// StudentShow default implementation.
func StudentShow(c buffalo.Context) error {
	// grab the id url parameter defined in app.go
	id := c.Param("id")
	// create a variable to receive the todo
	student := models.Student{}
	// grab the todo from the database
	err := models.DB.Find(&student, id)
	// handle possible error
	if err != nil {
		fmt.Println(err)
		return c.Render(http.StatusOK, r.JSON(err))
	}
	fmt.Println(student)
	//return the data as json
	return c.Render(http.StatusOK, r.JSON(student))
}

func StudentAdd(c buffalo.Context) error {

	//get item from url query
	name := c.Param("name")
	email := c.Param("email")
	username := c.Param("username")
	password := c.Param("password")
	// id, _ := strconv.Atoi(c.Param("id"))

	//create new instance of student
	student := &models.Student{Name: name, Email: email, Username: username, Password: password}

	//Create a fruit without running validations
	//err := models.DB.Create(student)
	fmt.Println(models.DB)

	fmt.Println(student)
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("transaction failed")
	}
	err := tx.Create(student)
	fmt.Println(student)

	// handle error
	if err != nil {
		fmt.Println("HHHHHHHHHHHHHH")
		fmt.Println(err)
		return c.Render(http.StatusUnprocessableEntity, r.JSON(err))
	}

	//return new todo as json
	return c.Render(http.StatusOK, r.JSON(student))
}

func StudentDelete(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	student := &models.Student{}
	if err := tx.Find(student, c.Param("id")); err != nil {
		return c.Error(404, err)
	}
	if err := tx.Destroy(student); err != nil {
		//
	}

	return c.Redirect(302, "/students/index")
}

// func StudentUpdate(c buffalo.Context) error {
//     // grab the id url parameter defined in app.go
//      name := c.Param("name")
//     id, _ := strconv.Atoi(c.Param("id"))
//     // create a variable to receive the todo
//     student := &models.Student{Name: name, ID: id}
//     // grab the todo from the database
//     err := models.DB.Find(&student, id)
//     // handle possible error
//     fmt.Println(models.DB)

//     fmt.Println(student)
//     tx, ok := c.Value("tx").(*pop.Connection)
//     if !ok {
//         return fmt.Errorf("transaction failed")
//     }

//      err := tx.Update(student)
//     if err != nil {
//         fmt.Println(err)
//         return c.Render(http.StatusOK, r.JSON(err))
//     }
//     fmt.Println(student)
//     //return the data as json
//     return c.Render(http.StatusOK, r.JSON(student))
// }

func StudentUpdate(c buffalo.Context) error {

	//get item from url query
	name := c.Param("name")
	id, _ := strconv.Atoi(c.Param("id"))

	//create new instance of student
	student := &models.Student{Name: name, ID: id}

	//Create a fruit without running validations
	//err := models.DB.Create(student)
	fmt.Println(models.DB)

	fmt.Println(student)
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("transaction failed")
	}
	err := tx.Update(student)
	fmt.Println(student)

	// handle error
	if err != nil {
		fmt.Println("HHHHHHHHHHHHHH")
		fmt.Println(err)
		return c.Render(http.StatusUnprocessableEntity, r.JSON(err))
	}

	//return new todo as json
	return c.Render(http.StatusOK, r.JSON(student))
}

// func StudentUpdate(c buffalo.Context) error {
//     name := c.Param("name")
//     tx := c.Value("tx").(*pop.Connection)
//     student := &models.Student{Name: name}
//     if err := tx.Find(student, c.Param("id")); err != nil {
//         return c.Error(404, err)
//     }
//     if err := tx.Update(student); err != nil {
//        //
//     }

//     return c.Redirect(302, "/students/index")
// }
