package controllers

import (
    "net/http"
    "github.com/go-chi/render"
    "github.com/smc181002/blog-golang/models"
)

// Fetch all the blogs from database and render the data
// as JSON
func GetBlogs(w http.ResponseWriter, r *http.Request) {
    users := []models.User{
        {Name: "smc181002"},
    }
    render.JSON(w, r, users)
}
