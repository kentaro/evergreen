package evergreen

import (
	"reflect"
	"regexp"
	"testing"
)

func TestNewUuid(t *testing.T) {
	actual, _ := newUuid()
	expected := regexp.MustCompile("[a-z0-9]+-[a-z0-9]+-[a-z0-9]+-[a-z0-9]+")

	if !expected.MatchString(actual) {
		t.Errorf("UUID generation error")
	}
}

func TestNewRequest(t *testing.T) {
	req := newRequest()

	actual := reflect.TypeOf(req).String()
	expected := "*http.Request"

	if actual != expected {
		t.Errorf("Request generation error: %s", actual)
	}
}
