package day_two

import (
	"fmt"
	"strings"

	"github.com/bobstrecansky/AOC2023/internal"
)

type gameSum struct {
	gameId int
	color  struct {
		rMax int
		gMax int
		bMax int
	}
}

func partOne(file string) int {
	input := internal.ReadInput(file)
	gs := parseGameInput(input)
	var idSum int
	for _, v := range gs {
		if v.color.rMax <= 12 && v.color.gMax <= 13 && v.color.bMax <= 14 {
			idSum += v.gameId + 1 // games are 0 indexed
		}
	}
	fmt.Println("Part 1:", idSum)
	return idSum
}

func partTwo(file string) int {
	var res int
	input := internal.ReadInput(file)
	gs := parseGameInput(input)
	for _, v := range gs {
		res += v.color.rMax * v.color.gMax * v.color.bMax
	}
	fmt.Println("Part 2:", res)
	return res
}

func parseGameInput(input []string) []gameSum {
	var color string
	var count int
	g := make([]gameSum, len(input))
	for gameNo, v := range input {
		g[gameNo].gameId = gameNo
		// trim Game #: from the input strings
		game := strings.Index(v, ":")
		v = v[game+1:]

		for _, game := range strings.Split(v, ";") {
			for _, gameDetails := range strings.Split(game, ";") {
				gameDetails = strings.TrimSpace(gameDetails)
				for _, colorDetails := range strings.Split(gameDetails, ", ") {
					fmt.Sscanf(colorDetails, "%d %s", &count, &color)
					switch color {
					case "red":
						if count > g[gameNo].color.rMax {
							g[gameNo].color.rMax = count
						}
					case "green":
						if count > g[gameNo].color.gMax {
							g[gameNo].color.gMax = count
						}
					case "blue":
						if count > g[gameNo].color.bMax {
							g[gameNo].color.bMax = count
						}
					}
				}
			}
		}

	}
	return g
}
