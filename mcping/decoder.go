package mcping

import (
	"encoding/json"
	"github.com/iverly/go-mcping/api/types"
	"github.com/jmoiron/jsonq"
	"strings"
)

func decodeResponse(response string) *types.PingResponse  {
	d := map[string]interface{}{}
	dec := json.NewDecoder(strings.NewReader(response))
	dec.Decode(&d)
	jq := jsonq.NewQuery(d)

	presp := &types.PingResponse{}

	presp.Sample = decodePlayersSimple(jq)
	presp.Motd = decodeMotd(jq)

	presp.Online, _ = jq.Int("players", "online")
	presp.Max, _ = jq.Int("players", "max")
	presp.Protocol, _ = jq.Int("version", "protocol")
	presp.Favicon, _ = jq.String("favicon")
	presp.Version, _ = jq.String("version", "name")

	return presp
}

func decodePlayersSimple(jq *jsonq.JsonQuery) []types.PlayerSample {
	psm, _ := jq.ArrayOfObjects("players", "sample")
	var playerSamples []types.PlayerSample
	for k := range psm {
		sample := types.PlayerSample{}
		sample.UUID = psm[k]["id"].(string)
		sample.Name = psm[k]["name"].(string)
		playerSamples = append(playerSamples, sample)
	}
	return playerSamples
}

func decodeMotd(jq *jsonq.JsonQuery) string {
	tm, _ := jq.ArrayOfObjects("description", "extra")
	var sb strings.Builder
	for k := range tm {
		sb.WriteString(tm[k]["text"].(string))
	}
	return sb.String()
}
