package conway

import (
	"bytes"
	"image"
	"image/color"
	"math/rand"
)

type Board [][]uint8

var (
	BoardPalette = color.Palette{color.Gray{0}, color.Gray{255}}
)

func BlankBoard(n int) Board {
	b := make(Board, n)
	grid := make([]uint8, n*n)
	for i := range b {
		b[i], grid = grid[:n], grid[n:]
	}
	return b
}

func RandomBoard(n, numEntries int) Board {
	b := BlankBoard(n)
	rand.Seed(42)
	for i := 0; i < numEntries; i++ {
		b[rand.Intn(n)][rand.Intn(n)] = 1
	}
	return b
}

func (b Board) ToImage() *image.Paletted {
	n := len(b)
	im := image.NewPaletted(image.Rect(0, 0, n, n), BoardPalette)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if b[i][j] == 0 {
				im.SetColorIndex(i, j, 0)
			} else {
				im.SetColorIndex(i, j, 1)
			}
		}
	}
	return im
}

type ConwayGame struct {
	Board
	Count int
}

func NewGame(n, numEntries int) ConwayGame {
	g := ConwayGame{
		Board: RandomBoard(n, numEntries),
	}

	dim := len(g.Board)
	for i := 0; i < dim; i++ {
		for j := 0; j < dim; j++ {
			if g.Board[i][j] > 0 {
				g.Count++
			}
		}
	}

	return g
}

func (g *ConwayGame) TakeTurn() {

	n := len(g.Board)
	newBoard := BlankBoard(n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			newBoard[i][j] = g.Board[i][j]
			im1 := index(i-1, n)
			ip1 := index(i+1, n)
			jm1 := index(j-1, n)
			jp1 := index(j+1, n)
			sum := g.Board[im1][jm1] + g.Board[i][jm1] + g.Board[ip1][jm1] +
				g.Board[im1][j] + g.Board[ip1][j] +
				g.Board[im1][jp1] + g.Board[i][jp1] + g.Board[ip1][jp1]
			if g.Board[i][j] == 1 {
				if sum < 2 || sum > 3 {
					newBoard[i][j] = 0
				}
			} else if sum == 3 {
				newBoard[i][j] = 1
			}
		}
	}
	g.Count = 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			g.Board[i][j] = newBoard[i][j]
			if g.Board[i][j] > 0 {
				g.Count++
			}
		}
	}
}

func (g ConwayGame) String() string {
	var b bytes.Buffer
	for _, row := range g.Board {
		for _, val := range row {
			if val > 0 {
				b.WriteString("x ")
			} else {
				b.WriteString("o ")
			}
		}
		b.WriteString("\n")
	}
	return b.String()
}

func index(i, n int) int {
	if i < 0 {
		return i + n
	}
	return i % n
}
