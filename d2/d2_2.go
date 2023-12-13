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
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	scanner := bufio.NewScanner(file)

	var gamesNotToAdd []int
	var gamesToAdd []int
	addGame := true

	for scanner.Scan() {

		//Split by line
		line := scanner.Text()

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

					switch {
					case num > 12 && color == "red":
						fmt.Println("Do not add game #", gameNum, splitByWord)
						gamesNotToAdd = append(gamesNotToAdd, gameNum)
					case num > 13 && color == "green":
						fmt.Println("Do not add game #", gameNum, splitByWord)
						gamesNotToAdd = append(gamesNotToAdd, gameNum)
					case num > 14 && color == "blue":
						fmt.Println("Do not add game #", gameNum, splitByWord)
						gamesNotToAdd = append(gamesNotToAdd, gameNum)
					} //end of switch statement

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
			gamesToAdd = append(gamesToAdd, gameNum)
		}

	} //end of loop by lines

	total := 0
	for _, v := range gamesToAdd {
		total += v
	}
	fmt.Printf("%d is all of the games added together", total)

} //end of main
