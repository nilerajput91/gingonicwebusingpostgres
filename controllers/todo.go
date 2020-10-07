package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v9"
	orm "github.com/go-pg/pg/v9/orm"
	guuid "github.com/google/uuid"
)

// Todo struct for todo information filled
type Todo struct {
	ID string `json:"id"`

	Title string `json:"title"`

	Body string `json:"body"`

	Completed string `json:"completed"`

	CreatedAt time.Time `json:"created_at"`

	UpdatedAt time.Time `json:"updated_at"`
}

// CreateTodoTable func for create the table on db
func CreateTodoTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}

	createError := db.CreateTable(&Todo{}, opts)

	if createError != nil {
		log.Printf("Error while creating the todo table,Reason: %v\n", createError)
		return createError
	}

	log.Printf("Todo table is created")

	return nil
}

var dbConnect *pg.DB

// InitiateDB Initialize DB connections (to avoid the too many connection)
func InitiateDB(db *pg.DB) {
	dbConnect = db
}

// GetAllTodos func for getting the all todos from db
func GetAllTodos(c *gin.Context) {
	var todos []Todo
	err := dbConnect.Model(&todos).Select()

	if err != nil {
		log.Printf("Error while getting all todos,Reason:%v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "something went wrong ",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All Todos",
		"data":    todos,
	})

	return
}

// CreateTodo func to create the todo
func CreateTodo(c *gin.Context) {
	var todo Todo

	c.BindJSON(&todo)
	title := todo.Title
	body := todo.Body
	completed := todo.Completed
	id := guuid.New().String()

	insertError := dbConnect.Insert(&Todo{
		ID:        id,
		Title:     title,
		Body:      body,
		Completed: completed,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if insertError != nil {
		log.Printf("Error while inserting new todo into db,Reason:%v\n ", insertError)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "something wet wrong ",
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Todo Created Successfully",
	})

	return

}

// GetSingleTodo for get single todo
func GetSingleTodo(c *gin.Context) {
	todoId := c.Param("todoId")
	todo := &Todo{ID: todoId}
	err := dbConnect.Select(todo)

	if err != nil {
		log.Printf("Error while getting the single todo,Reason %v/n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "single todo not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Single Todo",
		"data":    todo,
	})

	return
}

// EditTodo for editing the single todo
func EditTodo(c *gin.Context) {
	todoId := c.Param("todoId")

	var todo Todo

	c.BindJSON(&todo)
	completed := todo.Completed

	_, err := dbConnect.Model(&Todo{}).Set("completed=?", completed).Where("id=?", todoId).Update()

	if err != nil {
		log.Printf("Error,Reason:%v\n", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "something went wrong",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Todo Edited Successfully",
	})

	return
}

// DeleteTodo func for delete the single todo
func DeleteTodo(c *gin.Context) {

	todoId := c.Param("todoId")
	todo := &Todo{ID: todoId}

	err := dbConnect.Delete(todo)
	if err != nil {
		log.Printf("Error while deleting a single todo,Reason %v/n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Todo Deleted Successfully",
	})
	return
}
