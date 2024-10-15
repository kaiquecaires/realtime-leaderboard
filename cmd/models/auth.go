package models

type LoginParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (p *LoginParams) Validate() map[string]string {
	errors := map[string]string{}

	if len(p.Username) < 5 || len(p.Username) > 20 {
		errors["username"] = "username must have between 5 and 20 chars"
	}

	if len(p.Password) < 8 || len(p.Password) > 20 {
		errors["password"] = "password must have between 8 and 20 chars"
	}

	return errors
}

type Login struct {
	AccessToken string `json:"access_token"`
}
