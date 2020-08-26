package routes

import (
	"fmt"
	"localdev/main/services"
	"net/http"

	"google.golang.org/appengine"
)

func LoginURL(w http.ResponseWriter, r *http.Request) {

	ctx := appengine.NewContext(r)

	url := services.Locator.UserService().LoginURL(ctx, "/")

	fmt.Fprint(w, url)
}

func LogoutURL(w http.ResponseWriter, r *http.Request) {

	ctx := appengine.NewContext(r)

	url := services.Locator.UserService().LogoutURL(ctx, "/")

	fmt.Fprint(w, url)
}

func Who(w http.ResponseWriter, r *http.Request) {

	ctx := appengine.NewContext(r)

	id := services.Locator.UserService().GetCurrentUserID(ctx)

	fmt.Fprint(w, id)
}
