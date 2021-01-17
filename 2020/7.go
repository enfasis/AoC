package main

import (
     "fmt"
     "bufio"
     "os"
     s "strings"
     "strconv"
)

type Bag struct {
     Color string
     Contains map[*Bag]int
     Parents map[*Bag]bool
}

func (p Bag) addChild(child *Bag, number int){
     p.Contains[child] = number
     (*child).Parents[&p] = true
}

func getParents(m *map[string]bool, n Bag) *map[string]bool{
     if n.Parents == nil{
          return m
     }
     for k, _ := range n.Parents {
          (*m)[(*k).Color] = true
          m = getParents(m, *k)
     }
     return m
}

func getChildren(n int, c Bag) int{
     if c.Contains == nil {
          return 0
     }
     for key, value := range c.Contains {
          n += value*getChildren(1, *key)
     }
     return n
}

func newBag(color string) Bag{
     bag := Bag{Color: color}
     bag.Contains = map[*Bag]int{}
     bag.Parents = map[*Bag]bool{}
     return bag
}

func main() {
     f, _ := os.Open("i_07.txt")
     scanner := bufio.NewScanner(f)
     bags := map[string]*Bag{}
     for scanner.Scan() {
          data := s.Split(scanner.Text(), " contain ")
          parentColor := s.Replace(data[0], " bags", "", 1)
          if _, ok := bags[parentColor]; !ok {
               parent := newBag(parentColor)
               bags[parentColor] = &parent
          }
          parent := *bags[parentColor]

          childrenData := s.Split(data[1], ", ")
          for _, childData:= range childrenData{
               d := s.Split(childData, " ")
               if d[0] == "no" {
                    continue
               }
               childColor:= d[1] + " " + d[2]
               number, _:= strconv.Atoi(d[0])
               if _, ok := bags[childColor]; !ok {
                    child := newBag(childColor)
                    bags[childColor] = &child
               }
               parent.addChild(bags[childColor], number)
          }
     }
     m := map[string]bool{}
     result1 := len(*getParents(&m, *bags["shiny gold"]))
     result2 := getChildren(0, *bags["shiny gold"])
     fmt.Println(result1, result2)
}
