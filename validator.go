package validator

import (
	"fmt"
	"net"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Validator func(name string, value string, inputs map[string]string, params []string) bool

var Validators = map[string]Validator{
	"accepted":       ValidateAccepted,
	"active_url":     ValidateActiveUrl,
	"alpha":          ValidateAlpha,
	"alpha_dash":     ValidateAlphaDash,
	"alpha_num":      ValidateAlphaNumeric,
	"boolean":        ValidateBoolean,
	"chars":          ValidateChars,
	"chars_between":  ValidateCharsBetween,
	"confirmed":      ValidateConfirmed,
	"date":           ValidateDate,
	"different":      ValidateDifferent,
	"digits":         ValidateDigits,
	"digits_between": ValidateDigitsBetween,
	"email":          ValidateEmail,
	"in":             ValidateIn,
	"integer":        ValidateInteger,
	"ip":             ValidateIp,
	"max_chars":      ValidateMaxChars,
	"max_digits":     ValidateMaxDigits,
	"max_value":      ValidateMaxValue,
	"min_chars":      ValidateMinChars,
	"min_digits":     ValidateMinDigits,
	"min_value":      ValidateMinValue,
	"not_in":         ValidateNotIn,
	"numeric":        ValidateNumeric,
	"regex":          ValidateRegex,
	"required":       ValidateRequired,
	"same":           ValidateSame,
	"url":            ValidateUrl,
	"value":          ValidateValue,
	"value_between":  ValidateValueBetween,
}

var Messages = map[string]string{
	"accepted":       "The %s must be accepted.",
	"active_url":     "The %s is not a valid URL.",
	"alpha":          "The %s may only contain letters.",
	"alpha_dash":     "The %s may only contain letters, numbers, and dashes.",
	"alpha_num":      "The %s may only contain letters and numbers.",
	"boolean":        "The %s field must be true or false.",
	"chars":          "The %s field must have %s characters.",
	"chars_between":  "The %s field must have between %s characters.",
	"confirmed":      "The %s confirmation does not match.",
	"date":           "The %s is not a valid date.",
	"different":      "The %s and %s must be different.",
	"digits":         "The %s must have %s digits.",
	"digits_between": "The %s must have between %s and %s digits.",
	"email":          "The %s must be a valid email address.",
	"in":             "The selected %s is invalid.",
	"max_chars":      "The %s must have fewer than %s characters.",
	"max_digits":     "The %s must have fewer than %s digits.",
	"max_value":      "The %s must be less than %s.",
	"min_chars":      "The %s must have more than %s characters.",
	"min_digits":     "The %s must have more than %s digits.",
	"min_value":      "The %s must be greater than %s.",
	"integer":        "The %s must be an integer.",
	"ip":             "The %s must be a valid IP address.",
	"not_in":         "The selected %s is invalid.",
	"numeric":        "The %s must be a number.",
	"regex":          "The %s format is invalid.",
	"required":       "The %s field is required.",
	"same":           "The %s and %s must match.",
	"url":            "The %s format is invalid.",
	"value":          "The %s must %s.",
	"value_between":  "The %s must be between %s and %s.",
}

func ValidateAccepted(name string, value string, inputs map[string]string, params []string) bool {
	valid := []string{"1", "true", "yes", "on"}
	return stringInSlice(value, valid)
}

func ValidateActiveUrl(name string, value string, inputs map[string]string, params []string) bool {
	lc := strings.ToLower(value)
	if validScheme := strings.HasPrefix(lc, "http://") || strings.HasPrefix(lc, "https://"); !validScheme {
		return false
	}
	// trim schema and then check dns
	lc = strings.TrimPrefix(lc, "http://")
	lc = strings.TrimPrefix(lc, "https://")
	_, err := net.LookupHost(lc)
	return err == nil
}

func ValidateAlpha(name string, value string, inputs map[string]string, params []string) bool {
	return regexp.MustCompile("^[a-zA-Z]+$").MatchString(value)
}

func ValidateAlphaDash(name string, value string, inputs map[string]string, params []string) bool {
	return regexp.MustCompile("^[a-zA-Z0-9-_]+$").MatchString(value)
}

func ValidateAlphaNumeric(name string, value string, inputs map[string]string, params []string) bool {
	return regexp.MustCompile("^[a-zA-Z0-9]+$").MatchString(value)
}

func ValidateBoolean(name string, value string, inputs map[string]string, params []string) bool {
	_, err := strconv.ParseBool(value)
	return err == nil
}

func ValidateChars(name string, value string, inputs map[string]string, params []string) bool {
	if len(params) == 1 {
		charCount, err := strconv.ParseInt(params[0], 10, 16)
		return err == nil && int64(len(value)) == charCount
	}
	return false
}

func ValidateCharsBetween(name string, value string, inputs map[string]string, params []string) bool {
	if len(params) == 2 {
		p1 := []string{params[0]}
		p2 := []string{params[1]}
		return ValidateMinChars(name, value, inputs, p1) && ValidateMaxChars(name, value, inputs, p2)
	}
	return false
}

func ValidateConfirmed(name string, value string, inputs map[string]string, params []string) bool {
	fieldValue, fieldExists := inputs[name+"_confirmation"]
	return fieldExists && fieldValue == value
}

func ValidateDate(name string, value string, inputs map[string]string, params []string) bool {
	if len(params) == 1 {
		_, err := time.Parse(params[0], value)
		return err == nil
	}
	return false
}

func ValidateDifferent(name string, value string, inputs map[string]string, params []string) bool {
	return !ValidateSame(name, value, inputs, params)
}

func ValidateDigits(name string, value string, inputs map[string]string, params []string) bool {
	if len(params) == 1 && regexp.MustCompile("^[0-9]+$").MatchString(value) {
		digitCount, err := strconv.ParseInt(params[0], 10, 16)
		return err == nil && int64(len(value)) == digitCount
	}
	return false
}

func ValidateDigitsBetween(name string, value string, inputs map[string]string, params []string) bool {
	if len(params) == 2 {
		p1 := []string{params[0]}
		p2 := []string{params[1]}
		return ValidateMinDigits(name, value, inputs, p1) && ValidateMaxDigits(name, value, inputs, p2)
	}
	return false
}

func ValidateEmail(name string, value string, inputs map[string]string, params []string) bool {
	return regexp.MustCompile("^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$").MatchString(value)
}

func ValidateIn(name string, value string, inputs map[string]string, params []string) bool {
	return stringInSlice(value, params)
}

func ValidateInteger(name string, value string, inputs map[string]string, params []string) bool {
	_, err := strconv.ParseInt(value, 10, 64)
	return err == nil
}

func ValidateIp(name string, value string, inputs map[string]string, params []string) bool {
	return net.ParseIP(value) != nil
}

func ValidateMinChars(name string, value string, inputs map[string]string, params []string) bool {
	if len(params) == 1 {
		minChars, err := strconv.ParseInt(params[1], 10, 16)
		return err == nil && int64(len(value)) >= minChars
	}
	return false
}

func ValidateMinDigits(name string, value string, inputs map[string]string, params []string) bool {
	if len(params) == 1 && regexp.MustCompile("^[0-9]+$").MatchString(value) {
		minDigits, err := strconv.ParseInt(params[0], 10, 16)
		return err == nil && int64(len(value)) >= minDigits
	}
	return false
}

func ValidateMinValue(name string, value string, inputs map[string]string, params []string) bool {
	if len(params) == 1 {
		minValue, mvErr := strconv.ParseFloat(params[0], 64)
		floatValue, fvErr := strconv.ParseFloat(value, 64)
		return mvErr == nil && fvErr == nil && floatValue >= minValue
	}
	return false
}

func ValidateMaxChars(name string, value string, inputs map[string]string, params []string) bool {
	if len(params) == 1 {
		maxChars, err := strconv.ParseInt(params[1], 10, 16)
		return err == nil && int64(len(value)) <= maxChars
	}
	return false
}

func ValidateMaxDigits(name string, value string, inputs map[string]string, params []string) bool {
	if len(params) == 1 && regexp.MustCompile("^[0-9]+$").MatchString(value) {
		maxDigits, err := strconv.ParseInt(params[0], 10, 16)
		return err == nil && int64(len(value)) <= maxDigits
	}
	return false
}

func ValidateMaxValue(name string, value string, inputs map[string]string, params []string) bool {
	if len(params) == 1 {
		maxValue, mvErr := strconv.ParseFloat(params[0], 64)
		floatValue, fvErr := strconv.ParseFloat(value, 64)
		return mvErr == nil && fvErr == nil && floatValue <= maxValue
	}
	return false
}

func ValidateNotIn(name string, value string, inputs map[string]string, params []string) bool {
	return !stringInSlice(value, params)
}

func ValidateNumeric(name string, value string, inputs map[string]string, params []string) bool {
	_, err := strconv.ParseFloat(value, 64)
	return err == nil
}

func ValidateRegex(name string, value string, inputs map[string]string, params []string) bool {
	if len(params) == 1 {
		rx, err := regexp.Compile(params[0])
		return err == nil && rx.MatchString(value)
	}
	return false
}

func ValidateRequired(name string, value string, inputs map[string]string, params []string) bool {
	return value != ""
}

func ValidateSame(name string, value string, inputs map[string]string, params []string) bool {
	if len(params) == 1 {
		fieldValue, fieldExists := inputs[params[0]]
		return fieldExists && fieldValue == value
	}
	return false
}

func ValidateUrl(name string, value string, inputs map[string]string, params []string) bool {
	lc := strings.ToLower(value)
	if validScheme := strings.HasPrefix(lc, "http://") || strings.HasPrefix(lc, "https://"); validScheme {
		return true // todo
	}
	return false
}

func ValidateValue(name string, value string, inputs map[string]string, params []string) bool {
	if len(params) == 1 {
		expectedValue, evErr := strconv.ParseFloat(params[0], 16)
		actualValue, avErr := strconv.ParseFloat(value, 16)
		return evErr == nil && avErr == nil && expectedValue == actualValue
	}
	return false
}

func ValidateValueBetween(name string, value string, inputs map[string]string, params []string) bool {
	if len(params) == 2 {
		p1 := []string{params[0]}
		p2 := []string{params[1]}
		return ValidateMinValue(name, value, inputs, p1) && ValidateMaxValue(name, value, inputs, p2)
	}
	return false
}

func Validate(inputs map[string]string, rules map[string]string) (bool, map[string]string) {
	// initialize an error messages map
	messages := make(map[string]string)
	for fieldName, fieldRulesRaw := range rules {
		// process each rule
		// start by extracting relevant field info
		fieldValue, fieldExists := inputs[fieldName]
		fieldRules := strings.Split(fieldRulesRaw, "|")
		fieldIsRequired := stringInSlice("required", fieldRules)
		alwaysValidate := stringInSlice("always", fieldRules)
		if fieldIsRequired && !fieldExists {
			// add message saying field is required
			// don't worry about the value, that will be handled below
			messages[fieldName] = buildErrorMessage(fieldName, "required", []string{})
		} else if fieldIsRequired || alwaysValidate || (!fieldIsRequired && fieldValue != "") {
			// process the rules
			for i := 0; i < len(fieldRules); i++ {
				rule, params := splitRuleParams(fieldRules[i])
				if vFn, vExists := Validators[rule]; vExists {
					// specified validator exists, call it
					if !vFn(fieldName, fieldValue, inputs, params) {
						// validation failed, add a message
						messages[fieldName] = buildErrorMessage(fieldName, rule, params)
					}
				}
			}
		}
	}
	// validation succeeded if there are no error messages
	return len(messages) == 0, messages
}

func RulesFromStruct(s interface{}) map[string]string {
	var rules = make(map[string]string)
	rv := reflect.ValueOf(s)
	for i := 0; i < rv.NumField(); i++ {
		rti := rv.Type().Field(i)
		jsonName := rti.Tag.Get("json")
		validate := rti.Tag.Get("validate")
		if jsonName != "" && validate != "" {
			rules[jsonName] = validate
		}
	}
	return rules
}

func buildErrorMessage(fieldName string, rule string, params []string) string {
	message, exists := Messages[rule]
	if !exists || strings.Count(message, "%s") != 1+len(params) {
		return fmt.Sprintf("The %s is invalid.", fieldName)
	}
	args := make([]interface{}, 1+len(params))
	args[0] = fieldName
	for i := 0; i < len(params); i++ {
		args[i+1] = params[i]
	}
	return fmt.Sprintf(message, args...)
}

func splitRuleParams(ruleWithParams string) (string, []string) {
	var params []string
	if split := strings.Split(ruleWithParams, ":"); len(split) == 2 {
		// params were passed
		rule := split[0]
		params := strings.Split(split[1], ",")
		return rule, params
	}
	return ruleWithParams, params
}

func stringInSlice(needle string, haystack []string) bool {
	for _, val := range haystack {
		if val == needle {
			return true
		}
	}
	return false
}
