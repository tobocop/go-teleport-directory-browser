package user

type Store interface {
	GetUserPassword(username string) (bool, string)
}

type StaticUserStore struct {
	Users map[string]string
}

func (s *StaticUserStore) GetUserPassword(username string) (bool, string)  {
	val, ok := s.Users[username]
	return ok, val
}
