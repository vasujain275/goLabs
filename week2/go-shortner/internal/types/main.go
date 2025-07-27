package types

type ShortenRequest struct{
	Url string `json:"url"`
}

type ShortenResponse struct{
	ShortCode string `json:"short_code"`
}
