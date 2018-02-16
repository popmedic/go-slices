package slices

import "sync"

type IItem interface {
	Get() interface{}
	Set(v interface{})
	Dup() IItem
	Compare(IItem) bool
}

type StringItem struct {
	value string
	lock  sync.RWMutex
}

func NewStringItem(s string) IItem {
	return &StringItem{
		value: s,
		lock:  sync.RWMutex{},
	}
}

func (s *StringItem) Get() interface{} {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.value
}

func (s *StringItem) Set(value interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if v, ok := value.(string); ok {
		s.value = v
	}
}

func (s *StringItem) Dup() IItem {
	if v, ok := s.Get().(string); ok {
		return NewStringItem(v)
	}
	return NewStringItem("")
}

func (s *StringItem) Compare(i IItem) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()

	if v, ok := i.Get().(string); ok {
		return v == s.value
	}
	return false
}

func (s *StringItem) String() string {
	return s.value
}
