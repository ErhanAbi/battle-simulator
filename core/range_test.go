package core

import (
	"math/rand"
	"testing"
	"time"
)

func TestRange(t *testing.T) {
	type args struct {
		min float64
		max float64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "creates a random value between given range (10-20)",
			args: args{
				min: 10,
				max: 20,
			},
		},
		{
			name: "creates a random value between (0.1-0.3)",
			args: args{
				min: 0.1,
				max: 0.3,
			},
		},
		{
			name: "creates a random value between (100-200)",
			args: args{
				min: 100,
				max: 200,
			},
		},
	}

	rand.Seed(time.Now().UnixNano())

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Range(tt.args.min, tt.args.max); got <= tt.args.min || got > tt.args.max {
				t.Errorf("Range() = %v, want between %v - %v", got, tt.args.min, tt.args.max)
			}
		})
	}
}
