package evergreen

import (
	"bytes"
	"net/http/httptest"
	"testing"
)

func TestEmptyGifSet(t *testing.T) {
	emptyGif := &EmptyGif{}

	writer := httptest.NewRecorder()
	emptyGif.Set(writer)

	{
		actual := writer.Body.Bytes()
		expected := emptyGif.newEmptyGif()

		if !bytes.Equal(actual, expected) {
			t.Errorf("Empty GIF is not set correctly")
		}
	}

	{
		actual := writer.Header().Get("Content-Type")
		expected := "image/gif"

		if actual != expected {
			t.Errorf("Content-Type is not set correctly")
		}
	}

	{
		actual := writer.Header().Get("Content-Length")
		expected := "43"

		if actual != expected {
			t.Errorf("Content-Length is not set correctly")
		}
	}
}
