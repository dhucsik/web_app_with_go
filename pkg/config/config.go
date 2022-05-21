package config

import "html/template"

//AppConfig holds the application config
type AppConfig struct {
	UserCache     bool
	TemplateCache map[string]*template.Template
}
