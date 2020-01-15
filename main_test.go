package main

import (
	"fmt"
	"testing"
)

func TestMerge1(t *testing.T) {
	ret := fmt.Sprintf("%v", merge([]interval{{1, 35}, {4, 15}, {10, 30}}))
	want := "[[1 35]]"
	if ret != want {
		t.Errorf("TestMerge error: got '%v' want '%v'", ret, want)
	}
}

func TestMerge2(t *testing.T) {
	ret := fmt.Sprintf("%v", merge([]interval{{1, 20}, {10, 30}}))
	want := "[[1 30]]"
	if ret != want {
		t.Errorf("TestMerge error: got '%v' want '%v'", ret, want)
	}
}

func TestMerge3(t *testing.T) {
	ret := fmt.Sprintf("%v", merge([]interval{{47, 60}, {1, 3}, {17, 18}, {20, 24}, {24, 25}, {7, 11}, {45, 50}, {10, 17}, {30, 32}, {31, 35}, {4, 15}}))
	want := "[[1 3] [4 18] [20 25] [30 35] [45 60]]"
	if ret != want {
		t.Errorf("TestMerge error: got '%v' want '%v'", ret, want)
	}
}

func TestMergeEmpty(t *testing.T) {
	ret := fmt.Sprintf("%v", merge([]interval{}))
	want := "[]"
	if ret != want {
		t.Errorf("TestMerge error: got '%v' want '%v'", ret, want)
	}
}
