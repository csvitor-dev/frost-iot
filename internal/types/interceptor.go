package types

type Interceptor interface {
	Invoke([]byte) ([]byte, error)
}
