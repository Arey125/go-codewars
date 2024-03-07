package kata

func generate(generated map[string]bool, str []byte, l int) map[string]bool {
	if l == 1 {
    generated[string(str)] = true
		return generated
	}

	generated = generate(generated, str, l-1)

	for ind := range str[:l] {
		if str[ind] == str[l-1] {
			continue
		}
		str[ind], str[l-1] = str[l-1], str[ind]
		generated = generate(generated, str, l-1)
		str[ind], str[l-1] = str[l-1], str[ind]
	}

	return generated
}

func Permutations(s string) []string {
	str := []byte(s)

  resmap := generate(make(map[string]bool, 0), str, len(str))

  res := make([]string, 0, len(resmap))
  for k := range resmap {
    res = append(res, k)
  }

  return res
}
