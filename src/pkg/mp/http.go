/**
 * @Time : 3/30/21 5:27 PM
 * @Author : solacowa@gmail.com
 * @File : http
 * @Software: GoLand
 */

package mp

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/nsini/cardbill/src/encode"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func MakeHTTPHandler(s Service, dmw []endpoint.Middleware, opts []kithttp.ServerOption) http.Handler {
	ems := []endpoint.Middleware{}

	ems = append(ems, dmw...)

	eps := NewEndpoint(s, map[string][]endpoint.Middleware{
		//"RecentRepay": ems,
		"BankList": ems,
		//"CreditCards": ems,
		//"Record": ems,
		//"BusinessTypes": ems,
		//"Statistics": ems,
	})

	r := mux.NewRouter()

	r.Handle("/recent-repay", kithttp.NewServer(
		eps.RecentRepayEndpoint,
		decodeRecentRepayRequest,
		encode.JsonResponse,
		opts...,
	)).Methods(http.MethodGet)
	r.Handle("/recent-repay", kithttp.NewServer(
		eps.RecentRepayEndpoint,
		decodeRecentRepayRequest,
		encode.JsonResponse,
		opts...,
	)).Methods(http.MethodGet)
	r.Handle("/credit-cards", kithttp.NewServer(
		eps.CreditCardsEndpoint,
		kithttp.NopRequestDecoder,
		encode.JsonResponse,
		opts...,
	)).Methods(http.MethodGet)
	r.Handle("/statistics", kithttp.NewServer(
		eps.StatisticsEndpoint,
		kithttp.NopRequestDecoder,
		encode.JsonResponse,
		opts...,
	)).Methods(http.MethodGet)
	r.Handle("/business-types", kithttp.NewServer(
		eps.BusinessTypesEndpoint,
		kithttp.NopRequestDecoder,
		encode.JsonResponse,
		opts...,
	)).Methods(http.MethodGet)
	r.Handle("/record", kithttp.NewServer(
		eps.RecordEndpoint,
		decodeMpRecordRequest,
		encode.JsonResponse,
		opts...,
	)).Methods(http.MethodGet)
	r.Handle("/record/{id:[0-9]+}", kithttp.NewServer(
		eps.RecordDetailEndpoint,
		decodeRecordDetailRequest,
		encode.JsonResponse,
		opts...,
	)).Methods(http.MethodGet)
	r.Handle("/record", kithttp.NewServer(
		eps.RecordAddEndpoint,
		decodeRecordAddRequest,
		encode.JsonResponse,
		opts...,
	)).Methods(http.MethodPost)
	r.Handle("/banks", kithttp.NewServer(
		eps.BankListEndpoint,
		decodeBankListRequest,
		encode.JsonResponse,
		opts...,
	)).Methods(http.MethodGet)
	r.Handle("/login", kithttp.NewServer(
		eps.LoginEndpoint,
		decodeMpLoginRequest,
		encode.JsonResponse,
		opts...,
	)).Methods(http.MethodPost)
	r.Handle("/make-token", kithttp.NewServer(
		eps.MakeTokenEndpoint,
		decodeMpMakeTokenRequest,
		encode.JsonResponse,
		opts...,
	)).Methods(http.MethodPost)

	return r
}

func decodeRecordDetailRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, encode.InvalidParams.Error()
	}
	var req recordDetailRequest
	recordId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, encode.InvalidParams.Wrap(err)
	}

	req.Id = recordId
	return req, nil
}

func decodeRecordAddRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req recordAddRequest

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(body, &req); err != nil {
		return nil, err
	}
	req.Rate /= 100

	if req.TmpTime != "" {
		if t, err := time.Parse("2006-01-02 15:04:05", req.TmpTime); err == nil {
			tt := t.Local()
			req.SwipeTime = &tt
		} else {
			return nil, err
		}
	}
	return req, nil
}

func decodeMpRecordRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req recordRequest
	if cardId, err := strconv.ParseInt(r.URL.Query().Get("cardId"), 10, 64); err == nil {
		req.CardId = cardId
	}
	if bankId, err := strconv.ParseInt(r.URL.Query().Get("bankId"), 10, 64); err == nil {
		req.BankId = bankId
	}
	if p, err := strconv.Atoi(r.URL.Query().Get("page")); err == nil {
		req.Page = p
	} else {
		req.Page = 1
	}
	if p, err := strconv.Atoi(r.URL.Query().Get("pageSize")); err == nil {
		req.PageSize = p
	} else {
		req.PageSize = 10
	}

	start := r.URL.Query().Get("start")
	end := r.URL.Query().Get("end")
	if t, err := time.Parse("2006-01-02", start); err == nil {
		req.Start = &t
	}

	if t, err := time.Parse("2006-01-02", end); err == nil {
		req.End = &t
	}

	return req, nil
}

func decodeMpMakeTokenRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req makeTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, encode.InvalidParams.Wrap(err)
	}
	return req, nil
}

func decodeMpLoginRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req mpLoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, encode.InvalidParams.Wrap(err)
	}
	if req.Code == "" {
		return nil, encode.ErrAuthMPLoginCode.Error()
	}
	return req, nil
}

func decodeBankListRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req bankRequest
	req.bankName = r.URL.Query().Get("bankName")

	return req, nil
}

func decodeRecentRepayRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req recentRepayRequest
	recent, _ := strconv.Atoi(r.URL.Query().Get("recent"))
	if recent <= 0 {
		recent = 10
	}
	req.recent = recent

	return req, nil
}
