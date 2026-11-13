package oauthproviderstructs

type OAuthProvider struct {
	ProviderId  uint
	ProviderKey string
	AuthUrl     string
	TokenUrl    string
	AppId       uint // We need this to link our internal stuff to the customer stuff!
}

type ProviderScopes struct {
	ScopeId    uint
	Scope      string
	ProviderId uint
}
