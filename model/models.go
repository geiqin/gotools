package model

type DbBool int //Custom Bool : 1 = true , 2 = false

func (b DbBool) MarshalJSON() ([]byte, error) {
	if b == 1 {
		return []byte("true"), nil
	}
	return []byte("false"), nil
}

func (b DbBool) GetBool() bool {
	if b == 1 {
		return true
	}
	return false
}
