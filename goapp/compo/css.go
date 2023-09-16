package compo

import (
	"github.com/mlctrez/goappnew/goapp/global"
	s "github.com/oetherington/smetana"
)

func AddCSS(css *s.StyleSheet) {
	css.AddClass(global.ContentEl, s.CssProps{
		{"font-size", "25px"},
		{"text-align", "center"},
		{"margin", "50px"},
		{"color", "white"},
		{"background-color", "rgba(255, 255, 128, .5)"},
	})
	css.AddClass(global.UpEl, s.CssProps{
		{"margin", "20px"},
		{"background-color", "green"},
	})
	css.AddClass(global.CenterEl, s.CssProps{
		{"margin", "20px"},
		{"color", "blue"},
		{"background-color", "yellow"},
	})
	css.AddClass(global.FooterEl, s.CssProps{
		{"margin", "20px"},
		{"background-color", "blue"},
	})
	css.AddClass(global.BtnEl, s.CssProps{
		{"margin", "20px"},
		{"color", "rgb(31, 74, 78)"},
		{"background-color", "rgb(205, 226, 228)"},
		{"cursor", "grab"},
	})
	css.AddClass(global.MsgEl, s.CssProps{
		{"margin", "50px"},
		{"color", "black"},
		{"background-color", "rgb(6, 226, 246)"},
	})

}
