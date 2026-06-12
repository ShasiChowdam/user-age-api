type Config struct {
    AppPort    string
    DBHost     string
    DBPort     string
    DBUser     string
    DBPassword string
    DBName     string
    DBSSLMode  string
}
func LoadConfig() (*Config, error)