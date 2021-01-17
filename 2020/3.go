package main

import (
     "fmt"
     "bufio"
     "os"
)

func main() {
     f, _ := os.Open("i_03.txt")
     scanner := bufio.NewScanner(f)
     columns := [4]int{0}
     results := [4]int{0}
     slopes := [4]int{1, 3, 5, 7}
     row := 0
     col2 := 0
     res2:= 0
     /*
     for scanner.Scan() {
          line := scanner.Text()
          if string(line[column]) == "#"{
               result++
          }
          column = (column + 3)%31
     }
     */
     for scanner.Scan() {
          line := scanner.Text()
          for key, slope := range slopes {
               if string(line[columns[key]]) == "#"{
                    results[key]++
               }
               columns[key] = (columns[key] + slope)%31
          }

          if row % 2 == 0 {
               if string(line[col2]) == "#"{
                    res2++
               }
               col2 = (col2 + 1)%31
          }
          row++
     }
     total := 1

     for _, val := range results {
          total *= val
     }
     fmt.Println(total*res2)
}
