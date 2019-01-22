package concurrency_pattern

import (
	"net/http"
	"fmt"
)

// deal with the error during the goroutine runningg
// should treat error as important as common response

type httpResp struct {
	Response *http.Response
	Error error
}

func httpRequest(done <-chan interface{}, urls []string) <-chan httpResp {
	results := make(chan httpResp)
	go func() {
		defer close(results)

		for _, url := range urls {
			resp, err := http.Get(url)
			result := httpResp{Response:resp, Error:err}
			select {
			case <-done:
				return
			case results <- result:
			}
		}
	}()

	return results
}

func ErrorHandlerDemo() {
	done := make(chan interface{})
	defer close(done)

	urls := []string{"http://www.baidu.com", "http://badcase1", "http://badcase2", "http://www.sina.com"}
	results := httpRequest(done, urls)

	errCnt := 0
	for result := range results {
		if result.Error != nil {
			fmt.Printf("http request error %s\n", result.Error.Error())
			errCnt++
			if errCnt > 1 {
				fmt.Println("too many errors")
				break
			}
			continue
		}
		fmt.Printf("http response status %v\n", result.Response.Status)
	}
}
