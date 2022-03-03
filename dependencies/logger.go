package dependencies

type Logger interface {
	Error(error)
	Warning(string)
	Info(string)
}
