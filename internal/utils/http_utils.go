package utils

import (
	"github.com/alexanderlesser/golf-tracker/internal/response"
	"github.com/alexanderlesser/golf-tracker/pkg/httpstatus"
	"github.com/gin-gonic/gin"
)

// BindJSONRequest binds the JSON request to the given struct and handles errors.
func BindJSONRequest(c *gin.Context, req interface{}) bool {
	if err := c.ShouldBindJSON(req); err != nil {
		response.Error(err, "Invalid request body", httpstatus.BadRequest, c)
		return false // Indicate failure
	}
	return true // Indicate success
}
