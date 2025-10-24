package calcb

import "testing"

func floatsAlmostEqual(a, b, epsilon float64) bool {
	diff := a - b
	if diff < 0 {
		diff = -diff
	}
	return diff <= epsilon
}

func Test_rsiFunc(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		values []float64
		period int
		want   float64
	}{
		{
			name:   "Test RSI funcs equal 1",
			values: testValues,
			period: 14,
			want:   55.369383092155466,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			funcs := []func([]float64, int) float64{
				rsiDirty,
				rsiClean,
				rsiOpt,
			}
			for _, fun := range funcs {
				got := fun(tt.values, tt.period)
				if !floatsAlmostEqual(got, tt.want, 1e-9) {
					t.Errorf("%T() = %v, want %v", fun, got, tt.want)
				}
			}

		})
	}
}
