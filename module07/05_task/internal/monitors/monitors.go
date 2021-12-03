package monitors

type Monitor22 interface {
	Type() string
	Run() error
}
