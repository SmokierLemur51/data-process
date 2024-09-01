package handler

import (
	"net/http"

	"github.com/SmokierLemur51/data-process/db"
)

// Testing stage only
func (h *Handler) CommandHandler(w http.ResponseWriter, r *http.Request) {
	// Take command path value and execute on database
	switch r.PathValue("command") {
	case "create-schema":
		CheckErr(db.CreateSchema(h.DB, "db/scripts/schema.sql"))
	  http.Redirect(w, r, "/", http.StatusSeeOther)
  // During testing if we want to drop and recreate the database
	// without having to restart or rebuild the application.
	case "drop-tables":
		CheckErr(db.DeleteDatabase(h.DB))
		// All good send back to homepage
		http.Redirect(w, r, "/", http.StatusSeeOther)
	case "create-db-file":
		// Does not require any params
		CheckErr(db.CreateDatabaseFile())
	}
}

func (h *Handler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Success"))
}

func (h *Handler) ExampleHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.PathValue("value")))
}
