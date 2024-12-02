package val

import (
	"testing"

	log "github.com/sirupsen/logrus"

	"github.com/d97brooks/golio/api"
	"github.com/d97brooks/golio/internal"
	"github.com/d97brooks/golio/internal/mock"
)

func TestNewClient(t *testing.T) {
	c := NewClient(internal.NewClient(api.RegionBrasil, "key", mock.NewStatusMockDoer(200), log.StandardLogger()))
	if c == nil {
		t.Error("returned nil")
	}
}
