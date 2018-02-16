package slices

import "sync"

type MaxSlice struct {
	*SafeSlice
	max int
}

func NewMaxSlice(max int, items ...IItem) ISlice {
	if i := len(items) - max; i > 0 {
		items = items[i:]
	}
	if v, ok := NewSafeSlice(items...).(*SafeSlice); ok {
		return &MaxSlice{
			SafeSlice: v,
			max:       max,
		}
	}
	return &MaxSlice{
		SafeSlice: &SafeSlice{
			lock: sync.RWMutex{},
		},
		max: max,
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
