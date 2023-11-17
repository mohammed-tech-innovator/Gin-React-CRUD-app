package helpers

type LoginForm struct {
	Email          string
	Password       string
	RememberDevice bool
}

func CheckLoginForm(form LoginForm) (string, string, bool) {
	if !isValidEmail(form.Email) {
		return "bad Email", "⚠️ Please enter a valid Email📧.", true
	} else if IsNumeric(form.Password) {

		return "Password isn't valid", "⚠️Password is entirly numeric 🔢.", true

	} else if IsAlphabetic(form.Password) {

		return "Password isn't valid", "⚠️Password is entirly alphabetic 🔢.", true

	} else if len(form.Password) < 6 {

		return "short password", "⚠️ The password should be at least 6 characters 🔤.", true

	} else if len(form.Password) > 30 {

		return "Long password", "⚠️ Password should be 30 characters maximum.", true

	} else {
		return "", "", false
	}
}
