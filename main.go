package main
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CostMatrix struct {
	Matrix [][]int `json:"matrix"`
}

func main() {
	r := gin.Default()


	// Endpoint for Transpotation Problem 
	r.POST("/transpotation",)
	r.Run(":8081") 
}