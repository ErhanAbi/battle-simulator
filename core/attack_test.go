package core

import (
	"reflect"
	"testing"
)

func TestNewHit(t *testing.T) {
	type args struct {
		strength float64
	}
	tests := []struct {
		name string
		args args
		want Hit
	}{
		{
			name: "creates a new hit with potential damage equal to given strength",
			args: args{
				strength: 80,
			},
			want: Hit{
				PotentialDamage:     80,
				UsedOffensiveSkills: []string{},
				UsedDefensiveSkills: []string{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHit(tt.args.strength); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAttack(t *testing.T) {
	type args struct {
		strength float64
	}
	tests := []struct {
		name string
		args args
		want *Attack
	}{
		{
			name: "creates a new attack with a hit damage equal to provided strength",
			args: args{
				strength: 80,
			},
			want: &Attack{
				Hits:                []Hit{Hit{PotentialDamage: 80, UsedOffensiveSkills: []string{}, UsedDefensiveSkills: []string{}}},
				UsedOffensiveSkills: []string{},
				UsedDefensiveSkills: []string{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAttack(tt.args.strength); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAttack() = %v, want %v", got, tt.want)
			}
		})
	}
}
