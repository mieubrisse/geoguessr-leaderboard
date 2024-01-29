package server

import "os"

const (
	configEnvVar = "KURTOSIS_SERVER_CONFIG"
)

func NewServerFromEnv() {
	os.LookupEnv()

}
