package main

import (
     "fmt"
     "bufio"
     "os"
     "strconv"
     s "strings"
     "math"
)

type Gps struct {
     x int
     y int
     dir string
}

func NewGps() Gps{
     return Gps{x:0, y:0, dir: "E"}
}

func (g Gps) md() int{
     return int(math.Abs(float64(g.x))+ math.Abs(float64(g.y)))
}

func (g *Gps) sum2Dir(d string, value int ){
     switch d{
          case "N":
          g.y += value
          case "S":
          g.y -= value
          case "E":
          g.x += value
          case "W":
          g.x -= value
          case "F":
          g.sum2Dir(g.dir, value)
          case "R":
          d := "NESW"
          r := ((value/90)%4 + s.Index(d, g.dir))%4
          g.dir = string(d[r])
          case "L":
          d := "NESW"
          r := (s.Index(d, g.dir)+ 4 - (value/90)%4)%4
          g.dir = string(d[r])
     }
}

func (g *Gps) rotate(value int){
     x := float64((*g).x)
     y := float64((*g).y)
     deg := float64(value)*math.Pi/180.0
     (*g).x = int(math.Round(x*math.Cos(deg) - y*math.Sin(deg)))
     (*g).y = int(math.Round(x*math.Sin(deg) + y*math.Cos(deg)))
}

type Ship struct{
     wp Gps
     x int
     y int
}

func (g Ship) md() int{
     return int(math.Abs(float64(g.x))+ math.Abs(float64(g.y)))
}

func NewShip() Ship{
     return Ship{wp: Gps{x:10, y:1, dir:"E"}, x:0, y:0}
}

func (b *Ship) sum2Dir(d string, value int){
     switch d{
          case "N", "S", "E", "W":
          (*b).wp.sum2Dir(d, value)
          case "F":
          (*b).x += value*(*b).wp.x
          (*b).y += value*(*b).wp.y
          case "R":
          (*b).wp.rotate(-1*value)
          case "L":
          (*b).wp.rotate(value)
     }
}

func main() {
     f, _ := os.Open("i_12.txt")
     scanner := bufio.NewScanner(f)
     gps := NewGps()
     ship := NewShip()
     for scanner.Scan() {
          line := scanner.Text()
          value, _ := strconv.Atoi(line[1:])
          gps.sum2Dir(string(line[0]), value)
          ship.sum2Dir(string(line[0]), value)
     }
     fmt.Println(gps.md())
     fmt.Println(ship.md(), ship.x, ship.y)
}
