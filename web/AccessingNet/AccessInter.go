package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Location struct {
	Loc string `xml:"loc"`
}

type siteMapIndex struct {
	Locations []Location `xml:"sitemap"`
}

func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}

func main() {
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")

	bytes, _ := ioutil.ReadAll(resp.Body)
	/*string_body := string(bytes)
	fmt.Println(resp)
	fmt.Println(string_body)*/
	resp.Body.Close()

	var s siteMapIndex
	xml.Unmarshal(bytes, &s)
	fmt.Println(s.Locations)
}
