package admin

type GenerateTokenReqDto struct {
	UserId string `json:"user_id" binding:"required"`
}

type GenerateTokenRespDto struct {
	Token string `json:"token"`
}

type ValidateTokenReqDto struct {
	Token string `json:"token" binding:"required"`
}

type ValidateTokenRespDto struct {
	UserId    string `json:"user_id"`
	ExpiredAt string `json:"expired_at"`
}
