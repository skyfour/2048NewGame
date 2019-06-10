package work

import (
	"gotrain/game/utils"
	"strconv"
	"strings"
)

/**
 * @author skyfour
 * @date 2019-06-06
 * @email skyzhang@easemob.com
 */
const (
	DIRECTION_ALL   = 0
	DIRECTION_UP    = 1
	DIRECTION_DOWN  = 2
	DIRECTION_LEFT  = 3
	DIRECTION_RIGHT = 4
	DECOLLATOR      = "=="
)

/**
找到连接的数字
*/
func FindList(data [][]int) (bool, map[string]int) {

	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			tMap := make(map[string]int)
			if data[i][j] != 0 {
				findMaxList(data, i, j, data[i][j], DIRECTION_ALL, &tMap)
			}
			//fmt.Printf("i:%d,j:%d,tMap:%v\n", i, j, tMap)
			if len(tMap) >= 3 {
				return true, tMap
			}
		}

	}
	return false, nil
}

/*
将连接好的变成0，并且返回第一个的位置
*/
func ConnectChangeZero(data [][]int, temp map[string]int) (int, int, int) {
	ir, jr, value := -1, -1, 1000000000000
	for k, v := range temp {
		l := strings.Split(k, DECOLLATOR)
		if len(l) != 2 {
			return value, ir, jr
		}

		i, _ := strconv.ParseInt(l[0], 10, 64)
		j, _ := strconv.ParseInt(l[1], 10, 64)
		if v <= value {
			ir, jr, value = int(i), int(j), data[i][j]
		}
		data[i][j] = 0
	}
	return value, ir, jr
}

/**
数字下落，将0覆盖
*/
func Fall(data [][]int) {
	if len(data) == 0 {
		return
	}
	length := len(data)
	for j := 0; j < length; j++ {
		for i := 0; i < length; i++ {
			if data[i][j] == 0 && i > 0 && data[i-1][j] != 0 {
				temp := data[i][j]
				data[i][j] = data[i-1][j]
				data[i-1][j] = temp
				i = 0
			}
		}

	}
}

/**
将0通过随机数填充
*/
func Fill(data [][]int, n int) {
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			if data[i][j] == 0 {
				data[i][j] = utils.GetRandom(n)
			}
		}
	}
}

/**
direction 方向 0 全都 1 上 2 下 3 左 4 右
*/
//func findMaxList(data [][]int, i, j, temp int, list *modlue.List) {
func findMaxList(data [][]int, i, j, temp, direction int, m *map[string]int) {
	//fmt.Printf("%d,%d,%d\n", i, j, temp)
	key := strconv.FormatInt(int64(i), 10) + DECOLLATOR + strconv.FormatInt(int64(j), 10)
	//横向加一
	if i < len(data) && j+1 < len(data[i]) && direction != DIRECTION_RIGHT {
		if temp == data[i][j+1] && temp == data[i][j] {
			key1 := strconv.FormatInt(int64(i), 10) + DECOLLATOR + strconv.FormatInt(int64(j+1), 10)
			if (*m)[key] > 0 && (*m)[key1] > 0 {
				return
			}
			(*m)[key] = (*m)[key] + 1
			(*m)[key1] = (*m)[key] + 2
			findMaxList(data, i, j+1, temp, DIRECTION_LEFT, m)
		}

	}
	//横向减一
	if i < len(data) && j-1 >= 0 && direction != DIRECTION_LEFT {
		if temp == data[i][j-1] && temp == data[i][j] {
			key1 := strconv.FormatInt(int64(i), 10) + DECOLLATOR + strconv.FormatInt(int64(j-1), 10)
			if (*m)[key] > 0 && (*m)[key1] > 0 {
				return
			}
			(*m)[key] = (*m)[key] + 1
			(*m)[key1] = (*m)[key] + 2
			findMaxList(data, i, j-1, temp, DIRECTION_RIGHT, m)
		}

	}
	//纵向加一
	if i+1 < len(data) && direction != DIRECTION_UP {
		if temp == data[i+1][j] && temp == data[i][j] {
			key1 := strconv.FormatInt(int64(i+1), 10) + DECOLLATOR + strconv.FormatInt(int64(j), 10)
			if (*m)[key] > 0 && (*m)[key1] > 0 {
				return
			}
			(*m)[key] = (*m)[key] + 1
			(*m)[key1] = (*m)[key] + 2
			findMaxList(data, i+1, j, temp, DIRECTION_DOWN, m)
		}
	}
	//纵向减一
	if i-1 > 0 && direction != DIRECTION_DOWN {
		if temp == data[i-1][j] && temp == data[i][j] {
			key1 := strconv.FormatInt(int64(i-1), 10) + DECOLLATOR + strconv.FormatInt(int64(j), 10)
			if (*m)[key] > 0 && (*m)[key1] > 0 {
				return
			}
			(*m)[key] = (*m)[key] + 1
			(*m)[key1] = (*m)[key] + 2
			findMaxList(data, i-1, j, temp, DIRECTION_UP, m)
		}
	}
}
