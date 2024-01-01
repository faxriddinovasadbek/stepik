type testStruct struct {
	On bool
	Ammo int
	Power int
}

func (a *testStruct) Shoot() bool {
	if a.On == false{
		return false
	}else if a.Ammo > 0{
		a.Ammo -= 1
		return true
	}
	return false
}

func (a *testStruct) RideBike() bool {
	if a.On == false {
		return false
	}else if a.Power > 0{
		a.Power -= 1
		return true
	}
	return false
}


func main() {
    a := testStruct{true,1,2}
    testStruct := &a
}