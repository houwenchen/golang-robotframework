package db

import (
	"fmt"
	"sync"
)

var Storage *Cache

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	UserInfo Info   `json:"userinfo"`
}

type Info struct {
	Address string `json:"address"`
}

type Cache struct {
	m     sync.RWMutex
	Users map[string]User
}

func init() {
	Storage = NewCache()
}

func NewCache() *Cache {
	return &Cache{
		Users: make(map[string]User),
	}
}

func (c *Cache) Add(user User) error {
	c.m.Lock()
	defer c.m.Unlock()

	if !c.IsExists(user) {
		c.Users[user.Name] = user
		return nil
	}

	return fmt.Errorf("user: %#v already exists", user)
}

func (c *Cache) Delete(user User) error {
	c.m.Lock()
	defer c.m.Unlock()

	if c.IsExists(user) {
		delete(c.Users, user.Name)
		return nil
	}

	return fmt.Errorf("user: %#v doesn't exists", user)
}

func (c *Cache) Update(user User) error {
	c.m.Lock()
	defer c.m.Unlock()

	if c.IsExists(user) {
		c.Users[user.Name] = user
		return nil
	}

	return fmt.Errorf("user: %#v doesn't exists", user)
}

func (c *Cache) GetByName(name string) (User, bool) {
	c.m.RLock()
	defer c.m.RUnlock()

	user, ok := c.Users[name]
	return user, ok
}

func (c *Cache) IsExists(user User) bool {
	c.m.RLock()
	defer c.m.RUnlock()

	_, ok := c.Users[user.Name]
	return ok
}

func NewUser(name, passwd string, info Info) *User {
	return &User{
		Name:     name,
		Password: passwd,
		UserInfo: info,
	}
}

func NewInfo(address string) *Info {
	return &Info{
		Address: address,
	}
}
