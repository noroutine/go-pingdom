package pingdom

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var detailedCheckJson = `
{
	"id" : 85975,
	"name" : "My check 7",
	"resolution" : 1,
	"sendnotificationwhendown" : 0,
	"notifyagainevery" : 0,
	"notifywhenbackup" : false,
	"created" : 1240394682,
	"type" : {
		"http" : {
			"url" : "/",
			"port" : 80,
			"requestheaders" : {
				"User-Agent" : "Pingdom.com_bot_version_1.4_(http://www.pingdom.com/)",
				"Prama" : "no-cache"
			}
		}
	},
	"hostname" : "s7.mydomain.com",
	"status" : "up",
	"lasterrortime" : 1293143467,
	"lasttesttime" : 1294064823,
	"tags": []
}
`

var detailedCheckDnsJson = `
{
	"id" : 85975,
	"name" : "My check 8",
	"resolution" : 1,
	"sendnotificationwhendown" : 0,
	"notifyagainevery" : 0,
	"notifywhenbackup" : false,
	"created" : 1240394682,
	"type" : {
		"dns" : {
			"expectedip" : "1.1.1.1"
		}
	},
	"hostname" : "s7.mydomain.com",
	"status" : "up",
	"lasterrortime" : 1293143467,
	"lasttesttime" : 1294064823,
	"tags": []
}
`

func TestPingdomError(t *testing.T) {
	pe := PingdomError{StatusCode: 400, StatusDesc: "Bad Request", Message: "Missing param foo"}
	want := "400 Bad Request: Missing param foo"
	assert.Equal(t, want, pe.Error())
}

func TestCheckResponseUnmarshal(t *testing.T) {
	var ck CheckResponse
	err := json.Unmarshal([]byte(detailedCheckJson), &ck)
	assert.NoError(t, err)
	assert.Equal(t, "http", ck.Type.Name)
	assert.NotNil(t, ck.Type.HTTP)
	assert.Equal(t, 2, len(ck.Type.HTTP.RequestHeaders))
}

func TestCheckResponseUnmarshalDns(t *testing.T) {
	var ck CheckResponse
	err := json.Unmarshal([]byte(detailedCheckDnsJson), &ck)
	assert.NoError(t, err)
	assert.Equal(t, "dns", ck.Type.Name)
	assert.NotNil(t, ck.Type.DNS)
	assert.Equal(t, "1.1.1.1", ck.Type.DNS.ExpectedIp)
}
