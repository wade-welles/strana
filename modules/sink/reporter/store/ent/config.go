// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/dialect"
)

// Option function to configure the client.
type Option func(*config)

// Config is the configuration for the client and its builder.
type config struct {
	// driver used for executing database requests.
	driver dialect.Driver
	// debug enable a debug logging.
	debug bool
	// log used for logging on debug mode.
	log func(...interface{})
	// hooks to execute on mutations.
	hooks *hooks
}

// hooks per client, for fast access.
type hooks struct {
	Action       []ent.Hook
	Alias        []ent.Hook
	App          []ent.Hook
	Browser      []ent.Hook
	Campaign     []ent.Hook
	Connectivity []ent.Hook
	Device       []ent.Hook
	Event        []ent.Hook
	Extra        []ent.Hook
	Group        []ent.Hook
	Library      []ent.Hook
	Location     []ent.Hook
	Network      []ent.Hook
	OSContext    []ent.Hook
	Page         []ent.Hook
	Referrer     []ent.Hook
	Screen       []ent.Hook
	Session      []ent.Hook
	Timing       []ent.Hook
	User         []ent.Hook
	Viewport     []ent.Hook
}

// Options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...interface{})) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}
