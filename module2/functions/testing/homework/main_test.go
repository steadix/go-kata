package main

import (
	"fmt"
	"testing"
	"unicode"
)

func TestGreet(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "english name",
			args: args{name: "John"},
			want: "Hello John, you welcome!",
		},
		{
			name: "russian name",
			args: args{name: "Боря"},
			want: "Привет Боря, добро пожаловать!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Greet(tt.args.name); got != tt.want {
				t.Errorf("Greet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Greet(name string) string {

	for i := 0; i < len(name); i++ {
		if name[i] > unicode.MaxASCII {
			return fmt.Sprintf("Привет %s, добро пожаловать!", name)
		} else {
			return fmt.Sprintf("Hello %s, you welcome!", name)
		}
	}
	return ""
}
