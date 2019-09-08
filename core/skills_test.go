package core

import (
	"encoding/json"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func Test_Skill_GetModifier(t *testing.T) {
	type inargs struct {
		player        *Player
		attackFactory func() *Attack
		skill         Skill
	}

	type testargs struct {
		name          string
		inargs        inargs
		wantedAttacks []Attack
	}

	tests := []testargs{
		{
			name: "Critical Strike should duplicate the attack",
			inargs: inargs{
				player:        &Player{Name: "Frenzy Creep", PlayerStats: PlayerStats{Strength: 50.00}},
				skill:         &CriticalStrike{DoubleStrikeChance: 1, TripleStrikeChance: 0},
				attackFactory: func() *Attack { return NewAttack(50) },
			},
			wantedAttacks: []Attack{
				Attack{
					Hits:                []Hit{NewHit(50), NewHit(50)},
					UsedOffensiveSkills: []string{"CriticalStrike(2x)"},
					UsedDefensiveSkills: []string{},
				},
				Attack{
					Hits:                []Hit{NewHit(50), NewHit(50)},
					UsedOffensiveSkills: []string{"CriticalStrike(2x)"},
					UsedDefensiveSkills: []string{},
				},
			},
		},
		{
			name: "Critical Strike should add 2 more attacks (3x)",
			inargs: inargs{
				player:        &Player{Name: "Crazy Creep", PlayerStats: PlayerStats{Strength: 50.00}},
				skill:         &CriticalStrike{DoubleStrikeChance: 1, TripleStrikeChance: 1},
				attackFactory: func() *Attack { return NewAttack(50) },
			},
			wantedAttacks: []Attack{
				Attack{
					Hits:                []Hit{NewHit(50), NewHit(50), NewHit(50)},
					UsedOffensiveSkills: []string{"CriticalStrike(3x)"},
					UsedDefensiveSkills: []string{},
				},
				Attack{
					Hits:                []Hit{NewHit(50), NewHit(50), NewHit(50)},
					UsedOffensiveSkills: []string{"CriticalStrike(3x)"},
					UsedDefensiveSkills: []string{},
				},
			},
		},
		{
			name: "Critical Strike should add 2 more attacks (3x)",
			inargs: inargs{
				player:        &Player{Name: "Crazy Creep", PlayerStats: PlayerStats{Strength: 50.00}},
				skill:         &CriticalStrike{DoubleStrikeChance: 1, TripleStrikeChance: 1},
				attackFactory: func() *Attack { return NewAttack(50) },
			},
			wantedAttacks: []Attack{
				Attack{
					Hits:                []Hit{NewHit(50), NewHit(50), NewHit(50)},
					UsedOffensiveSkills: []string{"CriticalStrike(3x)"},
					UsedDefensiveSkills: []string{},
				},
				Attack{
					Hits:                []Hit{NewHit(50), NewHit(50), NewHit(50)},
					UsedOffensiveSkills: []string{"CriticalStrike(3x)"},
					UsedDefensiveSkills: []string{},
				},
			},
		},
		{
			name: "Resilience should block 10% damage every other attack",
			inargs: inargs{
				player:        &Player{Name: "Poor Shielded Creep"},
				skill:         &Resilience{DamageReduction: 0.1, Chance: 1},
				attackFactory: func() *Attack { return NewAttack(50) },
			},
			wantedAttacks: []Attack{
				Attack{
					Hits:                []Hit{NewHit(45)},
					UsedOffensiveSkills: []string{},
					UsedDefensiveSkills: []string{"Resilience(blocked 10.00% damage)"},
				},
				*NewAttack(50),
				Attack{
					Hits:                []Hit{NewHit(45)},
					UsedOffensiveSkills: []string{},
					UsedDefensiveSkills: []string{"Resilience(blocked 10.00% damage)"},
				},
				*NewAttack(50),
			},
		},
		{
			name: "Resilience should block 90% damage every other attack",
			inargs: inargs{
				player:        &Player{Name: "Super Shielded Creep"},
				skill:         &Resilience{DamageReduction: 0.9, Chance: 1},
				attackFactory: func() *Attack { return NewAttack(100) },
			},
			wantedAttacks: []Attack{
				Attack{
					Hits:                []Hit{NewHit(10)},
					UsedOffensiveSkills: []string{},
					UsedDefensiveSkills: []string{"Resilience(blocked 90.00% damage)"},
				},
				*NewAttack(100),
				Attack{
					Hits:                []Hit{NewHit(10)},
					UsedOffensiveSkills: []string{},
					UsedDefensiveSkills: []string{"Resilience(blocked 90.00% damage)"},
				},
				*NewAttack(100),
			},
		},
		{
			name: "Luck should evade all hits when it's at 100%",
			inargs: inargs{
				player: &Player{Name: "Super Lucky Creep"},
				skill:  &Luck{Chance: 1},
				attackFactory: func() *Attack {
					return &Attack{
						Hits:                []Hit{NewHit(100), NewHit(100)},
						UsedDefensiveSkills: []string{},
						UsedOffensiveSkills: []string{},
					}
				},
			},
			wantedAttacks: []Attack{
				Attack{
					Hits: []Hit{
						Hit{PotentialDamage: 0, UsedOffensiveSkills: []string{}, UsedDefensiveSkills: []string{"Got Lucky (you missed)"}},
						Hit{PotentialDamage: 0, UsedOffensiveSkills: []string{}, UsedDefensiveSkills: []string{"Got Lucky (you missed)"}},
					},
					UsedDefensiveSkills: []string{},
					UsedOffensiveSkills: []string{},
				},
			},
		},
		{
			name: "No Luck, No miss",
			inargs: inargs{
				player: &Player{Name: "Creep"},
				skill:  &Luck{Chance: 0},
				attackFactory: func() *Attack {
					return &Attack{
						Hits:                []Hit{NewHit(100), NewHit(100)},
						UsedOffensiveSkills: []string{},
						UsedDefensiveSkills: []string{},
					}
				},
			},
			wantedAttacks: []Attack{
				Attack{
					Hits: []Hit{
						NewHit(100), NewHit(100),
					},
					UsedDefensiveSkills: []string{},
					UsedOffensiveSkills: []string{},
				},
			},
		},
	}

	rand.Seed(time.Now().UnixNano())

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mod := tt.inargs.skill.GetModifier(tt.inargs.player)

			for i := range tt.wantedAttacks {
				wantedAttack := &tt.wantedAttacks[i]
				baseAttack := tt.inargs.attackFactory()
				attack := mod(baseAttack)

				if !reflect.DeepEqual(attack, wantedAttack) {
					attackPretty, _ := json.MarshalIndent(attack, "", "  ")
					wantedAttackPretty, _ := json.MarshalIndent(wantedAttack, "", "  ")
					t.Errorf(`Got: %v
					Wanted: %v`, string(attackPretty), string(wantedAttackPretty))
				}
			}
		})
	}
}
