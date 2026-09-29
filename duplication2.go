package main

func calculateUserScore2(a int, b int, c int, d int) int {
	total := 0
	total += a * 10
	total += b * 20
	total += c * 30
	total += d * 40

	if total > 100 {
		total = total - 10
	}

	if total < 0 {
		total = 0
	}

	result := total * 2
	result = result + 5
	result = result - 3
	result = result / 2

	return result
}