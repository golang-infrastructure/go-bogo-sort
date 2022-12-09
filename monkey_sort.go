package monkey_sort

import (
	"github.com/golang-infrastructure/go-gtypes"
	"github.com/golang-infrastructure/go-shuffle"
)

// Sort 对切片使用猴子排序
func Sort[T gtypes.Ordered](slice []T) int {
	count := 0
	for !isSorted(slice) {
		shuffle.FisherYatesKnuthShuffle(slice)
		count++
	}
	return count
}

// 判断切片是否已经有序
func isSorted[T gtypes.Ordered](slice []T) bool {
	if len(slice) <= 1 {
		return true
	}
	for index := 1; index < len(slice); index++ {
		if slice[index] < slice[index-1] {
			return false
		}
	}
	return true
}
