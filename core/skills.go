package core

import (
	"fmt"
	"math/rand"
)

// CriticalStrike represents the critical strike skill
type CriticalStrike struct {
	DoubleStrikeChance float64
	TripleStrikeChance float64
}

// GetDescription returns the long description of the skill
func (cs *CriticalStrike) GetDescription() string {
	return fmt.Sprintf(`Critical Strike(%f.2%% chance for 2x; %f.2%% chance for 3x)`, cs.DoubleStrikeChance*100, cs.TripleStrikeChance*100)
}

// GetBattleDescription returns the short (in battle) description of the triggered skill
func (cs *CriticalStrike) GetBattleDescription(multiplier int) string {
	return fmt.Sprintf(`CriticalStrike(%dx)`, multiplier)
}

// GetModifier returns the skill in a chainable form
func (cs *CriticalStrike) GetModifier(player *Player) AttackModifier {
	modifier := func(attack *Attack) *Attack {
		if rand.Float64() <= cs.DoubleStrikeChance {
			multipler := 2
			hits := append(attack.Hits, Hit{PotentialDamage: player.Strength})

			if rand.Float64() <= cs.TripleStrikeChance {
				multipler = 3
				hits = append(hits, Hit{PotentialDamage: player.Strength})
			}

			attack.UsedOffensiveSkills = append(attack.UsedOffensiveSkills, cs.GetBattleDescription(multipler))
			attack.Hits = hits

		}
		return attack
	}
	return modifier
}

// Resilience is a defensive skill
// it has a given chance to block a percentage of the potential damage
// for every hit within an attack
type Resilience struct {
	Chance          float64
	DamageReduction float64
}

// GetDescription returns the long description of this skill
func (r *Resilience) GetDescription() string {
	return fmt.Sprintf(`Resilience (%f.2%% chance to block %f.2%% damage)`, r.Chance, r.DamageReduction)
}

// GetBattleDescription returns the short (battle) description of the skill
func (r *Resilience) GetBattleDescription() string {
	return fmt.Sprintf(`Resilience(blocked %f.2%% damage)`, r.DamageReduction)
}

// GetModifier converts the skill to a (chainable) attack modifier
func (r *Resilience) GetModifier(*Player) AttackModifier {
	usedLastTurn := false
	modifier := func(attack *Attack) *Attack {
		if usedLastTurn {
			usedLastTurn = false
			return attack
		}

		if rand.Float64() <= r.Chance {
			usedLastTurn = true
			attack.UsedDefensiveSkills = append(attack.UsedDefensiveSkills, r.GetBattleDescription())
			for i := 0; i < len(attack.Hits); i++ {
				hit := &attack.Hits[i]
				hit.PotentialDamage = hit.PotentialDamage - r.DamageReduction*hit.PotentialDamage
			}
		}

		return attack
	}
	return modifier
}

// Luck is a defensive skill
// it has a chance to evade a hit
type Luck struct {
	Chance float64
}

// GetDescription returns the skill's long description
func (l *Luck) GetDescription() string {
	return fmt.Sprintf(`Luck (%f.2%% chance to evade hits)`, l.Chance)
}

// GetBattleDescription returns the short description of the skill
func (l *Luck) GetBattleDescription() string {
	return fmt.Sprintf(`Got Lucky (you missed)`)
}

// GetModifier returns the attack modifier
func (l *Luck) GetModifier(*Player) AttackModifier {
	modifier := func(attack *Attack) *Attack {
		for i := 0; i < len(attack.Hits); i++ {
			hit := &attack.Hits[i]
			if rand.Float64() < l.Chance {
				hit.PotentialDamage = 0
				hit.UsedDefensiveSkills = append(hit.UsedDefensiveSkills, l.GetBattleDescription())
			}
		}

		return attack
	}

	return modifier
}
