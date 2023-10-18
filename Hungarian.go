package main
import (
	"math"
	"encoding/json"
)

type JSONData struct {
    Data [][]string `json:"data"`
}
/*
jsonStr := `{"data": [["row1col1", "row1col2"], ["row2col1", "row2col2"]]}`

var jsonData JSONData

err := json.Unmarshal([]byte(jsonStr), &jsonData)
if err != nil {
        fmt.Println("Error:", err)
        return
    }
twoDArray := jsonData.Data
fmt.Println(twoDArray)
*/

func Assignment(matrix CostMatrix.Matrix) int,error {
	mat:=makeSquareMatrix(&matrix)
	
	//perform row minimum
	for i:=0;i<len(matrix);i++{
		min:=findMinRow(mat[i])
		for j:=0;j<len(mat[i]);j++{
			mat[i][j]-=min
		}
	}

	//perform coloum minimum
    for i:=0;i<len(mat[i]);i++{
		min:=findMinRow(mat[][i])
		for j:=0;j<len(matrix);j++{
			mat[j][i]-=min
		}
	}
    
    //Assignment
	

}


//find min in the array 
func findMinRow(arr []int) int {
   min:=math. MaxInt64
   for x,_:=range arr {
      if min>x {
		min=x
	  }
   }
   return min
}

