package main

import (
	"goweb/internal/models"
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/form/v4"
)

type application struct {
	errorLog       *log.Logger
	infoLog        *log.Logger
	posts          models.PostModelInterface
	comments       models.CommentModelInterface
	likes          models.LikesModelInterface
	users          models.UserModelInterface
	templateCache  map[string]*template.Template
	formDecoder    *form.Decoder
	sessionManager *scs.SessionManager
}
