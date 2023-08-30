package data

// Game represents a game that is being played. It contains a list of players with their scores, a name and a flag to indicate if the game has ended.
type Game struct {
	Name         string
	Ended        bool
	PlayerScores []*Score
}

// NewGame creates a new game with the given name.
func NewGame(name string) *Game {
	return &Game{
		Name:         name,
		PlayerScores: []*Score{},
	}
}

// AddPlayers adds the given players to the game.
func (g *Game) AddPlayers(p ...*Player) {
	for _, player := range p {
		g.PlayerScores = append(g.PlayerScores, &Score{Player: player})
	}
}

// RemovePlayer removes the given player from the game. It's based on the player's name.
func (g *Game) RemovePlayer(p *Player) {
	for i, score := range g.PlayerScores {
		if score.Player.Name == p.Name {
			g.PlayerScores = append(g.PlayerScores[:i], g.PlayerScores[i+1:]...)
		}
	}
}

// End sets the game's ended flag to true.
func (g *Game) End() {
	g.Ended = true
}

// GetPlayer returns the player with the given name. If no player is found, nil is returned.
func (g *Game) GetPlayer(name string) *Player {
	for _, score := range g.PlayerScores {
		if score.Player.Name == name {
			return score.Player
		}
	}
	return nil
}
