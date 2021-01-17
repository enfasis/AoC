package main

import (
     "fmt"
     "io/ioutil"
     s "strings"
     "strconv"
)

func Reverse(s string) string {
     r := []rune(s)
     for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
          r[i], r[j] = r[j], r[i]

     }
     return string(r)
}

type Tile struct{
     data []string
     id int
     x int
     y int
}

func  (t * Tile) getImage() []string{
     r := []string{}
     for _, v := range t.data[1:len(t.data)-1]{
          r = append(r, v[1:len(v)-1])
     }
     return r
}

func (t *Tile) getBorder() []string{
     right := ""
     left := ""
     for i:=0; i < len(t.data); i++{
          right += string(t.data[i][len(t.data)-1])
          left += string(t.data[i][0])
     }
     top := t.data[0]
     bottom := Reverse(t.data[len(t.data)-1])
     left = Reverse(left)
     result := []string{top, right, bottom, left, Reverse(top), Reverse(left), Reverse(bottom), Reverse(right)}
     return result
}

func (t *Tile) rotate(n int){
     if n == 0 {
          return
     }
     r := []string{}
     for i:=0; i < len(t.data[0]); i++{
          ns := ""
          for j:=len(t.data)-1; j >= 0; j--{
               ns += string(t.data[j][i])
          }
          r = append(r, ns)
     }
     t.data = r
     t.rotate(n-1)
}


func (t *Tile) flip() {
     r := []string{}
     for j:=0; j <len(t.data); j++{
          r = append(r, Reverse(t.data[j]))
     }
     t.data = r
}

func solve(tile *Tile, tiles []*Tile) []*Tile{
     con := []*Tile{}
     noCon := []*Tile{}
     for _, t := range tiles {
          if connect(tile, t) {
               con = append(con, t)
          } else {
               noCon = append(noCon, t)
          }
     }
     for _, t := range con {
          noCon = solve(t, noCon)
     }
     return noCon
}

func connect(t1, t2 *Tile) bool {
     found := false
     for k1, v1 := range t1.getBorder()[:4] {
          for k2, v2 := range t2.getBorder() {
               if Reverse(v1) == v2 {
                    switch k1 {
                    case 0:
                         t2.x = t1.x
                         t2.y = t1.y+1
                    case 1:
                         t2.y = t1.y
                         t2.x = t1.x+1
                    case 2:
                         t2.x = t1.x
                         t2.y = t1.y-1
                    case 3:
                         t2.y = t1.y
                         t2.x = t1.x-1
                    }
                    if k2>3 {
                         t2.flip()
                    }
                    r := (8+(k1+2)%4 - (k2%4))%4
                    t2.rotate(r)
                    found = true
                    break
               }
          }
          if found {
               break
          }
     }
     return found
}

func getGrid(r,c int) [][]*Tile {
     grid := [][]*Tile{}
     for i:=0; i < r;i++{
          n := []*Tile{}
          for j:=0 ; j< c; j++ {
               n = append(n, nil)
          }
          grid = append(grid, n)
     }
     return grid
}

func joinImages(i1 []string, i2 []string) []string{
     r := []string{}
     for k, _ := range i1{
          r = append(r, i1[k]+i2[k])
     }
     return r
}

func countElement(data []string, e string) int{
     count := 0
     for _, v := range data {
          for _, w := range v {
               if e == string(w){
                    count++
               }
          }
     }
     return count
}

func (t *Tile) findMonsters() int{
     a := []string{"                  # ","#    ##    ##    ###", " #  #  #  #  #  #   "}
     b := t.data
     count := 0
     oS := countElement(a, "#")
     for i:= 0; i < len(b)-len(a); i++ {
          for j:=0; j < len(b[0]) - len(a[0]); j++{
               match := true
               for k:=0; k < len(a); k++ {
                    for l:=0; l <len(a[0]); l++{
                         if a[k][l] == '#' &&  b[i+k][j+l] != '#'{
                              match = false
                              break
                         }
                    }
                    if !match {
                         break
                    }
               }
               if match {
                    count++
               }
          }
     }
     return oS * count
}

func main() {
     f, _ := ioutil.ReadFile("i_20.txt")
     data := s.Split(string(f), "\n\n")
     tiles := []*Tile{}
     for _, v := range data {
          tD :=  s.Split(string(v), ":\n")
          if len(tD) == 1 {
               continue
          }
          id, _ := strconv.Atoi(string(tD[0][len(tD[0])-4:]))
          t := Tile{data: s.Split(tD[1], "\n")[:10], id: id}
          tiles = append(tiles, &t)
     }
     solve(tiles[0], tiles[1:])
     minX := 0
     minY := 0
     maxX := 0
     maxY := 0
     for i, _ := range tiles {
          if maxX < tiles[i].x {
               maxX = tiles[i].x
          }
          if maxY < tiles[i].y {
               maxY = tiles[i].y
          }
          if minX > tiles[i].x {
               minX = tiles[i].x
          }
          if minY > tiles[i].y {
               minY = tiles[i].y
          }
     }
     lenY := 1+maxY-minY
     lenX := 1+maxX-minX
     grid := getGrid(lenY, lenX)
     for _, t := range tiles {
          grid[t.y - minY][t.x - minX] = t
     }
     fmt.Println(grid[0][0].id*grid[0][lenX-1].id*grid[lenY-1][lenX-1].id*grid[lenY-1][0].id)

     image := []string{}
     for j:=lenY-1 ; j >= 0; j--{
          r := grid[j][0].getImage()
          for i := 1; i < lenX; i++ {
               r = joinImages(r, grid[j][i].getImage())
          }
          image = append(image, r...)
     }
     iT := Tile{data:image}
     monster := 0
     for i := 0; i < 4; i++ {
          iT.rotate(i)
          monster = iT.findMonsters()
          if monster != 0 {
               break
          }
     }
     if monster == 0 {
          iT.flip()
          for i := 0; i < 4; i++ {
               iT.rotate(i)
               monster = iT.findMonsters()
               if monster != 0 {
                    break
               }
          }
     }
     fmt.Println(countElement(iT.data, "#")-monster)
}
