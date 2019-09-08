package core

import (
	"math/rand"
	"reflect"
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

func TestNewPlayer(t *testing.T) {
	type args struct {
		name   string
		stats  PlayerStats
		skills PlayerSkills
	}
	tests := []struct {
		name string
		args args
		want *Player
	}{
		{
			name: "creates a new player with the specified skill set and abilities",
			args: args{
				name: "Terminal - The Twisted Warlock",
				stats: PlayerStats{
					Health:   10000,
					Defence:  500,
					Strength: 1000,
					Speed:    1200,
					Luck:     0.3,
				},
				skills: PlayerSkills{
					OffensiveSkills: []Skill{},
					DefensiveSkills: []Skill{},
				},
			},
			want: &Player{
				Name: "Terminal - The Twisted Warlock",
				PlayerSkills: PlayerSkills{
					OffensiveSkills: []Skill{},
					DefensiveSkills: []Skill{},
				},
				PlayerStats: PlayerStats{
					Health:   10000,
					Defence:  500,
					Strength: 1000,
					Speed:    1200,
					Luck:     0.3,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewPlayer(tt.args.name, tt.args.stats, tt.args.skills)
			if got.Name != tt.want.Name {
				t.Errorf("NewPlayer name is not correct; got %s wanted %s", got.Name, tt.want.Name)
			}

			if !reflect.DeepEqual(got.PlayerStats, tt.want.PlayerStats) {
				t.Errorf("NewPlayer() didn't return the same player stats; got %v want %v", got.PlayerStats, tt.want.PlayerStats)
			}

			if !reflect.DeepEqual(got.PlayerSkills, tt.want.PlayerSkills) {
				t.Errorf("NewPlayer() didn't return the same player skills; got %v want %v", got.PlayerSkills, tt.want.PlayerSkills)
			}
		})
	}
}

func TestPlayer_IsDead(t *testing.T) {
	tests := []struct {
		name string
		p    *Player
		want bool
	}{
		{
			name: "should return true if the player's health is smaller or equal to 0",
			p:    &Player{PlayerStats: PlayerStats{Health: 0}},
			want: true,
		},
		{
			name: "should return false if the player's health is smaller or equal to 0",
			p:    &Player{PlayerStats: PlayerStats{Health: 1}},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.IsDead(); got != tt.want {
				t.Errorf("Player.IsDead() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlayer_GenerateAttack(t *testing.T) {
	tests := []struct {
		name string
		p    *Player
		want *Attack
	}{
		{
			name: "should generate a simple attack if the player has no skills",
			p: NewPlayer(
				"Foo Creep",
				PlayerStats{
					Health:   90,
					Strength: 20,
					Defence:  30,
					Speed:    10,
					Luck:     0,
				},
				PlayerSkills{
					OffensiveSkills: []Skill{},
					DefensiveSkills: []Skill{},
				},
			),

			want: NewAttack(20),
		},
		{
			name: "should generate an attack where all skills are applied",
			p: NewPlayer("Foo Hero", PlayerStats{
				Health:   90,
				Strength: 80,
				Defence:  50,
				Speed:    30,
				Luck:     0.2,
			}, PlayerSkills{
				OffensiveSkills: []Skill{&CriticalStrike{DoubleStrikeChance: 1, TripleStrikeChance: 1}},
				DefensiveSkills: []Skill{},
			}),
			want: &Attack{
				Hits:                []Hit{NewHit(80), NewHit(80), NewHit(80)},
				UsedOffensiveSkills: []string{"CriticalStrike(3x)"},
				UsedDefensiveSkills: []string{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.GenerateAttack(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Player.GenerateAttack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlayer_DefendAttack(t *testing.T) {
	type args struct {
		attack *Attack
	}
	tests := []struct {
		name       string
		p          *Player
		args       args
		wantHealth float64
	}{
		{
			name: "defend the attack and updates player's health accordingly",
			p: NewPlayer("Defensive Ronan", PlayerStats{
				Health:   100,
				Strength: 50,
				Defence:  75,
				Speed:    10,
				Luck:     0,
			}, PlayerSkills{}),
			args: args{
				attack: NewAttack(100),
			},
			wantHealth: 75,
		},
		{
			name: "applies skills to incoming attack and updates health accordingly",
			p: NewPlayer("Resilient Barry", PlayerStats{
				Health:   2,
				Strength: 1,
				Defence:  1,
				Speed:    1,
				Luck:     0,
			}, PlayerSkills{
				OffensiveSkills: []Skill{},
				DefensiveSkills: []Skill{&Resilience{Chance: 1, DamageReduction: 0.5}},
			}),
			args: args{
				attack: NewAttack(4),
			},
			wantHealth: 1, // 2 - (4 * 0.5 (damage after skills applied) - 1 (defence)) = 1
		},
		{
			name: "applies luck to incoming attack and upates health accordingly",
			p: NewPlayer("Lucky Jim", PlayerStats{
				Health:   1,
				Strength: 25,
				Defence:  30,
				Speed:    85,
				Luck:     1,
			}, PlayerSkills{
				OffensiveSkills: []Skill{},
				DefensiveSkills: []Skill{},
			}),
			args: args{
				attack: NewAttack(100),
			},
			wantHealth: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.DefendAttack(tt.args.attack)
			if tt.wantHealth != tt.p.Health {
				t.Errorf("Expected player's life to be %.2f after the attack but got %.2f", tt.wantHealth, tt.p.Health)
			}
		})
	}
}
