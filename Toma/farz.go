package Toma

func Calculate(summ int, currency string) (min, max int) {
	minPersent := 4
	maxPersent := 6
	min = (summ * minPersent / 12) / 100
	max = (summ * maxPersent / 12) / 100

	return min, max
}
