package main

import (
     "fmt"
     "bufio"
     "os"
)

type counter func(arr *[][]int, x int, y int, e int) int

func countElement(arr *[][]int, x int, y int, e int) int{
     occ := 0
     if (*arr)[x+1][y-1] == e { occ++ }
     if (*arr)[x][y-1] == e { occ++ }
     if (*arr)[x-1][y-1] == e { occ++ }

     if (*arr)[x+1][y+1] == e { occ++ }
     if (*arr)[x][y+1] == e { occ++ }
     if (*arr)[x-1][y+1] == e { occ++ }

     if (*arr)[x-1][y] == e { occ++ }
     if (*arr)[x+1][y] == e { occ++ }
     return occ
}

func countFirstElement(arr *[][]int, x int, y int, e int) int{
     occ := 0
     rows := len((*arr))
     cols := len((*arr)[0])
     // bottom
     for i:=x+1; i < rows; i++ {
          v := (*arr)[i][y]
          if v == e {
               occ++
               break
          } else if v != -1 {
               break
          }
     }
     // top
     for i:=x-1; 0 < i; i-- {
          v := (*arr)[i][y]
          if v == e {
               occ++
               break
          } else if v != -1 {
               break
          }
     }
     // right
     for i:=y+1; i < cols; i++ {
          v := (*arr)[x][i]
          if v == e {
               occ++
               break
          } else if v != -1 {
               break
          }
     }
     // left
     for i:=y-1; 0 < i; i-- {
          v := (*arr)[x][i]
          if v == e {
               occ++
               break
          } else if v != -1 {
               break
          }
     }
     // bottom right
     for i:=1; x+i < rows && y+i < cols; i++ {
          v := (*arr)[x+i][y+i]
          if v == e {
               occ++
               break
          } else if v != -1 {
               break
          }
     }
     // top left
     for i:=1; x-i > 0 && y-i > 0; i++ {
          v := (*arr)[x-i][y-i]
          if v == e {
               occ++
               break
          } else if v != -1 {
               break
          }
     }
     // top right
     for i:=1; x-i > 0 && y+i < cols; i++ {
          v := (*arr)[x-i][y+i]
          if v == e {
               occ++
               break
          } else if v != -1 {
               break
          }
     }
     // bottom left
     for i:=1; x+i < rows && y-i > 0; i++ {
          v := (*arr)[x+i][y-i]
          if v == e {
               occ++
               break
          } else if v != -1 {
               break
          }
     }

     return occ
}

func rule(arr *[][]int, x int, y int, fn counter, occ int) (int, bool){
     switch (*arr)[x][y] {
     case 0:
          if fn(arr, x, y, 1) == 0 {
               return 1, true
          }
     case 1:
          if fn(arr, x, y, 1) >= occ {
               return 0, true
          }
     }
     return (*arr)[x][y], false
}

func addTB(arr *[][]int){
     a := []int{}
     b := []int{}
     for i:= 0; i<len((*arr)[0]); i++{
          a = append(a, -1)
          b = append(b, -1)
     }
     *arr = append(*arr, a)
     *arr = append([][]int{b}, *(arr)...)
}

func applyRules(arr *[][]int, fn counter, occ int) ([][]int, bool){
     rows := len((*arr))
     cols := len((*arr)[0])
     newArr := [][]int{}
     changed := 0
     for i:=1; i < rows-1; i++{
          r := []int{-1}
          for j:=1; j < cols-1; j++{
               v, ch := rule(arr, i, j, fn, occ)
               if ch {
                    changed++
               }
               r = append(r, v)
          }
          r = append(r, -1)
          newArr = append(newArr, r)
     }
     addTB(&newArr)
     return newArr, changed > 0
}


func countOccupied(arr *[][]int) int{
     rows := len((*arr))
     cols := len((*arr)[0])
     sum := 0
     for i:=1; i < rows-1; i++{
          for j:=1; j < cols-1; j++{
               if (*arr)[i][j] == 1 {
                    sum++
               }
          }
     }
     return sum
}


func main() {
     f, _ := os.Open("i_11.txt")
     scanner := bufio.NewScanner(f)
     seats := [][]int{}
     for scanner.Scan() {
          line := scanner.Text()
          s := []int{-1}
          for _, v := range line {
               n := 0
               switch string(v){
                    case ".":
                    n = -1
                    case "#":
                    n = 1
                    case "L":
                    n = 0
               }
               s = append(s, n)
          }
          s = append(s, -1)
          seats = append(seats, s)
     }
     addTB(&seats)
     c := seats
     changed := true
     for changed {
          c, changed = applyRules(&c, countElement, 4)
     }
     fmt.Println(countOccupied(&c), changed)
     changed = true
     c = seats
     for changed {
          c, changed = applyRules(&c, countFirstElement, 5)
     }
     fmt.Println(countOccupied(&c), changed)
}
