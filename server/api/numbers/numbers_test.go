package numbers

import (
	"fmt"
	"testing"

	A "github.com/snavarro89/stablyprime/app"
)

func TestZero(t *testing.T) {
	app := A.App{
		Data: A.Data{
			Numbers: NumbersModel{DB: nil},
		},
	}
	prime, err := app.Data.GetPrime(0)
	expect := 0
	if err != nil || prime != 0 {
		t.Fatalf(`GetPrime(0) = %b, %v, want match for %b, nil`, prime, err, expect)
	}
}

func TestOne(t *testing.T) {
	app := A.App{
		Data: A.Data{
			Numbers: NumbersModel{DB: nil},
		},
	}
	prime, err := app.Data.GetPrime(1)
	expect := 0
	expectError := fmt.Sprintf("There are no prime numbers before %d", 1)
	if err == nil || prime != 0 {
		t.Fatalf(`GetPrime(1) = %b, %v, want match for %b, %v`, prime, err, expect, expectError)
	}
}

func TestNegative(t *testing.T) {
	app := A.App{
		Data: A.Data{
			Numbers: NumbersModel{DB: nil},
		},
	}
	prime, err := app.Data.GetPrime(-2)
	expect := "Only posisitve numbers allowed"
	if err == nil {
		t.Fatalf(`GetPrime(-5) = %b, %v, want match for 0, %s`, prime, err, expect)
	}
}

func TestOdd(t *testing.T) {
	app := A.App{
		Data: A.Data{
			Numbers: NumbersModel{DB: nil},
		},
	}
	prime, err := app.Data.GetPrime(57)
	expect := 53
	if err != nil || prime != 53 {
		t.Fatalf(`GetPrime(53) = %b, %v, want match for %b, nil`, prime, err, expect)
	}
}

func TestEven(t *testing.T) {
	app := A.App{
		Data: A.Data{
			Numbers: NumbersModel{DB: nil},
		},
	}
	prime, err := app.Data.GetPrime(56)
	expect := 53
	if err != nil || prime != 53 {
		t.Fatalf(`GetPrime(53) = %d, %v, want match for %d, nil`, prime, err, expect)
	}
}

func TestPrime(t *testing.T) {
	app := A.App{
		Data: A.Data{
			Numbers: NumbersModel{DB: nil},
		},
	}
	prime, err := app.Data.GetPrime(53)
	expect := 47
	if err != nil || prime != 47 {
		t.Fatalf(`GetPrime(47) = %b, %v, want match for %b, nil`, prime, err, expect)
	}
}
