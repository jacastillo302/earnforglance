package domain

// CaptchaTypeID represents a type of reCAPTCHA
type CaptchaTypeID int

const (
	// CheckBoxReCaptchaV2 represents reCAPTCHA v2 check box
	CheckBoxReCaptchaV2 CaptchaTypeID = 10

	// ReCaptchaV3 represents reCAPTCHA v3
	ReCaptchaV3 CaptchaTypeID = 20
)
