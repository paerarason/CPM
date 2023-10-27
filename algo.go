
package main

import "fmt"

type Activity struct {
Name       string
Duration   int
Predecessors []*Activity
Successors   []*Activity
// You can include other fields as needed.
}


func main() {
// Create activities
A := &Activity{Name: "A", Duration: 3}
B := &Activity{Name: "B", Duration: 4}
C := &Activity{Name: "C", Duration: 2}
D := &Activity{Name: "D", Duration: 5}
E := &Activity{Name: "E", Duration: 2}

// Define dependencies
B.Predecessors = append(B.Predecessors, A)
C.Predecessors = append(C.Predecessors, A)
D.Predecessors = append(D.Predecessors, B, C)
E.Predecessors = append(E.Predecessors, D)

// Update successors
A.Successors = append(A.Successors, B, C)
B.Successors = append(B.Successors, D)
C.Successors = append(C.Successors, D)
D.Successors = append(D.Successors, E)

// Now you have a CPM graph where each activity can have multiple predecessors and successors.

fmt.Printf("Activity A Predecessors: %v, Successors: %v\n", A.Predecessors, A.Successors)
fmt.Printf("Activity D Predecessors: %v, Successors: %v\n", D.Predecessors, D.Successors)
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
         "time":1
      },
      {
         "name":"F",
         "previous":["D"],
         "time":1
      }]
   }
   */
   