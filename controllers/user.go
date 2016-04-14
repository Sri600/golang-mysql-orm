package controllers
// created by H.G Nuwan Indika
import (
  //"time"

  "github.com/op/go-logging"
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
  _ "github.com/go-sql-driver/mysql"

    "../models"
)

var loguser = logging.MustGetLogger("Todolist")

type TodolistControllerUser struct {
    DB *gorm.DB
}

func (ac *TodolistControllerUser) SetDB(d *gorm.DB) {
    ac.DB = d
    ac.DB.LogMode(true)
}

// Get all user
func (ac *TodolistControllerUser) ListUsers(c *gin.Context) {

   results := []models.Users{}
  err := ac.DB.Find(&results)

    if err != nil {
        loguser.Debugf("Error when looking up user, the error is '%v'", err)
        res := gin.H{
                "status": "404",
                "error": "No user found",
        }
        c.JSON(404, res)
        return
    }
    content := gin.H {
            "status": "200",
            "result": "Success",
            "User": results,
        }

  c.Writer.Header().Set("Content-Type", "application/json")
  c.JSON(200, content)
}

// Get a user
func (ac *TodolistControllerUser) GetUsers(c *gin.Context) {
    // Grab id
    id := c.Params.ByName("id")

    user := models.Users{}
    err := ac.DB.Where("ID=?", id).Find(&user)

    if err != nil {
        loguser.Debugf("Error when looking up user, the error is '%v'", err)
        res := gin.H{
                "status": "404",
                "error": "User not found",
        }
        c.JSON(404, res)
        return
    }

    content := gin.H{
            "status": "201",
            "result": "Success",
            "User": user,
        }

    c.Writer.Header().Set("Content-Type", "application/json")
    c.JSON(200, content)
}

// Create a user
func (ac *TodolistControllerUser) CreateUsers(c *gin.Context) {

  var user models.Users

  // This will infer what binder to use depending on the content-type header.
  c.Bind(&user)

    // Update Timestamps
    //user.CreateDate = time.Now().String()
    //user.UpdateDate = time.Now().String()

    err := ac.DB.Save(&user)
    if err != nil {
        loguser.Debugf("Error while creating a user, the error is '%v'", err)
        res := gin.H{
                "status": "403",
                "error": "Unable to create user",
        }
        c.JSON(404, res)
        return
    }

    content := gin.H{
            "status": "201",
            "result": "Success",
            "UserID": user.ID,
        }

  c.Writer.Header().Set("Content-Type", "application/json")
  c.JSON(201, content)
}

func (ac *TodolistControllerUser) UpdateUsers(c *gin.Context) {
    // Grab id
    id := c.Params.ByName("id")
  var user models.Users

  c.Bind(&user)

    // Update Timestamps
    //user.UpdateDate = time.Now().String()

    //err := ac.DB.Model(&models.auth_User).Where("user_id = ?", id).Updates(&cm)
    err := ac.DB.Where("ID = ?", id).Updates(&user)
    if err != nil {
        loguser.Debugf("Error while updating a user, the error is '%v'", err)
        res := gin.H{
                "status": "403",
                "error": "Unable to update user",
        }
        c.JSON(403, res)
        return
    }

    content := gin.H{
            "status": "201",
            "result": "Success",
            "UserID": user.ID,
        }

    c.Writer.Header().Set("Content-Type", "application/json")
    c.JSON(201, content)
}

func (ac *TodolistControllerUser) DeleteUsers(c *gin.Context) {
    // Grab id
    id := c.Params.ByName("id")
  var user models.Users

  c.Bind(&user)

    // Update Timestamps
    //user.UpdateDate = time.Now().String()

    err := ac.DB.Where("ID = ?", id).Delete(&user)
    if err != nil {
        loguser.Debugf("Error while deleting a user, the error is '%v'", err)
        res := gin.H{
                "status": "403",
                "error": "Unable to delete user",
        }
        c.JSON(403, res)
        return
    }

    content := gin.H {
            "result": "Success",
            "UserID": user.ID,
        }

  c.Writer.Header().Set("Content-Type", "application/json")
  c.JSON(201, content)
}