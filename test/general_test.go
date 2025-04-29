package test

import "testing"

func Test_With_Cleanup(t *testing.T) {
	// Some test code here
	t.Cleanup(func() {
		// cleanup logic
	})
	// more test code here
}
