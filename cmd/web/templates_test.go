package main

import (
	"testing"
	"time"
)

func TestHumanDate(t *testing.T) {

	tests := []struct {
		name	string
		tm		time.Time
		want	string
	}{
		{
			name:	"UTC",
			tm:		time.Date(2024, 3, 17, 10, 15, 0, 0, time.UTC),
			want:	"17. March 2024",
		},
		{
			name:	"empty",
			tm:		time.Time{},
			want:	"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hd := humanDate(tt.tm)

			if hd != tt.want {
				t.Errorf("got: %q; want: %q", hd, tt.want)
			}
		})
	}
}
