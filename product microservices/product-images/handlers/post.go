package handlers

import (
	"net/http"
	"path/filepath"
)

// swagger:route POST /images/{id}/{filename} ProductImages SaveImage
//  Saves Product image to file store
//
//  responses:
// 		201: noContentResponse
// 		400: errorResponse
//      500: errorResponse

// saveFile saves the contents of the requests to a file
func (f *Files) Save(w http.ResponseWriter, r *http.Request) {
	id, fn := f.getPATH(r)

	f.logger.Printf("[INFO] Validating Path, id: %s, filename :%s\n", id, fn)

	// check filepath is valid
	if id == "" || fn == "" {
		f.logger.Println("[ERROR] filename or id invalid", fn, id)
		w.WriteHeader(http.StatusBadRequest)
		ToJSON(w, &GenericError{Message: "Filename or product id invalid"})
		return
	}

	fp := filepath.Join(id, fn)
	err := f.store.Save(fp, r.Body)
	if err != nil {
		f.logger.Println("[ERROR] Unable to save file", err)
		w.WriteHeader(http.StatusInternalServerError)
		ToJSON(w, &GenericError{Message: err.Error()})
		return
	}
	w.WriteHeader(http.StatusCreated)
}
