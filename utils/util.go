package utils

import "math/rand"

/**
 * @author skyfour
 * @date 2019-06-04
 * @email skyzhang@easemob.com
 */
func GetRandom(n int) int {
	return rand.Intn(n) + 1
}
