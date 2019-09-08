package core

import (
	"fmt"
	"log"
	"strings"
)

// LogsCommentator pretty prints duel actions to console
type LogsCommentator struct {
}

// Start comments the starting of the duel
func (lc *LogsCommentator) Start() {
	log.Println("Welcome everyone to a new duel :)")
}

func (lc *LogsCommentator) getPlayerPresentation(p *Player) string {
	offensiveSkills := []string{}
	defensiveSkills := []string{}

	for _, skill := range p.OffensiveSkills {
		offensiveSkills = append(offensiveSkills, skill.GetDescription())
	}
	for _, skill := range p.DefensiveSkills {
		defensiveSkills = append(defensiveSkills, skill.GetDescription())
	}

	playerPresentation := fmt.Sprintf(`
		%s has the following stats:
		Health: %.2f
		Strength: %.2f
		Defence: %.2f
		Speed: %.2f
		Luck: %.2f%%
	`, p.Name,
		p.Health,
		p.Strength,
		p.Defence,
		p.Speed,
		p.Luck*100,
	)

	if len(offensiveSkills) > 0 || len(defensiveSkills) > 0 {
		playerPresentation += fmt.Sprintf("He also has the following skills\n")
	}

	if len(offensiveSkills) > 0 {
		playerPresentation = playerPresentation + fmt.Sprintf("Offensive: %s\n", strings.Join(offensiveSkills, ", "))
	}

	if len(defensiveSkills) > 0 {
		playerPresentation = playerPresentation + fmt.Sprintf("Defensive: %s\n", strings.Join(defensiveSkills, ", "))
	}

	return playerPresentation
}

// PresentPlayers presents the given players in order and also
// tells which one of them will hit first
func (lc *LogsCommentator) PresentPlayers(first, second *Player) {
	log.Printf("Our duelists are %s and %s\n", first.Name, second.Name)
	log.Printf("Let's loock a bit at those fighters' stats and skills\n")
	log.Printf("%s\n", lc.getPlayerPresentation(first))
	log.Printf("%s\n", lc.getPlayerPresentation(second))
	log.Printf("%s will hit first on each round\n", first.Name)
}

// PresentRound is used for presenting each round
func (lc *LogsCommentator) PresentRound(round int) {
	log.Printf("\n\nRound %d! Start!", round)
}

// PresentAttack presents an attack between 2 opponents
func (lc *LogsCommentator) PresentAttack(attack *Attack, attacker, defender *Player) {
	hits := "hits"
	if len(attack.Hits) == 1 {
		hits = "hit"
	}

	log.Printf("%s attacks %s. The attack contained %d %s\n", attacker.Name, defender.Name, len(attack.Hits), hits)

	for i, hit := range attack.Hits {
		log.Printf("Hit %d with %.2f potential damage on %s\n", i+1, hit.PotentialDamage, defender.Name)
		usedOffensiveSkills := strings.Join(hit.UsedOffensiveSkills, ", ")
		usedDefensiveSkills := strings.Join(hit.UsedDefensiveSkills, ", ")
		if usedOffensiveSkills != "" {
			log.Printf("%s used the following offensive skills for this hit: %s\n", attacker.Name, usedOffensiveSkills)
		}
		if usedDefensiveSkills != "" {
			log.Printf("%s used as defensive skills for this hit: %s\n", defender.Name, usedDefensiveSkills)
		}
	}

	if len(attack.UsedOffensiveSkills) > 0 {
		log.Printf("%s used the following offensive skills on this attack: %s\n", attacker.Name, strings.Join(attack.UsedOffensiveSkills, ", "))
	}

	if len(attack.UsedDefensiveSkills) > 0 {
		log.Printf("%s used the following defensive skills on this attack: %s\n", defender.Name, strings.Join(attack.UsedDefensiveSkills, ", "))
	}

	log.Printf("%s has %.2f remaining health\n\n", defender.Name, defender.Health)
}

// EndDuelKnockout ends the duel in case a knockout was registered
func (lc *LogsCommentator) EndDuelKnockout(round int, winner, loser *Player) {
	log.Printf("Knockout in round %d!! %s is dead! Congratulations to %s for winning the duel before the last round\n", round, loser.Name, winner.Name)
}

// EndDuelTie ends the duel after all rounds depleted
func (lc *LogsCommentator) EndDuelTie(round int, player1, player2 *Player) {
	log.Print("The duel finished with a tie! Congratulations to both players for this battle\n")
	log.Printf("After those %d rounds, %s remains with %.2f health while %s has %.2f health remaining\n",
		round, player1.Name, player1.Health, player2.Name, player2.Health)
}
