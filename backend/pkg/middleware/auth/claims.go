package auth

type AuthzClaims struct {
	Subject  string
	Action   string
	Resource string
	Project  string
}

func (c *AuthzClaims) GetSubject() string {
	return c.Subject
}

func (c *AuthzClaims) GetAction() string {
	return c.Action
}

func (c *AuthzClaims) GetResource() string {
	return c.Resource
}

func (c *AuthzClaims) GetProject() string {
	return c.Project
}
