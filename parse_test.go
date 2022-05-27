package main

import (
	"os"
	"testing"
)

func Test_parseArgs(t *testing.T) {
	t.Run("parseArgs", func(t *testing.T) {
		tests := []struct {
			name    string
			args    []string
			command string
			server  string
			token   string
			version string
		}{
			{"Empty", []string{}, "", "", "", ""},
			{"Program Path", []string{os.Args[0]}, "", "", "", ""},
			{"One String", []string{os.Args[0], "list"}, "list", "", "", ""},
			{"Two String", []string{os.Args[0], "del", "0"}, "del", "", "", ""},
			{"Login String", []string{os.Args[0], "oc", "login", "--server=hostname.com", "--token=SOMEHASH"}, "login", "hostname.com", "SOMEHASH", "4.X"},
			{"Login String OKD3.X", []string{os.Args[0], "oc", "login", "hostname.com", "--token=SOMEHASH"}, "login", "hostname.com", "SOMEHASH", "3.X"},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				command, server, token, version := parseArgs(tt.args)

				if command != tt.command {
					t.Fatalf("Command %s does not match expected %s", command, tt.command)
				}

				if server != tt.server {
					t.Fatalf("Server %s does not match expected %s", server, tt.server)
				}

				if token != tt.token {
					t.Fatalf("Token %s does not match expected %s", token, tt.token)
				}

				if version != tt.version {
					t.Fatalf("Version %s does not match expected %s", version, tt.version)
				}
			})
		}
	})
}
