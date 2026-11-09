package oauthproviderstructs

type OAuthProvider struct {
	ProviderId  uint
	ProviderKey string
	AuthUrl     string
	TokenUrl    string
}

type ProviderScopes struct {
	ScopeId    uint
	Scope      string
	ProviderId uint
}
