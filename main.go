package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/pfzero/battle-simulator/core"
)

func main() {
	t := time.Now()
	rand.Seed(t.UnixNano())

	// p1 := NewPlayer("Orange", PlayerStats{})
	p1 := core.NewPlayer("Orange", core.PlayerStats{
		Health:   core.Range(70, 100),
		Strength: core.Range(70, 80),
		Defence:  core.Range(45, 55),
		Speed:    core.Range(40, 50),
		Luck:     core.Range(0.1, 0.3),
	}, core.PlayerSkills{
		OffensiveSkills: []core.Skill{&core.CriticalStrike{DoubleStrikeChance: 0.6, TripleStrikeChance: 0.8}},
		DefensiveSkills: []core.Skill{&core.Resilience{Chance: 0.6, DamageReduction: 0.5}},
	})

	p2 := core.NewPlayer("Bluji The Weak", core.PlayerStats{
		Health:   core.Range(60, 90),
		Strength: core.Range(60, 90),
		Defence:  core.Range(40, 60),
		Speed:    core.Range(40, 60),
		Luck:     core.Range(0.25, 0.4),
	}, core.PlayerSkills{
		OffensiveSkills: []core.Skill{},
		DefensiveSkills: []core.Skill{},
	})

	for i := 0; i < 20; i++ {
		attack := p2.GenerateAttack()
		p1.DefendAttack(attack)

		fmt.Println(attack)
		fmt.Println(p1)
		fmt.Println(p2)
	}

}
