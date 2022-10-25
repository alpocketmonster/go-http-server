package validator

type validator struct {
	rule string
}

func New(new_rule string) *validator {
	var v validator
	v.rule = new_rule

	return &v
}

func (v *validator) IsValid(val string) bool {
	if val == v.rule {
		return true
	}
	return false
}

func (v *validator) GetRule() string {
	return v.rule
}
