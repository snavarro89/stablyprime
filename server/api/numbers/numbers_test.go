package numbers

import "testing"

func TestZero(t *testing.T) {
	prime, err := GetPrime(0)
	expect := 0
	if err != nil || prime != 0 {
		t.Fatalf(`GetPrime(0) = %b, %v, want match for %b, nil`, prime, err, expect)
	}
}

func TestOne(t *testing.T) {
	prime, err := GetPrime(1)
	expect := 1
	if err != nil || prime != 1 {
		t.Fatalf(`GetPrime(1) = %b, %v, want match for %b, nil`, prime, err, expect)
	}
}

func TestNegative(t *testing.T) {
	prime, err := GetPrime(-2)
	expect := "Incorrect input: Only posisitve numbers allowed"
	if err == nil {
		t.Fatalf(`GetPrime(-5) = %b, %v, want match for 0, %s`, prime, err, expect)
	}
}

func TestOdd(t *testing.T) {
	prime, err := GetPrime(57)
	expect := 53
	if err != nil || prime != 53 {
		t.Fatalf(`GetPrime(53) = %b, %v, want match for %b, nil`, prime, err, expect)
	}
}

func TestEven(t *testing.T) {
	prime, err := GetPrime(56)
	expect := 53
	if err != nil || prime != 53 {
		t.Fatalf(`GetPrime(53) = %b, %v, want match for %b, nil`, prime, err, expect)
	}
}

func TestPrime(t *testing.T) {
	prime, err := GetPrime(53)
	expect := 47
	if err != nil || prime != 47 {
		t.Fatalf(`GetPrime(47) = %b, %v, want match for %b, nil`, prime, err, expect)
	}
}
