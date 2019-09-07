package core

// Skill represents a single skill for a player
type Skill interface {
	GetDescription() string
	GetModifier(*Player) AttackModifier
}

// AttackModifier represents a chainable attack skill
type AttackModifier func(*Attack) *Attack

// Pipe chains multiple skills (attack modifiers)
func pipe(modifiers []AttackModifier) AttackModifier {
	return func(attack *Attack) *Attack {
		for _, modifier := range modifiers {
			attack = modifier(attack)
		}
		return attack
	}
}

func pipeSkills(p *Player, skills []Skill) AttackModifier {
	attackModifiers := []AttackModifier{}
	for _, skill := range skills {
		attackModifiers = append(attackModifiers, skill.GetModifier(p))
	}

	return pipe(attackModifiers)
}
