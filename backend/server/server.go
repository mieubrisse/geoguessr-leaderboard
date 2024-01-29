package server

const (
	defaultDbfilepath
)

type ServerOptions struct {
	Players    []string
	DbFilepath KurtosisFilepath
}

func GetDefaultServerOptions() ServerOptions {
	return ServerOptions{
		Players:    []string{},
		DbFilepath: "",
	}
}

func ValidateServerOptions(options ServerOptions) error {
	if dbFile
}

type Server struct {
	ServerOptions ServerOptions
}
