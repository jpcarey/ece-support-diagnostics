package eceAPI

import (
	"github.com/elastic/ece-support-diagnostics/config"
	"github.com/elastic/ece-support-diagnostics/store"
)

type testFileStore struct {
	store.ContentStore
	cfg *config.Config
}

// Rest defines HTTP URIs to be collected
type Rest struct {
	Method   string
	URI      string
	Filename string
	Headers  map[string]string
	Loop     string
	Sub      []Rest
}
