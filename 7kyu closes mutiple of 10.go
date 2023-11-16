// Given a number return the closest number to it that is divisible by 10.

// Example input:

// 22
// 25
// 37

// Expected output:

// 20
// 30
// 40

package kata

func ClosestMultipleOf10(n uint32) uint32 {
	if n%10 < 5 {
		return n - (n % 10)
	} else {
		return (n + 10) - (n % 10)
	}
}
