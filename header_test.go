package evergreen

import (
	"net/http/httptest"
	"testing"
)

func TestHeaderSet(t *testing.T) {
	header := &Header{}

	writer := httptest.NewRecorder()
	header.Set(writer, "test")

	actual := writer.Header().Get("X-Evergreen-Id")
	expected := "test"

	if actual != expected {
		t.Errorf("X-Evergreen-Id header is not set correctly")
	}
}
