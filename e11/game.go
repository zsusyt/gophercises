package blackjack

type GameState struct {

}

type Move func(GameState) GameState

func Hit(gs GameState) GameState {
	return gs
}

func Stand(gs GameState) GameState {
	return gs
}