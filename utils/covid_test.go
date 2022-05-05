package utils

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchGlabalCovidInfo(t *testing.T) {
	if os.Getenv("CI") == "true" {
		t.Skip()
	}
	globalCovid := new(GlabalCovid)
	err := FetchGlabalCovidInfo(globalCovid)
	if err != nil {
		t.Error("FetchGlabalCovidInfo FAIL")
	}
}

func TestGlobalCovidResponse(t *testing.T) {
	const srandardGlobalCovidResponse string = `{
		"0": {
		  "cases": "506,690,362",
		  "deaths": "6,232,767",
		  "cfr": "1.23%",
		  "countries": 199
		}
	  }`

	globalCovidResponse := new(GlabalCovid)
	json.NewDecoder(bytes.NewReader([]byte(srandardGlobalCovidResponse))).Decode(globalCovidResponse)

	assert.Equal(t, globalCovidResponse.Zero.Cases, "506,690,362")
	assert.Equal(t, globalCovidResponse.Zero.Deaths, "6,232,767")
	assert.Equal(t, globalCovidResponse.Zero.Cfr, "1.23%")
	assert.Equal(t, globalCovidResponse.Zero.Countries, 199)
}

func TestFetchTaiwanCovidInfo(t *testing.T) {
	if os.Getenv("CI") == "true" {
		t.Skip()
	}
	taiwanCovid := new(TaiwanCovid)
	err := FetchTaiwanCovidInfo(taiwanCovid)
	if err != nil {
		t.Error("FetchTaiwanCovidInfo FAIL")
	}
}

func TestTaiwanCovidResponse(t *testing.T) {
	const srandardTaiwanCovidResponse string = `{
		"0": {
		  "確診": "61,686",
		  "死亡": 856,
		  "送驗": "8,174,134",
		  "排除": "8,105,057",
		  "昨日確診": 5221,
		  "昨日排除": "57,972",
		  "昨日送驗": "64,648"
		}
	  }`

	taiwanCovidResponse := new(TaiwanCovid)
	json.NewDecoder(bytes.NewReader([]byte(srandardTaiwanCovidResponse))).Decode(taiwanCovidResponse)

	assert.Equal(t, taiwanCovidResponse.Zero.Case, "61,686")
	assert.Equal(t, taiwanCovidResponse.Zero.Death, 856)
	assert.Equal(t, taiwanCovidResponse.Zero.Inform, "8,174,134")
	assert.Equal(t, taiwanCovidResponse.Zero.Except, "8,105,057")
	assert.Equal(t, taiwanCovidResponse.Zero.LastCase, 5221)
	assert.Equal(t, taiwanCovidResponse.Zero.LastExcept, "57,972")
	assert.Equal(t, taiwanCovidResponse.Zero.LastInform, "64,648")
}
