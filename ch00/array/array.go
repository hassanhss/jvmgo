package array

type People struct {
	Name string
}

type Peoples []People

func newPeoples(count int) Peoples {
	if count > 0 {
		return make([]People, count)
	}
	return nil
}

func (self Peoples) SetPeople(index int, val People) {
	self[index] = val
}
