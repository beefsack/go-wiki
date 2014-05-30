package wiki

type Config struct {
	TemplatesDir string
	Persist      Persister
}

func (c Config) templatesDir() string {
	d := c.TemplatesDir
	if d == "" {
		return "templates"
	}
	return d
}
