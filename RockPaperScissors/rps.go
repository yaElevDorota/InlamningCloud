package rps

import (
	"math/rand"
	"time"
)

const (
	ROCK     = 0 // beats scissors. (scissors + 1) % 3 = 0
	PAPER    = 1 // beats rock. (rock + 1) % 3 = 1
	SCISSORS = 2 // beats paper. (paper + 1) % 3 = 2
	//PLAYERWINS = 1
	//COMPUTERWINS = 2
	//DRAW         = 3
)

type Round struct {
	//Winner         int    `json:"winner"`
	Message        string `json:"message"`
	ComputerChoice string `json:"computer_choice"`
	RoundResult    string `json:"round_result"`
}

// creating 3 data structures: slices
var winList = []string{
	"You should buy a lottery ticket.",
	"Great!",
	"Nice!",
}
var loseList = []string{
	"Too bad!",
	"Boohoo",
	"Opps",
}
var drawList = []string{
	"Great minds think alike",
	"Uh oh. Try again.",
	"Nobody wins, but you can try again.",
}

// this func receives playerValue int: the buttons hv da values of 0, 1 & 2 -rsp
// this func returns 1)int:if it was a tie or who won?, 2)str:da choice, 3)str:who won?
func PlayRound(playerValue int) Round {
	rand.Seed(time.Now().UnixNano())
	computerValue := rand.Intn(3) //moved cpu choice to 4 loop, so itz reset each tym thru
	messageValue := rand.Intn(3)
	computerChoice := ""
	roundResult := ""
	/* winner := 0 */
	message := ""

	switch computerValue {
	case ROCK:
		computerChoice = "Computer chose ROCK"
		break
	case PAPER:
		computerChoice = "Computer chose PAPER"
		break
	case SCISSORS:
		computerChoice = "Computer chose SCISSORS"
		break
	default:
	}

	if playerValue == computerValue {
		roundResult = "It's a draw"
		/* winner = DRAW */
		message = drawList[messageValue]
	} else if playerValue == (computerValue+1)%3 {
		roundResult = "Player wins!"
		/* winner = PLAYERWINS */
		message = winList[messageValue]
	} else {
		roundResult = "Computer wins!"
		/* winner = COMPUTERWINS */
		message = loseList[messageValue]
	}

	var result Round
	result.Message = message
	result.ComputerChoice = computerChoice
	result.RoundResult = roundResult
	return result
}