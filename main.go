package main

import (
	"math/rand"
	"time"

	"github.com/pfzero/battle-simulator/core"
)

func main() {
	t := time.Now()
	rand.Seed(t.UnixNano())

	p1 := core.NewPlayer("Na`arun The Wicked", core.PlayerStats{
		Health:   core.Range(70, 100),
		Strength: core.Range(70, 80),
		Defence:  core.Range(45, 55),
		Speed:    core.Range(40, 50),
		Luck:     core.Range(0.1, 0.3),
	}, core.PlayerSkills{
		OffensiveSkills: []core.Skill{&core.CriticalStrike{DoubleStrikeChance: 0.1, TripleStrikeChance: 0.01}},
		DefensiveSkills: []core.Skill{&core.Resilience{Chance: 0.2, DamageReduction: 0.5}},
	})

	p2 := core.NewPlayer("Peanut", core.PlayerStats{
		Health:   core.Range(60, 90),
		Strength: core.Range(60, 90),
		Defence:  core.Range(40, 60),
		Speed:    core.Range(40, 60),
		Luck:     core.Range(0.25, 0.4),
	}, core.PlayerSkills{
		OffensiveSkills: []core.Skill{},
		DefensiveSkills: []core.Skill{},
	})

	dm := &core.DuelMaster{
		Rounds:      20,
		RoundsDelay: time.Second,
		AttackDelay: 500 * time.Millisecond,

		PlayerOne: p1,
		PlayerTwo: p2,
	}

	dm.StartDuel(&core.LogsCommentator{})
}
