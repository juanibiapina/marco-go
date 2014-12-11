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

func (app *application) Equal(o Expr) bool {
	switch other := o.(type) {
	case *application:
		return app.List.Equal(other.List)
	default:
		return false
	}
}
