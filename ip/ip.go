package ip

type IP interface {
	Get() (string, error)
}
