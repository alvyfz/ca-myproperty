package openapi

type Ipapi struct {
	IP                 string `json:"ip"`
	Version            string `json:"version"`
	City               string `json:"city"`
	Region             string `json:"region"`
	RegionCode         string `json:"region_code"`
	CountryCode        string `json:"country_code"`
	CountryCodeIso3    string `json:"country_code_iso3"`
	CountryName        string `json:"country_name"`
	CountryCapital     string `json:"country_capital"`
	CountryTld         string `json:"country_tld"`
	ContinentCode      string `json:"continent_code"`
	Postal             string `json:"postal"`
	Timezone           string `json:"timezone"`
	Asn                string `json:"asn"`
	Org                string `json:"org"`
}
