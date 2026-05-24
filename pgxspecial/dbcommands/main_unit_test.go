//go:build !integration

package dbcommands_test

import "testing"

func runTestMain(m *testing.M) int {
	return m.Run()
}
