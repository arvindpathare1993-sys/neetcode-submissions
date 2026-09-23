func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}
	a := strings.Split(s, "")
	b := strings.Split(t, "")
	sort.Strings(a)
	sort.Strings(b)
	if strings.Join(a, "") == strings.Join(b, "") {
		return true
	}
	return false
}