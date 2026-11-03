package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// TODO: the Shape interface (Area() int and Name() string), the three
// shape types, their String methods, and the sortable slice type.

// num returns word i of a split line as an int (0 if the word is missing).
type Shape interface {
	Name() string
	Area() int
}
type rect struct{ W, H int }
type square struct{ S int }
type tri struct{ B, H int }

func (r rect) Area() int      { return r.W * r.H }
func (s square) Area() int    { return s.S * s.S }
func (t tri) Area() int       { return t.B * t.H / 2 }
func (r rect) Name() string   { return "rect" }
func (r square) Name() string { return "square" }
func (r tri) Name() string    { return "tri" }
func (r rect) String() string {
	return fmt.Sprintf("rect %dx%d area=%d", r.W, r.H, r.Area())
}
func (s square) String() string {
	return fmt.Sprintf(
		"square %d area=%d",
		s.S,
		s.Area(),
	)
}

func (t tri) String() string {
	return fmt.Sprintf(
		"tri %dx%d area=%d",
		t.B,
		t.H,
		t.Area(),
	)
}

type shapeList []Shape

func (s shapeList) Len() int {
	return len(s)
}
func (s shapeList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s shapeList) Less(i, j int) bool {
	if s[i].Area() != s[j].Area() {
		return s[i].Area() > s[j].Area()
	}

	return s[i].Name() < s[j].Name()
}
func num(f []string, i int) int {
	if i >= len(f) {
		return 0
	}
	n, _ := strconv.Atoi(f[i])
	return n
}

// keeps the sort import alive until you use it

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)
	var lines [][]string // each input line split into words: {"rect", "3", "4"}
	for sc.Scan() {
		f := strings.Fields(sc.Text())
		if len(f) > 0 {
			lines = append(lines, f)
		}
	}

	// TODO: turn each line into a Shape (or remember its unknown kind),
	var shapes shapeList
	var unknown []string
	rects := 0
	squares := 0
	tris := 0

	// sort, print the shapes, the count line, then the unknown kinds.
	for _, fields := range lines {
		switch fields[0] {
		case "rect":
			shapes = append(shapes, rect{
				W: num(fields, 1),
				H: num(fields, 2),
			})
			rects++

		case "square":
			shapes = append(shapes, square{
				S: num(fields, 1),
			})
			squares++

		case "tri":
			shapes = append(shapes, tri{
				B: num(fields, 1),
				H: num(fields, 2),
			})
			tris++

		default:
			unknown = append(unknown, fields[0])
		}
	}

	sort.Stable(shapes)

	for _, shape := range shapes {
		fmt.Println(shape)
	}

	fmt.Printf(
		"rects=%d squares=%d tris=%d\n",
		rects,
		squares,
		tris,
	)

	for _, kind := range unknown {
		fmt.Printf("unknown: %s\n", kind)
	}
}
