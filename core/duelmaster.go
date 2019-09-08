package core

import "time"

// Commentator represents an entity that can log / render / animate
// every event within a duel
type Commentator interface {
	Start()
	PresentPlayers(first, second *Player)
	PresentRound(int)
	PresentAttack(attack *Attack, attacker, defender *Player)
	EndDuelKnockout(round int, winner, loser *Player)
	EndDuelTie(round int, player1, player2 *Player)
}

// DuelMaster contains logic for the duel
type DuelMaster struct {
	Rounds      int
	RoundsDelay time.Duration
	AttackDelay time.Duration

	PlayerOne *Player
	PlayerTwo *Player
}

func (dm *DuelMaster) getPlayersInOrder() (*Player, *Player) {
	if dm.PlayerOne.Speed > dm.PlayerTwo.Speed {
		return dm.PlayerOne, dm.PlayerTwo
	}

	if dm.PlayerOne.Speed < dm.PlayerTwo.Speed {
		return dm.PlayerTwo, dm.PlayerOne
	}

	// equal speed

	if dm.PlayerOne.Luck < dm.PlayerTwo.Luck {
		return dm.PlayerTwo, dm.PlayerOne
	}

	return dm.PlayerOne, dm.PlayerTwo
}

// StartDuel contains the logic for the duel between 2 combatants
func (dm *DuelMaster) StartDuel(commentator Commentator) {

	player1, player2 := dm.getPlayersInOrder()

	commentator.Start()
	commentator.PresentPlayers(player1, player2)

	var round int
	var knockout bool
	for i := 1; i <= dm.Rounds; i++ {
		time.Sleep(dm.RoundsDelay)

		round = i
		commentator.PresentRound(round)
		attack := player1.GenerateAttack()
		player2.DefendAttack(attack)

		commentator.PresentAttack(attack, player1, player2)

		if player2.IsDead() {
			commentator.EndDuelKnockout(round, player1, player2)
			knockout = true
			break
		}

		time.Sleep(dm.AttackDelay)

		attack = player2.GenerateAttack()
		player1.DefendAttack(attack)

		commentator.PresentAttack(attack, player2, player1)
		if player1.IsDead() {
			commentator.EndDuelKnockout(round, player2, player1)
			knockout = true
			break
		}
	}

	if !knockout {
		commentator.EndDuelTie(round, player1, player2)
	}
}
