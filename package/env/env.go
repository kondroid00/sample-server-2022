package env

type EnvFunc interface {
	Getenv(key string) string
}

type EnvFuncImpl func(string) string

func (f EnvFuncImpl) Getenv(key string) string {
	return f(key)
}
