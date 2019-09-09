package core

import (
	"math"
)

// PlayerStats represents the stats of a player
type PlayerStats struct {
	Health   float64
	Strength float64
	Defence  float64
	Speed    float64
	Luck     float64
}

// PlayerSkills represents the player's skills
type PlayerSkills struct {
	OffensiveSkills []Skill
	DefensiveSkills []Skill
}

// Player represents a duel fighter
type Player struct {
	Name string

	PlayerStats
	PlayerSkills

	offensiveAttackModifier AttackModifier
	defensiveAttackModifier AttackModifier
}

// NewPlayer creates a new player based on the given stats and skills
func NewPlayer(name string, stats PlayerStats, skills PlayerSkills) *Player {
	p := &Player{
		Name:         name,
		PlayerStats:  stats,
		PlayerSkills: skills,
	}

	p.offensiveAttackModifier = pipeSkills(p, p.OffensiveSkills)
	p.defensiveAttackModifier = pipeSkills(p, append([]Skill{&Luck{Chance: p.Luck}}, p.DefensiveSkills...))

	return p
}

// IsDead checks wether the player has died
func (p *Player) IsDead() bool {
	return p.Health <= 0
}

// GenerateAttack generates player's attack
func (p *Player) GenerateAttack() *Attack {
	baseAttack := NewAttack(p.Strength)
	return p.offensiveAttackModifier(baseAttack)
}

// DefendAttack represents the logic for defending oponent player's
// attack
func (p *Player) DefendAttack(attack *Attack) {
	if p.IsDead() {
		return
	}

	attackAfterDefense := p.defensiveAttackModifier(attack)

	for _, hit := range attackAfterDefense.Hits {
		if p.IsDead() {
			break
		}

		damage := math.Max(0, hit.PotentialDamage-p.Defence)
		p.Health = math.Max(0, p.Health-damage)
	}
}
