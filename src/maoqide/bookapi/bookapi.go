package bookapi

import (
	"encoding/json"
	"errors"
	"fmt"
	//"io/ioutil"
	"net/http"
	"strings"

	"maoqide/entity"
)

const (
	SEARCH_URL      = "https://api.douban.com/v2/book/search"
	ISBN_SEARCH_URL = "https://api.douban.com/v2/book/isbn/"
)

//q			查询关键字		q和tag必传其一
//tag		查询的tag		q和tag必传其一
//start		取结果的offset	默认为0
//count		取结果的条数		默认为20，最大为100
func search(q, tag, start, count string) (books []entity.Book, err error) {

	if q == "" && tag == "" {

		return nil, errors.New("q & tag both null")
	}
	client := &http.Client{}

	//generate request
	request := SEARCH_URL + "?"
	if q != "" {
		request += "q=" + q + "&"
	}
	if tag != "" {
		request += "tag=" + tag + "&"
	}
	if start != "" {
		request += "start=" + start + "&"
	}
	if count != "" {
		request += "count=" + count
	}

	request = strings.TrimSuffix(request, "&")

	//http get
	resp, _ := client.Get(request)

	response := entity.Response{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		fmt.Errorf(err.Error())
		return
	}
	books = response.Books
	return
}

func keywordSearch(q string) (books []entity.Book, err error) {
	return search(q, "", "", "")
}

func isbnSearch(isbn string) (book entity.Book, err error) {

	client := &http.Client{}
	request := ISBN_SEARCH_URL + isbn

	//http get
	resp, _ := client.Get(request)

	response := entity.Book{}
	err = json.NewDecoder(resp.Body).Decode(&response)

	if err != nil {
		fmt.Errorf(err.Error())
		return
	}
	book = response

	return

}

func Test() {
	books, _ := keywordSearch("python")
	print(books[2].Title)
	b, _ := isbnSearch("7505715666")
	print(b.Title)
}
