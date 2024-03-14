package greetings

import (
	"regexp"
	"testing"
)

func Test_GreetingsHelloWhenisOk(t *testing.T) {
	name := "David"

	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("David")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("David") = %q, %v, quiere coincidenciapara %#q, nil`, msg, err, want)
	}

}

func Test_GreetingsHelloWhenIsEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, quiere un "", error`, msg, err)
	}

}
