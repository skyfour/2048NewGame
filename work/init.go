package work

import (
	"gotrain/game/utils"
)

/**
 * @author skyfour
 * @date 2019-06-04
 * @email skyzhang@easemob.com
 */
//初始化n*n方格数据
func InitData(n int) [][]int {
	data := [][]int{}

	for i := 0; i < n; i++ {
		temp := []int{}
		for j := 0; j < n; {
			a := utils.GetRandom(n)
			if j-2 >= 0 && a == temp[j-1] && a == temp[j-2] {
				continue
			}
			temp = append(temp, a)

			j++
		}
		data = append(data, temp)
	}
	return data
}
