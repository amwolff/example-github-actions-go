package demo

import (
	"math"
	"testing"
)

func TestNextInteger(t *testing.T) {
	tests := []struct {
		name    string
		i       int
		want    int
		wantErr bool
	}{
		{
			name:    "next is 1",
			i:       0,
			want:    1,
			wantErr: false,
		},
		{
			name:    "next is 2",
			i:       1,
			want:    2,
			wantErr: false,
		},
		{
			name:    "next is 3",
			i:       2,
			want:    3,
			wantErr: false,
		},
		{
			name:    "next is 11",
			i:       10,
			want:    11,
			wantErr: false,
		},
		{
			name:    "next is 101",
			i:       100,
			want:    101,
			wantErr: false,
		},
		{
			name:    "next is 1001",
			i:       1000,
			want:    1001,
			wantErr: false,
		},
		{
			name:    "next is integer overflow",
			i:       math.MaxInt64,
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NextInteger(tt.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("NextInteger() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NextInteger() got = %v, want %v", got, tt.want)
			}
		})
	}
}
