package slices

import "testing"

func TestSafeSliceRemove(t *testing.T) {
	exp := NewSafeSlice(NewStringItem("kevin"), NewStringItem("rules"))
	given := NewSafeSlice(
		NewStringItem("who"),
		NewStringItem("knows"),
		NewStringItem("that"),
		NewStringItem("kevin"),
		NewStringItem("the"),
		NewStringItem("great"),
		NewStringItem("rules"),
		NewStringItem("the"),
		NewStringItem("city"),
	)
	got := NewSafeSlice(given.All()...)
	got.Remove(
		NewStringItem("who"),
		NewStringItem("knows"),
		NewStringItem("that"),
		NewStringItem("the"),
		NewStringItem("great"),
		NewStringItem("the"),
		NewStringItem("city"),
	)
	if !got.Compare(exp) {
		t.Errorf("given %s expected %s got %s", given, exp, got)
	}
	exp = NewSafeSlice(NewStringItem("kevin"))
	got.RemoveAt(1)
	if !got.Compare(exp) {
		t.Errorf("given %s expected %s got %s", given, exp, got)
	}
	exp = NewSafeSlice()
	got.RemoveAll()
	if !got.Compare(exp) {
		t.Errorf("given %s expected %s got %s", given, exp, got)
	}
	expStr := "[ who knows that kevin the great rules the city ]"
	given = NewSafeSlice(
		NewStringItem("who"),
		NewStringItem("knows"),
		NewStringItem("that"),
		NewStringItem("kevin"),
		NewStringItem("the"),
		NewStringItem("great"),
		NewStringItem("rules"),
		NewStringItem("the"),
		NewStringItem("city"),
	)
	gotStr := given.String()
	if expStr != gotStr {
		t.Errorf("given %q expected %q got %q", given, expStr, gotStr)
	}
}

func TestMaxSlice(t *testing.T) {
	given := NewMaxSlice(
		10,
		NewStringItem("who"),
		NewStringItem("knows"),
		NewStringItem("that"),
		NewStringItem("kevin"),
		NewStringItem("the"),
		NewStringItem("great"),
		NewStringItem("rules"),
	)
	exp := NewMaxSlice(
		10,
		NewStringItem("who"),
		NewStringItem("knows"),
		NewStringItem("that"),
		NewStringItem("kevin"),
		NewStringItem("the"),
		NewStringItem("great"),
		NewStringItem("rules"),
		NewStringItem("the"),
		NewStringItem("city"),
	)
	got := NewMaxSlice(
		10,
		NewStringItem("who"),
		NewStringItem("knows"),
		NewStringItem("that"),
		NewStringItem("kevin"),
		NewStringItem("the"),
		NewStringItem("great"),
		NewStringItem("rules"),
	)
	got.Add(
		NewStringItem("the"),
		NewStringItem("city"),
	)
	if got.Compare(given) {
		t.Errorf("given %q expected %q got %q", given, exp, got)
	}
	exp = NewMaxSlice(
		10,
		NewStringItem("who"),
		NewStringItem("knows"),
		NewStringItem("that"),
		NewStringItem("kevin"),
		NewStringItem("the"),
		NewStringItem("great"),
		NewStringItem("rules"),
		NewStringItem("the"),
		NewStringItem("city"),
		NewStringItem("who"),
		NewStringItem("knows"),
	)
	got.Add(
		NewStringItem("who"),
		NewStringItem("knows"),
	)
	if got.Compare(given) {
		t.Errorf("given %q expected %q got %q", given, exp, got)
	}
}
