package calculator

import (
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 2.5, 3.5, 6.0},
		{"negative numbers", -2.5, -3.5, -6.0},
		{"mixed numbers", -2.5, 3.5, 1.0},
		{"zero addition", 0, 5.5, 5.5},
		{"both zero", 0, 0, 0},
		{"large numbers", 1e10, 1e10, 2e10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%.2f, %.2f) = %.2f, expected %.2f", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 5.0, 3.0, 2.0},
		{"negative numbers", -5.0, -3.0, -2.0},
		{"mixed numbers", -5.0, 3.0, -8.0},
		{"subtract zero", 5.0, 0, 5.0},
		{"subtract from zero", 0, 5.0, -5.0},
		{"same numbers", 5.0, 5.0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Subtract(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Subtract(%.2f, %.2f) = %.2f, expected %.2f", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 2.5, 4.0, 10.0},
		{"negative numbers", -2.5, -4.0, 10.0},
		{"mixed numbers", -2.5, 4.0, -10.0},
		{"multiply by zero", 5.0, 0, 0},
		{"multiply by one", 5.0, 1, 5.0},
		{"fractions", 0.5, 0.5, 0.25},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Multiply(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Multiply(%.2f, %.2f) = %.2f, expected %.2f", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name        string
		a, b        float64
		expected    float64
		expectError bool
	}{
		{"positive numbers", 10.0, 2.0, 5.0, false},
		{"negative numbers", -10.0, -2.0, 5.0, false},
		{"mixed numbers", -10.0, 2.0, -5.0, false},
		{"divide by one", 5.0, 1.0, 5.0, false},
		{"divide zero", 0, 5.0, 0, false},
		{"fractions", 1.0, 3.0, 0.3333333333333333, false},
		{"divide by zero", 5.0, 0, 0, true},
		{"divide zero by zero", 0, 0, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Divide(tt.a, tt.b)

			if tt.expectError {
				if err == nil {
					t.Errorf("Divide(%.2f, %.2f) expected error but got none", tt.a, tt.b)
				}
				return
			}

			if err != nil {
				t.Errorf("Divide(%.2f, %.2f) unexpected error: %v", tt.a, tt.b, err)
				return
			}

			if result != tt.expected {
				t.Errorf("Divide(%.2f, %.2f) = %.16f, expected %.16f", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestPower(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive base and exponent", 2.0, 3.0, 8.0},
		{"negative base, even exponent", -2.0, 2.0, 4.0},
		{"negative base, odd exponent", -2.0, 3.0, -8.0},
		{"power of zero", 5.0, 0, 1.0},
		{"power of one", 5.0, 1.0, 5.0},
		{"zero to positive power", 0, 3.0, 0},
		{"fractional exponent", 4.0, 0.5, 2.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Power(tt.a, tt.b)
			if math.Abs(result-tt.expected) > 1e-10 {
				t.Errorf("Power(%.2f, %.2f) = %.2f, expected %.2f", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestSqrt(t *testing.T) {
	tests := []struct {
		name        string
		a           float64
		expected    float64
		expectError bool
	}{
		{"perfect square", 9.0, 3.0, false},
		{"non-perfect square", 2.0, math.Sqrt(2), false},
		{"zero", 0, 0, false},
		{"fraction", 0.25, 0.5, false},
		{"large number", 100.0, 10.0, false},
		{"negative number", -4.0, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Sqrt(tt.a)

			if tt.expectError {
				if err == nil {
					t.Errorf("Sqrt(%.2f) expected error but got none", tt.a)
				}
				return
			}

			if err != nil {
				t.Errorf("Sqrt(%.2f) unexpected error: %v", tt.a, err)
				return
			}

			if math.Abs(result-tt.expected) > 1e-10 {
				t.Errorf("Sqrt(%.2f) = %.10f, expected %.10f", tt.a, result, tt.expected)
			}
		})
	}
}

func TestIsEven(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected bool
	}{
		{"even positive", 4, true},
		{"odd positive", 3, false},
		{"even negative", -4, true},
		{"odd negative", -3, false},
		{"zero", 0, true},
		{"large even", 1000, true},
		{"large odd", 1001, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsEven(tt.n)
			if result != tt.expected {
				t.Errorf("IsEven(%d) = %v, expected %v", tt.n, result, tt.expected)
			}
		})
	}
}

func TestAbs(t *testing.T) {
	tests := []struct {
		name     string
		a        float64
		expected float64
	}{
		{"positive number", 5.5, 5.5},
		{"negative number", -5.5, 5.5},
		{"zero", 0, 0},
		{"negative zero", -0.0, 0},
		{"large positive", 1e10, 1e10},
		{"large negative", -1e10, 1e10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Abs(tt.a)
			if result != tt.expected {
				t.Errorf("Abs(%.2f) = %.2f, expected %.2f", tt.a, result, tt.expected)
			}
		})
	}
}

// Benchmark tests for performance measurement
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(1.5, 2.5)
	}
}

func BenchmarkDivide(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Divide(10.0, 2.0)
	}
}

func BenchmarkPower(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Power(2.0, 10.0)
	}
}
