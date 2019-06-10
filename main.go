package main

import (
	"bufio"
	"fmt"
	"gotrain/game/work"
	"os"
	"strconv"
	"strings"
)

/**
 * @author skyfour
 * @date 2019-06-04
 * @email skyzhang@easemob.com
 */
func main() {

	reader := bufio.NewReader(os.Stdin)
	gameData := work.InitData(5)
	print(gameData)

	isInput := true
	//分数
	score := 0
	//可使用次数
	times := 5
	for {
		print(gameData)
		if isInput {
			isHave, d := work.FindList(gameData)
			if isHave {
				print(gameData)
				v, i, j := work.ConnectChangeZero(gameData, d)
				score += v * len(d)
				times += 1
				gameData[i][j] = v + 1
				work.Fall(gameData)
				work.Fill(gameData, 6)
				continue
			}
		}
		fmt.Printf("Your score is %d.\n", score)
		fmt.Printf("Your times is %d.\n", times)
		if times == 0 {
			fmt.Printf("Your are game over \n")
			break
		}

		isInput = false
		data, _, _ := reader.ReadLine()
		input := strings.Split(string(data), " ")
		if len(input) != 2 {
			fmt.Printf("Your input is invaild %s.\n", data)
			continue
		}
		i, err := strconv.ParseInt(input[0], 10, 64)
		if err != nil {
			fmt.Printf("Your input is invaild %s.\n", data)
			continue
		}
		j, err := strconv.ParseInt(input[1], 10, 64)
		if err != nil {
			fmt.Printf("Your input is invaild %s.\n", data)
			continue
		}
		gameData[i][j] = gameData[i][j] + 1
		times -= 1
		isInput = true

	}

}

func print(data [][]int) {
	fmt.Printf("===========\n")
	for _, v := range data {
		fmt.Printf("%v\n", v)
	}
}
