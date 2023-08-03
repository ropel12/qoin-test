package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Tostring(data []int) (res string) {

	if len(data) == 0 {
		return "_ (Berhenti bermain karena tidak memiliki dadu)"
	}
	for index, val := range data {
		if index != len(data)-1 {
			res += fmt.Sprintf("%d,", val)
		} else {
			res += fmt.Sprintf("%d", val)
		}
	}
	return res
}

func dice(players int, dices int) {
	rand.Seed(time.Now().UnixNano())
	playerpoint := map[int]int{}
	diceplayers := map[int][]int{}
	evaluatedice := map[int][]int{}
	lengthdice := map[int]int{}
	move := 1
	for {
		fmt.Printf("Pemain: %d , Dadu: %d\n =================== \n", players, dices)
		fmt.Printf("Giliran %d, lempar dadu: \n", move)
		for i := 1; i <= players; i++ {
			prevpoint := playerpoint[i]
			if move == 1 {
				prevpoint = 0
				for j := 0; j < dices; j++ {
					roll := rand.Intn(6) + 1
					if roll == 6 {
						playerpoint[i]++
					} else if roll == 1 {
						if i == players {
							evaluatedice[1] = append(evaluatedice[1], 1)
						} else {
							evaluatedice[i+1] = append(evaluatedice[i+1], 1)
						}
					} else {
						evaluatedice[i] = append(evaluatedice[i], roll)
					}
					diceplayers[i] = append(diceplayers[i], roll)
				}
			} else if move > 1 {
				diceplayers[i] = []int{}
				prevpoint = playerpoint[i]
				fmt.Println("pemain", i, "Isi", evaluatedice, move)
				for j := 0; j < lengthdice[i]; j++ {
					fmt.Println()
					roll := rand.Intn(6) + 1
					if roll == 6 {
						playerpoint[i]++
					} else if roll == 1 {
						if i == players {
							evaluatedice[1] = append(evaluatedice[1], 1)
						} else {
							if lengthdice[i+1] > 0 {
								evaluatedice[i+1] = append(evaluatedice[i+1], 1)
							}
						}
					} else {
						evaluatedice[i] = append(evaluatedice[i], roll)
					}
					diceplayers[i] = append(diceplayers[i], roll)
				}

			}
			fmt.Printf("\n\t Pemain #%d (%d): %s\n", i, prevpoint, Tostring(diceplayers[i]))
		}
		move++
		finish := 0
		fmt.Println("\tSetelah Evaluasi:")
		for player, dice := range evaluatedice {
			fmt.Println(dice)
			fmt.Printf("\t Pemain #%d (%d):%s \n", player, playerpoint[player], Tostring(dice))
			lengthdice[player] = len(dice)
			if len(dice) == 0 {
				finish++
			}
			evaluatedice[player] = []int{}
		}
		if finish == players-1 {
			break
		}
	}
}

func main() {
	var players, dices int
	fmt.Printf("Masukan jumlah player: ")
	fmt.Scanf("%d\n", &players)
	fmt.Printf("Masukan jumlah dadu: ")
	fmt.Scanf("%d\n", &dices)
	dice(players, dices)
}
