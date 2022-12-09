package monkey_sort

import "testing"

func TestSort(t *testing.T) {
	slice := []int{2, 3, 1, 4}
	count := Sort(slice)
	t.Log(count)
	t.Log(slice)
}
