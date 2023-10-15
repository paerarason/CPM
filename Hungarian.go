package main
import "math"
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

//stopped
func makeSquareMatrix(matrix *[][]int) *[][]int{
     for _,i:=range len(matrix){
		for _,j:=range len(matrix[i]){
			if len(matrix)==len(matrix[i]){
                
			}
		}
	 }   
}