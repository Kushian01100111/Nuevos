package kata

import "math"

func numsOfSquare(n int) int {
	arr := make([]int, n+1)
	for i := 1; i <= n; i++ {
		arr[i] = math.MaxInt32
	}

	for i := 1; i <= n; i++ {
		for j := 1; j*j <= i; j++ {
			arr[i] = min(arr[i], arr[i-j*j]+1)
		}
	}

	return arr[n]
}
