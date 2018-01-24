package intigua

import (
	"net/http"
	"log"
)

type Config struct{
	Host 		string
	Username	string
	ApiKey		string
}

func (c *Config) Client() (*Service, error) {
	service := NewService(&http.Client{
		Transport: &Transport{
			Username: c.Username,
			Password: c.ApiKey,
			UserAgent: "",
		},
	},
	c.Host,
	)

	log.Printf("[INFO] Intigua Client configured for user: %s", c.Username)

	return service, nil
}
