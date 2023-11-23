package kata

import "strconv"

func fizzBuzz(n int) []string {
	nums := [2]int{3, 5}
	print := [2]string{"Fizz", "Buzz"}
	arrStr := []string{}

	for i := 1; i <= n; i++ {
		result := ""
		for ite, num := range nums {
			for item, str := range print {
				if i%num == 0 {
					if ite == item {
						result += str
					}
				}
			}
		}
		if result == "" {
			result += strconv.Itoa(i)
			arrStr = append(arrStr, result)
		} else {
			arrStr = append(arrStr, result)
		}
	}
	return arrStr
}
