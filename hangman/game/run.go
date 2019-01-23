package game

func Run() {
	showLoading()
	category := selectCategory()
	word := randWord(categories[category])
	newGame(word).Start()
}
