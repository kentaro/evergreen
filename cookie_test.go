package evergreen

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCookieGet(t *testing.T) {
	{
		storage := &Cookie{key:"test"}

		req := newRequest()
		req.AddCookie(&http.Cookie{Name: "test", Value: "test"})

		actual, err := storage.Get(req)
		expected := "test"

		if actual != expected {
			t.Errorf(err.Error())
		}

		if err != nil {
			t.Errorf(err.Error())
		}
	}

	{
		storage := &Cookie{key:"invalid key"}

		req := newRequest()
		req.AddCookie(&http.Cookie{Name: "test", Value: "test"})

		actual, err := storage.Get(req)
		expected := ""

		if actual != expected {
			t.Errorf("Invalid return value")
		}

		if err.Error() != "Identity not found in cookie" {
			t.Errorf("Error message is not set correctly")
		}
	}
}

func TestCookieSet(t *testing.T) {
	storage := &Cookie{key:"test"}

	expires, _ := time.Parse("Mon, 2 Jan 2006 15:04:05 MST", "Tue, 31 Dec 2030 23:30:45 GMT")
	writer := httptest.NewRecorder()
	storage.Set(writer, "test")

	actual := writer.Header().Get("Set-Cookie")
	expected := (&http.Cookie{
		Name: "test",
		Value: "test",
		Expires: expires,
		MaxAge: 630720000,
	}).String()

	if actual != expected {
		t.Errorf("Cookie not set correctly")
	}
}
