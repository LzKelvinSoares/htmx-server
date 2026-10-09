package utils

import (
	"html/template"
	"htmx-server/shared/constants"
	"log"
	"net/http"
)

type GetErrorParams struct {
	W         http.ResponseWriter
	ErrorCode int
	CustomMsg string
	Err       error
}

func GetHtmlHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", constants.CONTENT_TYPE["html"])
}

func GetError(params GetErrorParams) {
	if params.Err != nil {
		log.Println("erro:", params.Err)
	}
	switch params.ErrorCode {
	case http.StatusBadRequest:
		http.Error(params.W, params.CustomMsg, http.StatusBadRequest)
		return
	case http.StatusNotFound:
		http.Error(params.W, params.CustomMsg, http.StatusNotFound)
		return
	default:
		http.Error(params.W, "erro interno", http.StatusInternalServerError)
		return
	}
}

func GetTemplate(w http.ResponseWriter, tpl *template.Template, templateName string, data any) {
	GetHtmlHeader(w)
	if err := tpl.ExecuteTemplate(w, templateName, data); err != nil {
		params := GetErrorParams{
			W:         w,
			ErrorCode: http.StatusInternalServerError,
			Err:       err,
		}
		GetError(params)
	}
}
