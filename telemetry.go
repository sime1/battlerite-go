package battlerite

type Telemetry struct {
	Cursor int `json:"cursor"`
	Type string `json:"type"`
	Data map[string]interface{} `json:"dataObject"`
}