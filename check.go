package main

import (
	"fmt"
	"log"
	"os/exec"
	"syscall"
)

func checkInternetStatus() (bool, error) {
	// nc -dzw1 domain.com 443
	// Source: https://stackoverflow.com/a/10385867/7432786
	cmd := exec.Command("nc", "-d", "-z", "-w1", nsDomain, nsPort)

	err := cmd.Start()
	if err != nil {
		log.Fatalf("cmd.Start: %v", err)
		return false, fmt.Errorf("failed to start the command due to %v", err)
	}

	err = cmd.Wait()
	if err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			// The program has exited with an exit code != 0

			// This works on both Unix and Windows. Although package
			// syscall is generally platform dependent, WaitStatus is
			// defined for both Unix and Windows and in both cases has
			// an ExitStatus() method with the same signature.
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				if status.ExitStatus() != 1 {
					log.Printf("Exit Status: %d\n", status.ExitStatus())
				}
				// Assume internet is disconnected
				return false, nil
			}
		} else {
			log.Fatalf("cmd.Wait: %v\n", err)
		}
	}
	return true, nil
}
