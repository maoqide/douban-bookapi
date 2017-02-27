package bookapi

import (
	"encoding/json"
	"errors"
	"fmt"
	//"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"maoqide/entity"
)

const (
	SEARCH_URL        = "https://api.douban.com/v2/book/search"
	ISBN_SEARCH_URL   = "https://api.douban.com/v2/book/isbn/"
	SERIES_SEARCH_URL = "https://api.douban.com/v2/book/series/"
)

var (
	client = &http.Client{}
)

//q			查询关键字		q和tag必传其一
//tag		查询的tag		q和tag必传其一
//start		取结果的offset	默认为0
//count		取结果的条数		默认为20，最大为100
func search(q, tag, start, count string) (response entity.Response, err error) {

	if q == "" && tag == "" {
		err = errors.New("q & tag both null")
		return
	}
	//client := &http.Client{}

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

	response = entity.Response{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		fmt.Errorf(err.Error())
		return
	}
	return
}

func keywordSearch(q string) (response entity.Response, err error) {
	return search(q, "", "", "")
}

func isbnSearch(isbn string) (book entity.Book, err error) {

	//client := &http.Client{}
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

func seriesSearch(id string) (response entity.Response, err error) {

	request := SERIES_SEARCH_URL + id + "/books"

	//http get
	resp, _ := client.Get(request)

	response = entity.Response{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		fmt.Errorf(err.Error())
		return
	}
	return

}

func searchAll(q, tag string) (books []entity.Book, err error) {
	start, COUNT := 0, 100
	books = make([]entity.Book, 0, 100)
	for {
		response, err1 := search(q, tag, strconv.Itoa(start), strconv.Itoa(COUNT))
		if err1 != nil {
			err = err1
			//err = errors.New("error when execute searchAll")
			return
		}
		bs := response.Books
		books = append(books, bs...)
		start = start + COUNT
		if len(bs) < 100 {
			break
		}
	}

	return
}

func Test() {
	//resp, _ := keywordSearch("python")
	//println(resp.Total)
	//b, _ := isbnSearch("7505715666")
	//println(b.Title)
	//resp2, _ := seriesSearch("2")
	//println(resp2.Books[0].Series.Title)
	bs, _ := searchAll("python", "python")
	//println(len(bs))
	out, _ := json.Marshal(bs)
	println(string(out))

}
