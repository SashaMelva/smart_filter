package entity

type Tokens struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken []byte `json:"refreshToken"`
}
type TokenModel struct {
	UserGUID     string `json:"user_id"`
	AccessToken  string `json:"accessToken"`
	RefreshToken []byte `json:"refreshToken"`
}

type RefreshToken struct {
	UserId       int    `json:"user_id"`
	RefreshToken []byte `json:"refreshToken"`
}
