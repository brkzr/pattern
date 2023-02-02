package main

import "fmt"

func main() {
	games := []Game{
		{
			name:     "Angry Birds",
			genre:    "Puzzle",
			platform: Mobile},
		{
			name:     "Age of Empires",
			genre:    "Strategy",
			platform: Pc},
		{
			name:     "Last of Us",
			genre:    "Action",
			platform: Console},
		{
			name:     "Catan",
			genre:    "Strategy",
			platform: Board},
	}

	for _, game := range games {
		game, err := createGame(game.name, game.genre, game.platform)
		if err == nil {
			printDetails(game)
		} else {
			fmt.Println(err.Error() + "\n---")
		}
	}

	// pcGame, _ := createGame("Age Of Empires", "Strategy", Pc)
	// mobileGame, _ := createGame("Angry Birds", "Puzzle", Mobile)
	// consoleGame, _ := createGame("Last of Us", "Action", Console)
	//boardGame, _ := createGame("Chess", "Stratecy", Board)

	// printDetails(pcGame)
	// printDetails(mobileGame)
	// printDetails(consoleGame)
	//printDetails(boardGame)
}

func printDetails(g IGame) {
	fmt.Printf("Game: %s\n", g.getName())
	fmt.Printf("Genre: %s\n", g.getGenre())
	fmt.Printf("Platform: %d\n---\n", g.getPlatform())
}
