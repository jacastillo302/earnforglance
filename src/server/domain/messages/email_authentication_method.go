package domain

// EmailAuthenticationMethod represents an authentication method
type EmailAuthenticationMethod int

const (
	// None represents email account does not require authentication
	None EmailAuthenticationMethod = 0

	// Ntlm represents authentication through the default network credentials
	Ntlm EmailAuthenticationMethod = 5

	// Login represents authentication through username and password
	Login EmailAuthenticationMethod = 10

	// GmailOAuth2 represents authentication through Google APIs Client with OAuth2
	GmailOAuth2 EmailAuthenticationMethod = 15

	// MicrosoftOAuth2 represents authentication through Microsoft Authentication Client with OAuth2
	MicrosoftOAuth2 EmailAuthenticationMethod = 20
)
