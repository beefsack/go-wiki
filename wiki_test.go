package wiki

import "testing"

func TestNewWiki(t *testing.T) {
	NewWiki(Config{
		Persist: NewInMemory(),
	})
}
