/**
 * @author [Gopi Karmakar]
 * @email [gopi.karmakar@monstar-lab.com]
 * @create date 2018-01-26 03:00:53
 * @modify date 2018-01-26 03:00:53
 * @desc [description]
 */
// A simple test

package main

import (
	"testing"
	"webapp/app/main"
)

func TestHello(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Hello Albert",
			args: args{name: "Albert"},
			want: "Hello Albert",
		},
		{
			name: "Test Hello Nicolai",
			args: args{name: "Nicolai"},
			want: "Hello Nicolai",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := main.Hello(tt.args.name); got != tt.want {
				t.Errorf("Hello() = %v, want %v", got, tt.want)
			}
		})
	}
}
