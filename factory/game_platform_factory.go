package main

import "fmt"

func createGame(name, genre string, platform Platform) (IGame, error) {
	if platform == Pc {
		return createPcGame(name, genre), nil
	} else if platform == Console {
		return createConsoleGame(name, genre), nil
	} else if platform == Mobile {
		return createMobileGame(name, genre), nil
	}
	return nil, fmt.Errorf("platform not available")
}
