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
				for j := 0; j < lengthdice[i]; j++ {
					roll := rand.Intn(6) + 1
					if roll == 6 {
						playerpoint[i]++
					} else if roll == 1 {
						if i == players {
							if lengthdice[i+1] > 0 {
								evaluatedice[1] = append(evaluatedice[1], 1)
							}
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
			fmt.Printf("\t Pemain #%d (%d):%s \n", player, playerpoint[player], Tostring(dice))
			lengthdice[player] = len(dice)
			if len(dice) == 0 {
				finish++
			}
			evaluatedice[player] = []int{}
			diceplayers[player] = dice
		}
		if finish == players-1 {
			break
		}
	}
	var winner, loser int
	for player, val := range playerpoint {
		if winner < val && len(diceplayers[player]) == 0 {
			winner = player
		}
		if len(diceplayers[player]) > 0 {
			loser = player
		}
	}
	fmt.Printf("Game berakhir karena hanya pemain #%d yang memiliki dadu.\n", loser)
	fmt.Printf("Game dimenangkan oleh pemain #%d karena memiliki poin lebih banyak dari pemain lainnya\n", winner)
}

func Channel() {

}

func main() {
	var players, dices int
	fmt.Printf("Masukan jumlah player: ")
	fmt.Scanf("%d\n", &players)
	fmt.Printf("Masukan jumlah dadu: ")
	fmt.Scanf("%d\n", &dices)
	dice(players, dices)

	// messages := make(chan string, 10)

	// for i := 0; i < 10; i++ {
	// 	req := uuid.New().String()
	// 	messages <- req
	// }
	// close(messages)

	// for val := range messages {
	// 	fmt.Println(val)
	// }

	// arr := []int{}
	// wg := &sync.WaitGroup{}
	// request := []int{1, 2, 3, 4, 5}
	// wg.Add(len(request))
	// numbers := make(chan int)
	// go func(numbers <-chan int) {
	// 	for val := range numbers {
	// 		arr = append(arr, val)
	// 	}
	// }(numbers)
	// for _, val := range request {
	// 	go func(wg *sync.WaitGroup, i int) {
	// 		defer wg.Done()
	// 		numbers <- i
	// 	}(wg, val)
	// }
	// wg.Wait()
	// close(numbers)
	// fmt.Println(arr)

}
