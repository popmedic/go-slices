package slices

type MaxSlice struct {
	*SafeSlice
	max int
}

func NewMaxSlice(max int, items ...IItem) *MaxSlice {
	ss := newMaxSafeSlice(max, items...)

	return &MaxSlice{
		SafeSlice: ss,
		max:       max,
	}
}

func newMaxSafeSlice(max int, items ...IItem) *SafeSlice {
	if i := len(items) - max; i > 0 {
		return NewSafeSlice(items[i:]...)
	} else {
		return NewSafeSlice(items...)
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

func (ms *MaxSlice) SetMax(max int) {
	ms.SafeSlice = newMaxSafeSlice(max, ms.SafeSlice.All()...)
}

func (ms *MaxSlice) GetMax() int {
	return ms.max
}
