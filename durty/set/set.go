package set

// SetInterface is the interface that all set types must implement
type SetInterface[T any] interface {
	Add(value T)
	Remove(value T)
	Contains(value T) bool
	Size() int
	Values() []T
	Clear()
	MarshalJSON() ([]byte, error)
	UnmarshalJSON(data []byte) error
}
