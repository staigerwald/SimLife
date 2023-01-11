package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width  = 80
	height = 15
)

type Universe [][]bool

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (u Universe) Fillon25Percent() {
	Probability := 0
	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			Probability = rand.Intn(4)
			if Probability == 3 {
				u[h][w] = true
			}
		}
	}
}
func (u Universe) Show() {
	for h := range u {
		for w := range u[h] {
			if u[h][w] {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Print("\x0c")
}
func (u Universe) Alive(h, w int) bool {
	for h >= height {
		h %= height
	}
	for h < 0 {
		h += height
	}
	for w < 0 {
		w += width
	}
	for w >= width {
		w %= width
	}
	return u[h][w]
}
func (u Universe) Next(h, w int) bool {
	if u.Alive(h, w) {
		switch {
		case u.Neighbors(h, w) > 3 || u.Neighbors(h, w) < 2:
			return false
		case u.Neighbors(h, w) == 3 || u.Neighbors(h, w) == 2:
			return true
		}
	} else {
		if u.Neighbors(h, w) == 3 {
			return true
		}
	}
	return false
}
func Step(a, b Universe) {
	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			b[h][w] = a.Next(h, w)
		}
	}
}
func (u Universe) Neighbors(x, y int) int {
	n := 0
	for v := -1; v <= 1; v++ {
		for h := -1; h <= 1; h++ {
			if !(v == 0 && h == 0) && u.Alive(x+h, y+v) {
				n++
			}
		}
	}
	return n
}
func NewUniverse() Universe {
	NewUniverse := make(Universe, height)
	for i := range NewUniverse {
		NewUniverse[i] = make([]bool, width)
	}
	return NewUniverse
}

func main() {
	NewUniverse, TempUniverse := NewUniverse(), NewUniverse()
	NewUniverse.Fillon25Percent()
	NewUniverse.Show()

	for {
		Step(NewUniverse, TempUniverse)
		NewUniverse.Show()
		time.Sleep(time.Second / 8)
		NewUniverse, TempUniverse = TempUniverse, NewUniverse
	}
}
