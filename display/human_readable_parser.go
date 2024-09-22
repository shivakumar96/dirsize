package display

func ParseBytesToHumanReadable(size float32) (float32, string) {
	suffix := []string{"B", "K", "M", "G", "T", "P"}
	suffixIndx := 0
	currSize := size

	for int64(currSize) >= 1024 {
		currSize = float32(int(currSize) >> 10) // divide by 1024, while retainging the value after decimal point
		suffixIndx++
	}
	return currSize, suffix[suffixIndx]

}

func ParseHumanReadableToBytes(size float32, humanRedableVal string) int64 {
	var mulval int64 = 1

	// multiply by 1024 to convert to bytes
	switch humanRedableVal {
	case "P":
		mulval <<= 10
		fallthrough
	case "T":
		mulval <<= 10
		fallthrough
	case "G":
		mulval <<= 10
		fallthrough
	case "M":
		mulval <<= 10
		fallthrough
	case "K":
		mulval <<= 10
		fallthrough
	default:
		mulval <<= 0
	}

	return int64(size * float32(mulval))

}
