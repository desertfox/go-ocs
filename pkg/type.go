package ocs

import "github.com/desertfox/ocs/pkg/config"

type Ocs struct {
	Host   config.Host
	Config *config.Ocsconfig
}
