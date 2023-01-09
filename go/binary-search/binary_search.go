package binarysearch

func SearchInts(list []int, key int) int {
	if len(list) == 1 {
		if list[0] == key {
			return 0
		} else {
			return -1
		}
	}
	if len(list) < 1 {
		return -1
	}

	middle := len(list) / 2

	switch val := list[middle]; {
	case val == key:
		return middle
	case val > key:
		pos := SearchInts(list[0:middle], key)
		if pos < 0 {
			return -1
		}
		return middle - len(list[0:middle]) + pos
	case val < key:
		pos := SearchInts(list[middle:], key)
		if pos < 0 {
			return -1
		}
		return middle + pos
	default:
		return -1
	}
}
