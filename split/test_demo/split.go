package test_demo

import "strings"

// split.go

func Split(s, sep string) (ret []string) {
	ret = make([]string, 0, strings.Count(s, sep)+1)
	idx := strings.Index(s, sep)
	for idx > -1 {
		ret = append(ret, s[:idx])
		s = s[idx+len(sep):]
		idx = strings.Index(s, sep)
	}
	ret = append(ret, s)
	return
}
