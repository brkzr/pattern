package main

type IGame interface {
	setName(n string)
	getName() string
	setGenre(g string)
	getGenre() string
	setPlatform(p Platform)
	getPlatform() Platform
}

type Platform int

const (
	Pc Platform = iota
	Console
	Mobile
	Board
)

type Game struct {
	name     string
	genre    string
	platform Platform
}

func (g *Game) setName(name string) {
	g.name = name
}

func (g *Game) getName() string {
	return g.name
}

func (g *Game) setGenre(genre string) {
	g.genre = genre
}

func (g *Game) getGenre() string {
	return g.genre
}

func (g *Game) getPlatform() Platform {
	return g.platform
}

func (g *Game) setPlatform(p Platform) {
	g.platform = p
}
