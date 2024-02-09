package foods

type Food struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Qty         int    `json:"qty"`
}

type Foods []Food
