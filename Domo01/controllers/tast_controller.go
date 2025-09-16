package controllers

import (
	"gocode/Domo01/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var tasks []models.Task
var nextID = 1

func GetTasks(c *gin.Context) {
	c.JSON(http.StatusOK, tasks)
}

func AddTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task.ID = nextID
	nextID++
	tasks = append(tasks, task)
	c.JSON(http.StatusOK, task)
}

func CompleteTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Done = true
			c.JSON(http.StatusOK, tasks[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

func DeleteTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}
