// AAAAABBBBBCCCCC -> A5,B5,C5
package run_length

type RunLength interface {
	Encode(input []byte) []byte
	Decode(input []byte) []byte
}

type firstRunLength struct{}

func (s *firstRunLength) Encode(input []byte) []byte {
	var result []byte
	var prev byte
	var count byte
	for i, b := range input {
		if i == 0 {
			count++
			prev = b
			continue
		}
		if b == prev {
			count++
		} else {
			result = append(result, prev, count)
			count = 1
		}
		prev = b
	}
	result = append(result, prev, count)

	return result
}

func (s *firstRunLength) Decode(input []byte) []byte {
	var char byte
	var result []byte
	for i, b := range input {
		if i%2 == 0 {
			char = b
		} else {
			for j := 0; byte(j) < b; j++ {
				result = append(result, char)
			}
		}
	}
	return result
}

type secondRunLength struct{}

// AAAAABBBBBCCCCC -> A5,B5,C5

func (s *secondRunLength) Encode(input []byte) []byte {
	var count byte = 1
	var result []byte
	for i := 1; i < len(input); i++ {
		if input[i] == input[i-1] {
			count++
		} else {
			result = append(result, input[i-1], count)
			count = 1
		}
	}
	result = append(result, input[len(input)-1], count)
	return result
}

// A5,B5,C5 -> AAAAABBBBBCCCCC

func (s *secondRunLength) Decode(input []byte) []byte {
	var result []byte
	for i := 0; i < len(input)-1; i = i + 2 {
		char := input[i]
		num := input[i+1]
		for j := 0; byte(j) < num; j++ {
			result = append(result, char)
		}
	}
	return result
}
