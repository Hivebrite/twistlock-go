package sdk

type TokenResponse struct {
	Token string `json:"token"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (c *Client) Authentication(username, password string) error {
	req, err := c.NewRequest("POST", "authenticate", &Credentials{
		Username: username,
		Password: password,
	})

	if err != nil {
		return err
	}

	token := TokenResponse{}
	_, err = c.Do(req, &token)
	if err != nil {
		return err
	}

	c.Token = token.Token

	return nil
}
