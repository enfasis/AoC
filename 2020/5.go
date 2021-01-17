package main

import (
     "fmt"
     "bufio"
     "os"
     "math"
)

func getValue(data string, top string, bottom string, max float64) int {
     t := max
     b := 0.0
     position := 0
     last := len(data)-1
     for i, value := range data {
          n := float64(last-i)
          if string(value) == top {
               t -= math.Pow(2, n)
         } else if string(value) == bottom {
               b += math.Pow(2, n)
         }
     }
     if string(data[last]) == top {
          position = int(b)
     } else {
          position = int(t)
     }
     return position
}


func getId(line string, arr *[128][8]int) int {
     row := getValue(line[0:7], "F", "B", 127.0)
     col := getValue(line[7:10], "L", "R", 7.0)
     arr[row][col] = 1
     return 8 * row + col
}

func main() {
     f, _ := os.Open("i_05.txt")
     scanner := bufio.NewScanner(f)
     max := 0
     arr := [128][8]int{}
     for scanner.Scan() {
          val := getId(scanner.Text(), &arr)
          if max <= val {
               max = val
          }
     }
     fmt.Println(max)
     isDone := false
     for i := 1; i < 126; i++ {
          for j := 0; j< 6; j++ {
               if arr[i][j] == 1 && arr[i][j+1] == 0 && arr[i][j+2] == 1 {
                    isDone = true
                    fmt.Println(8*i + j+1)
                    break
               }
          }
          if isDone {
               break
          }
     }
}
