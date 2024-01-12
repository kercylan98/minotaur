package collection

// SwapSlice 将切片中的两个元素进行交换
func SwapSlice[S ~[]V, V any](slice *S, i, j int) {
	if i < 0 || j < 0 || i >= len(*slice) || j >= len(*slice) {
		return
	}
	(*slice)[i], (*slice)[j] = (*slice)[j], (*slice)[i]
}
