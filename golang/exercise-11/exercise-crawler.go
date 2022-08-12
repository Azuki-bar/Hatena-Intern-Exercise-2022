package main

import (
	"fmt"
	"sync"
	"time"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type cache struct {
	c  map[string]payload
	mu sync.Mutex
}

func newCache() *cache {
	return &cache{c: make(map[string]payload)}
}
func (c *cache) isExist(url string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, ok := c.c[url]
	return ok
}
func (c *cache) add(url string, p payload) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.c[url] = p
}

type payload struct {
	url  string
	body string
	urls []string
	err  error
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	c := newCache()
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		rawCrawl(url, depth, fetcher, wg, c)
	}()
	wg.Wait()
}

func rawCrawl(url string, depth int, fetcher Fetcher, wg *sync.WaitGroup, c *cache) {
	if depth <= 0 {
		return
	}
	if c.isExist(url) {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	p := payload{url: url, body: body, urls: urls, err: err}
	c.add(url, p)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			rawCrawl(u, depth-1, fetcher, wg, c)
		}(u)
	}
}
func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	time.Sleep(1 * time.Second)
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
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
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
