package slices

type MaxSlice struct {
	*SafeSlice
	max int
}

func NewMaxSlice(max int, items ...IItem) *MaxSlice {
	return &MaxSlice{
		SafeSlice: NewSafeSlice(items...),
		max:       max,
	}
}

func (ms *MaxSlice) Add(items ...IItem) ISlice {
	if i := (ms.Len() + len(items)) - ms.max; i > 0 {
		ms.SafeSlice.lock.Lock()
		ms.SafeSlice.items = ms.SafeSlice.items[i:]
		ms.SafeSlice.lock.Unlock()
	}
	return ms.SafeSlice.Add(items...)
}
