package evergreen

import (
	"bytes"
	"github.com/lestrrat/go-tcputil"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"testing"
	"time"
)

func TestServerRun(t *testing.T) {
	port, err := tcputil.EmptyPort()

	if err != nil {
		t.Errorf("Failed to get an empty port: %v", err)
	}

	server := &Server{}
	go server.Run(map[string]string{"cookie_key": "test", "port": strconv.Itoa(port)})

	err = tcputil.WaitLocalPort(port, 10 * time.Second)

	if err != nil {
		t.Errorf("Failed to listen to the port %d: %v", err)
	}

	res, _ := http.Get("http://127.0.0.1:" + strconv.Itoa(port))
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
