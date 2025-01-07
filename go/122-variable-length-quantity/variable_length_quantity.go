package variablelengthquantity

import (
	"errors"
)

func EncodeVarint(input []uint32) []byte {
	var output []byte

	for _, x := range input {
		var tmp []byte
		tmp = append(tmp, byte(x%128))
		x /= 128
		for x != 0 {
			tmp = append([]byte{128 + byte(x%128)}, tmp...)
			x /= 128
		}
		output = append(output, tmp...)
	}

	if len(output) == 0 {
		output = []byte{0x0}
	}

	return output
}

func DecodeVarint(input []byte) ([]uint32, error) {
	if input[len(input)-1] > 0x7f {
		return []uint32{0x0}, errors.New("incomplete sequence")
	}

	var output []uint32
	var b10 uint32
	for i := 0; i < len(input); i++ {
		b10 *= uint32(128)
		b10 += uint32(input[i] % 128)
		if input[i] < 128 {
			output = append(output, b10)
			b10 = 0
		}
	}

	return output, nil
}
