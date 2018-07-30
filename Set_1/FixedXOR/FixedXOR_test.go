package main

import (
	"bytes"
	"testing"
)

func TestFixedXOR(t *testing.T) {
	tables := []struct {
		x string
		y string
		z string
	}{
		{"1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965", "746865206b696420646f6e277420706c6179"},
	}
	for _, table := range tables {
		result := fixedXOR(table.x, table.y)
		if !bytes.Equal(result, hexDecode(table.z)) {
			t.Errorf("fixedXOR(%s, %s) was incorrect, got: %s, want: %s.", table.x, table.y, table.z, result)
		}
	}
}
