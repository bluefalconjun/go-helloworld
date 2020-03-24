package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

/*
defer的用法

在某个函数调用前加上 defer 前缀会在 包含它的函数结束时 调用它。
有时你需要清理资源，例如关闭一个文件，在我们的案例中是关闭一个服务器，使它不再监听一个端口。
你想让它在函数结束时执行（关闭服务器），但要把它放在你创建服务器语句附近，以便函数内后面的代码仍可以使用这个服务器

httptest的用法

一种方便地创建测试服务器的方法，这样你就可以进行可靠且可控的测试。
使用和 net/http 相同的接口作为「真实的」服务器会和真实环境保持一致，并且只需更少的学习
*/

func TestRacer(t *testing.T) {

	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		server := makeDelayedServer(25 * time.Millisecond)

		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
