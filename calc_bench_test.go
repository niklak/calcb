package calcb

import (
	_ "embed"
	"encoding/json"
	"testing"
)

//go:embed test-data/test-values.json
var testValuesJSON []byte
var testValues []float64

func init() {
	if err := json.Unmarshal(testValuesJSON, &testValues); err != nil {
		panic(err)
	}
}

func BenchmarkRsiDirty(b *testing.B) {

	period := 14

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = rsiDirty(testValues, period)
	}
}

func BenchmarkRsiClean(b *testing.B) {

	period := 14

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = rsiClean(testValues, period)
	}
}

func BenchmarkRsiOpt(b *testing.B) {

	period := 14

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = rsiOpt(testValues, period)
	}
}
