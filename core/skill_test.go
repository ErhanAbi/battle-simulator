package core

import (
	"reflect"
	"testing"
)

func Test_pipe(t *testing.T) {
	type args struct {
		modifiers []AttackModifier
	}
	tests := []struct {
		name string
		args args
		want *Attack
	}{
		{
			name: "should chain attack modifiers",
			args: args{
				modifiers: []AttackModifier{
					func(attack *Attack) *Attack {
						attack.Hits = append(attack.Hits, NewHit(30), NewHit(40))
						return attack
					},
					func(attack *Attack) *Attack {
						attack.UsedOffensiveSkills = append(attack.UsedOffensiveSkills, "Critical Strike(2x)")
						return attack
					},
				},
			},
			want: &Attack{
				Hits:                []Hit{NewHit(30), NewHit(40)},
				UsedOffensiveSkills: []string{"Critical Strike(2x)"},
				UsedDefensiveSkills: []string{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pipe(tt.args.modifiers); !reflect.DeepEqual(got(&Attack{
				Hits: []Hit{}, UsedOffensiveSkills: []string{}, UsedDefensiveSkills: []string{},
			}), tt.want) {
				t.Errorf("pipe() = %v, want %v", got, tt.want)
			}
		})
	}
}

type doubleDamageAttack struct {
}

func (dda *doubleDamageAttack) GetDescription() string {
	return `Double Damage Skill. Every hit has double damage`
}

func (dda *doubleDamageAttack) GetModifier(p *Player) AttackModifier {
	return func(attack *Attack) *Attack {
		for i := range attack.Hits {
			hit := &attack.Hits[i]
			hit.PotentialDamage *= 2
		}
		attack.UsedOffensiveSkills = append(attack.UsedOffensiveSkills, "Double Damage")
		return attack
	}
}

func Test_pipeSkills(t *testing.T) {

	type args struct {
		p      *Player
		skills []Skill
	}

	tests := []struct {
		name string
		args args
		want *Attack
	}{
		{
			name: "should pipe the skills to get an AttackModifier",
			args: args{
				p: &Player{
					Name:         "Super-Weirdo Creep",
					PlayerSkills: PlayerSkills{OffensiveSkills: []Skill{&doubleDamageAttack{}}},
					PlayerStats: PlayerStats{
						Health:   10,
						Strength: 10,
						Defence:  10,
						Luck:     0.1,
						Speed:    10,
					},
				},
				skills: []Skill{&doubleDamageAttack{}},
			},
			want: &Attack{
				Hits:                []Hit{Hit{PotentialDamage: 20, UsedOffensiveSkills: []string{}, UsedDefensiveSkills: []string{}}},
				UsedOffensiveSkills: []string{"Double Damage"},
				UsedDefensiveSkills: []string{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pipeSkills(tt.args.p, tt.args.skills); !reflect.DeepEqual(got(NewAttack(tt.args.p.Strength)), tt.want) {
				t.Errorf("pipeSkills() = %v, want %v", got, tt.want)
			}
		})
	}
}
