package box

type deductible interface {
	deduct(balance int) int
}

type jail struct {
	name               string
	amountToBeDeducted int
}

type treasure struct {
	name            string
	amountToBeAdded int
}

func NewJail(name string, amountToBeDeducted int) *jail {
	return &jail{name: name, amountToBeDeducted: amountToBeDeducted}
}

// jail is a deductible which deducts amountToBeDeducted
func (j *jail) deduct(balance int) int {
	return balance - j.amountToBeDeducted
}

func NewTreasure(name string, amountToBeAdded int) *treasure {
	return &treasure{name: name, amountToBeAdded: amountToBeAdded}
}

// treasure is a deductible that actually adds the amountToBeAdded
func (t *treasure) deduct(balance int) int {
	return balance + t.amountToBeAdded
}
