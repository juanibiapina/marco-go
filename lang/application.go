package lang

type application struct {
	List Expr
}

func MakeApplication(list Expr) *application {
	return &application{list}
}

func (app *application) String() string {
	return "Application"
}
