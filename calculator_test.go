package calculator_test

import (
	"calculator"
	"math/rand"
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	var want float64 = 7
	got := calculator.Add(2, 2)

	if want != got {
		t.Errorf("want %f, got %f", want, got)
		t.Fatalf("want %f---, got %f", want, got)

	}

}

func TestSubtract(t *testing.T) {
	t.Parallel()
	var want float64 = 1
	got := calculator.Subtract(4, 2)
	if want != got {
		t.Fatalf("want %f, got %f", want, got)
	}

}

// func TestMultiply(t *testing.T) {
// 	t.Parallel()
// 	var want float64 = 2 * 5

// 	got := calculator.Multiply(2, 5)
// 	if want != got {
// 		t.Fatalf("want %f, go %f ", want, got)
// 	}
// }
func TestMultiply(t *testing.T) {
	testCases := []struct {
		desc string
		a    float64
		b    float64
		want float64
	}{
		{
			desc: "teste Ok",
			a:    2,
			b:    5,
			want: 10,
		},
		{
			desc: "teste OK II",
			a:    -3,
			b:    5,
			want: -15,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := calculator.Multiply(tC.a, tC.b)
			if tC.want != got {
				t.Errorf("want %f, got %f ", tC.want, got)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	testCases := []struct {
		desc        string
		a           float64
		b           float64
		errExpected bool
		want        float64
	}{
		{
			desc:        "Teste with a simple operation",
			a:           4,
			b:           2,
			errExpected: false,
			want:        2,
		},
		{
			desc:        "Teste divide by 0",
			a:           4,
			b:           0,
			errExpected: true,
			want:        0,
		},
	}

	// esperar error e dá error  -- não dá falha
	// esperar erro e não dá erro -  executa a falha
	// não espera erro e não dá erro - não dá falha
	// não espera erro e dá erro - executa a falha

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, err := calculator.Divide(tC.a, tC.b)
			errorReceived := err != nil
			if errorReceived != tC.errExpected {
				t.Fatalf("Undexpected error %s", err)
			}

			if tC.want != got {
				t.Errorf("want %f, got %f ", tC.want, got)
			}
		})
	}
}
func TestAddRandom(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	total := rand.Intn(120)
	for i := 0; i < total; i++ {
		numA := rand.Float64()
		numB := rand.Float64()
		want := numA + numB
		got := calculator.Add(numA, numB)
		if want != got {
			t.Errorf("want %f---, got %f", want, got)
		}

	}
}
