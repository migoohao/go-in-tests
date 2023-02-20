package hello

import (
	"testing"
)

type HelloStruct struct {
	TestName     string
	GreetingName string
	Language     string
	Expected     string
}

func TestHello(t *testing.T) {
	cases := []HelloStruct{
		{"saying hello to people in English", "Chris", "English", "Hello, Chris"},
		{"saying hello to people in French", "Chris", "French", "Bonjour, Chris"},
		{"saying hello to people in Spanish", "Chris", "Spanish", "Hola, Chris"},
		{"saying hello to people in Chinese", "李雷", "Chinese", "你好，李雷"},
		{"say hello world when an empty string is supplied", "", "", "Hello, World"},
	}
	for _, v := range cases {
		t.Run(v.TestName, func(t *testing.T) {
			got := Hello(v.GreetingName, v.Language)
			want := v.Expected
			assertEquals(got, want, t)
		})
	}
}

func assertEquals(got string, want string, t *testing.T) {
	t.Helper()
	if got != want {
		t.Errorf("got '%q' want '%q'", got, want)
	}
}
