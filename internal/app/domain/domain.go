package domain

type Account struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Metadata struct {
	LastModified string `json:"lastModified"`
	Entries      int    `json:"entries"`
}

type Domain struct {
	Metadata Metadata  `json:"metadata"`
	Accounts []Account `json:"accounts"`
}
