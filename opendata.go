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

