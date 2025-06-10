package models

type ContactInfo struct {
	Age   int    `json:"age"`
	Email string `json:"eamil"`
}

type User struct {
	Id          int         `json:"id" db:"id"`
	Firstname   string      `json:"firstname" db:"first_name"`
	Lastname    string      `json:"lastname" db:"last_name"`
	Email       string      `json:"email" db:"email"`
	Password    string      `json:"pwd" db:"pwd"`
	Roles       []string    `json:"roles"`
	ContactInfo ContactInfo `json:"contact"`
}

type Seller struct {
	Firstname   string      `json:"firstname"`
	Lastname    string      `json:"lastname"`
	ContactInfo ContactInfo `json:"contact"`
}

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

type Product struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
