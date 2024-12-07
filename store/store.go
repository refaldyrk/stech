package store

import "aidanwoods.dev/go-paseto"

type Store struct {
	Key paseto.V4SymmetricKey
}

func NewStore() *Store {
	return &Store{
		Key: paseto.NewV4SymmetricKey(),
	}
}

func (s *Store) GetKey() paseto.V4SymmetricKey {
	return s.Key
}
