package router

import (
	"github.com/pclubiitk/puppy-love/controllers"
	"github.com/pclubiitk/puppy-love/db"

	"github.com/kataras/iris"
)

func PuppyRoute(db db.PuppyDb) {

	iris.Get("/", func(ctx *iris.Context) {
		ctx.JSON(iris.StatusAccepted, "Hello from the other side!")
	})

	// User administration
	uPre := "/users"
	iris.Handle("POST", uPre+"/login/first", controllers.UserFirst{db})
	iris.Handle("POST", uPre+"/data/update", controllers.UserUpdateData{db})
	iris.Handle("POST", uPre+"/data/submit", controllers.UserSubmitTrue{db})

	iris.Handle("GET", uPre+"/data/info", controllers.UserLoginGet{db})
	iris.Handle("GET", uPre+"/get/:id", controllers.UserGet{db})

	iris.Get(uPre+"/mail/:id", controllers.UserMail)

	// Listing users
	lPre := "/list"
	iris.Handle("GET", lPre+"/gender/:gender", controllers.ListAll{db})
	iris.Handle("GET", lPre+"/compute", controllers.ComputeList{db})
	iris.Handle("GET", lPre+"/pubkey/:gender", controllers.PubkeyList{db})

	// Compute
	cPre := "/compute"
	iris.Handle("POST", cPre+"/token", controllers.ComputeStep{db, 0})
	iris.Handle("POST", cPre+"/result", controllers.ComputeStep{db, 1})
	iris.Handle("POST", cPre+"/value", controllers.ComputeStep{db, 2})

	// Declare
	iris.Handle("POST", "/declare/choices", controllers.DeclareStep{db})

	// Session administration
	sesPre := "/session"
	iris.Handle("POST", sesPre+"/login", controllers.SessionLogin{db})
	iris.Get(sesPre+"/logout", controllers.SessionLogout)

	// Admin
	aPre := "/admin"
	iris.Handle("GET", aPre+"/compute/drop", controllers.ComputeDelete{db})
	iris.Handle("GET", aPre+"/compute/list", controllers.ComputeListAdmin{db})
	iris.Handle("GET", aPre+"/compute/prepare", controllers.ComputePrepare{db})

	iris.Handle("GET", aPre+"/user/drop", controllers.UserDelete{db})
	iris.Handle("POST", aPre+"/user/new", controllers.UserNew{db})
}
