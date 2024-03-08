package pagination

import (
	"html"
	"html/template"
	"math"
	"net/url"
	"strconv"
)

const ON_EACH_SIDE int = 3

type Paginator struct {
	PerPage     int
	TotalAmount int
	CurrentPage int
	TotalPage   int
	BaseUrl     string

	// render parts
	FirstPart  []template.HTML
	MiddlePart []template.HTML
	LastPart   []template.HTML
}

func NewPaginator(totalAmount, perPage, currentPage int, baseUrl string) Paginator {
	if currentPage < 1 {
		currentPage = 1
	}

	totalPage := int(math.Ceil(float64(totalAmount) / float64(perPage)))
	if currentPage > totalPage {
		currentPage = totalPage
	}

	return Paginator{
		PerPage:     perPage,
		TotalAmount: totalAmount,
		CurrentPage: currentPage,
		TotalPage:   totalPage,
		BaseUrl:     baseUrl,
	}
}

func (p Paginator) HasPages() bool {
	return p.TotalPage > 1
}

func (p Paginator) Generate() Paginator {
	if !p.HasPages() {
		return p
	}
	if p.TotalPage < (ON_EACH_SIDE*2 + 6) {
		p.FirstPart = p.getUrlRange(1, p.TotalPage)
	} else {
		window := ON_EACH_SIDE * 2
		lastPage := p.TotalPage
		if p.CurrentPage < window {
			p.FirstPart = p.getUrlRange(1, window+2)
			p.LastPart = p.getUrlRange(lastPage-1, lastPage)
		} else if p.CurrentPage > (lastPage - window) {
			p.FirstPart = p.getUrlRange(1, 2)
			p.LastPart = p.getUrlRange(lastPage-(window+2), lastPage)
		} else {
			p.FirstPart = p.getUrlRange(1, 2)
			p.MiddlePart = p.getUrlRange(p.CurrentPage-ON_EACH_SIDE, p.CurrentPage+ON_EACH_SIDE)
			p.LastPart = p.getUrlRange(lastPage-1, lastPage)
		}
	}

	return p
}

func (p Paginator) getUrlRange(start, end int) []template.HTML {
	var ret []template.HTML
	for i := start; i <= end; i++ {
		ret = append(ret, p.getUrl(i, strconv.Itoa(i)))
	}
	return ret
}

func (p Paginator) getUrl(page int, text string) template.HTML {
	strPage := strconv.Itoa(page)
	if p.CurrentPage == page {
		return p.GetActivePageWrapper(strPage)
	} else {
		baseUrl, _ := url.Parse(p.BaseUrl)
		params := baseUrl.Query()
		delete(params, "page")
		strParam := ""
		for k, v := range params {
			strParam = strParam + "&" + k + "=" + v[0] // TODO
		}

		href := baseUrl.String() + "?page=" + strPage + strParam
		return p.GetAvailablePageWrapper(href, text)
	}
}

func (p Paginator) GetActivePageWrapper(text string) template.HTML {
	return safeHtml(`<li class="page-item active"><span class="page-link">` + text + `</span></li>`)
}

func (p Paginator) GetDisabledPageWrapper(text string) template.HTML {
	return safeHtml(`<li class="page-item disabled wtf-wrapper"><span class="page-link">` + text + `</span></li>`)
}

func (p Paginator) GetAvailablePageWrapper(href, page string) template.HTML {
	return safeHtml(`<li class="page-item"><a class="page-link" href="` + href + `">` + page + `</a></li>`)
}

func (p Paginator) GetDots() template.HTML {
	return safeHtml(`<li class="page-item disabled"><span class="page-link">...</span></li>`)
}

func (p Paginator) GetPreviousButton() template.HTML {
	text := "<"
	if p.CurrentPage <= 1 {
		return p.GetDisabledPageWrapper(text)
	}

	return p.getUrl(p.CurrentPage-1, text)
}

func (p Paginator) GetNextButton() template.HTML {
	text := ">"
	if p.CurrentPage == p.TotalPage {
		return p.GetDisabledPageWrapper(text)
	}
	return p.getUrl(p.CurrentPage+1, text)
}

func safeHtml(s string) template.HTML {
	return template.HTML(html.UnescapeString(s))
}
