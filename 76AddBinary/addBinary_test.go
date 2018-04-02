package _6AddBinary

import (
	"testing"
	"strings"
)

func Test_addBinary(t *testing.T) {
	if strings.Compare(addBinary("10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101","110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011"),
		"100") == 0 {
		t.Log("pass")
	} else {
		t.Error("can't pass")
	}
}