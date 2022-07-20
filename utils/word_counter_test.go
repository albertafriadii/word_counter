package utils

import (
	"reflect"
	"testing"
)

func TestWordCount(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{
			name: "normal case",
			args: args{s: "Go is a statically typed. Go is a programming language!"},
			want: map[string]int{"go": 2, "is": 2, "a": 2, "statically": 1, "typed": 1, "programming": 1, "language": 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WordCount(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WordCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
