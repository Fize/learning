package main

import "fmt"

func mid(array []int) float64 {
	l := len(array)
	if (l % 2) == 0 {
		return (float64(array[l/2]) + float64(array[l/2-1])) / 2
	}
	return float64(array[l/2])
}

func main() {
	num1 := []int{1, 4, 5, 7, 8, 9}
	num2 := []int{2, 4, 9}
	m := len(num1)
	n := len(num2)

	if m > n {
		temp := num1
		num1 = num2
		num2 = temp
		m = len(num1)
		n = len(num2)
	}
	k := (m + n)
	for {
		if m+n == 2 {
			s := []int{num1[0], num2[0]}
			fmt.Println(mid(s))
			break
		}
		if m+n == 3 {
			if m > n {
				if num1[1] > num2[0] && num2[0] > num1[0] {
					fmt.Println(float64(num2[0]))
				}
				if num2[0] > num1[1] {
					fmt.Println(float64(num1[1]))
				}
				if num1[0] > num2[0] {
					fmt.Println(float64(num1[0]))
				}
			}
			if m < n {
				if num2[1] > num1[0] && num1[0] > num2[0] {
					fmt.Println(float64(num1[0]))
				}
				if num1[0] > num2[1] {
					fmt.Println(float64(num2[1]))
				}
				if num2[0] > num1[0] {
					fmt.Println(float64(num2[0]))
				}
			}
			break
		}
		mid1 := mid(num1)
		mid2 := mid(num2)
		if mid1 <= mid2 {
			num1 = num1[m/2+1:]
			if n%2 == 0 {
				num2 = num2[:n/2]
			} else {
				num2 = num2[:n/2+1]
			}
			m = len(num1)
			n = len(num2)
		} else {
			if m%2 == 0 {
				num1 = num1[:m/2]
			} else {
				num1 = num1[:m/2+1]
			}
			num2 = num2[n/2+1:]
			m = len(num1)
			n = len(num2)
		}
		fmt.Println("num1:", num1, m)
		fmt.Println("num2:", num2, n)
	}
}
