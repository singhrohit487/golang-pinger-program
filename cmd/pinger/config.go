package main

import "fmt"

const DefaultInterface = "0.0.0.0"
const DefaultPort = 8000
const DefaultTargetHost = "localhost"
const DefaultTargetPath = "/"
const DefaultTargetPort = 8000
const DefaultTargetProto = "http"

// Config provides an interface for the configuration we will be using
// in this service
type Config struct {
	// Interface is the network interface we will listen on
	Interface string `json:"interface"`
	// Port is the port we will listen on
	Port uint16 `json:"port"`
	// TargetHost is the hostname of the target to ping
	TargetHost string `json:"target_host"`
	// TargetPath is the path URI of the target to ping
	TargetPath string `json:"target_path"`
	// TargetPort is the port which the target is listening on
	TargetPort uint16 `json:"target_port"`
	// TargetProto is the TCP protocol to use ("http" or "https")
	TargetProto string `json:"target_proto"`
}

// getTargetURL retrieves the exact URL which we can use to ping the
// target server
func (c *Config) getTargetURL() string {
	return fmt.Sprintf("%s://%s:%v%s", c.TargetProto, c.TargetHost, c.TargetPort, c.TargetPath)
}
