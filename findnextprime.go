package piscine

func IsPrimeNumb(nb int) bool {
	if nb <= 1 {
		return false
	}
	for i := 2; i <= nb/2; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

func FindNextPrime(nb int) int {
	if nb < 1 {
		return 2
	}
	is_prime := false
	for !is_prime {
		is_prime = IsPrimeNumb(nb)
		if is_prime {
			return nb
		}
		nb += 1
	}
	return 1
}
