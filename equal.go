package fuzztestingingo

func Equal(a []byte, b []byte) bool {
	// comment out this function to see failing test
	if len(a) != len(b) {
		return false
	}
 	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
