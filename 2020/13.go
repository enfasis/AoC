package main

import (
     "fmt"
     "bufio"
     "os"
     "strconv"
     s "strings"
     //"math"
)


func main() {
     f, _ := os.Open("i_13.txt")
     defer f.Close()
     scanner := bufio.NewScanner(f)
     data := []string{}
     for scanner.Scan() {
          data = append(data, scanner.Text())
     }
     time, _ := strconv.Atoi(data[0])
     buses := s.Split(s.ReplaceAll(data[1], "x,", ""), ",")
     busId := 0
     min := time
     for _, id := range buses {
          bus, _ := strconv.Atoi(id)
          mWait := bus - (time%bus)
          if min > mWait {
               min = mWait
               busId = bus
          }
     }
     fmt.Println(busId, min, busId*min)
     number := 1
     add := 1
     for k, id:= range s.Split(data[1], ","){
          if string(id) != "x"{
               bus, _ := strconv.Atoi(id)
               for (number + k) % bus != 0 {
                    number += add
               }
               add *= bus
          }
     }
     fmt.Println(number)
}
