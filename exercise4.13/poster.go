package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var queryUrl = "http://www.omdbapi.com/?apikey=%s&t=%s"

type QueryResponse struct {
	Poster   string
	Response string
	Error    string
}

func searchMovie(query, apikey string) (string, error) {
	query = url.PathEscape(query)

	resp, err := http.Get(fmt.Sprintf(queryUrl, apikey, query))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return "", err
	}
	_ = resp

	movie := &QueryResponse{}
	r, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(r, movie)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return "", err
	}
	if movie.Response != "True" {
		return "", fmt.Errorf(movie.Error)
	}
	return movie.Poster, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: poster QUERY")
		os.Exit(1)
	}

	url, err := searchMovie(strings.Join(os.Args[1:], " "), "whatacheapass")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println(url)
}
