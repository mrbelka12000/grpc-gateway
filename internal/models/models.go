package models

type Data struct {
	Ul []struct {
		Name    string `json:"name"`
		Inn     string `json:"inn"`
		CeoName string `json:"ceo_name"`
		Url     string `json:"url"`
	} `json:"ul"`
}
