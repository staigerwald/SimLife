package main

import (
	"fmt"
	"math/rand"
	"time"
	"os"
	"os/exec"
	"runtime"
)

const (
	width  = 80
	height = 15
)

type Universe [][]bool
var clear map[string]func() //create a map for storing clear funcs

func init() {
	rand.Seed(time.Now().UnixNano())
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
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
	CallClear()
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
func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok { //if we defined a clear func for that platform:
		value()  //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func main() {
	NewUniverse, TempUniverse := NewUniverse(), NewUniverse()
	NewUniverse.Fillon25Percent()
	NewUniverse.Show()

	for {
		Step(NewUniverse, TempUniverse)
		NewUniverse.Show()
		time.Sleep(time.Second / 200)
		NewUniverse, TempUniverse = TempUniverse, NewUniverse
	}
}
