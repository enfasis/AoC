package main

import (
     "fmt"
     "io/ioutil"
     s "strings"
     "strconv"
     "regexp"
)

func getRules(data string) []int{
     rules := []int{}
     reRules := regexp.MustCompile(`\d+`)
     tRules := reRules.FindAll([]byte(data), -1)
     for _, v := range tRules{
          n, _ := strconv.Atoi(string(v))
          rules = append(rules, n)
     }
     return rules
}

func getNumbers(data string) []int {
     values := s.Split(data, ",")
     tmp := []int{}
     for _, value := range values{
          n, _ := strconv.Atoi(string(value))
          tmp = append(tmp, n)
     }
     return tmp
}

func getTickets(data string) [][]int{
     tickets := [][]int{}
     rows := s.Split(data, "\n")
     for _, row := range rows {
          tmp  := getNumbers(row)
          if len(tmp) > 1{
               tickets = append(tickets, tmp)
          }
     }
     return tickets
}

func check4Rules(value int, i int, rules *[]int) bool{
     r1 := (*rules)[i] <= value && value <= (*rules)[i+1]
     r2 := (*rules)[i+2] <= value && value <= (*rules)[i+3]
     return r1 || r2
}

func validate(value int, rules *[]int) bool{
     for i:=0; i< len(*rules)/4; i++{
          if check4Rules(value, 4*i, rules) {
               return true
          }
     }
     return false
}

func scanningErrRate(rules *[]int, tickets *[][]int) (int, [][]int){
     possibleTickets := [][]int{}
     err := 0
     for _, row := range *tickets {
          hasErr := false
          for _, v := range row {
               if !validate(v, rules) {
                    err += v
                    hasErr = true
                    break
               }
          }
          if !hasErr{
               possibleTickets = append(possibleTickets, row)
          }
     }
     return err, possibleTickets
}

func removeElementsFrom(remove map[int]int, arr []int) []int{
     nC := []int{}
     for _, v := range arr {
          isIn := false
          for _, rE := range remove{
               if v == rE {
                    isIn = true
                    break
               }
          }
          if !isIn {
               nC = append(nC, v)
          }
     }
     return nC
}

func sortFields(fields *map[int][]int){
     founds := map[int]int{}
     notFounds := []int{}
     for rule, column := range *fields {
          if len(column) == 1{
               founds[rule] = column[0]
          } else {
               notFounds = append(notFounds, rule)
          }
     }
     if !(len(notFounds) == 0){
          for _, rule := range notFounds{
               (*fields)[rule] = removeElementsFrom(founds, (*fields)[rule])
          }
          sortFields(fields)
     }
}


func getFields(rules []int, tickets *[][]int) int64{
     m := map[int][]int{}
     for rule:=0; rule<len(rules)/4; rule++{
          fields := []int{}
          for col:=0; col < len((*tickets)[0]); col++{
               noSkip := true
               for row:=0; row < len((*tickets)) && noSkip; row++{
                    val := (*tickets)[row][col]
                    noSkip = check4Rules(val, 4*rule, &rules)
               }
               if noSkip {
                    fields = append(fields, col)
               }
          }
          m[rule] = fields
     }
     mul := int64(1)
     sortFields(&m)
     for i:=0;i<6;i++{
          mul *= int64((*tickets)[len(*tickets)-1][m[i][0]])
     }
     return mul
}

func main() {
     f, _ := ioutil.ReadFile("i_16.txt")
     data := s.Split(string(f), "\n\n")
     rules := getRules(data[0])
     tickets := getTickets(data[2])
     errRate, validated := scanningErrRate(&rules, &tickets)
     validated = append(validated, getNumbers(s.Split(data[1], ":\n")[1]))
     fmt.Println(errRate, getFields(rules, &validated))
}
