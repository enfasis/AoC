package main

import (
     "fmt"
     "bufio"
     "os"
     s "strings"
     "strconv"
)

type Ins struct {
     Id int
     Name string
     Value int
     Next *Ins
     Before *Ins
}


func (ins Ins) execute(acc *int) (*Ins, bool) {
     next := ins.Next
     err := false
     if ins.Next == nil {
          return nil, true
     }
     switch ins.Name {
     case "acc":
          *acc += ins.Value
     case "jmp":
          next, err = jump2(&ins, ins.Value)
          if err {
               return nil, true
          }
     }
     return next, false
}

func jump2(ins *Ins, step int) (*Ins, bool){
     o := 1
     if step < 0 {
          o = -1
     }
     next := ins
     for i := 0; i < o*step; i++ {
          if next == nil {
               return nil, true
          }
          if o < 0 {
               next = (*next).Before
          } else {
               next = (*next).Next
          }
     }
     return next, false
}

func findLoop(ins *Ins, acc *int) bool{
     tmp := ins
     err := false
     executed := map[int]bool{}
     for {
          tmp, err = (*tmp).execute(acc)
          if err {
               return false
          }
          if _, ok := executed[(*tmp).Id]; ok {
               return true
          } else {
               executed[(*tmp).Id] = true
          }
     }
     return false
}

func findEnding(ins *Ins, acc *int, lastId int) bool{
     tmp := ins
     err := false
     executed := map[int]bool{}
     for {
          lT := tmp
          tmp, err = (*tmp).execute(acc)
          if (*lT).Id == lastId{
               return true
          }
          if err {
               return false
          }
          if _, ok := executed[(*tmp).Id]; ok {
               return false
          } else {
               executed[(*tmp).Id] = true
          }
     }
     return false
}


func main() {
     f, _ := os.Open("i_08.txt")
     scanner := bufio.NewScanner(f)
     result1 := 0
     last := &Ins{}
     (*last).Next = &Ins{}
     first := Ins{}
     isFirst := true
     i := 0
     for scanner.Scan() {
          i++
          data := s.Split(scanner.Text(), " ")
          value, _ := strconv.Atoi(data[1])
          new := Ins{Id: i, Name: data[0], Value: value, Before: last, Next: &Ins{}}
          *(*new.Before).Next = new
          last = &new
          if isFirst == true {
               first = new
               isFirst = false
          }
     }
     if findLoop(&first, &result1){
          fmt.Println(result1)
     }
     tmp := &first
     er := false
     result2 := 0
     for {
          fmt.Println("loop for", (*tmp).Id, (*tmp).Name, (*tmp).Value, result2)
          tmp, er = (*tmp).execute(&result2)
          lT := tmp
          if er {
               break
          }
          tR := result2
          if (*tmp).Name == "nop" {
               tmp, _ = jump2(tmp, (*tmp).Value)
          } else if (*tmp).Name == "jmp" {
               tmp = (*tmp).Next
          }
          if findEnding(tmp, &tR, i){
               result2 = tR
               break
          } else {
               tmp = lT
          }
     }
     fmt.Println(result2)
}
