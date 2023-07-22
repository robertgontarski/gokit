package gokit

import "github.com/joho/godotenv"

var (
	Env map[string]string
)

func initEnv() {
	loadEnv()
}

func loadEnv() {
	if err := godotenv.Load(".env.local"); err != nil {
		if err := godotenv.Load(); err != nil {
			panic(err)
		}
	}

	localEnv, _ := godotenv.Read(".env.local")

	defaultEnv, err := godotenv.Read(".env")
	if err != nil {
		panic(err)
	}

	if 0 != len(defaultEnv) {
		Env = defaultEnv
	}

	if 0 == len(localEnv) {
		return
	}

	for key, value := range localEnv {
		Env[key] = value
	}
}
