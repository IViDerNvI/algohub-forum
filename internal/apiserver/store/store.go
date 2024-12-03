package store

var client Factory

type Factory interface {
	Users() UserStore
	Posts() PostStore
	Likes() LikeStore
	Tokens() TokenStore
	Problems() ProbStore
	Close() error
}

func NewFactory() Factory {
	return client
}

func SetFactory(f Factory) {
	client = f
}
