package tools

func SortArray(arr *[]int) {
	for i := 0; i < len(*arr); i++ {
		for j := i + 1; j < len(*arr)-1; j++ {
			if (*arr)[i] > (*arr)[j] {
				tmp := (*arr)[i]
				(*arr)[i] = (*arr)[j]
				(*arr)[j] = tmp
			}
		}
	}
}
