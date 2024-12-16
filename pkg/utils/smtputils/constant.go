package smtputils

import (
	"embed"
)

//go:embed templates/*.html
var EmailHTMLTemplates embed.FS

const (
	ResetPasswordSubject = "[domain-driven-design] Please reset your password"
	VerificationSubject  = "[domain-driven-design] Verify your account"
	PharmacistSubject    = "[domain-driven-design] Pharmacist account"
)

type emailTemplate string

const (
	ResetPasswordTemplate emailTemplate = "templates/forgot-password.html"
	VerificationTemplate  emailTemplate = "templates/verification.html"
	PharmacistTemplate    emailTemplate = "templates/pharmacist.html"
)
