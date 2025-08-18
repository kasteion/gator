package config

func (c *Config) SetUser(username string) error {
	c.CurrentUserName = username

	err := write(*c)
	if err != nil {
		return  err
	}

	return nil
}