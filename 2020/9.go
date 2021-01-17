package main

import (
     "fmt"
     "bufio"
     "os"
     "strconv"
)


func main() {
     f, _ := os.Open("i_09.txt")
     scanner := bufio.NewScanner(f)
     data := []int{}
     count := 0
     preamble := 25
     result1 := 0
     for scanner.Scan() {
          number, _ := strconv.Atoi(scanner.Text())
          found := false
          if count > preamble {
               for i := len(data)-preamble; i < len(data)-1 && !found; i++ {
                    for j := i+1;j < len(data) && !found ; j++ {
                         if data[i] + data[j] == number {
                              data = append(data, number)
                              found = true
                         }
                    }
               }
               if !found {
                    result1 = number
                    break
               }
          } else {
               data = append(data, number)
          }
          count++
     }
     found := false
     result2 := 0
     for i := 0; i < len(data)-1 && !found; i++ {
          sum := 0
          for j := i; j < len(data) && sum <= result1 && !found; j++ {
               sum += data[j]
               if sum == result1 {
                    min, max := data[i], data[i]
                    for k := i; k < j ; k++{
                         if min > data[k] {
                              min = data[k]
                         }
                         if max < data[k] {
                              max = data[k]
                         }
                    }
                    result2 = min+max
                    found = true
               }
          }
     }
     fmt.Println(result1, result2)
}
