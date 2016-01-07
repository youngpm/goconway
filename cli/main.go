package main

import (
	"fmt"
	"image"
	"image/gif"
	"log"
	"os"

	"github.com/youngpm/conway"
)

func main() {

	n := 150
	steps := 10000

	im := gif.GIF{LoopCount: 5, Config: image.Config{ColorModel: conway.BoardPalette, Width: n, Height: n}}
	game := conway.NewGame(n, n*n/2)
	var counts []int

	im.Image = append(im.Image, game.ToImage())
	im.Delay = append(im.Delay, 20)
	fmt.Printf("step = %d Count = %d\n", 0, game.Count)
	for i := 0; i < steps; i++ {
		game.TakeTurn()
		counts = append(counts, game.Count)
		fmt.Printf("step = %d Count = %d\n", i+1, game.Count)
		im.Image = append(im.Image, game.ToImage())
		im.Delay = append(im.Delay, 20)

		if i > 4 {
			countset := make(map[int]bool)
			for _, val := range counts[i-5:] {
				countset[val] = true
			}
			if len(countset) == 1 {
				break
			}
		}
	}

	f, err := os.Create("/home/pyoung/conway.gif")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = gif.EncodeAll(f, &im)
	if err != nil {
		log.Fatal(err)
	}

	// for i := 0; i < 20; i++ {
	// 	fmt.Println(game)
	// 	game.TakeTurn()
	// }
	// fmt.Println(game)

}
