package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

var XkcdUrl = "https://xkcd.com/%d/info.0.json"

func main() {
	for i := 1; ; i++ {
		resp, err := http.Get(fmt.Sprintf(XkcdUrl, i))
		fmt.Printf("Downloading xkcd #%d\n", i)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			break
		}
		if resp.StatusCode != 200 {
			fmt.Printf("#%d isn't out yet! quitting\n", i)
			break
		}
		f, err := os.OpenFile(fmt.Sprintf("xkcd-%d.json", i), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			break
		}
		_, err = io.Copy(f, resp.Body)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			break
		}
	}
}
