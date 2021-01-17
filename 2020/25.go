package main

import (
     "fmt"
     "io/ioutil"
     s "strings"
     "strconv"
)

func transform(sn, loopSize int) int {
     val := 1
     for i:= 0; i < loopSize; i++ {
          val = (val*sn)%20201227
     }
     return val
}

func getLoopSize(sn, pk int) int{
     ls := 0
     val := 1
     for val != pk {
          val = (val*sn)%20201227
          ls++
     }
     return ls
}

func main() {
     f, _ := ioutil.ReadFile("i_25.txt")
     data := s.Split(string(f), "\n")
     cardPk, _ := strconv.Atoi(data[0])
     doorPk, _ := strconv.Atoi(data[1])
     cardLs := getLoopSize(7, cardPk)
     doorLs := getLoopSize(7, doorPk)
     fmt.Println(transform(cardPk, doorLs))
}
