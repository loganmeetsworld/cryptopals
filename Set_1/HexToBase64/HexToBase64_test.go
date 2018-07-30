package main

import (
	"testing"
)

func TestHexToBase64(t *testing.T) {
	tables := []struct {
		x string
		y string
	}{
		{"49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d", "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"},
	}
	for _, table := range tables {
		result, _ := hexToBase64(table.x)
		if result != table.y {
			t.Errorf("hexToBase64(%s) was incorrect, got: %s, want: %s.", table.x, table.y, result)
		}
	}
}
