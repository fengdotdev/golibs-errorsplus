package ontry

type OnTry interface {
	Try(fn func())
}
