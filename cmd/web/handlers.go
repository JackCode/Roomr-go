package main

import "net/http"

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	rooms, err := app.rooms.GetAllRooms()
	if err != nil {
		app.serverError(nil, err)
	}

	app.render(w, r, "home.page.tmpl", &templateData{
		Rooms: rooms,
	})
}
