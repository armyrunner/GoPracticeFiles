package filehelper

type Accounts struct {
	Accounts []Account `json:"Accounts"`
}

type Account struct {
	Name     string `json:"Name"`
	URL      string `json:"URL"`
	Username string `json:"Username"`
	Password string `json:"Password"`
}
