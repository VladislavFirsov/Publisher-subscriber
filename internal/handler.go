package internal

type Handler interface {
	Home()
	AllClients()
	ClientById()
}
