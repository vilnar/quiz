package common

import (
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
	return sq
}
