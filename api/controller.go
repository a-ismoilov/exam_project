package api

import (
	"fmt"
	"github.com/Abdur-Rohman/exam_project/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type controllers struct {
	src IService
}

func NewController(service IService) *controllers {
	return &controllers{
		src: service,
	}
}

func (con *controllers) WriteWords() func(c *gin.Context) {
	return func(c *gin.Context) {
		var (
			body model.Items
		)
		if err := c.BindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("%v", err),
			})
			return
		}
		log.Println(body)

		if err := con.src.WriteWords(body); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("%v", err),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	}
}

func (con *controllers) ReadWords() func(c *gin.Context) {
	return func(c *gin.Context) {
		page, err := strconv.Atoi(c.Query("page"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("%v", err),
			})
			return
		}
		limit, err := strconv.Atoi(c.Query("limit"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("%v", err),
			})
			return
		}
		body, err := con.src.ReadWords(page, limit)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("%v", err),
			})
			return
		}

		c.JSON(http.StatusOK, body)
	}
}

type IServer interface {
	WriteWords() func(c *gin.Context)
	ReadWords() func(c *gin.Context)
}
