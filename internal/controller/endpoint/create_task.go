package endpoint

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/heloayer/to-do-list/pkg/mongo"
	"gopkg.in/mgo.v2/bson"
)

type Task struct {
	Title    string    `json:"title" binding:"required,max=200"`
	ActiveAt time.Time `json:"activeAt" binding:"required"`
}

// Обработчик POST запроса на /api/todo-list/tasks
func CreateTask(c *gin.Context) {

	mongo := &mongo.Mongo{}

	var task Task
	if err := c.ShouldBindJSON(&task); err != nil {
		// Возвращаем ошибку валидации
		if _, ok := err.(validator.ValidationErrors); ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "incorrect query"})
		}
		return
	}

	// Проверяем уникальность записи
	count, err := mongo.Collection.CountDocuments(context.Background(), bson.M{"title": task.Title, "activeAt": task.ActiveAt})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сервера"})
		return
	}
	if count > 0 {
		c.Status(http.StatusNoContent)
		return
	}

	// Добавляем запись в базу данных
	if _, err := mongo.Collection.InsertOne(context.Background(), task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сервера"})
		return
	}

	c.Status(http.StatusNoContent)

}
