package main

import (
     "fmt"
     "bufio"
     "os"
     s "strings"
     "strconv"
     "regexp"
)

type operation func(string) int

func operate(data string) int{
     numberString := regexp.MustCompile(`\d+`).FindAllString(data, -1)
     operators := regexp.MustCompile(`[\+\*]`).FindAllString(data, -1)
     numbers := []int{}
     for _, v := range numberString {
          number, _ := strconv.Atoi(v)
          numbers = append(numbers, number)
     }
     result := numbers[0]
     for i:=1; i<len(numbers);i++{
          switch operators[i-1]{
               case "+":
               result += numbers[i]
               case "*":
               result *= numbers[i]
          }
     }
     return result
}

func operatePrecedence(data string) int{
     re := regexp.MustCompile(`\d+ \+ \d+`)
     idx := re.FindStringIndex(data)
     for len(idx) != 0{
          numberString := s.Split(data[idx[0]:idx[1]], " + ")
          n1, _ := strconv.Atoi(numberString[0])
          n2, _ := strconv.Atoi(numberString[1])
          data = data[:idx[0]] + strconv.Itoa(n1+n2) + data[idx[1]:]
          idx = re.FindStringIndex(data)
     }
     numberString := regexp.MustCompile(`\d+`).FindAllString(data, -1)
     result := 1
     for _, v := range numberString{
          n, _ := strconv.Atoi(v)
          result *= n
     }
     return result
}

func calculate(data string, fn operation) int{
     re := regexp.MustCompile(`\([\d\+\*\s]+\)`)
     idx :=  re.FindStringIndex(data)
     for len(idx) != 0 {
          val := calculate(data[idx[0]+1:idx[1]-1], fn)
          data = data[:idx[0]] + strconv.Itoa(val) + data[idx[1]:]
          idx = re.FindStringIndex(data)
     }
     return fn(data)
}


func main() {
     f, _ := os.Open("i_18.txt")
     scanner := bufio.NewScanner(f)
     sum1 := 0
     sum2 := 0
     for scanner.Scan() {
          sum1 += calculate(scanner.Text(), operate)
          sum2 += calculate(scanner.Text(), operatePrecedence)
     }
     fmt.Println(sum1, sum2)
}
