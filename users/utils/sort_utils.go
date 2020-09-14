package utils

func BubbleSort(items []int) {
	swapped := false
	for i := 0; i < len(items)-1; i++ {
		if items[i] > items[i+1] {
			items[i+1], items[i] = items[i], items[i+1]
			swapped = true
		}
	}
	if swapped {
		BubbleSort(items)
	}
}
