package main

import (
     "fmt"
     "bufio"
     "os"
)

func removeValue(s []string, v string) ([]string, bool){
     size := len(s)
     i := 0
     for i = 0; i < size; i++ {
          if string(s[i]) == v {
               s[i] = s[size-1]
               return s[:size-1], true
          }
     }
     return s, false
}

func countAnswers(data []string) (int, int) {
     answers := []string{}
     for i := 97; i < 97+26; i++ {
          answers = append(answers, string(i))
     }
     isRemoved := false
     count := 0
     repeated := make(map[string]bool)
     for _, value := range data[0]{
          repeated[string(value)] = true
     }
     for _, person := range data {
          tRepeated := make(map[string]bool)
          for _, answer := range person {
               answers, isRemoved = removeValue(answers, string(answer))
               if isRemoved {
                    count++
               }
               for key, _ := range repeated {
                    if key == string(answer) {
                         tRepeated[key] = true
                         break
                    }
               }
          }
          repeated = tRepeated
     }
     return count, len(repeated)
}

func main() {
     f, _ := os.Open("i_06.txt")
     scanner := bufio.NewScanner(f)
     data := []string{}
     result1 := 0
     result2 := 0
     for scanner.Scan() {
          line := scanner.Text()
          if line == "" {
               r1, r2 := countAnswers(data)
               result1 += r1
               result2 += r2
               data = []string{}
               continue
          }
          data = append(data, line)
     }
     r1, r2 := countAnswers(data)
     result1 += r1
     result2 += r2
     fmt.Println(result1, result2)
}
