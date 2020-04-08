package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/eencogneeto/backend/api/auth"
	"github.com/eencogneeto/backend/api/models"
	"github.com/eencogneeto/backend/api/responses"
	"github.com/eencogneeto/backend/api/utils/formaterror"
	"github.com/gorilla/mux"
)

func (server *Server) CreateNoun(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	noun := models.Noun{}
	err = json.Unmarshal(body, &noun)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	noun.Prepare()
	err = noun.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	_ = uid
	// if uid != post.AuthorID {
	// 	responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
	// 	return
	// }
	nounCreated, err := noun.SaveNoun(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, nounCreated.ID))
	responses.JSON(w, http.StatusCreated, nounCreated)
}

func (server *Server) GetNouns(w http.ResponseWriter, r *http.Request) {

	noun := models.Noun{}

	nouns, err := noun.FindAllNouns(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, nouns)
}

func (server *Server) GetNoun(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	noun := models.Noun{}

	nounReceived, err := noun.FindNounByID(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, nounReceived)
}

func (server *Server) UpdateNoun(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Check if the post id is valid
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	// Check if the auth token is valid and get the user id from it
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	// Check if the post exist
	noun := models.Noun{}
	err = server.DB.Debug().Model(models.Noun{}).Where("id = ?", pid).Take(&noun).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Noun not found"))
		return
	}
	_ = uid
	// // If a user attempt to update a post not belonging to him
	// if uid != post.AuthorID {
	// 	responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
	// 	return
	// }
	// Read the data posted
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Start processing the request data
	nounUpdate := models.Noun{}
	err = json.Unmarshal(body, &nounUpdate)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	_ = uid
	// //Also check if the request user id is equal to the one gotten from token
	// if uid != postUpdate.AuthorID {
	// 	responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
	// 	return
	// }

	nounUpdate.Prepare()
	err = nounUpdate.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	nounUpdate.ID = noun.ID //this is important to tell the model the post id to update, the other update field are set above

	nounUpdated, err := nounUpdate.UpdateANoun(server.DB)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, nounUpdated)
}

func (server *Server) DeleteNoun(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Is a valid noun id given to us?
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	// Is this user authenticated?
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	// Check if the noun exist
	noun := models.Noun{}
	err = server.DB.Debug().Model(models.Noun{}).Where("id = ?", pid).Take(&noun).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Unauthorized"))
		return
	}
	_ = uid
	// // Is the authenticated user, the owner of this noun?
	// if uid != post.AuthorID {
	// 	responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
	// 	return
	// }
	_, err = noun.DeleteANoun(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", pid))
	responses.JSON(w, http.StatusNoContent, "")
}
