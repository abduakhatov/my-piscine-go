package hackathon

func ActiveBits(n int) uint {
	if n < 2 {
		return n
	}
	return (n % 2) + ActiveBits(n/2)

}
