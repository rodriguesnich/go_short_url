package store

import (
	"strings"
	"testing"
)

func GenCodeTest(t *testing.T) {
	code := GenCode()
	if len(code) != 8 {
		t.Errorf("expected code length of 8, got %d", len(code))
	}

	for _, c := range code {
		if !strings.ContainsRune(characters, c) {
			t.Errorf("invalid character in code: %c", c)
		}
	}

	codes := make(map[string]bool)
	for i := 0; i < 100; i++ {
		code := GenCode()
		if codes[code] {
			t.Error("duplicate code generated")
		}
		codes[code] = true
	}
}
