package files

import (
	"awesomwProject/lib/e"
	. "awesomwProject/storage"
	"context"
	"encoding/gob"
	"fmt"
	"os"
	"path/filepath"
)

const defaultPerm = 0774

var _ IStorage = new(Storage)

type Storage struct {
	basePath string
}

func New(basePath string) *Storage {
	return &Storage{basePath: basePath}
}

func (s Storage) Store(ctx context.Context, u *User) (err error) {
	defer func() { err = e.WrapIfErr("can't save", err) }()

	fPath := filepath.Join(s.basePath, u.Name)
	if err = os.MkdirAll(fPath, defaultPerm); err != nil {
		return err
	}

	fName, err := fileName(u)
	if err != nil {
		return err
	}

	fPath = filepath.Join(fPath, fName)

	file, err := os.Create(fPath)
	if err != nil {
		return err
	}

	defer func() { _ = file.Close() }()

	if err := gob.NewEncoder(file).Encode(u); err != nil {
		return err
	}
	return nil
}

func (s Storage) Remove(ctx context.Context, u *User) error {
	fileName, err := fileName(u)
	if err != nil {
		return e.Wrap("can't remove file", err)
	}

	path := filepath.Join(s.basePath, u.Name, fileName)

	if err := os.Remove(path); err != nil {
		msg := fmt.Sprintf("can't remove file %s", path)

		return e.Wrap(msg, err)
	}
	return nil
}

func fileName(u *User) (string, error) {
	return u.Hash()
}
