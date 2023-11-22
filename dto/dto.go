package dto

type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

type BrasilApi struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

type BrasilAberto struct {
	Result struct {
		Street         string `json:"street"`
		Complement     string `json:"complement"`
		District       string `json:"district"`
		DistrictID     int    `json:"districtId"`
		City           string `json:"city"`
		CityID         int    `json:"cityId"`
		IbgeID         int    `json:"ibgeId"`
		State          string `json:"state"`
		StateShortname string `json:"stateShortname"`
		Zipcode        string `json:"zipcode"`
	} `json:"result"`
}
