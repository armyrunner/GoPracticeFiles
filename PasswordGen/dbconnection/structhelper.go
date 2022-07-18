package dbconnection

type Accounts struct {
	Accounts []Account `json:"Accounts"`
}

type Account struct {
	Name     string `json:"Name"`
	URL      string `json:"URL"`
	Username string `json:"Username"`
	Password string `json:"Password"`
}

type AccountInformation struct{
	AcctInformation []AcctInfo `json:"AccountInformation"`
}


type AcctInfo struct{
	Username string `json:"Username"`
	Password string `json:"Password"`
}
