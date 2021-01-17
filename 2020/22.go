package main

import (
     "fmt"
     "io/ioutil"
     s "strings"
     "strconv"
)

type Game struct{
     d1 []int
     d2 []int
     rounds [][2][]int
}

func getDeck(data string) []int{
     r := []int{}
     for _, v := range s.Split(data, "\n")[1:]{
          n, err := strconv.Atoi(v)
          if err == nil { r = append(r, n) }
     }
     return r
}

func (g *Game) play(){
     c1 := g.d1[0]
     c2 := g.d2[0]
     if c1 > c2 {
          g.d1 = append(g.d1, c1)
          g.d1 = append(g.d1, c2)
     } else {
          g.d2 = append(g.d2, c2)
          g.d2 = append(g.d2, c1)
     }
     g.d1 = g.d1[1:]
     g.d2 = g.d2[1:]
     if len(g.d1) !=0 && len(g.d2) !=0 {
          g.play()
     }
}

func (g *Game) score() int {
     d := []int{}
     if len(g.d1) == 0 {
          d = g.d2
     }else{
          d = g.d1
     }
     r := 0
     for k, v := range d {
          r += (len(d)-k)*v
     }
     return r
}

func isEqual(arr1 *[]int, arr2 *[]int) bool {
     if len(*arr1) != len(*arr2){ return false }
     for k, v := range *arr1 {
          if (*arr2)[k] != v {
               return false
          }
     }
     return true
}

func (g *Game) checkLoop() bool {
     for _, v := range g.rounds {
          if isEqual(&v[0], &g.d1) && isEqual(&v[1], &g.d2){
               return true
          }
     }
     return false
}

func (g *Game) playRecursive() int {
     if len(g.d1) == 0 {
          return 2
     }
     if len(g.d2) == 0 {
          return 1
     }
     if g.checkLoop() {
          return 1
     }
     c1 := g.d1[0]
     c2 := g.d2[0]
     winner := 0
     if c1 <= len(g.d1)-1 && c2 <= len(g.d2)-1{
          nG := Game{d1: getDeckTill(&g.d1, c1), d2: getDeckTill(&g.d2, c2)}
          winner = nG.playRecursive()
     } else {
          if c1 > c2 {
               winner = 1
          } else {
               winner = 2
          }
     }
     g.rounds = append(g.rounds, [2][]int{g.d1, g.d2})
     switch winner {
     case 1:
          g.d1 = append(g.d1, c1)
          g.d1 = append(g.d1, c2)
     case 2:
          g.d2 = append(g.d2, c2)
          g.d2 = append(g.d2, c1)
     }
     g.d1 = g.d1[1:]
     g.d2 = g.d2[1:]
     return g.playRecursive()
}

func getDeckTill(d *[]int, n int) []int{
     r := []int{}
     for i:=1; i < n+1 && i < len(*d); i++{
          r = append(r, (*d)[i])
     }
     return r
}

func main() {
     f, _ := ioutil.ReadFile("i_22.txt")
     data := s.Split(string(f), "\n\n")
     g1 := Game{d1: getDeck(data[0]), d2: getDeck(data[1])}
     g2 := Game{d1: getDeck(data[0]), d2: getDeck(data[1])}
     g1.play()
     fmt.Println(g1.score())
     g2.playRecursive()
     fmt.Println(g2.score())
}
