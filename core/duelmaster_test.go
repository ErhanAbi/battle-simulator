package core

import (
	"testing"
)

func TestDuelMaster_getPlayersInOrder(t *testing.T) {
	tests := []struct {
		name   string
		dm     *DuelMaster
		first  string
		second string
	}{
		{
			name: "it should return the player with highest speed first",
			dm: &DuelMaster{
				PlayerOne: NewPlayer("Some Hero", PlayerStats{Speed: 50}, PlayerSkills{}),
				PlayerTwo: NewPlayer("Some Villain", PlayerStats{Speed: 49}, PlayerSkills{}),
			},
			first:  "Some Hero",
			second: "Some Villain",
		},
		{
			name: "it should return the player with highest luck if they have the same speed",
			dm: &DuelMaster{
				PlayerOne: NewPlayer("Some Hero", PlayerStats{Speed: 50, Luck: 10}, PlayerSkills{}),
				PlayerTwo: NewPlayer("Some Villain", PlayerStats{Speed: 50, Luck: 20}, PlayerSkills{}),
			},
			first:  "Some Villain",
			second: "Some Hero",
		},
		{
			name: "it should return the first passed player as first player if both speed and luck are the same",
			dm: &DuelMaster{
				PlayerOne: NewPlayer("Some Hero", PlayerStats{Speed: 50, Luck: 10}, PlayerSkills{}),
				PlayerTwo: NewPlayer("Some Villain", PlayerStats{Speed: 50, Luck: 10}, PlayerSkills{}),
			},
			first:  "Some Hero",
			second: "Some Villain",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p1, p2 := tt.dm.getPlayersInOrder()
			first, second := p1.Name, p2.Name

			if first != tt.first || second != tt.second {
				t.Errorf("Expected order hero to be: %s, %s but got %s, %s", tt.first, tt.second, first, second)
			}
		})
	}
}

// dummy implementation of the Commenter interface
type dummyCommentator struct {
}

func (dc *dummyCommentator) Start()                                                   {}
func (dc *dummyCommentator) PresentPlayers(first, second *Player)                     {}
func (dc *dummyCommentator) PresentRound(int)                                         {}
func (dc *dummyCommentator) PresentAttack(attack *Attack, attacker, defender *Player) {}
func (dc *dummyCommentator) EndDuelKnockout(int, *Player, *Player)                    {}
func (dc *dummyCommentator) EndDuelTie(int, *Player, *Player)                         {}

func TestDuelMaster_StartDuel(t *testing.T) {
	type args struct {
		commentator Commentator
	}
	tests := []struct {
		name          string
		dm            *DuelMaster
		args          args
		player1Health float64
		player2Health float64
	}{
		{
			name: "It should end if the first player knockouts second",
			dm: &DuelMaster{
				Rounds: 20,
				PlayerOne: NewPlayer("Winner", PlayerStats{
					Health:   100,
					Defence:  0,
					Strength: 100,
					Luck:     0,
					Speed:    100,
				}, PlayerSkills{}),
				PlayerTwo: NewPlayer("Loser", PlayerStats{
					Health:   100,
					Defence:  0,
					Strength: 100,
					Luck:     0,
					Speed:    90,
				}, PlayerSkills{}),
			},
			args:          args{commentator: &dummyCommentator{}},
			player1Health: 100,
			player2Health: 0,
		},
		{
			name: "It should end if the second player knockouts first",
			dm: &DuelMaster{
				Rounds: 20,
				PlayerOne: NewPlayer("Winner", PlayerStats{
					Health:   100,
					Defence:  0,
					Strength: 100,
					Luck:     0,
					Speed:    100,
				}, PlayerSkills{}),
				PlayerTwo: NewPlayer("Loser", PlayerStats{
					Health:   100,
					Defence:  100,
					Strength: 100,
					Luck:     0,
					Speed:    90,
				}, PlayerSkills{}),
			},
			args:          args{commentator: &dummyCommentator{}},
			player1Health: 0,
			player2Health: 100,
		},
		{
			name: "It should end after the number of rounds specified",
			dm: &DuelMaster{
				Rounds: 20,
				PlayerOne: NewPlayer("Winner", PlayerStats{
					Health:   100,
					Defence:  0,
					Strength: 2,
					Luck:     0,
					Speed:    100,
				}, PlayerSkills{}),
				PlayerTwo: NewPlayer("Loser", PlayerStats{
					Health:   100,
					Defence:  2,
					Strength: 2,
					Luck:     0,
					Speed:    90,
				}, PlayerSkills{}),
			},
			args:          args{commentator: &dummyCommentator{}},
			player1Health: 60,
			player2Health: 100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.dm.StartDuel(tt.args.commentator)
			if tt.dm.PlayerOne.Health != tt.player1Health {
				t.Errorf("Expected first player's health to be %.2f but got %.2f", tt.player1Health, tt.dm.PlayerOne.Health)
			}
			if tt.dm.PlayerTwo.Health != tt.player2Health {
				t.Errorf("Expected second player's health to be %.2f but got %.2f", tt.player2Health, tt.dm.PlayerTwo.Health)
			}
		})
	}
}
