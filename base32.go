package base32

var encodeMapping = map[byte]byte{
	0:  '0',
	1:  '1',
	2:  '2',
	3:  '3',
	4:  '4',
	5:  '5',
	6:  '6',
	7:  '7',
	8:  '8',
	9:  '9',
	10: 'A',
	11: 'B',
	12: 'C',
	13: 'D',
	14: 'E',
	15: 'F',
	16: 'G',
	17: 'H',
	18: 'J',
	19: 'K',
	20: 'M',
	21: 'N',
	22: 'P',
	23: 'Q',
	24: 'R',
	25: 'S',
	26: 'T',
	27: 'V',
	28: 'W',
	29: 'X',
	30: 'Y',
	31: 'Z',
}

func Encode(source []byte) (text string) {
	runes := []byte{}

	for len(source) > 0 {
		var buf [8]byte
		var result [8]byte

		switch len(source) {
		default:
			buf[7] = (source[4] & 0x1F)
			buf[6] = (source[4] >> 5)
			fallthrough
		case 4:
			buf[6] |= ((source[3] << 3) & 0x1F)
			buf[5] = ((source[3] >> 2) & 0x1F)
			buf[4] = (source[3] >> 7)
			fallthrough
		case 3:
			buf[4] |= ((source[2] << 1) & 0x1F)
			buf[3] = ((source[2] >> 4) & 0x1F)
			fallthrough
		case 2:
			buf[3] |= ((source[1] << 4) & 0x1F)
			buf[2] = ((source[1] >> 1) & 0x1F)
			buf[1] = ((source[1] >> 6) & 0x1F)
			fallthrough
		case 1:
			buf[1] |= ((source[0] << 2) & 0x1F)
			buf[0] = (source[0] >> 3)
		}

		for i := 0; i < len(buf); i += 1 {
			result[i] = encodeMapping[buf[i]&31]
		}

		switch len(source) {
		default:
			runes = append(runes, result[:8]...)
		case 4:
			runes = append(runes, result[:7]...)
		case 3:
			runes = append(runes, result[:5]...)
		case 2:
			runes = append(runes, result[:4]...)
		case 1:
			runes = append(runes, result[:2]...)
		}

		if len(source) < 5 {
			break
		}

		source = source[5:]
	}

	text = string(runes)
	return
}

var decodeMapping = map[byte]byte{
	'0': 0, 'o': 0, 'O': 0,
	'1': 1, 'i': 1, 'I': 1, 'l': 1, 'L': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'a': 10, 'A': 10,
	'b': 11, 'B': 11,
	'c': 12, 'C': 12,
	'd': 13, 'D': 13,
	'e': 14, 'E': 14,
	'f': 15, 'F': 15,
	'g': 16, 'G': 16,
	'h': 17, 'H': 17,
	'j': 18, 'J': 18,
	'k': 19, 'K': 19,
	'm': 20, 'M': 20,
	'n': 21, 'N': 21,
	'p': 22, 'P': 22,
	'q': 23, 'Q': 23,
	'r': 24, 'R': 24,
	's': 25, 'S': 25,
	't': 26, 'T': 26,
	'v': 27, 'V': 27,
	'w': 28, 'W': 28,
	'x': 29, 'X': 29,
	'y': 30, 'Y': 30,
	'z': 31, 'Z': 31,
}

func Decode(text string) (bytes []byte) {
	source := []byte(text)

	for len(source) > 0 {
		var buf [8]byte
		var result [5]byte
		var size int

		size = len(source)

		if size > 8 {
			size = 8
		}

		for i := 0; i < size; i += 1 {
			buf[i] = decodeMapping[source[i]]
		}

		switch len(source) {
		default:
			result[4] = (buf[6] << 5) | (buf[7])
			fallthrough
		case 7:
			result[3] = (buf[4] << 7) | (buf[5] << 2) | (buf[6] >> 3)
			fallthrough
		case 5:
			result[2] = (buf[3] << 4) | (buf[4] >> 1)
			fallthrough
		case 4:
			result[1] = (buf[1] << 6) | (buf[2] << 1) | (buf[3] >> 4)
			fallthrough
		case 2:
			result[0] = (buf[0] << 3) | (buf[1] >> 2)
		}

		switch len(source) {
		default:
			bytes = append(bytes, result[:5]...)
		case 7:
			bytes = append(bytes, result[:4]...)
		case 5:
			bytes = append(bytes, result[:3]...)
		case 4:
			bytes = append(bytes, result[:2]...)
		case 2:
			bytes = append(bytes, result[:1]...)
		}

		if len(source) < 8 {
			break
		}

		source = source[8:]
	}

	return bytes
}
