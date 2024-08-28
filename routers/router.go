package routers

import (
	"catapiproject/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	//beego.Router("/", &controllers.MainController{})
	beego.Router("/", &controllers.CatController{})
	beego.Router("/breeds", &controllers.BreedsController{})
	beego.Router("/favorites", &controllers.FavoritesController{})

}
