package maps

import "testing"

func TestNonRepeating(t *testing.T) {
	tests := []struct {
		a string
		b int
	}{
		{"abc", 3},
		{"abcdabdcbd", 4},
		{"黑灰化肥发黑发灰会挥发", 5},
	}
	for _, tt := range tests {
		if actual := NonRepeating(tt.a); actual != tt.b {
			t.Errorf("NonRepeating(%s) Got %d Expected %d", tt.a, actual, tt.b)
		}
	}

}
