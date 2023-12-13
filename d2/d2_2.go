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

} //end of main

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

		//Split by line
		line := scanner.Text()

		largestRed := 0
		largestGreen := 0
		largestBlue := 0

		//Split by Game and Picks
		game := strings.Split(line, ":")[0]
		picks := strings.Split(line, ":")[1]

		//Split by just number
		gameNum, err := strconv.Atoi(strings.Fields(game)[1])
		if err != nil {
			log.Fatalf("error: %s", err)
		}
		//Split the second half of the line by all of the words
		rounds := strings.Split(picks, ";")

		for _, v := range rounds {
			splitByWord := strings.Fields(v)

			for i := 0; i < len(splitByWord); i++ {
				if i%2 == 0 { //Should always be num

					num, err := strconv.Atoi(splitByWord[i])
					if err != nil {
						log.Fatalf("error: %s", err)
					}
					color := strings.Trim(splitByWord[i+1], ",.!?")
					/* PART ONE */
					switch {
					case num > 12 && color == "red":
						// fmt.Println("Do not add game #", gameNum, splitByWord)
						gamesNotToAdd = append(gamesNotToAdd, gameNum)
					case num > 13 && color == "green":
						// fmt.Println("Do not add game #", gameNum, splitByWord)
						gamesNotToAdd = append(gamesNotToAdd, gameNum)
					case num > 14 && color == "blue":
						// fmt.Println("Do not add game #", gameNum, splitByWord)
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

				} //end of if statement checking for modulo 2

			} //end of loop by each individual word

		} //end of loop by rounds

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

	} //end of loop by lines

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
