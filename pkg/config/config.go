package config

import (
	"log"
	"text/template"

	"github.com/alexedwards/scs/v2"
)

// Appconfig holds application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}
