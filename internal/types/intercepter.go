package types

type Intercepter interface {
	Invoke([]byte) ([]byte, error)
}
