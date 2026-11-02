package main

func rotateRunes(s string, shift int) string {
	r := []rune(s)
	length := len(r)
	if length == 0 {
		return ""
	}
	turns := shift % length
	if turns < 0 {
		r = append(r[length+turns:], r[:length+turns]...)
	} else if turns > 0 {
		r = append(r[turns:], r[:turns]...)
	}
	res := string(r)
	return res
}
