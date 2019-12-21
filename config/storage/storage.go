package storage

type Storage string

func (s Storage) String() string {
	return string(s)
}

func FromString(value string) Storage {
	return Storage(value)
}
