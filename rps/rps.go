package rps

import (
	"math/rand"
	"time"
)

const (
	ROCK     = 0 // beats scissors. (scissors + 1) % 3 = 0
	PAPER    = 1 // beats rock. (rock + 1) % 3 = 1
	SCISSORS = 2 // beats paper. (paper + 1) % 3 = 2
	//***Comment next  three line as it is not required
	// PLAYERWINS   = 1
	// COMPUTERWINS = 2
	// DRAW         = 3
)

type Round struct {
	//***changed winner to message
	//Winner         int    `json:"winner"`
	Message        string `json:"message"`
	ComputerChoice string `json:"computer_choice"`
	RoundResult    string `json:"round_result"`
}

//***created a three slice

var wonMessage = []string{
	"Good job",
	"Nice Work",
	"you won a jackpot",
}

var drawMessage = []string{
	"it a draw",
	"ohh..try again",
	"u can try again",
}
var loseMessage = []string{
	"it has been fade away",
	"Try  harder",
	"All lost",
}

func PlayRound(playerValue int) Round {
	//func PlayRound(playerValue int) (int, string, string) {

	rand.Seed(time.Now().UnixNano())
	computerValue := rand.Intn(3)
	computerChoice := ""
	rountResult := ""
	//***commented the winner as it not been used
	//winner := 0

	switch computerValue {
	case ROCK:
		computerChoice = "Computer choose ROCK"
		break
	case PAPER:
		computerChoice = "Computer chosse PAPER"
		break
	case SCISSORS:
		computerChoice = "Computer choose Scissors"
		break
	default:
	}
	//*** Generate a var which will help to generate a random message betn 1 - 2
	messageInt := rand.Intn(3)
	message := ""

	if playerValue == computerValue {
		rountResult = "It is Draw"
		//winner = DRAW
		//***populate message for Draw message
		message = drawMessage[messageInt]

	} else if playerValue == (computerValue+1)%3 {
		rountResult = "Player Wins"
		//winner = PLAYERWINS
		//***populate message for won message
		message = wonMessage[messageInt]

	} else {
		rountResult = "Computer wins"
		//winner = COMPUTERWINS
		//***populate message for lose message
		message = loseMessage[messageInt]
	}

	//	return winner, computerChoice, result

	var result Round
	//result.Winner = winner
	result.Message = message
	result.ComputerChoice = computerChoice
	result.RoundResult = rountResult

	return result
}
