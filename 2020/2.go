package main

import (
     "fmt"
     "bufio"
     "os"
     s "strings"
     "strconv"
)

func main() {
     f, _ := os.Open("i_02.txt")
     scanner := bufio.NewScanner(f)
     result := 0
     for scanner.Scan() {
          dat := s.Split(scanner.Text(), " ")
          policy := s.Split(dat[0], "-")
          min, _ := strconv.Atoi(policy[0])
          max, _ := strconv.Atoi(policy[1])
          letter := string(dat[1][0])
          password := dat[2]
          /*
          number := s.Count(password, letter)
          if  min <= number && number <= max {
               result++
          }
          */
          if (string(password[min-1]) == letter) != (string(password[max-1]) == letter) {
               result++
          }
     }
     fmt.Println(result)
}
