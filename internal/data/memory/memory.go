package memory

import (
	"github.com/shoriwe/pivot/internal/data"
	"github.com/shoriwe/pivot/internal/data/objects"
	"sync"
)

type Memory struct {
	users      map[string]*objects.User
	usersMutex *sync.Mutex
}

func (m *Memory) Register(user *objects.User) (succeed bool, err error) {
	m.usersMutex.Lock()
	defer m.usersMutex.Unlock()
	_, found := m.users[user.Email]
	if found {
		return false, nil
	}
	m.users[user.Email] = user
	return true, nil
}

func (m *Memory) Login(email, password string) (user *objects.User, succeed bool, err error) {
	m.usersMutex.Lock()
	defer m.usersMutex.Unlock()
	user, succeed = m.users[email]
	if succeed {
		if data.CheckPasswords(password, user.Password) {
			return user, true, nil
		}
	}
	return nil, false, nil
}

func NewMemory() data.Database {
	return &Memory{
		users: map[string]*objects.User{
			"admin@upb.motors.co": {
				Email:    "admin@upb.motors.co",
				Password: "5a38afb1a18d408e6cd367f9db91e2ab9bce834cdad3da24183cc174956c20ce35dd39c2bd36aae907111ae3d6ada353f7697a5f1a8fc567aae9e4ca41a9d19d", // admin
			},
		},
		usersMutex: new(sync.Mutex),
	}
}
