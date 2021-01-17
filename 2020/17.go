package main

import (
     "fmt"
     "io/ioutil"
     s "strings"
     //"strconv"
     //"regexp"
)

type Space struct{
     data [][][]int
     cycle int
     xS int
     yS int
     zS int
}

func alterState(from *[][][]int, to *[][][]int, x,y,z int) {
     active := 0
     state := (*from)[z][x][y]
     for i := -1; i < 2; i++ {
          for j := -1; j < 2; j++ {
               for k := -1; k < 2; k++ {
                    if k == 0 && j == 0 && i == 0  {
                         continue
                    }
                    if (*from)[z+k][x+i][y+j] == 1 {
                         active++
                    }
               }
          }
     }
     if state == 1 && !(active == 2 || active == 3){
          state = 0
     } else if state == 0 && active == 3 {
          state = 1
     }
     (*to)[z][x][y] = state
}

func (sp *Space) proccess(){
     if sp.cycle == 1{
          return
     }
     nDS := createDataSpace(len(sp.data[0]), len(sp.data[0][0]), len(sp.data))
     for z := -1 ; z < sp.zS +1 ; z++{
          for x := -1 ; x < sp.xS + 1; x++{
               for y := -1 ; y < sp.yS +1 ; y++{
                    alterState(&sp.data, &nDS,  x+sp.cycle, y+sp.cycle, z+sp.cycle)
               }
          }
     }
     sp.cycle--
     sp.xS+=2
     sp.yS+=2
     sp.zS+=2
     sp.data = nDS
     sp.proccess()
}

func (sp *Space) countActive() int{
     active := 0
     for z := -1 ; z < sp.zS +1 ; z++{
          for x := -1 ; x < sp.xS + 1; x++{
               for y := -1 ; y < sp.yS +1 ; y++{
                    if sp.data[z+sp.cycle][y+sp.cycle][x+sp.cycle] == 1{
                         active++
                    }
               }
          }
     }
     return active
}

func createDataSpace(xS, yS, zS int) [][][]int{
     xyzC := [][][]int{}
     for z := 0 ; z < zS; z++{
          xyC := [][]int{}
          for x := 0 ; x < xS; x++{
               yC := []int{}
               for y := 0 ; y < yS; y++{
                    yC = append(yC, 0)
               }
               xyC = append(xyC, yC)
          }
          xyzC = append(xyzC, xyC)
     }
     return xyzC
}

func loadData(data string, cycle int)  Space{
     cycle++
     cubes := s.Split(data, "\n")
     xS := len(cubes) - 1
     yS := len(cubes[0])
     zS := 1
     dataSpace := createDataSpace(xS+2*cycle, yS+2*cycle, zS+2*cycle)
     for x, rows := range cubes {
          for y, state := range rows {
               if state == '#'{
                    dataSpace[cycle][cycle+x][cycle+y] = 1
               }
          }
     }
     return Space{data: dataSpace, cycle: cycle, xS: xS, yS: yS, zS: zS}
}

func main() {
     f, _ := ioutil.ReadFile("i_17.txt")
     cycles := 6
     space := loadData(string(f), cycles)
     space.proccess()
     fmt.Println(space.countActive())
}
