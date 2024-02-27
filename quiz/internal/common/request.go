package common

import (
	"fmt"
	"net/http"
	"strings"
	"unicode/utf8"
)

func GetPageFromRequest(r *http.Request) int {
	p := r.Form.Get("page")
	if p == "" {
		return 1
	}
	page := StringToInt(p)
	if page < 1 {
		page = 1
	}
	return page
}

func GetSearchQueryFromRequest(r *http.Request) string {
	sq := r.Form.Get("search_query")
	if utf8.RuneCountInString(sq) < 2 {
		return ""
	}
	sq = strings.Trim(sq, "\"")
	sq = NormalizeString(sq)
	return sq
}

func DebugRequest(r *http.Request) string {
	var request []string
	// Add the request string
	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)
	// Add the host
	request = append(request, fmt.Sprintf("Host: %v", r.Host))
	// Loop through headers
	// for name, headers := range r.Header {
	// 	name = strings.ToLower(name)
	// 	for _, h := range headers {
	// 		request = append(request, fmt.Sprintf("%v: %v", name, h))
	// 	}
	// }

	// If this is a POST, add post data
	if r.Method == "POST" {
		r.ParseForm()
		request = append(request, "\n")
		request = append(request, r.Form.Encode())
	}
	// Return the request as a string
	return strings.Join(request, "\n")
}
