package kata

func ValidISBN10(isbn string) bool {
	if len(isbn) != 10 {
		return false
	}

	sum := 0

	for ind, c := range isbn {
		if c == 'X' || c == 'x' {
      if ind != 9 {
        return false
      }

			sum += 10 * (ind + 1)
			continue
		}

    if c < '0' || c > '9' {
      return false
    }
		sum += int(c-'0') * (ind + 1)
	}

	return sum%11 == 0
}
