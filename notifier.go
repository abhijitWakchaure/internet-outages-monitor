package main

// Notifier ...
type Notifier interface {
	Register() error
	Notify(message string) error
}
