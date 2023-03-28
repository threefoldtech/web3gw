package main

import "testing"

type testStruct struct {
	a *int
	b int
}

func TestStateManagerLoadDefault(t *testing.T) {
	sm := StateManager[testStruct]{}
	res, exists := sm.Get("ab")
	if exists {
		t.Fatal("StateManager claims state exists but it doesn't")
	}
	if res.a != nil || res.b != 0 {
		t.Fatal("Expected default result, got ", res)
	}
}

func TestStateManagerSetAndLoad(t *testing.T) {
	sm := StateManager[testStruct]{}
	a := 5
	sm.Set("ab", testStruct{a: &a, b: 6})
	res, exists := sm.Get("ab")
	if !exists {
		t.Fatal("StateManager claims state doesn't exist but it does")
	}
	if res.a != &a || res.b != 6 {
		t.Fatal("Expected default result, got ", res)
	}
}
