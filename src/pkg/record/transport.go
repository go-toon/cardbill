/**
 * @Time: 2019-08-18 11:34
 * @Author: solacowa@gmail.com
 * @File: transport
 * @Software: GoLand
 */

package record

import (
	"context"
	"encoding/json"
	kitjwt "github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/nsini/cardbill/src/middleware"
	"github.com/nsini/cardbill/src/util/encode"
	"io/ioutil"
	"net/http"
	"strconv"
)

type endpoints struct {
	PostEndpoint endpoint.Endpoint
	ListEndpoint endpoint.Endpoint
}

func MakeHandler(svc Service, logger log.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encode.EncodeError),
		kithttp.ServerBefore(kithttp.PopulateRequestContext),
		kithttp.ServerBefore(kitjwt.HTTPToContext()),
	}

	eps := endpoints{
		PostEndpoint: makePostEndpoint(svc),
		ListEndpoint: makeListEndpoint(svc),
	}

	ems := []endpoint.Middleware{
		middleware.CheckLogin(logger), // 2
		//kitjwt.NewParser(kpljwt.JwtKeyFunc, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory), // 1
	}

	mw := map[string][]endpoint.Middleware{
		"Post": ems,
		"List": ems,
	}

	for _, m := range mw["Post"] {
		eps.PostEndpoint = m(eps.PostEndpoint)
	}
	for _, m := range mw["List"] {
		eps.ListEndpoint = m(eps.ListEndpoint)
	}

	r := mux.NewRouter()
	r.Handle("/record", kithttp.NewServer(
		eps.PostEndpoint,
		decodePostRequest,
		encode.EncodeResponse,
		opts...,
	)).Methods("POST")

	r.Handle("/record", kithttp.NewServer(
		eps.ListEndpoint,
		decodeListRequest,
		encode.EncodeResponse,
		opts...,
	)).Methods("GET")

	return r
}

func decodeListRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))

	if pageSize == 0 {
		pageSize = 10
	}

	return listRequest{page, pageSize}, nil
}

func decodePostRequest(_ context.Context, r *http.Request) (request interface{}, err error) {

	var req postRequest

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal([]byte(body), &req); err != nil {
		return nil, err
	}
	req.Rate /= 10000

	return req, nil

	//amount, err := strconv.ParseFloat(req.Amount, 10)
	//if err != nil {
	//	return nil, err
	//}
	//businessType, err := strconv.ParseInt(req.BusinessType, 10, 64)
	//if err != nil {
	//	return nil, err
	//}
	//rate, err := strconv.ParseFloat(req.Rate, 10)
	//if err != nil {
	//	return nil, err
	//}
	//cardId, err := strconv.ParseInt(req.CardId, 10, 64)
	//if err != nil {
	//	return nil, err
	//}

	//return postRequest{
	//	Amount:       amount,
	//	BusinessName: req.BusinessName,
	//	BusinessType: businessType,
	//	Rate:         rate / 10000,
	//	CardId:       cardId,
	//}, nil
}
