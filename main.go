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

	// Define a route for deserializing the JSON data.
	r.POST("/cpm", func(c *gin.Context) {
		var network Network
		// Parse the JSON from the request body into the 'network' struct.
		if err := c.ShouldBindJSON(&network); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Now 'network' contains the deserialized data.
		//graph:=network.Serilaize()
		c.JSON(http.StatusOK, network)
	})

	// Endpoint for Transpotation Problem 
	r.POST("/transpotation",func(c *gin.Context) {
		var matrix CostMatrix
		 if err := c.ShouldBindJSON(&matrix); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
        result,err:=Assignment(matrix.Matrix)
		c.JSON(http.StatusOK,result)
	})

	r.Run(":8081") 
}