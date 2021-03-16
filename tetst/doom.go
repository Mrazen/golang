package main

type doom struct {
	On    bool
	Ammo  int
	Power int
}

func (d *doom) Shoot() bool {
	if d.On == true && d.Ammo > 0 {
		d.Ammo--
		return true
	}
	return false

}

func (d *doom) RideBike() bool {
	if d.On == true && d.Power > 0 {
		d.Power--
		return true
	}
	return false

}

func main() {
	testStruct := new(doom)
	// d.Shoot()
	// d.RideBike()

}
