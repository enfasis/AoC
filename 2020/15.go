package main

import (
     "fmt"
     "bufio"
     "os"
     "strconv"
     s "strings"
)

func addNumber(m *map[int][]int, number int, turn int) int{
     newNumber := 0
     arr := (*m)[number]
     if len(arr) == 1 {
          (*m)[0] = append((*m)[0], turn)
          return 0
     }
     newNumber = arr[len(arr)-1] - arr[len(arr)-2]
     if arrN, ok := (*m)[newNumber]; ok {
          if len(arrN) == 1{
               (*m)[newNumber] = append(arrN, turn)
          } else {
               (*m)[newNumber] = append(arrN[1:], turn)
          }
     } else {
          (*m)[newNumber] = []int{turn}
     }
     return newNumber
}

func main() {
     f, _ := os.Open("i_15.txt")
     defer f.Close()
     scanner := bufio.NewScanner(f)
     data := map[int][]int{}
     turn := 1
     lastNumber := 0
     for scanner.Scan() {
          for _, v := range s.Split(scanner.Text(), ","){
               n, _ := strconv.Atoi(v)
               data[n] = []int{turn}
               lastNumber = n
               turn++
          }
     }
     for turn <= 2020 {
          lastNumber = addNumber(&data, lastNumber, turn)
          turn++
     }
     fmt.Println(lastNumber)
     for turn <= 30000000  {
          lastNumber = addNumber(&data, lastNumber, turn)
          turn++
     }
     fmt.Println(lastNumber)
}
