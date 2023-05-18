package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func fetchURL(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func scrapeWebsite(url string, resultChan chan<- string) {
	content, err := fetchURL(url)
	if err != nil {
		fmt.Printf("Error fetching %s: %s\n ", url, err.Error())
		resultChan <- ""
		return
	}

	resultChan <- content
}

//concurrent process
func conc(urls []string) (time.Duration, error) {

	startTime := time.Now()

	resultChan := make(chan string)

	for _, url := range urls {
		go scrapeWebsite(url, resultChan)
	}

	for range urls {
		// content := <-resultChan
		// fmt.Printf("Fetched Content: %s\n", content)
	}

	endTime := time.Now()                 // Record end time
	elapsedTime := endTime.Sub(startTime) // Calculate elapsed time
	return elapsedTime, nil               // Display execution time
}

//no concurrency implemented, sequential
func seq(urls []string) (time.Duration, error) {

	startTime := time.Now()

	for _, url := range urls {
		_, err := fetchURL(url) //content

		if err != nil {
			fmt.Printf("Error fetching %s: %s\n ", url, err.Error())
			continue
		}

		// fmt.Printf("Fetched Content: %s\n", content)
	}
	endTime := time.Now()                 // Record end time
	elapsedTime := endTime.Sub(startTime) // Calculate elapsed time
	return elapsedTime, nil               // Display execution time
}

func main() {

	urls := []string{
		"https://gift-py.netlify.app/2022/05/09/face_mask_detection/",
		"https://gift-py.netlify.app/2022/05/18/heart-disease-model/",
		"https://gift-py.netlify.app/2022/05/18/multilabel-classification/",
	}

	timeDurationForConcurrent, err := conc(urls)

	if err != nil {
		fmt.Printf(err.Error())
	}

	timeDurationForSequential, err := seq(urls)

	if err != nil {
		fmt.Printf(err.Error())
	}

	fmt.Printf("Execution time for Sequential Process: %s\n", timeDurationForSequential)
	fmt.Printf("Execution time Concurrent Process: %s\n", timeDurationForConcurrent)

	speedUp := float64(timeDurationForSequential) / float64(timeDurationForConcurrent)
	percentageImprovement := float64(timeDurationForSequential) / float64(timeDurationForConcurrent) * 100

	fmt.Printf("Speed-up: %.0f x\n", speedUp)
	fmt.Printf("Percentage Improvement: %.0f %%\n", percentageImprovement)

}
