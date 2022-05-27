package ocs

import "time"

type host struct {
	Server  string    `yaml:"Server"`
	Token   string    `yaml:"Token"`
	Version string    `yaml:"Version"`
	Created time.Time `yaml:"Created"`
}

func NewHost(server, token, version string) host {
	return host{
		Server:  server,
		Token:   token,
		Version: version,
		Created: time.Now(),
	}
}
