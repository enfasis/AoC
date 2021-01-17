package main

import (
     "fmt"
     "io/ioutil"
     s "strings"
     //"strconv"
     "regexp"
)

func sortMem(mem *map[string]string, rule string, d, max_depth int) string {
     if d > max_depth {
          return ""
     }
     result := ""
     data := s.Split((*mem)[rule], "|")
     for _, v := range data {
          numbers := regexp.MustCompile(`\d+`).FindAllString(v, -1)
          result += "|"
          if len(numbers) == 0 {
               result += s.Trim(v, `"`)
          } else{
               d++
               for _, number := range numbers {
                    result += sortMem(mem, number, d, max_depth)
               }
          }

     }
     return "(" + result[1:]  +")"
}

func main() {
     f, _ := ioutil.ReadFile("i_19.txt")
     data := s.Split(string(f), "\n\n")
     mem1 := map[string]string{}
     mem2 := map[string]string{}
     max_depth := 0
     for _, v := range s.Split(data[0], "\n") {
          data := s.Split(string(v), ": ")
          mem1[string(data[0])] = string(data[1])
          mem2[string(data[0])] = string(data[1])
     }
     for _, v := range s.Split(data[1], "\n"){
          if max_depth < len(string(v)){
               max_depth = len(string(v))
          }
     }
     sum := 0
     re := regexp.MustCompile( "^" + sortMem(&mem1, "0", 1, max_depth) + "$")
     for _, v := range s.Split(data[1], "\n"){
          if re.MatchString(string(v)){
               sum++
          }
     }
     fmt.Println(sum)

     mem2["8"] = "42 | 42 8"
     mem2["11"] = "42 31 | 42 11 31"
     sum = 0
     re = regexp.MustCompile( "^" + sortMem(&mem2, "0", 1, max_depth) + "$")
     for _, v := range s.Split(data[1], "\n"){
          if re.MatchString(string(v)){
               sum++
          }
     }
     fmt.Println(sum)
}
