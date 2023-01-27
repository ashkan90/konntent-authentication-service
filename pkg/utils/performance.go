package utils

const (
	StartOffsetValue          = 0
	MinimumOptimalOffsetValue = 9999
	MedianOptimalOffsetValue  = MinimumOptimalOffsetValue * 10
	MaximumOptimalOffsetValue = MedianOptimalOffsetValue * 10
)

var (
	optimalOffsetMap = map[int64]int{
		StartOffsetValue:          500,
		MinimumOptimalOffsetValue: 2500,
		MedianOptimalOffsetValue:  5000,
		MaximumOptimalOffsetValue: 25000,
	}
)

func Boundary(ln, size int) (bounds map[int]int) {
	bounds = make(map[int]int)
	j := 0
	for i := 0; i < ln; i += size {
		j += size
		if j > ln {
			j = ln
		}
		bounds[i] = j
	}

	return
}

func Offset(ln int) int {
	return OrderedComparison(optimalOffsetMap, int64(ln)).(int)
}
