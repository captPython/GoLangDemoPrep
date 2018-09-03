package main

import (
	"encoding/xml"
	"fmt"
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

type NewsMap struct {
	Keyword  string
	Location string
}

func main() {
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")

	bytes, _ := ioutil.ReadAll(resp.Body)
	/*string_body := string(bytes)
	fmt.Println(resp)
	fmt.Println(string_body)*/
	resp.Body.Close()

	var s siteMapIndex
	var n News
	news_map := make(map[string]NewsMap)

	xml.Unmarshal(bytes, &s)

	for _, Location := range s.Locations {
		//fmt.Printf("\n%s", Location)
		resp, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)

		for idx, _ := range n.Titles {

			news_map[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}
			fmt.Println(idx)
			fmt.Println(n.Titles)
		}
	}

	for idx, data := range news_map {
		fmt.Println("\n\n\n", idx)
		fmt.Println("\n", data.Keyword)
		fmt.Println("\n", data.Location)
	}
}
