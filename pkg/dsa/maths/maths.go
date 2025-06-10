package maths

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

/**
  a=bq+r
  a*b=lcm(a,b) *gcd(a,b)
  sum{1.. n} = (n*(n+1))/2
  s= (n*(n+1)*(2n+1))/6 // Sum of Squres
*/

/**
Let m be a positive integer and let a and b be integers. <br>
* a is congruent to b under module m .     m | a-b  <=> a%m=b%m=r. Same reminder.<br>
* a-b=mk => a=mk+b;<br>

 If a ≡ b (mod m) and c ≡ d (mod m),
 then 1. a + c ≡ b + d (mod m)
      2. ac ≡ bd (mod m).<br><br>
      3. (a + b)%m = ((a%m) + (b%m))%m
      4. (a*b)% m = ((a%m)*(b%m))%m.
      5. (a-b)%m = ((a%m)-(b%m)+m)%m
	  6. (a/b)%m =((a%m)*((1/b)%m))%m

	Extended Eucledan algorithm : gcd(a,b)=ax+by - linear combination.
	Coefficients x and y are used to find modular inverse.
*/

func FactorsArr(x int) []int {
	if x > 0 {
		factors := make([]int, 0, x/2)
		factors = append(factors, 1)
		for i := 2; i <= x/2; i++ {
			if x%i == 0 {
				factors = append(factors, i)
			}
		}
		factors = append(factors, x)
		return factors
	}
	return nil
}

func Factors(x int) string {
	if x > 0 {
		var sb strings.Builder
		sb.WriteString("1, ")
		for i := 2; i <= x/2; i++ {
			if x%i == 0 {
				sb.WriteString(strconv.Itoa(i) + ", ")
			}
		}
		sb.WriteString(strconv.Itoa(x))
		return sb.String()
	}
	return ""
}

func PrimeFactors(n int) string {
	// Print the number of 2s that divide n
	fmt.Printf("Prime factor of= %v => ", n)
	for n%2 == 0 {
		fmt.Print("2 ")
		n /= 2
	}

	// n must be odd at this point.  So we can
	// skip one element (Note i = i +2)
	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		// While i divides n, print i and divide n
		for n%i == 0 {
			fmt.Print(i)
			n /= i
		}
	}

	// This condition is to handle the case when
	// n is a prime number greater than 2
	if n > 2 {
		fmt.Print(n)
	}
	return ""

}

// Exponentiation
// O(n)
func PowerR(x int, exponent int) int {
	if exponent == 0 {
		return 1
	} else {
		return x * PowerR(x, exponent-1)
	}
}

func PowerRBinary(x int, exponent int) int {
	if exponent == 0 {
		return 1
	} else if exponent%2 == 0 {
		return PowerR(x*x, exponent/2)
	} else {
		return x * PowerR(x*x, (exponent-1)/2)
	}
}

// ** Modular **//

func IsCongruent(a int, b int, modulo int) bool {
	fmt.Printf(" a mod m = b mod m => %v = %v ", a%modulo, b%modulo)
	return (a-b)%modulo == 0
}

// a  +(mod m) b = a+b (mod m)
func ModAdd(input []int, modulo int) int {
	modSum := 0
	for _, val := range input {
		modSum = (modSum + val) % modulo
	}
	return modSum
}

// a  *(mod m) b = a*b (mod m)
func ModMulti(input []int, modulo int) int {
	modProduct := 1
	for _, val := range input {
		modProduct = (modProduct * val) % modulo
	}
	return modProduct
}

/**
 *   power(5,25) (mod 4)
 */

func ModExponent(base int, exponent int, modulo int) int {
	result := 1
	powerToCarry := base % modulo
	for exponent != 0 {
		number := exponent & 1
		fmt.Printf("Binary digit : %v", number)
		if number == 1 {
			result = (result * powerToCarry) % modulo
			fmt.Printf("result : %v \n", result)
		}
		powerToCarry = (powerToCarry * powerToCarry) % modulo
		fmt.Printf("Power to carry : %v", powerToCarry)
		exponent = exponent >> 1
	}
	return result
}

/**
 * Modular exponentiation .
 */
func ModuloPower(x int, y int, p int) int {
	res := 1 // Initialize result
	for y > 0 {
		// If y is odd, multiply x with result
		if (y & 1) != 0 {
			res = res * x
		}
		// y must be even now
		y = y >> 1 // y = y/2
		x = x * x  // Change x to x^2
	}
	return res % p
}
