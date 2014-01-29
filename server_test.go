package evergreen

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"regexp"
	"strconv"
	"testing"
	"time"
)

func TestServerRun(t *testing.T) {
	port := strconv.Itoa(emptyPort())
	server := &Server{}

	go server.Run(map[string]string{"cookie_key": "test", "port": port})
	time.Sleep(1 * time.Second) // XXX

	res, _ := http.Get("http://127.0.0.1:" + port)
	etagRegexp := regexp.MustCompile("^[a-z0-9]+-[a-z0-9]+-[a-z0-9]+-[a-z0-9]+-[a-z0-9]+$")

	{
		actual := res.StatusCode
		expected := 200

		if actual != expected {
			t.Errorf("Invalid response code: %v", actual)
		}
	}

	{
		var actual string
		for _, cookie := range res.Cookies() {
			if cookie.Name == "test" {
				actual = cookie.Value
				break
			}
		}
		expected := etagRegexp

		if !expected.MatchString(actual) {
			t.Errorf("Cookie error")
		}
	}

	{
		actual := res.Header.Get("ETag")
		expected := etagRegexp

		if !expected.MatchString(actual) {
			t.Errorf("ETag error")
		}
	}

	{
		actual := res.Header.Get("X-Evergreen-Id")
		expected := etagRegexp

		if !expected.MatchString(actual) {
			t.Errorf("Header error")
		}
	}

	{
		actual, _ := ioutil.ReadAll(res.Body)
		expected := (&EmptyGif{}).newEmptyGif()

		if !bytes.Equal(actual, expected) {
			t.Errorf("Response body error")
		}
	}
}

func emptyPort() int {
	for port := 10000; port < 20000; port++ {
		addr := fmt.Sprintf("localhost:%d", port)
		l, err := net.Listen("tcp", addr)
		if err == nil {
			defer l.Close()
			return port
		}
	}
	panic("can't listen empty port")
}
