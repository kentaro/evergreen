package evergreen

import (
	"net/http/httptest"
	"testing"
)

func TestETagGet(t *testing.T) {
	storage := &ETag{}

	{
		req := newRequest()
		req.Header.Set("If-None-Match", "test")
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
		req := newRequest()
		actual, err := storage.Get(req)
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
	storage := &ETag{}

	writer := httptest.NewRecorder()
	storage.Set(writer, "test")

	actual := writer.Header().Get("ETag")
	expected := "test"

	if actual != expected {
		t.Errorf("Cookie not set correctly")
	}
}
