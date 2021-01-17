package main

import (
     "fmt"
     "bufio"
     "os"
     s "strings"
     "strconv"
     "regexp"
)

func check(count int, hasCid bool) bool {
     if count == 8 {
          return true
     } else if count == 7 && !hasCid {
          return true
     } else {
          return false
     }
}
func isBetween(val string, n1 int, n2 int) bool {
     v, err := strconv.Atoi(val)
     if err != nil {
          return false
     }
     return n1 <= v && v <= n2

}

func checkLine(line string) bool {
     dat := s.Split(s.Trim(line, " "), " ")
     for _, value := range dat {
          ndat := s.Split(value, ":")
          if len(ndat) != 2{
               return false
          }
          code := ndat[1]
          f := true
          switch ndat[0] {
          case "byr":
                f = isBetween(code, 1920, 2002)
          case "iyr":
                f = isBetween(code, 2010, 2020)
          case "eyr":
                f = isBetween(code, 2020, 2030)
          case "hgt":
               t := code[len(code)-2:]
               n := code[:len(code)-2]
               if  t == "cm" {
                    f = isBetween(n, 150, 193)
               } else if t == "in" {
                    f = isBetween(n, 59, 76)
               } else {
                    f = false
               }
          case "hcl":
               validExp := regexp.MustCompile(`#[0-9a-f]{6}`)
               f = validExp.MatchString(code)
          case "ecl":
               colors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
               f = false
               for _, v:= range colors {
                    if v == code {
                         f = true
                         break
                    }
               }
          case "pid":
               validExp := regexp.MustCompile(`[0-9]{9}`)
               f = validExp.MatchString(code) && len(code) == 9
          }
          if !f {
               return false
          }

     }
     return true
}


func main() {
     f, _ := os.Open("i_04.txt")
     scanner := bufio.NewScanner(f)
     result := 0
     count := 0
     result2 := 0
     hasCid := false
     line := ""
     for scanner.Scan() {
          text := scanner.Text()
          if text == "" {
               if check(count, hasCid) {
                    result++
                    if checkLine(line) {
                         result2++
                    }
               }
               line = ""
               count = 0
               hasCid = false
               continue
          } else {
               line += (" " + text)
          }
          if !hasCid {
               hasCid = s.Contains(text, "cid:")
          }
          count += s.Count(text, ":")
     }
     if check(count, hasCid) {
          result++
          if checkLine(line){
               result2++
          }
     }
     fmt.Println(result, result2)
}
