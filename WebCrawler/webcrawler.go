package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// SafeCache is a thread-safe cache for fetched urls
type SafeCache struct {
	mu    sync.Mutex
	cache map[string]bool
}

// This creates a new SafeCache instance
func newSafeCache() *SafeCache {
	return &SafeCache{
		cache: make(map[string]bool),
	}
}

// This code adds a url to the cache
func (c *SafeCache) Add(url string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[url] = true
}

// This code checks if a URL is in the cache
func (c *SafeCache) Contains(url string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, ok := c.cache[url]
	return ok
}

// Crawl uses Fetcher to recursively crawl pages starting with url, to a max of depth
func crawl(url string, depth int, fetcher Fetcher, cache *SafeCache, wg *sync.WaitGroup) {
	defer wg.Done()

	if depth <= 0 {
		return
	}

	//Check if url is already fetched
	if cache.Contains(url) {
		return
	}

	body, urls, err := fetcher.Fetch(url)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)

	//Add url to cache
	cache.Add(url)

	var innerWg sync.WaitGroup
	for _, u := range urls {
		innerWg.Add(1)
		go func(u string) {
			defer innerWg.Done()
			crawl(u, depth-1, fetcher, cache, &innerWg)
		}(u)
	}
	innerWg.Wait()
}

func main() {
	cache := newSafeCache()
	var wg sync.WaitGroup
	wg.Add(1)
	crawl("https://golang.org/", 4, fetcher, cache, &wg)
	wg.Wait()
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go programming language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/fmt/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org",
			"https://golang.org/pkg/",
		},
	},
}
