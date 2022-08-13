package routes

import (

	"github.com/go-chi/chi/v5"
    "github.com/smc181002/blog-golang/controllers"
)

// declaring blog routes for the user
func Blogs(r chi.Router) {
    r.Get("/blogs", controllers.GetBlogs)
}
