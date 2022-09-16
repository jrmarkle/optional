package optional

import "testing"

func TestOptionalString(t *testing.T) {
	s1 := Some("hello")
	if !s1.Is() {
		t.Error("wrong value for Is")
	}
	if s1.Get() != "hello" {
		t.Error("failed to get string value")
	}

	s2 := None[string]()
	if s2.Is() {
		t.Error("wrong value for Is")
	}
	if s2.Get() != "" {
		t.Error("should get empty string")
	}
	if s2.GetOr("fallback") != "fallback" {
		t.Error("didn't get fallback value")
	}

	s2 = s1
	if !s1.Is() {
		t.Error("wrong value for Is")
	}
	if s1.Get() != "hello" {
		t.Error("failed to get string value")
	}

	s2 = Some("world")
	if s1.Get() != "hello" {
		t.Error("failed to get original string value")
	}
	if s2.Get() != "world" {
		t.Error("failed to get string value")
	}
}

func TestOptionalMap(t *testing.T) {
	m1 := Some(map[string]int{"foo": 11, "bar": 22})
	if !m1.Is() {
		t.Error("wrong value for Is")
	}
	if m1.Get()["foo"] != 11 {
		t.Error("failed to get map value")
	}

	m2 := None[map[string]int]()
	if m2.Is() {
		t.Error("wrong value for Is")
	}
	if m2.Get() != nil {
		t.Error("should get nil map")
	}
	if m2.GetOr(map[string]int{"foo": 111})["foo"] != 111 {
		t.Error("didn't get fallback value")
	}

	m2 = m1
	if !m2.Is() {
		t.Error("wrong value for Is")
	}
	if m2.Get()["bar"] != 22 {
		t.Error("failed to get copied map value")
	}
}

func TestFromToPtr(t *testing.T) {
	var i *int
	none := FromPtr(i)

	v := 10
	i = &v
	ten := FromPtr(i)

	if none.Is() {
		t.Error("wrong value for Is")
	}
	if !ten.Is() {
		t.Error("wrong value for Is")
	}

	if none.Get() != 0 {
		t.Error("should get zero")
	}
	if ten.Get() != 10 {
		t.Error("should get ten")
	}

	if none.ToPtr() != nil {
		t.Error("pointer should be nil")
	}
	tenPtr := ten.ToPtr()
	if tenPtr == nil || *tenPtr != 10 {
		t.Error("should get pointer to ten")
	}
	if tenPtr == i {
		t.Error("pointer should not be copied")
	}
}
