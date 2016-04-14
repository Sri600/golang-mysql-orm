package controllers
// created by H.G Nuwan Indika
import (

  "github.com/op/go-logging"
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
  _ "github.com/go-sql-driver/mysql"

    "../models"
)

var logtodolist = logging.MustGetLogger("Todolist")

type TodolistController struct {
    DB *gorm.DB
}

func (ac *TodolistController) SetDB(d *gorm.DB) {
    ac.DB = d
    ac.DB.LogMode(true)
}

// Get all todolist
func (ac *TodolistController) ListTodolist(c *gin.Context) {

   results := []models.TodoList{}
  err := ac.DB.Find(&results)

    if err != nil {
        logtodolist.Debugf("Error when looking up todolist, the error is '%v'", err)
        res := gin.H{
                "status": "404",
                "error": "No todolist found",
        }
        c.JSON(404, res)
        return
    }
    content := gin.H {
            "status": "200",
            "result": "Success",
            "Todolist": results,
        }

  c.Writer.Header().Set("Content-Type", "application/json")
  c.JSON(200, content)
}

// Get a todolist
func (ac *TodolistController) GetTodolist(c *gin.Context) {
    // Grab id
    id := c.Params.ByName("id")

    todolist := models.TodoList{}
    err := ac.DB.Where("id=?", id).Find(&todolist)

    if err != nil {
        logtodolist.Debugf("Error when looking up todolist, the error is '%v'", err)
        res := gin.H{
                "status": "404",
                "error": "Todolist not found",
        }
        c.JSON(404, res)
        return
    }

    content := gin.H{
            "status": "201",
            "result": "Success",
            "Todolist": todolist,
        }

    c.Writer.Header().Set("Content-Type", "application/json")
    c.JSON(200, content)
}

// Create a todolist
func (ac *TodolistController) CreateTodolist(c *gin.Context) {

  var todolist models.TodoList

  // This will infer what binder to use depending on the content-type header.
  c.Bind(&todolist)

    // Update Timestamps
    //todolist.CreateDate = time.Now().String()
    //todolist.UpdateDate = time.Now().String()

    err := ac.DB.Save(&todolist)
    if err != nil {
        logtodolist.Debugf("Error while creating a todolist, the error is '%v'", err)
        res := gin.H{
                "status": "403",
                "error": "Unable to create todolist",
        }
        c.JSON(404, res)
        return
    }

    content := gin.H{
            "status": "201",
            "result": "Success",
            "TodolistID": todolist.ID,
        }

  c.Writer.Header().Set("Content-Type", "application/json")
  c.JSON(201, content)
}

func (ac *TodolistController) UpdateTodolist(c *gin.Context) {
    // Grab id
    id := c.Params.ByName("id")
  var todolist models.TodoList

  c.Bind(&todolist)

    // Update Timestamps
    //todolist.UpdateDate = time.Now().String()

    //err := ac.DB.Model(&models.auth_Todolist).Where("todolist_id = ?", id).Updates(&cm)
    err := ac.DB.Where("id = ?", id).Updates(&todolist)
    if err != nil {
        logtodolist.Debugf("Error while updating a todolist, the error is '%v'", err)
        res := gin.H{
                "status": "403",
                "error": "Unable to update todolist",
        }
        c.JSON(403, res)
        return
    }

    content := gin.H{
            "status": "201",
            "result": "Success",
            "TodolistID": todolist.ID,
        }

    c.Writer.Header().Set("Content-Type", "application/json")
    c.JSON(201, content)
}

func (ac *TodolistController) DeleteTodolist(c *gin.Context) {
    // Grab id
    id := c.Params.ByName("id")
  var todolist models.TodoList

  c.Bind(&todolist)

    // Update Timestamps
    //todolist.UpdateDate = time.Now().String()

    err := ac.DB.Where("id = ?", id).Delete(&todolist)
    if err != nil {
        logtodolist.Debugf("Error while deleting a todolist, the error is '%v'", err)
        res := gin.H{
                "status": "403",
                "error": "Unable to delete todolist",
        }
        c.JSON(403, res)
        return
    }

    content := gin.H {
            "result": "Success",
            "TodolistID": todolist.ID,
        }

  c.Writer.Header().Set("Content-Type", "application/json")
  c.JSON(201, content)
}