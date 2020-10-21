package main

func Burbuja(s []int64) {
	iteracion := 0
	permutacion := true
	var actual int64
	for permutacion {
		permutacion = false
		iteracion++
		for actual = 0; actual < int64(len(s)-1); actual++ {
			if s[actual] > s[actual+1] {
				permutacion = true
				temp := s[actual]
				s[actual] = s[actual+1]
				s[actual+1] = temp
			}
		}
	}
}

func main() {
}
