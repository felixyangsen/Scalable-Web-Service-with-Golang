package palindrom

import "testing"

func Test_IsPalindrom_PositifCond(t *testing.T) {
	t.Run("result palindrome", func(t *testing.T) {
		var expected = true
		actual := IsPalindrom("malam")

		if actual != expected {
			t.Fatalf("expected %v, found %v", expected, actual)
		}
	})
}

func Test_IsPalindrom_NegatifCond(t *testing.T) {
	var expected = false
	actual := IsPalindrom("pagi")

	if actual != expected {
		t.Fatalf("expected %v, found %v", expected, actual)
	}
}
