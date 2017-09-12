package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

type xkcdComic struct {
	Num        int    `json:"num"`
	SafeTitle  string `json:"safe_title"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Transcript string `json:"transcript"`
}

func (c xkcdComic) String() string {
	return fmt.Sprintf(`xkcd #%d
Title: %s
Image: %s
Transcript:
%s
Alt text: %s
`, c.Num, c.SafeTitle, c.Img, c.Transcript, c.Alt)
}

var (
	num   = flag.Int("n", 0, "xkcd comic number")
	query = flag.String("q", "", "search terms")
)

func getComic(num int) (*xkcdComic, error) {
	var comic xkcdComic

	f, err := os.Open(fmt.Sprintf("xkcd-%d.json", num))
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &comic)
	if err != nil {
		return nil, err
	}
	return &comic, nil
}

func main() {
	flag.Parse()
	if *num != 0 {
		comic, err := getComic(*num)
		if err != nil {
			fmt.Println(*comic)
		}
	} else if *query != "" {
		for i := 1; ; i++ {
			// 404 is a joke and is missing. Skip it.
			if i == 404 {
				i++
			}
			comic, err := getComic(i)
			if err != nil {
				break
			}
			if match, _ := regexp.MatchString(*query, comic.Transcript); match {
				fmt.Println(comic)
			}
		}
	} else {
		flag.Usage()
	}
}
