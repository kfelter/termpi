package tags

import "strings"

func GetTagV(tags []string, k string) string {
	for _, t := range tags {
		ss := strings.SplitN(t, ":", 2)
		if len(ss) == 2 && ss[0] == k {
			return ss[1]
		}
	}
	return ""
}

func ReadOnly(tags []string) bool {
	if GetTagV(tags, "read_only") == "true" {
		return true
	}
	return false
}

func NoDestroy(tags []string) bool {
	if GetTagV(tags, "no_destroy") == "true" {
		return true
	}
	return false
}
