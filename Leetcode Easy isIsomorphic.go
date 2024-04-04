package kata

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	} else if s == t {
		return true
	}

	mapped := make(map[byte]byte)
	used := make(map[byte]bool)
	i := 0
	for i < len(s) && i < len(t) {
		if mapp, exist := mapped[s[i]]; exist {
			if mapp != t[i] {
				return false
			}
		} else {
			if used[t[i]] {
				return false
			}
			mapped[s[i]] = t[i]
			used[t[i]] = true
		}
		i++
	}
	return true
}
