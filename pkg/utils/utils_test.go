package utils

import "testing"

func TestEg(t *testing.T) {
	if "" != "" {
		t.Errorf("Never happen")
	}
}
