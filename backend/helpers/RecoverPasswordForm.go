package helpers

type RecoverPasswordForm struct {
	Email string
}

func CheckRecoverPasswordForm(form RecoverPasswordForm) (string, string, bool) {
	if !isValidEmail(form.Email) {
		return "bad Email", "⚠️ Please enter a valid Email📧.", true
	} else {
		return "", "", false
	}
}
