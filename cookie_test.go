package evergreen

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCookieGet(t *testing.T) {
	{
		identifier := &Cookie{key:"test"}

		req := newRequest()
		req.AddCookie(&http.Cookie{Name: "test", Value: "test"})

		actual, err := identifier.Get(req)
		expected := "test"

		if actual != expected {
			t.Errorf(err.Error())
		}

		if err != nil {
			t.Errorf(err.Error())
		}
	}

	{
		identifier := &Cookie{key:"invalid key"}

		req := newRequest()
		req.AddCookie(&http.Cookie{Name: "test", Value: "test"})

		actual, err := identifier.Get(req)
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
	identifier := &Cookie{key:"test"}

	writer := httptest.NewRecorder()
	identifier.Set(writer, "test")

	actual := writer.Header().Get("Set-Cookie")
	expected := (&http.Cookie{
		Name: "test",
		Value: "test",
		RawExpires: "Tue, 31 Dec 2030 23:30:45 GMT",
		MaxAge: 630720000,
	}).String()

	if actual != expected {
		t.Errorf("Cookie not set correctly")
	}
}
