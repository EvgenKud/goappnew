package css

import (
	"fmt"
	"github.com/mlctrez/goappnew/goapp/compo"
	"github.com/mlctrez/goappnew/goapp/global"
	s "github.com/oetherington/smetana"
	"os"
)

func Gen() error {
	fmt.Println("generate css styles")

	css := s.NewStyleSheet()
	compo.AddCSS(&css)
	gCss := s.RenderCss(css, s.Palette{})

	return os.WriteFile(global.AppCss, []byte(gCss), 0777)
}
