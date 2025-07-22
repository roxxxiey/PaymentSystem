package gin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetLast(c *gin.Context) {

	var request struct {
		CountTrans int `json:"count_trans" binding:"required, gt=0"`
	}
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Некорректные данные: " + err.Error(),
		})
		return
	}

	//Написать логику запроса из БД
	c.JSON(http.StatusOK, gin.H{
		"status":       "success",
		"count":        count,
		"transactions": transactions,
	})
}

func GetBalance(c *gin.Context) {

	var request struct {
		Address string  `json:"address" binding:"required"`
		Amount  float32 `json:"amount" binding:"required, gt=0"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Некорректные данные: " + err.Error(),
		})
		return
	}

	// Написать логику обработки получения баланса

	c.JSON(200, gin.H{
		"status":  "ok",
		"balance": balance,
	})
}
