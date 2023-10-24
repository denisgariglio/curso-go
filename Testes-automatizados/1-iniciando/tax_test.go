package tax

import (
	"testing"
)

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expected := 5.0

	result := CalculateTax(amount)

	if result != expected {
		t.Errorf("Expected %f but got %f", expected, result)
	}

}

func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		amount, expect float64
	}

	table := []calcTax{
		{500.0, 5.0},
		{1000.0, 10.0},
		{5000.0, 10.0},
		{0.0, 0.0},
	}

	for _, item := range table {
		result := CalculateTax(item.amount)
		if result != item.expect {
			t.Errorf("Expected %f but got %f", item.expect, result)
		}
	}
}

// para verificar a cobertura
// go test -coverprofile=coveragr.out

// para mostrar qual a cobertura que vc tem em formato html
// go tool cover -html=coverage.out

func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500.0)
	}
}

func BenchmarkCalculateTax2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax2(500.0)
	}
}

// go test -bench=.
// rodar somente o bench
// go test -bench=. -run=^#
// go test -bench=. -run=^# -count=10 --mostra a cada 10
// go test -bench=. -run=^# -count=10 -benchtime=3s --roda por apenas 3segundos
// go test -bench=. -run=^# -benchmem --verifica alocação de memoria

func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-1, -2, -2.5, 500.0, 1000.0, 1501.0}
	for _, amount := range seed {
		f.Add(amount)
	}
	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)
		if amount <= 0 && result != 0 {
			t.Errorf("Received %f but expected 0", result)
		}
		if amount > 20000 && result != 20 {
			t.Errorf("Received %f but expected 20", result)
		}

	})
}

// go test -fuzz=.
// go test -fuzz=. -run=^#
// go test -fuzz=. -fuzztime=5s -run=^#
