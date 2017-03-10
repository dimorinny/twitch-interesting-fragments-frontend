package configuration

//noinspection ALL
type Configuration struct {
	StorageHost string `env:"STORAGE_HOST"`
	Host        string `env:"HOST" envDefault:"127.0.0.1"`
	Port        int    `env:"PORT" envDefault:"8000"`
}
