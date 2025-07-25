package gin

import (
	logic "PaymentSystem/internal/handlers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetLast(c *gin.Context) {

	count := c.Query("count")
	if count == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "count parameter is required"})
		return
	}

	transactions, err := logic.GetLastTransactions(count)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":       "ok",
		"count":        count,
		"transactions": transactions,
	})
}

func GetBalance(c *gin.Context) {
	address := c.Param("wname")

	balance, err := logic.GetBalance(address)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  "ok",
		"balance": balance,
	})
}
