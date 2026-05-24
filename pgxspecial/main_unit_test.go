//go:build !integration

package pgxspecial_test

import "testing"

func runTestMain(m *testing.M) int {
	return m.Run()
}
