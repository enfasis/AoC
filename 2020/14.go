package main

import (
     "fmt"
     "bufio"
     "os"
     "strconv"
     s "strings"
     "regexp"
)

func sumMem(m map[int64]int64) int64 {
     sum := int64(0)
     for _, v := range m {
          sum += v
     }
     return sum
}

func applyMask(address int64, mask string) []int64{
     addresses := []int64{}
     if len(mask) == 0 {
          addresses = append(addresses, address)
          return addresses
     }
     change_bit := int64(1) << (len(mask)-1)
     tails := applyMask(address, mask[1:])
     switch mask[0] {
     case '0':
          addresses = tails
     case '1':
          for _, v:= range tails {
               addresses = append(addresses, v | change_bit)
          }
     case 'X':
          for _, v:= range tails {
               addresses = append(addresses, v | change_bit)
               addresses = append(addresses, v &^ change_bit)
          }
     }
     return addresses
}

func main() {
     f, _ := os.Open("i_14.txt")
     defer f.Close()
     scanner := bufio.NewScanner(f)
     mask := ""
     mem1 := map[int64]int64{}
     mem2 := map[int64]int64{}
     for scanner.Scan() {
          line := scanner.Text()
          if line[:4] == "mask"{
               mask = s.Split(line, "= ")[1]
          } else {
               reM := regexp.MustCompile(`\d+`)
               data := reM.FindAll([]byte(line), -1)
               address, _ := strconv.ParseInt(string(data[0]), 10, 64)
               value, _ := strconv.ParseInt(string(data[1]), 10, 64)

               // Part 2
               addresses := applyMask(address, mask)
               for _, v := range addresses{
                    mem2[v] = value
               }

               // Part 1

               shift := 0
               for i := len(mask)-1; i>=0; i-- {
                    if mask[i] == '0' {
                         value &^= 1 << shift
                    } else if mask[i] == '1' {
                         value |= 1 << shift
                    }
                    shift++
               }
               mem1[address] = value
          }
     }
     fmt.Println(sumMem(mem1))
     fmt.Println(sumMem(mem2))
}
