package animal

type Animal struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Species string `json:"species"`
	Age     int    `json:"age"`
}

type Animals []Animal
