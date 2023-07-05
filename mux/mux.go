/*
Copyright 2022 The deepauto-io LLC.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package mux

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/asaskevich/govalidator"

	"github.com/chatgpt-accesstoken/render"

	akt "github.com/chatgpt-accesstoken"
	"github.com/gin-gonic/gin"
)

type Server struct {
	openAuthSvc akt.OpenaiAuthService
	proxySvc    akt.ProxyService
}

func New(openAuthSvc akt.OpenaiAuthService, proxySvc akt.ProxyService) *Server {
	return &Server{
		openAuthSvc: openAuthSvc,
		proxySvc:    proxySvc,
	}
}

func (s Server) Handler() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Any("/health", s.Healthy)

	ag := r.Group("/auth")
	{
		ag.POST("/", s.handlerPostAccessToken) // support [潘多拉]
		ag.POST("/puid", s.handlerPostPUID)
		ag.POST("/all", s.handlerPostAll)
	}

	pg := r.Group("/proxy")
	{
		pg.GET("/", s.handlerGetProxy)
		pg.POST("/", s.handlerPostProxy)
		pg.DELETE("/:id", s.handlerDeleteProxy)
	}
	return r
}

func (s Server) Healthy(ctx *gin.Context) {
	ctx.Writer.WriteHeader(http.StatusNoContent)
}

func (s Server) handlerPostAccessToken(ctx *gin.Context) {
	in := new(akt.OpenaiAuthRequest)
	if err := ctx.BindJSON(in); err != nil {
		render.BadRequest(ctx.Writer, err)
		return
	}

	if govalidator.IsNull(in.Email) {
		render.BadRequest(ctx.Writer, errors.New("api: cannot find email"))
		return
	}

	if govalidator.IsNull(in.Password) {
		render.BadRequest(ctx.Writer, errors.New("api: cannot find password"))
		return
	}

	res, err := s.openAuthSvc.AccessToken(ctx, in)
	if err != nil {
		render.InternalError(ctx.Writer, err)
		return
	}

	render.JSON(ctx.Writer, struct {
		Default string `json:"default"`
	}{
		Default: res.AccessToken,
	}, http.StatusOK)
}

func (s Server) handlerPostPUID(ctx *gin.Context) {
	in := new(akt.OpenaiAuthRequest)
	if err := ctx.BindJSON(in); err != nil {
		render.BadRequest(ctx.Writer, err)
		return
	}

	if govalidator.IsNull(in.AccessToken) {
		render.BadRequest(ctx.Writer, errors.New("api: cannot access token"))
		return
	}

	res, err := s.openAuthSvc.PUID(ctx, in)
	if err != nil {
		render.InternalError(ctx.Writer, err)
		return
	}
	render.JSON(ctx.Writer, res, http.StatusOK)
}

func (s Server) handlerPostAll(ctx *gin.Context) {
	in := new(akt.OpenaiAuthRequest)
	if err := ctx.BindJSON(in); err != nil {
		render.BadRequest(ctx.Writer, err)
		return
	}

	if govalidator.IsNull(in.Email) {
		render.BadRequest(ctx.Writer, errors.New("api: cannot find email"))
		return
	}

	if govalidator.IsNull(in.Password) {
		render.BadRequest(ctx.Writer, errors.New("api: cannot find password"))
		return
	}

	res, err := s.openAuthSvc.AccessToken(ctx, in)
	if err != nil {
		render.InternalError(ctx.Writer, err)
		return
	}
	render.JSON(ctx.Writer, res, http.StatusOK)
}

func (s Server) handlerGetProxy(ctx *gin.Context) {
	list, err := s.proxySvc.List(ctx)
	if err != nil {
		render.InternalError(ctx.Writer, err)
		return
	}

	render.JSON(ctx.Writer, list, 20)
}

type proxyRequest struct {
	Proxy string `json:"proxy"`
}

func (s Server) handlerPostProxy(ctx *gin.Context) {
	in := new(proxyRequest)
	if err := ctx.BindJSON(in); err != nil {
		render.BadRequest(ctx.Writer, err)
		return
	}

	if govalidator.IsNull(in.Proxy) {
		render.BadRequest(ctx.Writer, errors.New("api: cannot find proxy"))
		return
	}

	if err := s.proxySvc.Add(ctx, in.Proxy); err != nil {
		render.InternalError(ctx.Writer, err)
		return
	}

	ctx.Writer.WriteHeader(http.StatusNoContent)
}

func (s Server) handlerDeleteProxy(ctx *gin.Context) {
	ip := ctx.Param("ip")
	if govalidator.IsNull(ip) {
		render.BadRequest(ctx.Writer, errors.New("api: cannot find ip"))
		return
	}

	if err := s.proxySvc.Delete(ctx, ip); err != nil {
		render.InternalError(ctx.Writer, err)
		return
	}
	ctx.Writer.WriteHeader(http.StatusNoContent)
}

// JSON writes the json-encoded error message to the response
// with a 400 bad request status code.
func JSON(w http.ResponseWriter, v interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	enc := json.NewEncoder(w)
	enc.Encode(v)
}
