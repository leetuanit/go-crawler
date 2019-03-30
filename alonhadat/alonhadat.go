package alonhadat

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"crawler/helper"
)

func crawlCat(catLink string) (bool, string, *regexp.Regexp) {
	resp, err := http.Get(catLink)
	if err != nil {
		fmt.Println("http get error.")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("http read error")
	}
	src := string(body)
	regex, _ := regexp.Compile(`/(?:\w+-)+\d{3,7}.html`)
	match := regex.MatchString(src)
	return match, src, regex
}
func crawlConditional(page int16) bool {
	catLink := "https://alonhadat.com.vn/can-ban-nha.htm"
	if page != 1 {
		catLink = fmt.Sprintf("https://alonhadat.com.vn/can-ban-nha/trang-%d.htm", page)
	}
	match, src, regex := crawlCat(catLink)
	if match {
		array := regex.FindAllString(src, -1)
		array = helper.Unique(array)
		for _, entry := range array {
			fmt.Println(entry)
		}
		page++
		crawlConditional(page)
	}
	fmt.Println(match, page)
	return false
}
func Run(){
	crawlConditional(1)
}
