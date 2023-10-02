package main
import (
"log"
)

type Activity struct{
	Name  string   `json:"name"`
	Previous [] string `json:previous`
	Time float32 `json:"time"`
   }

type Network struct{
	Activity []Activity  `json:"Activity"`
   }

type Graph struct {
      vertex [] *Vertex
   }

type Vertex struct {
   name string 
   timeTaken float32
	adjacent  []*Vertex
   earlyStart float32
   earlyFinish float32
   } 

//A kind of serialisation of the  Network struct to Graph Struct for study 
func (Net *Network) Serilaize() (string,float32){
      g:=&Graph{}
      vertices:=make(map[string]*Vertex)
      
      //Iterate Over the Actiivity 
      for _,r:=range Net.Activity {
            vert:=VertexFactory(r.Name,r.Time)
            append(g.vertex,&vert)
            vertices[r.Name]=vert
         }
      
      
      //ITERATE OVER THE VERTEX AND ADD ADJACENT 
      for _,act:= range Net.Activity{
         for _,pre:=range act.Previous{
            log.Println(pre," ")
            vert1:=vertices[act.Name]
            append(vert1.adjacent,vertices[pre])
            }
         }
      
      //path finder  Algorithm 
      
      //return  path and  TimeTaken 
      }

// Create Activity
func VertexFactory(name string,Times float32) *Vertex{
   return &Vertex{
      name:name,
      timeTaken:Times}
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
   