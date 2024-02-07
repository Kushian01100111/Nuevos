package kata

import "sort"

func frequencySort(s string) string {
	arr := map[string]int{}

	for _, str := range s {
		temp := string(str)
		if _, present := arr[temp]; !present {
			arr[temp] = 1
		} else {
			arr[temp]++
		}
	}

	sli := make([]string, 0, len(arr))

	for val := range arr {
		sli = append(sli, val)
	}

	sort.SliceStable(sli, func(i, j int) bool {
		return arr[sli[i]] > arr[sli[j]]
	})

	str := ""
	for _, byt := range sli {
		num := arr[byt]
		for i := 0; i < num; i++ {
			str += byt
		}
	}

	return str
}
