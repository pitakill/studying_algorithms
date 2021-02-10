package maximum_subarray

import (
	"fmt"
	"testing"
)

func Test_maxSubArray(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{1}, 1},
		{[]int{0}, 0},
		{[]int{-1}, -1},
		{[]int{-100_000}, -100_000},
	}

	for _, test := range tests {
		output := maxSubArray(test.input)

		if output != test.want {
			t.Errorf("maxSubArray outputs %d but wants %d", output, test.want)
		}
	}
}

func benchmark_maxSubArray(b *testing.B, input []int) {
	for i := 0; i < b.N; i++ {
		maxSubArray(input)
	}
}

func Benchmark_maxSubArray(b *testing.B) {
	benchmarks := []struct {
		input []int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}},
		{[]int{1}},
		{[]int{0}},
		{[]int{-1}},
		{[]int{-100_000}},
	}

	for _, benchmark := range benchmarks {
		name := fmt.Sprintf("%#v", benchmark.input)
		b.Run(name, func(b *testing.B) {
			benchmark_maxSubArray(b, benchmark.input)
		})
	}
}
