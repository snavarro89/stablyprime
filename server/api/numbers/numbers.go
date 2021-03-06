package numbers

import (
	"errors"
	"fmt"
)

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

	//If the number is either 0 or 2 then there are no prime numbers before this numbers
	if number < 3 {
		return 0, errors.New(fmt.Sprintf("There are no prime numbers before %d", number))
	}

	//If number is even it is not prime, if the input number is even just decrease by 1 its value.
	//I will hardcode the three in here to save one iteration when entering an odd number or else I would have to
	//iterate the first even number as well. The problem is that 2 is the only even prime number in the list.
	if number%2 == 0 || number == 3 {
		number--
	} else {
		number -= 2
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
