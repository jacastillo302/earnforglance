package domain

// CaptchaType represents a type of reCAPTCHA
type CaptchaType int

const (
	// CheckBoxReCaptchaV2 represents reCAPTCHA v2 check box
	CheckBoxReCaptchaV2 CaptchaType = 10

	// ReCaptchaV3 represents reCAPTCHA v3
	ReCaptchaV3 CaptchaType = 20
)
