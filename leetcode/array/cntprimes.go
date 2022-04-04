package array

func countPrimes(n int) int {
	isPrime := make([]bool, n)
	for i := 2; i < n; i++ {
		isPrime[i] = true
	}

	for i := 2; i*i < n; i++ {
		if isPrime[i] == true {
			for j := i * i; j < n; j += i {
				isPrime[j] = false
			}
		}
	}

	res := 0
	for i := 2; i < n; i++ {
		if isPrime[i] == true {
			res += 1
		}
	}

	return res
}
