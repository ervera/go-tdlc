package web

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type response struct {
	Data interface{} `json:"data"`
}

type errorResponse struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func Response(c *gin.Context, status int, data interface{}) {
	c.JSON(status, data)
}

func Success(c *gin.Context, status int, data interface{}) {
	Response(c, status, response{Data: data})
}

// NewErrorf creates a new error with the given status code and the message
// formatted according to args and format.
// func Error(c *gin.Context, status int, format string, args ...interface{}) {
// 	err := errorResponse{
// 		Code:    strings.ReplaceAll(strings.ToLower(http.StatusText(status)), " ", "_"),
// 		Message: fmt.Sprintf(format, args...),
// 		Status:  status,
// 	}

// 	Response(c, status, err)
// }

// // NewErrorf creates a new error with the given status code and the message
// // formatted according to args and format.
func Error(c *gin.Context, status int, format string, args ...interface{}) {
	err := errorResponse{
		Code:    strings.ReplaceAll(strings.ToLower(http.StatusText(status)), " ", "_"),
		Message: format,
		Status:  status,
	}
	if args != nil {
		aux := ""
		for _, v := range args {
			aux = fmt.Sprintf(aux, v)
		}
		err.Message = fmt.Sprintf(format, aux)
	}
	Response(c, status, err)
}
