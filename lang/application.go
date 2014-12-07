package lang

type application struct {
	List Expr
}

func MakeApplication(list Expr) *application {
	return &application{list}
}

func (application *application) String() string {
	return "TODO"
}
