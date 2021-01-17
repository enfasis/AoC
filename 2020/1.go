package main

import (
     "fmt"
     "bufio"
     "os"
     "strconv"
)

func main() {
     f, _ := os.Open("i_01.txt")
     scanner := bufio.NewScanner(f)
     dat := make([]int, 0)
     for scanner.Scan() {
          val, _ := strconv.Atoi(scanner.Text())
          dat = append(dat, val)
     }
     result := 0
     for i := 0; i < len(dat)-1; i++ {
          for j:=i+1; j <len(dat); j++ {
               v:= dat[i] + dat[j]
               if  v == 2020 {
                    result = dat[i] * dat[j]
                    break
               }
          }
          if result != 0 {
               break
          }
     }
     result = 0
     for i := 0; i < len(dat)-2; i++ {
          for j:=i+1; j <len(dat)-1; j++ {
               for k:=j+1; k < len(dat); k++ {
                    v:= dat[i] + dat[j] + dat[k]
                    if  v == 2020 {
                         fmt.Println(i, j, k)
                         result = dat[i] * dat[j] * dat[k]
                         break
                    }
               }
               if result != 0 {
                    break
               }
          }
          if result != 0 {
               break
          }
     }
     fmt.Println(result)
}
