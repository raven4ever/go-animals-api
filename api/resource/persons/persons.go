package persons

type Person struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	NickName  string `json:"nickName"`
	Age       int    `json:"age"`
}

type Persons []Person
