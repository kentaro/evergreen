package evergreen

import (
	"net/http"
	"testing"
)

func TestCookieDetector(t *testing.T) {
	detector := &CookieDetector{}

	{
		req := newRequest()
		req.AddCookie(&http.Cookie{Name:"test", Value:"test"})
		actual, err := detector.Detect("test", req)
		expected := "test"

		if actual != expected {
			t.Errorf(err.Error())
		}

		if err != nil {
			t.Errorf(err.Error())
		}
	}

	{
		req := newRequest()
		req.AddCookie(&http.Cookie{Name:"test", Value:"test"})
		actual, err := detector.Detect("invalid key", req)
		expected := ""

		if actual != expected {
			t.Errorf("Invalid return value")
		}

		if err.Error() != "Identity not found in cookie" {
			t.Errorf("Error message is not set correctly")
		}
	}
}
