package recursive

import "testing"

func TestRecursiveFactorial(t *testing.T) {
	t.Run("given 1 factorial 1", func(t *testing.T) {
		given := 1
		want := 1

		get := Fac(given)
		if want != get {
			t.Errorf("given %v want %v factorial %v", given, want, get)
		}

	})
	t.Run("given 2 factorial 2", func(t *testing.T) {
		given := 2
		want := 2

		get := Fac(given)
		if want != get {
			t.Errorf("given %v want %v factorial %v", given, want, get)
		}

	})
	t.Run("given 3 factorial 6", func(t *testing.T) {
		given := 3
		want := 6

		get := Fac(given)
		if want != get {
			t.Errorf("given %v want %v factorial %v", given, want, get)
		}

	})
	t.Run("given 4 factorial 6", func(t *testing.T) {
		given := 4
		want := 24

		get := Fac(given)
		if want != get {
			t.Errorf("given %v want %v factorial %v", given, want, get)
		}

	})
	t.Run("given 5 factorial 120", func(t *testing.T) {
		given := 5
		want := 120

		get := Fac(given)
		if want != get {
			t.Errorf("given %v want %v factorial %v", given, want, get)
		}

	})
}

func TestRecursivePow(t *testing.T) {
	t.Run("given 2 Pow 3", func(t *testing.T) {
		x := 2
		n := 3

		want := 8

		get := Pow(x, n)
		if want != get {
			t.Errorf("given %v want %v Pow %v", x, n, get)
		}

	})
	t.Run("given 2 Pow 4", func(t *testing.T) {
		x := 2
		n := 4

		want := 16

		get := Pow(x, n)
		if want != get {
			t.Errorf("given %v want %v Pow %v", x, n, get)
		}

	})
}
