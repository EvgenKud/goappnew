package compo

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goappnew/goapp/global"
)

var _ app.AppUpdater = (*Root)(nil)
var _ app.Mounter = (*Root)(nil)

type Root struct {
	app.Compo
	showMsg bool
}

func (r *Root) Render() app.UI {
	return app.Div().Class(global.ContentEl).Body(
		app.Div().Class(global.UpEl).Text("Hello"),
		app.Div().Class(global.CenterEl).Text("Have a good day"),
		app.Div().Class(global.FooterEl).Text("Goodbye"),
		app.Div().Class(global.BtnEl).Text("Press Me!").
			OnClick(r.clickBtn),
		app.If(r.showMsg, app.Div().Class(global.MsgEl).Text("Message here!")),
	)
}

func (r *Root) OnAppUpdate(ctx app.Context) {
	if app.Getenv("DEV") != "" && ctx.AppUpdateAvailable() {
		ctx.Reload()
	}
}

func (r *Root) OnMount(ctx app.Context) {

}

func (r *Root) clickBtn(ctx app.Context, e app.Event) {
	r.showMsg = !r.showMsg
}
