package main

import (
	"reflect"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	got := GroupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	want := [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GroupAnagrams = %v, want %v", got, want)
	}

	if got := GroupAnagrams([]string{""}); !reflect.DeepEqual(got, [][]string{{""}}) {
		t.Errorf("empty string: %v", got)
	}
	if got := GroupAnagrams([]string{"a"}); !reflect.DeepEqual(got, [][]string{{"a"}}) {
		t.Errorf("single: %v", got)
	}
	// Repeated letters must distinguish: aab vs abb.
	got = GroupAnagrams([]string{"aab", "abb", "baa"})
	want = [][]string{{"aab", "baa"}, {"abb"}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("multiplicity: %v, want %v", got, want)
	}
	if got := GroupAnagrams(nil); len(got) != 0 {
		t.Errorf("nil input: %v", got)
	}
}
