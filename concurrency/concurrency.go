package concurrency

type WebsiteChecker func(string) bool

type Websites []string

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	// Create goroutine for each url
	// and write result to channel
	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	// For all provided urls, read from the
	// result channel when ready and set the
	// result to whether the url is ok or not
	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
