package heavy

type Heavy struct {
	Data [10000]int
}

func ByValue(h Heavy) int {
	return h.Data[0]
}

func ByReferenc(h *Heavy) int {
	return h.Data[0]
}
