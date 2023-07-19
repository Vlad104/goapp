package entities

type LoginDto struct{
	Email string `json:"email"`
	Password string `json:"password"`
}

type AuthDto struct{
	AccessToken string `json:"accessToken"`
}