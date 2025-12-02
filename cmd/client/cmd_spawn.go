package main

func cmdSpawn(state *clientState, words []string) error {
	return state.gameState.CommandSpawn(words)
}
