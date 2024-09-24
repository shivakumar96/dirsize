package display

func ParseBytesToHumanReadable(size int64) (float32, string) {
	suffix := []string{"B", "K", "M", "G", "T", "P"}
	suffixIndx := 0
	currSize := float32(size)

	for int(currSize) >= 1024 {
		currSize /= float32(1024)
		suffixIndx++
	}
	return currSize, suffix[suffixIndx]

}

// added for future use
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
