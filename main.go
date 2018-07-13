package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func calculator(c *gin.Context) {
	ope := c.Param("operator")
	so_1, err := strconv.ParseFloat(c.Query("so1"), 64)
	so_2, err := strconv.ParseFloat(c.Query("so2"), 64)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "error input"})
	} else {
		//c.String(http.StatusOK, "kq=%f", so_1+so_2)
		if ope == "cong" {
			c.JSON(http.StatusOK, gin.H{
				"input_1": so_1,
				"input_2": so_2,
				"sum":     so_1 + so_2})
		}
		if ope == "tru" {
			c.JSON(http.StatusOK, gin.H{
				"input_1": so_1,
				"input_2": so_2,
				"sum":     so_1 - so_2})
		}
		if ope == "nhan" {
			c.JSON(http.StatusOK, gin.H{
				"input_1": so_1,
				"input_2": so_2,
				"sum":     so_1 * so_2})
		}
		if ope == "chia" {
			c.JSON(http.StatusOK, gin.H{
				"input_1": so_1,
				"input_2": so_2,
				"sum":     so_1 / so_2})
		}
	}
}

func main() {
	router := gin.Default()
	router.GET("/api/:operator", calculator)
	router.Run(":8080")
}
