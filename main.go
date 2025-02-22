package main

import (
    "io"
    "log"
    "net/http"
    "os"
    "fmt"
)

func httpGet(url string) {

	resp, err := http.Get(url)

    if err != nil {
      log.Fatal(err)
    }

    defer resp.Body.Close()

    _, err = io.Copy(os.Stdout, resp.Body)

    if err != nil {
      log.Fatal(err)
    }
}
func httpSPAM1(httpURL string) {
	for {
	httpGet(httpURL)
}
}


func main() {
	var choice int
	fmt.Println("Single GET or Spam (1/2)")
	fmt.Scan(&choice)
	if choice == 1 {
		var geturl string
		fmt.Println("Enter URL:")
		fmt.Scan(&geturl)
		httpGet(geturl)
		defer os.Exit(1)
	} else if choice == 2 {
		var geturl string
		fmt.Println("Enter URL:")
		fmt.Scan(&geturl)
		httpSPAM(geturl)
		defer os.Exit(2)
	} else {
		os.Exit(3)
	}
}
