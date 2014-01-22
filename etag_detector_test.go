package evergreen

import (
	"testing"
)

func TestETagDetector(t *testing.T) {
	detector := &ETagDetector{}

	{
		req := newRequest()
		req.Header.Set("If-None-Match", "test")
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
		actual, err := detector.Detect("invalid key", req)
		expected := ""

		if actual != expected {
			t.Errorf("Invalid return value")
		}

		if err.Error() != "Identity not found in ETag" {
			t.Errorf("Error message is not set correctly")
		}
	}
}
