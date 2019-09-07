package core

// Hit represents the description of a single hit
type Hit struct {
	PotentialDamage     float64
	UsedOffensiveSkills []string
	UsedDefensiveSkills []string
}

// NewHit creates a new hit object
func NewHit(strength float64) Hit {
	return Hit{
		PotentialDamage:     strength,
		UsedOffensiveSkills: []string{},
		UsedDefensiveSkills: []string{},
	}
}

// Attack represents a player's attack
type Attack struct {
	Hits                []Hit
	UsedOffensiveSkills []string
	UsedDefensiveSkills []string
}

// NewAttack creates a new attack object
func NewAttack(strength float64) *Attack {
	return &Attack{
		Hits:                []Hit{NewHit(strength)},
		UsedOffensiveSkills: []string{},
		UsedDefensiveSkills: []string{},
	}
}
