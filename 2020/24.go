package main

import (
     "fmt"
     "os"
     "bufio"
)

type Position struct {
     x int
     y int
}

func updatePosition(tiles *map[Position]bool, data string) {
     pos := Position{}
     for len(data) != 0 {
          cut := 0
          switch data[0] {
          case 'e':
               pos.x += 2
               cut = 1
          case 'w':
               pos.x -= 2
               cut = 1
          }
          if len(data) != 1{
               switch data[0:2] {
               case "se":
                    pos.x += 1
                    pos.y -= 1
                    cut = 2
               case "sw":
                    pos.x -= 1
                    pos.y -= 1
                    cut = 2
               case "ne":
                    pos.x += 1
                    pos.y += 1
                    cut = 2
               case "nw":
                    pos.x -= 1
                    pos.y += 1
                    cut = 2
               }
          }
          data = data[cut:]
     }
     if val, ok := (*tiles)[pos]; ok {
          (*tiles)[pos] = !val
     } else {
          (*tiles)[pos] = true
     }
}

func getNeighbours(pos Position) []Position {
     return []Position{
          Position{x: pos.x+2, y: pos.y},
          Position{x: pos.x-2, y: pos.y},
          Position{x: pos.x+1, y: pos.y+1},
          Position{x: pos.x+1, y: pos.y-1},
          Position{x: pos.x-1, y: pos.y-1},
          Position{x: pos.x-1, y: pos.y+1}}
}

func checkRule(tiles *map[Position]bool, nTiles *map[Position]bool, neighTiles *map[Position]bool, pos Position){
     neighbours := getNeighbours(pos)
     blacks := 0
     for _, neigh := range neighbours {
          if val, ok := (*tiles)[neigh]; ok {
               if val { blacks++ }
          } else {
               (*neighTiles)[neigh] = false
          }
     }
     isBlack := false
     if val, ok := (*tiles)[pos]; ok {
          isBlack = val
     }
     if isBlack && (blacks == 0 || blacks >2) {
          (*nTiles)[pos] = false
     } else if !isBlack && blacks == 2 {
          (*nTiles)[pos] = true
     } else {
          (*nTiles)[pos] = isBlack
     }
}

func getBlacks(tiles *map[Position]bool) int {
     sum := 0
     for _, v := range *tiles {
          if v == true{ sum++ }
     }
     return sum
}

func main() {
     f, _ := os.Open("i_24.txt")
     scanner := bufio.NewScanner(f)
     tiles := map[Position]bool{}
     for scanner.Scan(){
          updatePosition(&tiles, scanner.Text())
     }
     fmt.Println(getBlacks(&tiles))
     for i := 0; i < 100; i++ {
          nTiles := map[Position]bool{}
          neighTiles := map[Position]bool{}
          for pos, _ := range tiles {
               checkRule(&tiles, &nTiles, &neighTiles, pos)
          }
          for pos, _ := range neighTiles {
               checkRule(&tiles, &nTiles, &map[Position]bool{}, pos)
          }
          tiles = nTiles
     }
     fmt.Println(getBlacks(&tiles))
}
