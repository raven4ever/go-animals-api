package animals

type NewAnimal struct {
	Name    string `form:"name" validate:"required"`
	Species string `form:"species" validate:"required"`
	Age     int64  `form:"age" validate:"required,gte=0,lte=130"`
}

type AnimalById struct {
	ID string `param:"id" validate:"required,uuid"`
}

type DeleteAnimalById struct {
	ID string `param:"id" validate:"required,uuid"`
}
