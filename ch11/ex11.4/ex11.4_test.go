package word

import (
	"fmt"
	"math/rand"
	"time"
	"unicode"
	"unicode/utf8"
)

import "testing"

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"detartrated", true},
		{"A man, a plan, a canal: Panama", true},
		{"Evil I did dwell; lewd did I live.", true},
		{"Able was I ere I saw Elba", true},
		{"été", true},
		{"Et se resservir, ivresse reste.", true},
		{"palindrome", false}, // non-palindrome
		{"desserts", false},   // semi-palindrome
	}
	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.want {
			t.Errorf("IsPalindrome(%q) = %v", test.input, got)
		}
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome("A man, a plan, a canal: Panama")
	}
}

func ExampleIsPalindrome() {
	fmt.Println(IsPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(IsPalindrome("palindrome"))
	// Output:
	// true
	// false
}

// randomPalindrome returns a palindrome whose length and contents
// are derived from the pseudo-random number generator rng.
func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24
	runes := make([]rune, n)
	for i, j := 0, n; i < j; {
		r := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
		runes[i] = r
		if unicode.IsSpace(r) || unicode.IsPunct(r) {
			i++
			continue
		}
		runes[j-1] = r
		i++
		j--
	}
	return string(runes)
}

func randomNotPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		for {
			c := rng.Intn(0x999)
			r := rune(c) // random rune up to '\u0999'
			r2 := rune(c + 1)
			if unicode.IsLetter(r) == true && unicode.IsLetter(r2) == true && unicode.ToLower(r) != unicode.ToLower(r2) {
				runes[i] = r
				runes[n-1-i] = r2
				break
			}
		}

	}
	return string(runes)
}

func TestRandomPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}

func TestRandomNotPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomNotPalindrome(rng)
		if utf8.RuneCountInString(p) > 1 && IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = true", p)
		}
	}
}

func TestSingleNotPalidromes(t *testing.T) {
	var p string
	p = "Ĕĕ"
	if !IsPalindrome(p) {
		t.Errorf("IsPalindrome(%q) = false", p)
	}
}
