package lolw

import "testing"

func Test_lengthOfLastWord(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"a ", 1},
		{"hello world", 5},
		{" ", 0},
	}

	for _, test := range tests {
		output := lengthOfLastWord(test.input)

		if output != test.want {
			t.Errorf("lengthOfLastWord outputs %d but want %d", output, test.want)
		}
	}
}

func benchmark_lengthOfLastWord(b *testing.B, input string, fn func(string) int) {
	for i := 0; i < b.N; i++ {
		fn(input)
	}
}

func Benchmark_lengthOfLastWord(b *testing.B) {
	benchmarks := []struct {
		input string
	}{
		{"a "},
		{"hello world"},
		{" "},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.input, func(b *testing.B) {
			benchmark_lengthOfLastWord(b, benchmark.input, lengthOfLastWord)
		})
	}
}

func Test_lengthOfLastWord_0(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"a ", 1},
		{"hello world", 5},
		{" ", 0},
	}

	for _, test := range tests {
		output := lengthOfLastWord_0(test.input)

		if output != test.want {
			t.Errorf("lengthOfLastWord outputs %d but want %d", output, test.want)
		}
	}
}

func Benchmark_lengthOfLastWord_0(b *testing.B) {
	benchmarks := []struct {
		input string
	}{
		{"a "},
		{"hello world"},
		{" "},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.input, func(b *testing.B) {
			benchmark_lengthOfLastWord(b, benchmark.input, lengthOfLastWord_0)
		})
	}
}
