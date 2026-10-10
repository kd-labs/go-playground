package main

func minEatingSpeed(piles []int, h int) int {
	l, r := 1, 1

	for _, b := range piles {
		r = max(r, b)
	}

	for l < r {
		mid := l + (r-l)/2

		t := timeToFinish(piles, mid)
		if t <= h {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func timeToFinish(piles []int, rate int) int {
	var time int
	for _, pile := range piles {
		if pile%rate == 0 {
			time += pile / rate
		} else {
			time += (pile / rate) + 1
		}
	}
	return time
}
