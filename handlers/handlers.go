package handlers

import (
	"strconv"

	"github.com/Evelon44/MyTests/GormAPI/helpers"
	"github.com/Evelon44/MyTests/GormAPI/models"
	"github.com/gin-gonic/gin"
)

func GetAllToDo(c *gin.Context) {
	var todoList []models.ToDo
	err := models.GetAllToDo(&todoList)
	if err != nil {
		helpers.WriteJSON(c, 404, todoList)
		return
	}
	helpers.WriteJSON(c, 200, todoList)
}

func PostNewToDo(c *gin.Context) {
	var todo models.ToDo
	err := c.BindJSON(&todo)
	if err != nil {
		helpers.WriteJSON(c, 400, "json is invalid")
		return
	}
	err = models.AddNewToDo(&todo)
	if err != nil {
		helpers.WriteJSON(c, 404, todo)
		return
	}
	helpers.WriteJSON(c, 201, todo)
}

func GetToDoById(c *gin.Context) {
	var todo models.ToDo
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helpers.WriteJSON(c, 404, "invalid id format. Try again")
		return
	}
	err = models.GetToDoById(&todo, id)
	if err != nil {
		helpers.WriteJSON(c, 404, todo)
		return
	}
	helpers.WriteJSON(c, 200, todo)
}

func UpdateToDoById(c *gin.Context) {
	var todo models.ToDo
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helpers.WriteJSON(c, 404, "invalid id format. Try again")
		return
	}
	err = c.BindJSON(&todo)
	if err != nil {
		helpers.WriteJSON(c, 400, "json is invalid")
		return
	}
	err = models.UpdateToDoById(&todo, id)
	if err != nil {
		helpers.WriteJSON(c, 404, todo)
		return
	}
	helpers.WriteJSON(c, 202, todo)
}

func DeleteToDoById(c *gin.Context) {
	var todo models.ToDo
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helpers.WriteJSON(c, 404, "invalid id format. Try again")
		return
	}
	err = models.DeleteToDoById(&todo, id)
	if err != nil {
		helpers.WriteJSON(c, 404, todo)
		return
	}
	helpers.WriteJSON(c, 202, todo)
}
