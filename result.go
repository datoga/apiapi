package apiapi

type Result struct {
	HTML    string  `json:"html"`
	Network Network `json:"network"`
}

type Network struct {
	Headers map[string][]string `json:"headers"`
	Cookies map[string]string   `json:"cookies"`
}
