package utils

import "testing"

func TestAbs(t *testing.T) {
	testCases := []struct {
		name string
		got  int64
		want int64
	}{
		{"Abs(-1)", -10, 10},
		{"Abs(50)", 50, 50},
		{"Abs(0)", 0, 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := Abs(tc.got)
			if got != tc.want {
				t.Errorf("got %d; want %d", tc.got, tc.want)
			}
		})
	}
}

func TestLimitMaxValue(t *testing.T) {
	testCases := []struct {
		name string
		got  int64
		want int64
	}{
		{"LimitMaxValue(100, 95)", LimitMaxValue(100, 95), 95},
		{"LimitMaxValue(50, 95)", LimitMaxValue(50, 95), 50},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.got != tc.want {
				t.Errorf("got %d; want %d", tc.got, tc.want)
			}
		})
	}
}

func TestLimitMinValue(t *testing.T) {
	testCases := []struct {
		name string
		got  int64
		want int64
	}{
		{"LimitMinValue(4, 10)", LimitMinValue(4, 10), 10},
		{"LimitMinValue(15, 10)", LimitMinValue(15, 10), 15},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.got != tc.want {
				t.Errorf("got %d; want %d", tc.got, tc.want)
			}
		})
	}
}

func TestUint32(t *testing.T) {
	testCases := []struct {
		name string
		got  uint32
		want uint32
	}{
		{"Uint32(5)", Uint32("5"), 5},
		{"Uint32(10)", Uint32("10"), 10},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.got != tc.want {
				t.Errorf("got %d; want %d", tc.got, tc.want)
			}
		})
	}
}
