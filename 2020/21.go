package main

import (
     "fmt"
     "bufio"
     "os"
     s "strings"
     "sort"
)

type Product struct {
     ingredients []string
     allergens []string
}

func (p *Product) removeIngredient(val string){
     removeFromArr(&p.ingredients, val)
}

func (p *Product) removeAllergen(val string){
     removeFromArr(&p.allergens, val)
}

func isInArray(arr []string, val string) bool {
     for _, v := range arr {
          if v == val {
               return true
          }
     }
     return false
}

func removeFromArr(a *[]string, val string){
     for k, v := range *a {
          if v == val {
               (*a)[k] = (*a)[len(*a)-1]
               (*a)[len(*a)-1] = ""
               (*a) = (*a)[:len(*a)-1]
               break
          }
     }
}


func main() {
     f, _ := os.Open("i_21.txt")
     scanner := bufio.NewScanner(f)
     products := []*Product{}
     allergenSet := map[string]bool{}
     traductor := map[string]string{}
     for scanner.Scan() {
          d := s.Split(scanner.Text(), " (contains ")
          p := Product{ingredients: s.Split(d[0], " "), allergens: s.Split(d[1][:len(d[1])-1], ", ")}
          for _, v := range p.allergens {
               allergenSet[v] = true
          }
          products = append(products, &p)
     }
     allergens := []string{}
     for k, _ := range allergenSet { allergens = append(allergens, k) }
     i := 0
     for len(allergens) != 0 {
          allergen := allergens[i]
          lists := [][]string{}
          for _, p := range products {
               if isInArray(p.allergens, allergen){
                    lists = append(lists, p.ingredients)
               }
          }
          ingredients := []string{}
          for _, ingredient := range lists[0] {
               isAllergen := true
               for _, list := range lists[1:] {
                    if !isInArray(list, ingredient) {
                         isAllergen = false
                         break
                    }
               }
               if isAllergen {
                    ingredients = append(ingredients, ingredient)
               }
          }
          if len(ingredients) == 1 {
               i = 0
               for _, product := range products {
                    product.removeIngredient(ingredients[0])
                    product.removeAllergen(allergen)
               }
               traductor[allergen] = ingredients[0]
               removeFromArr(&allergens, allergen)
          } else {
               i++
          }
     }
     count := 0
     for _, p := range products{
          count += len(p.ingredients)
     }
     fmt.Println(count)
     for k, _ := range allergenSet { allergens = append(allergens, k) }
     sort.Strings(allergens)
     r2 := ""
     for _, v := range allergens {
          r2 += "," + traductor[v]
     }
     fmt.Println(r2[1:])
}
