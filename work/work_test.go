package work

import "testing"

/**
 * @author skyfour
 * @date 2019-06-10
 * @email skyzhang@easemob.com
 */

func TestFindList(t *testing.T) {
	gameData := [][]int{
		{6, 4, 2, 1, 6,},
		{4 ,6 ,4 ,3 ,1,},
		{4 ,1 ,1 ,5, 4},
		{6 ,1 ,1 ,3 ,2},
		{1 ,3 ,4 ,4 ,3},

	}
	isHave, d := FindList(gameData)
	if isHave {
		print(gameData)
		v, i, j := ConnectChangeZero(gameData, d)

		gameData[i][j] = v + 1
		Fall(gameData)
		Fill(gameData, 6)

	}
}
