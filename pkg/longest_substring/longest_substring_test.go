package longest_substring

import "testing"

func TestLongestSubstring(t *testing.T) {
	s1 := "abcabckJ"
	r1 := LengthOfLongestSubstring(s1)
	except1 := 5
	t.Logf("input: '%s', except: %d, result: %d", s1, except1, r1)

	s2 := "a"
	r2 := LengthOfLongestSubstring(s2)
	except2 := 1
	t.Logf("input: '%s', except: %d, result: %d", s2, except2, r2)
}
