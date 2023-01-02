package main

// Notifier ...
type Notifier interface {
	Register(version string) error
	Notify(message string) error
}
