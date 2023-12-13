package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	handleInput("input.txt")
}

func handleInput(fn string) {

	file, err := os.Open(fn)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	var gamesNotToAdd []int
	var partOne []int
	var partTwo []int
	addGame := true
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		//SPLIT BY LINE
		line := scanner.Text()

		//SPLIT THE LINE BY GAME AND ALL PICKS
		game := strings.Split(line, ":")[0]
		picks := strings.Split(line, ":")[1]

		//SPLIT BY JUST NUMBER
		gameNum, err := strconv.Atoi(strings.Fields(game)[1])
		if err != nil {
			log.Fatalf("error: %s", err)
		}
		//SPLIT THE SECOND HALF OF THE LINE BY ALL OF THE WORDS
		rounds := strings.Split(picks, ";")

		//FOR PART TWO :
		largestRed := 0
		largestGreen := 0
		largestBlue := 0

		for _, v := range rounds {
			splitByWord := strings.Fields(v) //THIS SPLITS EACH ROUND INTO A NUMBER AND A COLOR

			for i := 0; i < len(splitByWord); i++ {
				if i%2 == 0 { //EVENS SHOULD ALWAYS BE AN INT

					num, err := strconv.Atoi(splitByWord[i])
					if err != nil {
						log.Fatalf("error: %s", err)
					}
					color := strings.Trim(splitByWord[i+1], ",.!?") //CHECK FOR SPECIAL CHARS AND REMOVE BEFORE SWITCH

					/* PART ONE */
					switch {
					case num > 12 && color == "red":
						gamesNotToAdd = append(gamesNotToAdd, gameNum)
					case num > 13 && color == "green":
						gamesNotToAdd = append(gamesNotToAdd, gameNum)
					case num > 14 && color == "blue":
						gamesNotToAdd = append(gamesNotToAdd, gameNum)
					}
					/* PART TWO */
					switch {
					case num > largestRed && color == "red":
						largestRed = num
					case num > largestGreen && color == "green":
						largestGreen = num
					case num > largestBlue && color == "blue":
						largestBlue = num
					}

				} //END OF IF STATEMENT CHECKING FOR MODULO 2

			} //END OF LOOP BY EACH INDIVIDUAL WORD

		} //END OF LOOP BY ROUNDS

		for _, v := range gamesNotToAdd {
			if gameNum != v {
				addGame = true
			} else {
				addGame = false
			}
		}
		if addGame {
			partOne = append(partOne, gameNum)
		}

		largestCubed := largestRed * largestGreen * largestBlue
		partTwo = append(partTwo, largestCubed)

	} //END OF LOOP BY LINES

	total := 0
	for _, v := range partOne {
		total += v
	}
	fmt.Printf("Part One: %d is all of the games added together\n", total)

	total = 0
	for _, v := range partTwo {
		total += v
	}
	fmt.Printf("Part Two: %d is all of the mins cubed\n", total)

}
