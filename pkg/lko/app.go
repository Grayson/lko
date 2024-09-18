package lko

type App struct {
}

func InitApp() *App {
	return &App{}
}

func (app *App) Run() error {
	println("running")
	return nil
}
