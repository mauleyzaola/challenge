package business

type StringConstants []string

func (this StringConstants) Len() int {
	return len(this)
}

func (this StringConstants) Less(i, j int) bool {
	return len(this[i]) > len(this[j])
}

func (this StringConstants) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
