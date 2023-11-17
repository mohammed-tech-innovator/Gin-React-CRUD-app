package helpers

type SignupForm struct {
	Fname          string
	Lname          string
	Email          string
	Password       string
	RPassword      string
	RememberDevice bool
}

func CheckSignupForm(form SignupForm) (string, string, bool) {

	if len(form.Fname) < 1 {

		return "No frist name", "⚠️ please enter your first name.", true

	} else if len(form.Lname) < 1 {

		return "No last name", "⚠️ please enter your last name.", true
	} else if len(form.Fname) > 20 || len(form.Lname) > 20 {

		return "too long first name", "⚠️ maximum name lenght is 20 characters.", true
	} else if !isValidEmail(form.Email) {

		return "bad Email", "⚠️ Please enter a valid Email📧.", true
	} else if form.Password != form.RPassword {

		return "Password mismatch", "⚠️ Please enter the same password.", true

	} else if len(form.Password) < 6 {

		return "short password", "⚠️ The password should be at least 6 characters 🔤.", true

	} else if IsNumeric(form.Password) {

		return "Password isn't valid", "⚠️Password is entirly numeric 🔢.", true

	} else if IsAlphabetic(form.Password) {

		return "Password isn't valid", "⚠️Password is entirly alphabetic 🔢.", true

	} else if len(form.Password) > 30 {

		return "Long password", "⚠️ Password should be 30 characters maximum.", true

	} else {
		return "", "", false
	}
}
