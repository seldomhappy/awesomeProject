package storage

import (
	"awesomwProject/lib/e"
	"context"
	"crypto/sha1"
	"fmt"
	"io"
	"strconv"
)

// Checks interface.
//var _ storage.MessageStorage = new(MessagePostgresStorage)

type IStorage interface {
	Store(ctx context.Context, u *User) error
	Remove(ctx context.Context, u *User) error
}

type User struct {
	Name string
	Age  int
}

func (u User) Hash() (string, error) {
	h := sha1.New()
	if _, err := io.WriteString(h, u.Name); err != nil {
		return "", e.Wrap("can't calculate hash", err)
	}
	if _, err := io.WriteString(h, strconv.Itoa(u.Age)); err != nil {
		return "", e.Wrap("can't calculate hash", err)
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
