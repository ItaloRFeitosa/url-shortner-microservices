package http_helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type invalidArgument struct {
	Field  string `json:"field"`
	Value  string `json:"value"`
	Reason string `json:"reason"`
}

func ValidationErrorResponse(c *gin.Context, err error) {
	if errs, ok := err.(validator.ValidationErrors); ok {
		var invalidArgs []invalidArgument

		for _, err := range errs {
			invalidArgs = append(invalidArgs, invalidArgument{
				err.Field(),
				err.Value().(string),
				err.Field() + " is " + err.Tag(),
			})
		}

		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":  "ValidationError",
			"reason": invalidArgs,
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"error": err,
	})
	c.Abort()
}
