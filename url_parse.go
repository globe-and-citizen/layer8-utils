package utilities

import "strings"

func ParseURLPath(url string) (path string, query string) {
	qmIndex := strings.Index(url, "?")
	if qmIndex == -1 {
		return url, ""
	}
	return url[:qmIndex], url[qmIndex+1:]
}

func ParseQueryParams(query string) map[string]string {
	params := make(map[string]string)
	pairs := strings.Split(query, "&")
	for _, pair := range pairs {
		kv := strings.Split(pair, "=")
		if len(kv) == 2 {
			params[kv[0]] = kv[1]
		}
	}
	return params
}