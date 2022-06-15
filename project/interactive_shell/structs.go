package main

type SeDForm struct {
	Action    string `json:"action"`
	Searchkey string `json:"searchkey"`
	User      string `json:"user"`
	Password  string `json:"password"`
}

type ShDForm struct {
	Action   string `json:"action"`
	Domain   string `json:"domain"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type STForm struct {
	Action     string `json:"action"`
	Domainname string `json:"domainname"`
	TTL        string `json:"ttl"`
	Reason     string `json:"reason"`
	User       string `json:"user"`
	Password   string `json:"password"`
}

type respSeDForm struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Domains map[string]interface{} `json:"Domains"`
}

type respShDForm struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Domain  struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Type     string `json:"type"`
		Master   string `json:"master"`
		Accounts string `json:"accounts"`
		Records  []struct {
			ID       int    `json:"id"`
			DomainID int    `json:"domain_id"`
			Name     string `json:"name"`
			Type     string `json:"type"`
			Content  string `json:"content"`
			TTL      int    `json:"ttl"`
			Prio     int    `json:"prio"`
		} `json:"records"`
	} `json:"domain"`
}
