package main

import (
	"fmt"
	"os"
	"strconv"
)

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func gameTwoPlayers(sigP1, sigP2 string, numP1, numP2 int, mapWins map[int]string) int {
	//If sig is equals
	if sigP1 == sigP2 {
		//Pass minus number
		if numP1 < numP2 {
			if mapWins[numP1] == "" {
				mapWins[numP1] = strconv.FormatInt(int64(numP2), 10)
			} else {
				mapWins[numP1] = mapWins[numP1] + " " + strconv.FormatInt(int64(numP2), 10)
			}
			return numP1
		}
		if mapWins[numP1] == "" {
			mapWins[numP2] = strconv.FormatInt(int64(numP1), 10)
		} else {
			mapWins[numP2] = mapWins[numP2] + " " + strconv.FormatInt(int64(numP1), 10)
		}
		return numP2
	} else {
		if sigP1 == "C" && sigP2 == "P" ||
			sigP1 == "P" && sigP2 == "R" ||
			sigP1 == "R" && sigP2 == "L" ||
			sigP1 == "L" && sigP2 == "S" ||
			sigP1 == "S" && sigP2 == "C" ||
			sigP1 == "C" && sigP2 == "L" ||
			sigP1 == "L" && sigP2 == "P" ||
			sigP1 == "P" && sigP2 == "S" ||
			sigP1 == "S" && sigP2 == "R" ||
			sigP1 == "R" && sigP2 == "C" {
			if mapWins[numP1] == "" {
				mapWins[numP1] = strconv.FormatInt(int64(numP2), 10)
			} else {
				mapWins[numP1] = mapWins[numP1] + " " + strconv.FormatInt(int64(numP2), 10)
			}
			return numP1
		}
		if sigP2 == "C" && sigP1 == "P" ||
			sigP2 == "P" && sigP1 == "R" ||
			sigP2 == "R" && sigP1 == "L" ||
			sigP2 == "L" && sigP1 == "S" ||
			sigP2 == "S" && sigP1 == "C" ||
			sigP2 == "C" && sigP1 == "L" ||
			sigP2 == "L" && sigP1 == "P" ||
			sigP2 == "P" && sigP1 == "S" ||
			sigP2 == "S" && sigP1 == "R" ||
			sigP2 == "R" && sigP1 == "C" {
			if mapWins[numP1] == "" {
				mapWins[numP2] = strconv.FormatInt(int64(numP1), 10)
			} else {
				mapWins[numP2] = mapWins[numP2] + " " + strconv.FormatInt(int64(numP1), 10)
			}
			return numP2
		}
	}
	return -1
}

func gameRound(numPlay []int, m map[int]string, mapWin map[int]string, N int) []int {
	var numWinPlay []int
	//Divide number of the players in the middle
	fmt.Fprintln(os.Stderr, "Left key:")
	for i := 0; i < N/2; i += 2 {
		log := "Player " + strconv.FormatInt(int64(numPlay[i]), 10) + " | Letter " + m[numPlay[i]]
		fmt.Fprintln(os.Stderr, log)
		log = "Player " + strconv.FormatInt(int64(numPlay[i+1]), 10) + " | Letter " + m[numPlay[i+1]]
		fmt.Fprintln(os.Stderr, log)
		//If sig is equals
		result := gameTwoPlayers(m[numPlay[i]], m[numPlay[i+1]], numPlay[i], numPlay[i+1], mapWin)
		//result := gameTwoPlayers(sigPlay[i],sigPlay[i+1], numPlay[i], numPlay[i+1])
		log = "----------------------- Winner Left key | Player: " + strconv.FormatInt(int64(result), 10) + " -----------------------"
		fmt.Fprintln(os.Stderr, log)

		numWinPlay = append(numWinPlay, result)
	}

	if N > 2 {
		fmt.Fprintln(os.Stderr, "Right key:")
		for i := N - (N / 2); i < N; i += 2 {
			log := "Player " + strconv.FormatInt(int64(numPlay[i]), 10) + " | Letter " + m[numPlay[i]]
			fmt.Fprintln(os.Stderr, log)
			log = "Player " + strconv.FormatInt(int64(numPlay[i+1]), 10) + " | Letter " + m[numPlay[i+1]]
			fmt.Fprintln(os.Stderr, log)
			result := gameTwoPlayers(m[numPlay[i]], m[numPlay[i+1]], numPlay[i], numPlay[i+1], mapWin)
			log = "----------------------- Winner Right key | Player: " + strconv.FormatInt(int64(result), 10) + " -----------------------"
			fmt.Fprintln(os.Stderr, log)
			numWinPlay = append(numWinPlay, result)
		}
	}
	return numWinPlay
}

func main() {
	var N int
	fmt.Scan(&N)

	fmt.Fprintln(os.Stderr, "********************")

	fmt.Fprintln(os.Stderr, "Rules...")

	fmt.Fprintln(os.Stderr, "Scissors cuts Paper")
	fmt.Fprintln(os.Stderr, "Paper covers Rock")
	fmt.Fprintln(os.Stderr, "Rock crushes Lizard")
	fmt.Fprintln(os.Stderr, "Lizard poisons Spock")
	fmt.Fprintln(os.Stderr, "Spock smashes Scissor")
	fmt.Fprintln(os.Stderr, "Scissors decapitates Lizard")
	fmt.Fprintln(os.Stderr, "Lizard eats Paper")
	fmt.Fprintln(os.Stderr, "Paper disproves Spock")
	fmt.Fprintln(os.Stderr, "Spock vaporizes Rock")
	fmt.Fprintln(os.Stderr, "Rock crushes Scissors")
	fmt.Fprintln(os.Stderr, "********************")
	fmt.Fprintln(os.Stderr, "Reading all players...")

	var numPlay []int
	var sigPlay []string
	m := make(map[int]string)
	mapWins := make(map[int]string)
	for i := 0; i < N; i++ {
		var NUMPLAYER int
		var SIGNPLAYER string
		fmt.Scan(&NUMPLAYER, &SIGNPLAYER)
		sigPlay = append(sigPlay, SIGNPLAYER)
		numPlay = append(numPlay, NUMPLAYER)

		m[NUMPLAYER] = SIGNPLAYER

		log := "Player " + strconv.FormatInt(int64(numPlay[i]), 10) + " | Letter " + sigPlay[i]
		fmt.Fprintln(os.Stderr, log)
	}

	var roundNum = 1

	//Loop infinite
	for {
		log := "Round " + strconv.FormatInt(int64(roundNum), 10)
		fmt.Fprintln(os.Stderr, "********************")
		fmt.Fprintln(os.Stderr, log)
		roundNum++

		numPlay = gameRound(numPlay, m, mapWins, len(numPlay))
		if len(numPlay) == 1 {
			log := "Final Winner" + strconv.FormatInt(int64(numPlay[0]), 10)
			fmt.Fprintln(os.Stderr, log)

			fmt.Println(strconv.FormatInt(int64(numPlay[0]), 10)) // Write answer to stdout
			fmt.Println(mapWins[numPlay[0]])                      // Write answer to stdout
			break
		}
	}

}
