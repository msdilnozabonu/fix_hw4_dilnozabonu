package product

import (
	"fmt"
	"testing"
)

func ExampleCalculate() {
	calculate := Calculate(1_000_00)
	fmt.Println(calculate)
	//Output:
	//8333
}

func ExampleCalculate_negative() {
	calculate := Calculate(-1_000_00)
	fmt.Println(calculate)
	//Output:
	//-8333
}

func ExampleCalculate_zero() {
	calculate := Calculate(0)
	fmt.Println(calculate)
	//Output:
	//0
}

func ExampleCalculate_large_amount() {
	calculate := Calculate(12_000_000_00)
	fmt.Println(calculate)
	//Output:
	//100000000
}

func ExampleCalculate_with_tiyin() {
	calculate := Calculate(1_000_99)
	fmt.Println(calculate)
	//Output:
	//8341
}

func ExampleCalculate_minimal() {
	calculate := Calculate(12)
	fmt.Println(calculate)
	//Output:
	//1
}

func ExampleCalculate_eleven() {
	calculate := Calculate(11)
	fmt.Println(calculate)
	//Output:
	//0
}

func TestCalculate_various_amounts(t *testing.T) {
	tests := []struct {
		name  string
		price int
		want  int
	}{
		{
			name:  "Standard price",
			price: 120_000_00,
			want:  10_000_00,
		},
		{
			name:  "Price with tiyin",
			price: 120_000_50,
			want:  10_000_04,
		},
		{
			name:  "Very large price",
			price: 999_999_999_99,
			want:  83_333_333_33,
		},
		{
			name:  "Small price",
			price: 600_00,
			want:  50_00,
		},
		{
			name:  "Price just below rounding",
			price: 599_99,
			want:  49_99,
		},
		{
			name:  "Exact division",
			price: 12_000_00,
			want:  1_000_00,
		},
		{
			name:  "One sum",
			price: 1_00,
			want:  8,
		},
		{
			name:  "Maximum int32 value",
			price: 2147483647,
			want:  178956970,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Calculate(tt.price)
			if got != tt.want {
				t.Errorf("Calculate(%d) = %d, want %d", tt.price, got, tt.want)
			}
		})
	}
}

func TestCalculate_rounding(t *testing.T) {
	tests := []struct {
		name  string
		price int
		want  int
	}{
		{
			name:  "Rounds down when exact",
			price: 12,
			want:  1,
		},
		{
			name:  "Rounds down below 0.5",
			price: 17,
			want:  1,
		},
		{
			name:  "Rounds down at 0.5",
			price: 18,
			want:  1,
		},
		{
			name:  "Rounds down above 0.5",
			price: 23,
			want:  1,
		},
		{
			name:  "Next whole number",
			price: 24,
			want:  2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Calculate(tt.price)
			if got != tt.want {
				t.Errorf("Calculate(%d) = %d, want %d", tt.price, got, tt.want)
			}
		})
	}
}
