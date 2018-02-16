package slices

import (
	"fmt"
	"sync"
)

type SafeSlice struct {
	items []IItem
	lock  sync.RWMutex
}

func NewSafeSlice(items ...IItem) *SafeSlice {
	n := &SafeSlice{lock: sync.RWMutex{}}
	n.Add(items...)
	return n
}

func (ss *SafeSlice) Add(items ...IItem) ISlice {
	ss.lock.Lock()
	defer ss.lock.Unlock()

	ss.items = append(ss.items, items...)
	return ss
}

func (ss *SafeSlice) Remove(items ...IItem) {
	ss.lock.Lock()
	defer ss.lock.Unlock()

	for _, item := range items {
		for idx, i := range ss.items {
			if i.Compare(item) {
				ss.removeAt(idx)
			}
		}
	}
}

func (ss *SafeSlice) removeAt(i int) {
	l := len(ss.items)
	if i < l {
		if i == l-1 {
			ss.items = ss.items[:l-1]
		} else if i == 0 {
			ss.items = ss.items[1:]
		} else {
			ss.items = append(ss.items[:i], ss.items[i+1:]...)
		}
	}
}

func (ss *SafeSlice) RemoveAt(i int) {
	ss.lock.Lock()
	defer ss.lock.Unlock()

	ss.removeAt(i)
}

func (ss *SafeSlice) RemoveAll() {
	ss.lock.Lock()
	defer ss.lock.Unlock()

	ss.items = []IItem{}
}

func (ss *SafeSlice) Get(i int) IItem {
	ss.lock.RLock()
	defer ss.lock.RUnlock()

	return ss.items[i].Dup()
}

func (ss *SafeSlice) All() []IItem {
	ss.lock.RLock()
	defer ss.lock.RUnlock()

	n := make([]IItem, len(ss.items))
	for x, i := range ss.items {
		n[x] = i.Dup()
	}
	return n
}

func (ss *SafeSlice) Len() int {
	ss.lock.RLock()
	defer ss.lock.RUnlock()

	return len(ss.items)
}

func (ss *SafeSlice) Contains(item IItem) bool {
	ss.lock.RLock()
	defer ss.lock.RUnlock()

	for _, i := range ss.items {
		if i.Compare(item) {
			return true
		}
	}
	return false
}

func (ss *SafeSlice) Compare(ss2 ISlice) bool {
	for _, i := range ss.items {
		if !ss2.Contains(i) {
			return false
		}
	}
	for i := 0; i < ss2.Len(); i++ {
		if !ss.Contains(ss2.Get(i)) {
			return false
		}
	}
	return true
}

func (ss *SafeSlice) String() string {
	items := ss.All()
	n := "["
	for _, i := range items {
		n = fmt.Sprintf("%s %q", n, i)
	}
	return n + " ]"
}
