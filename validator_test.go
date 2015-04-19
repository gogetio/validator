package validator

import "testing"

func testRunValidator(rules map[string]string, valid map[string]string, invalid map[string]string, t *testing.T) {
	if ok, messages := Validate(valid, rules); !ok {
		t.Error("Expected valid, got invalid for: ", rules, valid, messages)
	}
	if ok, messages := Validate(invalid, rules); ok {
		t.Error("Expected invalid, got valid for: ", rules, invalid, messages)
	}
}

func TestValidateAccepted(t *testing.T) {

	rules := make(map[string]string)
	valid := make(map[string]string)
	invalid := make(map[string]string)

	rules["1"] = "accepted"
	rules["2"] = "accepted"
	rules["3"] = "accepted"
	rules["4"] = "accepted"

	valid["1"] = "1"
	valid["2"] = "yes"
	valid["3"] = "true"
	valid["4"] = "on"

	invalid["1"] = "0"
	invalid["2"] = "no"
	invalid["3"] = "false"
	invalid["4"] = "off"

	testRunValidator(rules, valid, invalid, t)
}

func TestValidateActiveUrl(t *testing.T) {

}

func TestValidateAlpha(t *testing.T) {

}

func TestValidateAlphaDash(t *testing.T) {

}

func TestValidateAlphaNumeric(t *testing.T) {

}

func TestValidateBoolean(t *testing.T) {

}

func TestValidateChars(t *testing.T) {

}

func TestValidateCharsBetween(t *testing.T) {

}

func TestValidateConfirmed(t *testing.T) {

}

func TestValidateDate(t *testing.T) {

}

func TestValidateDifferent(t *testing.T) {

}

func TestValidateDigits(t *testing.T) {

}

func TestValidateDigitsBetween(t *testing.T) {

}

func TestValidateEmail(t *testing.T) {

}

func TestValidateIn(t *testing.T) {

}

func TestValidateInteger(t *testing.T) {

}

func TestValidateIp(t *testing.T) {

}

func TestValidateMaxChars(t *testing.T) {

}

func TestValidateMaxDigits(t *testing.T) {

}

func TestValidateMaxValue(t *testing.T) {

}

func TestValidateMinChars(t *testing.T) {

}

func TestValidateMinDigits(t *testing.T) {

}

func TestValidateMinValue(t *testing.T) {

}

func TestValidateNotIn(t *testing.T) {

}

func TestValidateNumeric(t *testing.T) {

}

func TestValidateRegex(t *testing.T) {

}

func TestValidateRequired(t *testing.T) {

}

func TestValidateSame(t *testing.T) {

}

func TestValidateUrl(t *testing.T) {

}

func TestValidateValue(t *testing.T) {

}

func TestValidateValueBetween(t *testing.T) {

}
