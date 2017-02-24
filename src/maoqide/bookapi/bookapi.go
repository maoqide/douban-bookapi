package bookapi

import (
	"encoding/json"
	"errors"
	//"io/ioutil"
	"net/http"
	"strings"

	"maoqide/entity"
)

const (
	SEARCH_URL = "https://api.douban.com/v2/book/search"
)

//q			查询关键字		q和tag必传其一
//tag		查询的tag		q和tag必传其一
//start		取结果的offset	默认为0
//count		取结果的条数		默认为20，最大为100
func search(q, tag, start, count string) (response entity.Response, err error) {

	if q == "" && tag == "" {

		return entity.Response{}, errors.New("q & tag both null")
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
	//request := SEARCH_URL + "?q=" + q + "&tag=" + tag + "&start=" + start + "&count=" + count

	//print(request)
	resp, _ := client.Get(request)

	response, _ = convertResp(resp)

	//out, _ := json.Marshal(response)
	//print(string(out))
	return
}

func keywordSearch(q string) (response entity.Response, err error) {
	return search(q, "", "", "")
}

func convertResp(resp *http.Response) (response entity.Response, err error) {

	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {

	//}

	//respJson := string(body)

	response := entity.Response{}
	err = json.NewDecoder(resp.Body).Decode(&response)

	return
}

func Test() {
	search("python", "", "", "")
}
