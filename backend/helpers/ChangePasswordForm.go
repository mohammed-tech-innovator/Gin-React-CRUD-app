package helpers

type ChangePasswordForm struct {
	Email        string
	Password     string
	Newpassword1 string
	Newpassword2 string
}

func CheckChangePasswordForm(form ChangePasswordForm) (string, string, bool) {

	if !isValidEmail(form.Email) {

		return "bad Email", "⚠️ Please enter a valid Email📧.", true

	} else if form.Newpassword1 != form.Newpassword2 {

		return "Password mismatch", "⚠️ Please enter the same password.", true

	} else if len(form.Newpassword1) < 6 {

		return "short password", "⚠️ The password should be at least 6 characters 🔤.", true

	} else if IsNumeric(form.Newpassword1) {

		return "Password isn't valid", "⚠️Password is entirly numeric 🔢.", true

	} else if IsAlphabetic(form.Newpassword1) {

		return "Password isn't valid", "⚠️Password is entirly alphabetic 🔢.", true

	} else if len(form.Newpassword1) > 30 {

		return "Long password", "⚠️ Password should be 30 characters maximum.", true

	} else {
		return "", "", false
	}
}
