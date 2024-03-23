package kata

type LinkedList struct {
	Val  int
	Next *LinkedList
}

func reOrderList(head *LinkedList) {
	list := []int{}
	temp := head

	for temp != nil {
		list = append(list, temp.Val)
		temp = temp.Next
	}

	i, j := 0, len(list)-1
	switchVal := true
	for head != nil {
		if switchVal {
			head.Val = list[i]
			i++
			head = head.Next
			switchVal = false
		} else {
			head.Val = list[j]
			j--
			head = head.Next
			switchVal = true
		}
	}
}

// SALIO DE VUELTA
