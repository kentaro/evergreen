package evergreen

import (
	"net/http"
	"testing"
)

func TestCookie(t *testing.T) {
	identifier := &Cookie{}

	{
		req := newRequest()
		req.AddCookie(&http.Cookie{Name: "test", Value: "test"})
		actual, err := identifier.Get("test", req)
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
		req.AddCookie(&http.Cookie{Name: "test", Value: "test"})
		actual, err := identifier.Get("invalid key", req)
		expected := ""

		if actual != expected {
			t.Errorf("Invalid return value")
		}

		if err.Error() != "Identity not found in cookie" {
			t.Errorf("Error message is not set correctly")
		}
	}
}
