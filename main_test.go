package main

import "testing"

func TestReverseWords(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"the sky is blue", "blue is sky the"},
		{"  hello world  ", "world hello"},
		{"a good   example", "example good a"},
		{"single", "single"},
		{"", ""},
		{"   ", ""},
	}
	for _, c := range cases {
		if got := ReverseWords(c.in); got != c.want {
			t.Errorf("ReverseWords(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
