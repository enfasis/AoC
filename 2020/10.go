package main

import (
     "fmt"
     "bufio"
     "os"
     "sort"
     "strconv"
)

func main() {
     f, _ := os.Open("i_10.txt")
     scanner := bufio.NewScanner(f)
     data := []int{}
     for scanner.Scan() {
          number, _ := strconv.Atoi(scanner.Text())
          data = append(data, number)
     }
     sort.Ints(data)
     d1 := 0
     d3 := 1
     last := 0
     for _, value := range data {
          switch value - last{
               case 1:
               d1++
               case 3:
               d3++
          }
          last = value
     }
     fmt.Println(d1*d3)
     data = append(data, data[len(data)-1]+3)
     data = append([]int{0}, data...)
     mem := map[int]int{}
     mem[0] = 1
     for _, v := range data {
          mem[v] += mem[v-1] + mem[v-2] + mem[v-3]
     }
     fmt.Println(mem[data[len(data)-1]])
}
