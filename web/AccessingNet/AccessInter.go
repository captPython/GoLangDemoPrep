package main

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

type siteMapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles    []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

/*
type siteMapIndex struct {
	Locations []Location `xml:"sitemap"`
}

type Location struct {
	Loc string `xml:"loc"`
}

func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}
*/

func main() {
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")

	bytes, _ := ioutil.ReadAll(resp.Body)
	/*string_body := string(bytes)
	fmt.Println(resp)
	fmt.Println(string_body)*/
	resp.Body.Close()

	var s siteMapIndex
	var n News
	xml.Unmarshal(bytes, &s)

	for _, Location := range s.Locations {
		//fmt.Printf("\n%s", Location)
		resp, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)
	}
	//fmt.Println(s.Locations)
}
