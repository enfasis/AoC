package main

import (
     "fmt"
     "strconv"
)

type Cup struct{
     val int
     next *Cup
}

type CupCrabs struct {
     current *Cup
     mem map[int]*Cup
     max int
}

func (g *CupCrabs) move(){
     dest := g.current.val-1
     a := g.current.next
     b := a.next
     c := b.next
     for a.val == dest || b.val == dest || c.val == dest || dest == 0 {
          if dest == 0 { dest = g.max
          } else { dest-- }
     }
     g.current.next = c.next
     g.current = c.next
     tmpDest := g.mem[dest]
     c.next = tmpDest.next
     tmpDest.next = a
}

func (g *CupCrabs) result1() string {
     r := ""
     cup := g.mem[1].next
     for i:=0;i<8; i++{
          r+= strconv.Itoa(cup.val)
          cup = cup.next
     }
     return r
}

func (g *CupCrabs) result2() int {
     cup := g.mem[1].next
     return cup.val * cup.next.val
}

func arrangeGame1(data []int, move int) CupCrabs{
     game := CupCrabs{current: &Cup{val: data[0]}, max: 9, mem: map[int]*Cup{}}
     c := game.current
     game.mem[data[0]] = c
     for i:=1; i < len(data); i++ {
          c.next = &Cup{val: data[i]}
          game.mem[data[i]] = c.next
          c = c.next
     }
     c.next = game.current
     for i:=0;i<move;i++{
          game.move()
     }
     return game
}

func arrangeGame2(data []int, move int) CupCrabs{
     max := 1000000
     game := CupCrabs{current: &Cup{val: data[0]}, max: max, mem: map[int]*Cup{}}
     c := game.current
     game.mem[data[0]] = c
     i := 1
     for i < len(data) {
          c.next = &Cup{val: data[i]}
          game.mem[data[i]] = c.next
          c = c.next
          i++
     }
     for i < max {
          i++
          c.next = &Cup{val: i}
          game.mem[i] = c.next
          c = c.next
     }
     c.next = game.current
     for i:=0;i<move;i++{
          game.move()
     }
     return game
}

func main() {
     data := []int{2,1,5,6,9,4,7,8,3}
     g1 := arrangeGame1(data, 100)
     g2 := arrangeGame2(data, 10000000)
     fmt.Println(g1.result1(), g2.result2())
}
