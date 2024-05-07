package flyweight

// pieces 缓存的棋子
var pieces = map[string]*ChessPiece{
	"black": {"black"},
	"white": {"white"},
}

// ChessPieceFactory 棋子工厂
type ChessPieceFactory struct {
}

func (factory ChessPieceFactory) Create(color string) *ChessPiece {
	return pieces[color]
}
