package sa

type SASort interface {
	Sort([]byte) []int
	SortString(string) []int
}
