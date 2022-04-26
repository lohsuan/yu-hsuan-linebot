package main

// 中央氣象局開放資料平臺之資料擷取API https://opendata.cwb.gov.tw/dist/opendata-swagger.html
const cwdOpendataUrl = "https://opendata.cwb.gov.tw/api/v1/rest/datastore/F-C0032-001"

type WeatherResponse struct {
	Success string `json:"success"`
	Result  struct {
		ResourceID string `json:"resource_id"`
		Fields     []struct {
			ID   string `json:"id"`
			Type string `json:"type"`
		} `json:"fields"`
	} `json:"result"`
	Records struct {
		DatasetDescription string `json:"datasetDescription"`
		Location           []struct {
			LocationName   string `json:"locationName"`
			WeatherElement []struct {
				ElementName string `json:"elementName"`
				Time        []struct {
					StartTime string `json:"startTime"`
					EndTime   string `json:"endTime"`
					Parameter struct {
						ParameterName  string `json:"parameterName"`
						ParameterValue string `json:"parameterValue"`
					} `json:"parameter"`
				} `json:"time"`
			} `json:"weatherElement"`
		} `json:"location"`
	} `json:"records"`
}

// 疾管署 covid19 資料擷取API
const globalCovidUrl = "https://covid19dashboard.cdc.gov.tw/dash2"
const taiwanCovidUrl = "https://covid19dashboard.cdc.gov.tw/dash3"

type GlabalCovid struct {
	Zero struct {
		Cases     string `json:"cases"`
		Deaths    string `json:"deaths"`
		Cfr       string `json:"cfr"`
		Countries int    `json:"countries"`
	} `json:"0"`
}

type TaiwanCovid struct {
	Zero struct {
		Case       string `json:"確診"`
		Death      int    `json:"死亡"`
		Inform     string `json:"送驗"`
		Except     string `json:"排除"`
		LastCase   int    `json:"昨日確診"`
		LastExcept string `json:"昨日排除"`
		LastInform string `json:"昨日送驗"`
	} `json:"0"`
}
