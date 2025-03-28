package url

import "time"

type RegisterUrlReqDto struct {
	OriginalURL string `json:"original_url"`
	TTLinSec    int    `json:"ttl_in_sec"`
}

type RegisterUrlRespDto struct {
	OriginalURL string    `json:"original_url"`
	ShortURL    string    `json:"short_url"`
	ExpiredAt   time.Time `json:"expired_at"`
}

type RedirectUrlReqDto struct {
	ShortUrl string `json:"short_url"`
}
