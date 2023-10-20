package main
import (
	"math"
	"encoding/json"
)

type JSONData struct {
    Data [][] int `json:"data"`
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




func makeSquareMatrix(matrix [][] byte) [][]byte{
	var square bool 
	square=true
	max_length:=math.MinInt64

	if len(matrix)>max_length {
		max_length=len(matrix)
	}
	//iterate over the matrix to make the array  square matrix 
	for i:=0;i<len(matrix);i++{
		if len(matrix[0])!=len(matrix[i]){
			max_length=math.max(max_length,len(matrix[i]))
			square=false
		}
	}

    if square!=true{
        //coloum setting up 
		for i:=0;i<len(matrix);i++{
		 if len(matrix[i])<max_length{
			for j:=0;j<(max_length-matrix[i]);j++{
				append(matrix[i],0)
			}
		   }
	     }
		
		// row setup 
		for i:=0;i<max_length-matrix[i];i++{
			for j:=0;j<max_length;j++{
				append(matrix[len(matrix)+i],0)
			}
		}
	   return matrix
	       }	
	  return matrix
}