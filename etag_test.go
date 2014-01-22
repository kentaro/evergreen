package evergreen

import (
	"net/http/httptest"
	"testing"
)

func TestETagGet(t *testing.T) {
	identifier := &ETag{}

	{
		req := newRequest()
		req.Header.Set("If-None-Match", "test")
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
		req := newRequest()
		actual, err := identifier.Get(req)
		expected := ""

		if actual != expected {
			t.Errorf("Invalid return value")
		}

		if err.Error() != "Identity not found in ETag" {
			t.Errorf("Error message is not set correctly")
		}
	}
}

func TestETagSet(t *testing.T) {
	identifier := &ETag{}

	writer := httptest.NewRecorder()
	identifier.Set(writer, "test")

	actual := writer.Header().Get("ETag")
	expected := "test"

	if actual != expected {
		t.Errorf("Cookie not set correctly")
	}
}
