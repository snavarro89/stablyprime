package numbers

import "errors"

type NumbersModel struct {
	DB interface{} // <-- This can be any database you want to run queries against.
}

func (NumbersModel) GetPrime(number int) (int, error) {
	//We will only process positive integers.
	//Another approach would be to allow negative numbers and use the absolute value. For demostration purposes
	//This api will not allow negative numbers.
	if number < 0 {
		return 0, errors.New("Only positive numbers allowed")
	}

	//If the number is either 0-3 then just return the same value as they are their own primes.
	if number <= 3 {
		return number, nil
	}

	number--

	//If number is even it is not prime, if the input number is even just decrease by 1 its value.
	if number%2 == 0 {
		number--
	}

	result := -1
	isPrime := true
	//First Approach: Start iterating from the top for each number starting with n-1 (or n-2 in case it was an odd number) I
	// will determine if the number is prime, the first prime number to be found is the closest to n
	for i := number; i > 0; i -= 2 {
		for j := 2; j < i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			result = i
			break
		}
		isPrime = true
	}

	return result, nil

}
