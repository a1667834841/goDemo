package integers

import (
	"reflect"
	"testing"
)

/*
统计测试覆盖率的命令【go test -cover】
*/
func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}

	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

}

func TestAllSum(t *testing.T) {
	allSum := AllSum([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(allSum, want) {
		t.Errorf("allSum %v, want %v", allSum, want)
	}
}

/*
简单求和
*/
func Sum(arr []int) int {

	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

/*
利用range循环，range会将键值对返回，如果不想用索引可以用 “_”忽略
*/
func SumWithRange(arr []int) int {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return sum
}

/**
make 可以帮助创建给定长度的切片
*/
func AllSum(arrToSum ...[]int) (sums []int) {
	lengthOfArr := len(arrToSum)
	sums = make([]int, lengthOfArr)

	for i, arr := range arrToSum {
		sums[i] = Sum(arr)
	}
	return sums
}
