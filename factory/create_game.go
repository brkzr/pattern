package main

type PcGame struct {
	Game
}

type ConsoleGame struct {
	Game
}

type MobileGame struct {
	Game
}

func createPcGame(name, genre string) IGame {
	return &PcGame{
		Game: Game{
			name:     name,
			genre:    genre,
			platform: Pc,
		},
	}
}

func createConsoleGame(name, genre string) IGame {
	return &PcGame{
		Game: Game{
			name:     name,
			genre:    genre,
			platform: Console,
		},
	}
}

func createMobileGame(name, genre string) IGame {
	return &PcGame{
		Game: Game{
			name:     name,
			genre:    genre,
			platform: Mobile,
		},
	}
}
