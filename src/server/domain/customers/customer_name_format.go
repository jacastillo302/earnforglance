package domain

// CustomerNameFormat represents the customer name formatting enumeration.
type CustomerNameFormat int

const (
	// ShowEmails represents showing emails.
	ShowEmails CustomerNameFormat = 1

	// ShowUsernames represents showing usernames.
	ShowUsernames CustomerNameFormat = 2

	// ShowFullNames represents showing full names.
	ShowFullNames CustomerNameFormat = 3

	// ShowFirstName represents showing the first name.
	ShowFirstName CustomerNameFormat = 10
)
