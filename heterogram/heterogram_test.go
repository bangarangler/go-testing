package heterogram

import "testing"

func TestHeterogram(t *testing.T) {
	// // input := "toto"
	// input := "Alpha"
	// // input :=
	// expected := false
	// IsHeterogram(input)
	// if IsHeterogram(input) != expected {
	// 	t.Errorf("\n Word %s expected to be %t, but got %t \n", input, expected, !expected)
	// }
	for _, c := range testCases {
		if IsHeterogram(c.input) != c.expected {
			t.Errorf("\n Word %s expected to be %t, but is %t \n", c.input, c.expected, !c.expected)
		}
	}
}
