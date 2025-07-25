package gin

import (
	logic "PaymentSystem/internal/handlers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostSend(c *gin.Context) {

	var request struct {
		FromTrans string  `json:"from_trans" binding:"required"`
		ToTrans   string  `json:"to_trans" binding:"required"`
		Amount    float32 `json:"amount" binding:"required,gt=0"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Некорректные данные: " + err.Error(),
		})
		return
	}

	err := logic.SendMoney(request.FromTrans, request.ToTrans, request.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  "ok",
		"message": "Transaction success",
		"details": request,
	})
}
