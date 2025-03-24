package response

import (
	"github.com/alexanderlesser/golf-tracker/pkg/httpstatus"
	"github.com/gin-gonic/gin"
)

// Response defines the structure of a standard API response.
type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`  //omitempty will not show data if data is nil.
	Error   interface{} `json:"error,omitempty"` //omitempty will not show error if error is nil.
}

func Error(data interface{}, message string, statusCode int, c *gin.Context) {
	c.JSON(statusCode, Response{
		Status:  httpstatus.StatusNames[statusCode],
		Error:   data,
		Message: message,
	})
}

func Success(data interface{}, message string, statusCode int, c *gin.Context) {
	response := Response{
		Status:  httpstatus.StatusNames[statusCode],
		Message: message,
		Data:    data,
	}
	c.JSON(statusCode, response)
}
