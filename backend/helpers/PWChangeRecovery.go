package helpers

type PWCR struct {
	NPassword1 string
	NPassword2 string
}

func CheckPWCR(form PWCR) (string, string, bool) {

	if form.NPassword1 != form.NPassword2 {

		return "Password mismatch", "⚠️ Please enter the same password.", true

	} else if len(form.NPassword1) < 6 {

		return "short password", "⚠️ The password should be at least 6 characters 🔤.", true

	} else if IsNumeric(form.NPassword1) {

		return "Password isn't valid", "⚠️Password is entirly numeric 🔢.", true

	} else if IsAlphabetic(form.NPassword1) {

		return "Password isn't valid", "⚠️Password is entirly alphabetic 🔢.", true

	} else if len(form.NPassword1) > 30 {

		return "Long password", "⚠️ Password should be 30 characters maximum.", true

	} else {
		return "", "", false
	}
}
