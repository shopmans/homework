package pkg

func IndexOf(data []byte, sub []byte) int {
	if len(data) <= 0 || len(sub) <= 0 {
		return -1
	}
	subLength := len(sub)

	for i := 0; i < len(data); i++ {
		if data[i] != sub[0] {
			continue
		}

		match := true
		for k := 1; k < subLength; k++ {
			if data[i+k] != sub[k] {
				match = false
				break
			}
		}

		if match {
			return i
		}
	}

	return -1
}
