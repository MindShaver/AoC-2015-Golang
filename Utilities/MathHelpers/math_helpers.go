package MathHelpers

func Min(values []int) int {
	min := values[0]
	for _, v := range values {
		if v < min {
			min = v
		}
	}

	return min
}

func Sum(values []int) int {
	sum := 0
	for _, v := range values {
		sum += v
	}

	return sum
}

func Area(dimOne, dimTwo int) int {
	return dimOne * dimTwo
}

func SurfaceArea(dimOne, dimTwo int) int {
	return 2 * (dimOne * dimTwo)
}

func Perimeter(length, width int) int {
	return length + length + width + width
}

func Volume(length, width, height int) int {
	return length * width * height;
}
