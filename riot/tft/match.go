package tft

import (
	"fmt"

	"github.com/d97brooks/golio/internal"
	log "github.com/sirupsen/logrus"
)

// MatchClient provides methods for match endpoints of the League of Legends TFT API.
type MatchClient struct {
	c *internal.Client
}

// GetMatchesByPUUID returns a list of match ids by PUUID
func (mc *MatchClient) GetMatchesByPUUID(puuid string) ([]string, error) {
	logger := mc.logger().WithField("method", "GetMatchesByPUUID")
	url := fmt.Sprintf(endpointMatchesByPUUID, puuid)
	var out []string
	if err := mc.c.GetInto(url, &out); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return out, nil
}

// GetMatchByMatchID returns a match by matchID
func (mc *MatchClient) GetMatchByMatchID(matchId string) (*Match, error) {
	logger := mc.logger().WithField("method", "GetMatchByMatchID")
	url := fmt.Sprintf(endpointMatchByMatchID, matchId)
	var out *Match
	if err := mc.c.GetInto(url, &out); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return out, nil
}

func (mc *MatchClient) logger() log.FieldLogger {
	return mc.c.Logger().WithField("category", "match")
}
