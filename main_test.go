package main

import (
	"fmt"
	"testing"
)

func TestGenerateQRCodeReturnsValue(t *testing.T) {
	// result := hola()
	result := 2
	fmt.Println(result)
	if result == 2 {
		t.Errorf("Generated QRCode is nil")
	}
}
