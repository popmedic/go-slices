package slices

type ISlice interface {
	Add(items ...IItem) ISlice
	Remove(items ...IItem)
	RemoveAt(i int)
	RemoveAll()
	Get(i int) IItem
	All() []IItem
	Len() int
	Contains(item IItem) bool
	Compare(ss2 ISlice) bool
	String() string
}
