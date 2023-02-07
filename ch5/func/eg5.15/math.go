package math

import "errors"

// 编写类似sum的可变参数函数max和min。
// 考虑不传参时，max和min该如何处理，再编写至少接收1个参数的版本。
func max(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, errors.New("Parameter cannot be empty")
	}
	var max int
	for _, v := range vals {
		if v > max {
			max = v
		}
	}
	return max, nil
}

func min(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, errors.New("Parameter cannot be empty")
	}
	var min int
	for _, v := range vals {
		if v < min {
			min = v
		}
	}
	return min, nil
}

// 选择排序
func selectsort(vals []int) []int {
	for i := 0; i < len(vals); i++ {
		min := i
		for j := i; j < len(vals); j++ {
			if vals[min] > vals[j] {
				min = j
			}
		}
		vals[i], vals[min] = vals[min], vals[i]
	}
	return vals
}

// 快速排序
func quicksort(vals []int) []int {
	qsort(vals, 0, len(vals)-1)
	return (vals)
}

func qsort(vals []int, low, high int) {
	if low >= high {
		return
	}
	i := patition(vals, low, high)

	qsort(vals, low, i-1)
	qsort(vals, i+1, high)

}

// 划分数组，返回pviot下标
func patition(vals []int, low, high int) int {
	pivot := vals[low]
	i := low + 1
	for j := i; j <= high; j++ {
		if vals[j] < pivot {
			vals[j], vals[i] = vals[i], vals[j]
			i++
		}
	}
	i--
	vals[i], vals[low] = vals[low], vals[i]
	return i
}
