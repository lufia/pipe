package pipe_test

import (
	"fmt"
	"net/url"

	"github.com/lufia/pipe"
)

func WithPath(s string) func(u *url.URL) *url.URL {
	return func(u *url.URL) *url.URL {
		u.Path = s
		return u
	}
}

func WithParam(k, v string) func(u *url.URL) *url.URL {
	return func(u *url.URL) *url.URL {
		q := u.Query()
		q.Set(k, v)
		u.RawQuery = q.Encode()
		return u
	}
}

func ExampleValue_url() {
	u, _ := pipe.TryFrom(pipe.Value("https://example.com"), url.Parse).
		Chain(WithPath("/query")).
		Chain(WithParam("key", "value")).
		Eval()
	fmt.Println(u.String())
	// Output: https://example.com/query?key=value
}
