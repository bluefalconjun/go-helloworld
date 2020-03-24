package concurrency

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

/*
goroutines 是 Go 的基本并发单元，它让我们可以同时检查多个网站。
anonymous functions（匿名函数），我们用它来启动每个检查网站的并发进程。
channels，用来组织和控制不同进程之间的交流，使我们能够避免 race condition（竞争条件） 的问题。
the race detector（竞争探测器） 帮助我们调试并发代码的问题。
*/
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		result := <-resultChannel
		results[result.string] = result.bool
	}

	return results
}
