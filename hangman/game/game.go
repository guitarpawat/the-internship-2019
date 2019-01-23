package game

import (
	"fmt"
	"strings"
)

type Game struct {
	Word

	display    string
	wrong      []string
	wrongLimit int
	score      int
}

func parseWord(wordStr string) string {
	result := ""
	for _, v := range wordStr {
		if (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') {
			result += "_"
		} else {
			result += string(v)
		}
	}
	return result
}

func formatDisplay(display string) string {
	return strings.Join(strings.Split(display, ""), " ")
}

func newGame(word Word) *Game {
	return &Game{
		Word:       word,
		display:    parseWord(word.Word),
		wrong:      []string{},
		wrongLimit: 8,
		score:      20,
	}
}

func (game *Game) Start() {
	fmt.Println("Hint:", game.Hint)
	for {
		fmt.Println()
		game.showDisplay()

		if game.isWin() {
			fmt.Println("You Win!")
			return
		}

		if game.isLost() {
			fmt.Println("You Lost, the word is:", game.Word.Word)
			return
		}

		var input string
		fmt.Scanln(&input)
		for _, v := range input {
			game.guess(string(v))
		}
	}
}

func (game *Game) guess(char string) {
	changePts := 0
	for i, v := range game.Word.Word {
		str := string(v)
		if strings.ToUpper(str) == strings.ToUpper(char) {
			game.display = game.display[:i] + string(game.Word.Word[i]) + game.display[i+1:]
			changePts += 5
		}
	}

	if changePts == 0 {
		changePts = -5
		game.wrong = append(game.wrong, char)
	}
	game.score += changePts
}

func (game Game) showDisplay() {
	wrongLeft := game.wrongLimit - len(game.wrong)
	displayWord := formatDisplay(game.display)
	displayWord += fmt.Sprintf("\t score:%d, remaining wrong guess %d", game.score, wrongLeft)

	if len(game.wrong) > 0 {
		wrongLetter := strings.Join(game.wrong, ", ")
		displayWord += fmt.Sprintf(", wrong guessed: %s", wrongLetter)
	}

	fmt.Println(displayWord)
}

func (game Game) isWin() bool {
	return !strings.Contains(game.display, "_")
}

func (game Game) isLost() bool {
	return game.wrongLimit-len(game.wrong) <= 0 || game.score < 0
}
