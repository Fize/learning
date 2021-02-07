package add_digits

import "testing"

func TestAddDigits(t *testing.T) {
	num := 38
	except := 2
	result := AddDigits(num)
	t.Logf("input: %d, except: %d, result: %d\n", num, except, result)
}
