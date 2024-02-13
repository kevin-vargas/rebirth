package alive

type Aliver interface {
	IsAlive(string) (bool, error)
}
