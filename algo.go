package main
import "fmt"

type Activity struct {
Name       string
Duration   int
Predecessors []*Activity
Successors   []*Activity
}


//place predecessors to the particular task 
func (task *Activity) AddPredecessor(pre *Activity){
     append(task.Predecessors,pre)
   }

/*{
		
      "Activity":[
      {
        "name":"A",
        "previous":[],
       "time":4
     },
     {
        "name":"B",
         "previous":["A"],
         "time":58
      },
      {
         "name":"C",
         "previous":["B"],
         "time":2
      },
      {
         "name":"D",
         "previous":["B"],
         "time":2
      },
      {,
         "name":"E",
         "previous":["C"],
      
      {
         "name":"F",
         "previous":["D"],
         "time":1
      }]
  
}
   */
   
