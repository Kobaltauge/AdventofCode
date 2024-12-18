package main

import (
	"container/heap"
	"fmt"
	"strings"
)

var sample = board(
	B, C, B, D,
	A, D, C, A,
)

var win = board(
	A, B, C, D,
	A, B, C, D,
)

var input = board(
	A, D, A, B,
	B, C, D, C,
)

type Board [11 + 4 + 4]Pod

type Pod uint8

const (
	Empty Pod = iota
	A
	B
	C
	D
)

func (p Pod) String() string { return ".ABCD"[p : p+1] }

const EmptyBoard = `
#############
#...........#
###.#.#.#.###
  #.#.#.#.#
  #########
`

func (b Board) String() string {
	var args []any
	for _, r := range b {
		args = append(args, r)
	}
	return fmt.Sprintf(strings.ReplaceAll(EmptyBoard, ".", "%v"), args...)
}

func board(p ...Pod) Board {
	var b Board
	copy(b[11:], p)
	return b
}

var cost = [D + 1]int{A: 1, B: 10, C: 100, D: 1000}
var entry = [D + 1]int{A: 2, B: 4, C: 6, D: 8}
var up = [len(Board{})]int{
	11: 2, 12: 4, 13: 6, 14: 8,
	15: 11, 16: 12, 17: 13, 18: 14,
}

var down [len(Board{})]int

func init() {
	for i, u := range up {
		if u != 0 {
			down[u] = i
		}
	}
}

func move(b Board, c int, f func(Board, int)) {
Rooms:
	for i, p := range b {
		if p == Empty {
			continue
		}
		if i >= 11 {
			// into the hall
			d := 0
			e := i
			for up[e] != 0 {
				e = up[e]
				d++
				if b[e] != Empty {
					continue Rooms
				}
			}
			// Left
			for j := e - 1; j >= 0 && b[j] == Empty; j-- {
				if down[j] != 0 {
					continue
				}
				b := b
				b[i] = Empty
				b[j] = p
				f(b, c+(d+e-j)*cost[p])
			}
			// right
			for j := e + 1; j < 11 && b[j] == Empty; j++ {
				if down[j] != 0 {
					continue
				}
				b := b
				b[i] = Empty
				b[j] = p
				f(b, c+(d+j-e)*cost[p])
			}
		} else {
			// in hallway
			d := 0
			e := entry[p]
			bottom := e
			for j := e; j != 0; j = down[j] {
				if b[j] == p {
					for k := down[j]; k != 0; k = down[k] {
						if b[j] != p {
							continue Rooms
						}
					}
					break
				}
				if b[j] != Empty {
					continue Rooms
				}
				bottom = j
				d++
			}
			dx := 1
			if i > e {
				dx = -1
			}
			for j := i + dx; j != e; j += dx {
				if b[j] != Empty {
					continue Rooms
				}
				d++
			}
			b := b
			b[i] = Empty
			b[bottom] = p
			f(b, c+d*cost[p])
		}
	}
}

type Work struct {
	heap []heapEntry
	pos  map[Board]int
	prev map[Board]Board
}

type heapEntry struct {
	b    Board
	cost int
}

func (w *Work) Add(prev, b Board, c int) {
	if i, ok := w.pos[b]; ok {
		if i < 0 || w.heap[i].cost <= c {
			return
		}
		w.heap[i].cost = c
		heap.Fix((*byCost)(w), i)
	} else {
		heap.Push((*byCost)(w), heapEntry{b, c})
	}
	w.prev[b] = prev
}

func (w *Work) Next() (Board, int) {
	e := heap.Pop((*byCost)(w)).(heapEntry)
	return e.b, e.cost
}
func (w *Work) Empty() bool {
	return len(w.heap) == 0
}

func (w *Work) Path(b Board) []Board {
	prev := w.prev[b]
	if prev == b {
		return []Board{b}
	}
	return append(w.Path(prev), b)
}
func NewWork() *Work {
	return &Work{
		pos:  make(map[Board]int),
		prev: make(map[Board]Board),
	}
}

type byCost Work

func (w *byCost) Len() int           { return len(w.heap) }
func (w *byCost) Less(i, j int) bool { return w.heap[i].cost < w.heap[j].cost }

func (w *byCost) fix(i int) {
	w.pos[w.heap[i].b] = i
}

func (w *byCost) Swap(i, j int) {
	w.heap[i], w.heap[j] = w.heap[j], w.heap[i]
	w.fix(i)
	w.fix(j)
}

func (w *byCost) Push(x any) {
	w.heap = append(w.heap, x.(heapEntry))
	w.fix(len(w.heap) - 1)
}
func (w *byCost) Pop() any {
	x := w.heap[len(w.heap)-1]
	w.heap = w.heap[:len(w.heap)-1]
	w.pos[x.b] = -1
	return x
}

func main() {
	w := NewWork()
	w.Add(input, input, 0)
	for !w.Empty() {
		b, c := w.Next()
		if b == win {
			fmt.Println(c)
			for _, b := range w.Path(b) {
				fmt.Println(b)
			}
			return
		}
		prev := b
		move(b, c, func(b Board, c int) {
			w.Add(prev, b, c)
		})
	}
}
