package config

type Configuration interface {
	GetString(name string) (configValue string, found bool)
	GetSection(sectionName string) (section Configuration, found bool)
}
