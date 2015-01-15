package main

import (
	"fmt"
	"log"
	"net/http"

	"git.qrawl.net/qdoc/core"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func ApiUserGetMe(req *http.Request, r render.Render) {
	user, err := JwtGetUser(req)
	if err != nil {
		log.Println("TOKEN ERROR:", err)
	}
	r.JSON(200, user)
}

func ApiDocumentsGetOne(r render.Render, params martini.Params, req *http.Request) {
	user, err := JwtGetUser(req)
	if err != nil || user.ID.Hex() == "" {
		r.JSON(http.StatusUnauthorized, nil)
		return
	}
	doc, err := user.GetDocumentById(params["id"])
	if err != nil {
		r.JSON(404, nil)
		return
	}
	r.JSON(200, doc)
	return
}

func ApiDocumentsGetAll(req *http.Request, r render.Render) {
	user, err := JwtGetUser(req)
	if err != nil || user.ID.Hex() == "" {
		r.JSON(http.StatusUnauthorized, nil)
		return
	}
	docs, err := user.Documents()
	if err != nil {
		log.Println("[web.ApiDocumentsGetAll] ERROR:", err)
	}
	r.JSON(200, docs)
}

func ApiDocumentsPost(req *http.Request, w http.ResponseWriter, r render.Render, doc core.Document) {
	user, err := JwtGetUser(req)
	if err != nil || user.ID.Hex() == "" {
		r.JSON(http.StatusUnauthorized, nil)
		return
	}
	err = user.AddDocument(&doc)
	if err != nil {
		log.Println("[web.ApiDocumentsPost] ERROR:", err)
		r.JSON(500, nil)
		return
	}
	w.Header().Set("Location", fmt.Sprint("/documents/", doc.ID.Hex()))
	r.JSON(201, doc)
}

func ApiDocumentsDelete(req *http.Request, r render.Render, params martini.Params) {
	user, err := JwtGetUser(req)
	if err != nil || user.ID.Hex() == "" {
		r.JSON(http.StatusUnauthorized, nil)
		return
	}
	candidateDocument := core.Document{
		EnteredId: params["id"],
		Owner:     user.ID,
	}
	err = user.DeleteDocument(&candidateDocument)
	if err != nil {
		r.JSON(500, nil)
		return
	}
	r.JSON(204, nil)
}

func ApiDocumentsPut(req *http.Request, w http.ResponseWriter, r render.Render, candidateDoc core.Document, params martini.Params) {
	user, err := JwtGetUser(req)
	if err != nil || user.ID.Hex() == "" {
		r.JSON(http.StatusUnauthorized, nil)
		return
	}
	targetDoc, err := user.GetDocumentById(params["id"])
	if err != nil {
		r.JSON(500, nil)
		return
	}
	candidateDoc.ID = targetDoc.ID
	err = user.PutDocument(&candidateDoc)
	if err != nil {
		r.JSON(500, nil)
		return
	}
	r.JSON(http.StatusFound, nil)
}

/*
func ApiTopicsGetAll(req *http.Request, r render.Render) {
	user, err := JwtGetUser(req)
	if err != nil || user.ID.Hex() == "" {
		r.JSON(http.StatusUnauthorized, nil)
		return
	}
	docsByTopic, err := user.DocumentsByTopic()
	if err != nil {
		r.JSON(204, docsByTopic)
		return
	}
	r.JSON(200, docsByTopic)
}
*/
