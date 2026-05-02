// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"net/http"

	"github.com/boltz-bio/boltz-api-go/internal/apijson"
	"github.com/boltz-bio/boltz-api-go/internal/requestconfig"
	"github.com/boltz-bio/boltz-api-go/option"
	"github.com/boltz-bio/boltz-api-go/packages/param"
	"github.com/boltz-bio/boltz-api-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type CursorPage[T any] struct {
	Data    []T    `json:"data"`
	FirstID string `json:"first_id" api:"nullable"`
	LastID  string `json:"last_id" api:"nullable"`
	HasMore bool   `json:"has_more"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		FirstID     respjson.Field
		LastID      respjson.Field
		HasMore     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r CursorPage[T]) RawJSON() string { return r.JSON.raw }
func (r *CursorPage[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *CursorPage[T]) GetNextPage() (res *CursorPage[T], err error) {
	if len(r.Data) == 0 {
		return nil, nil
	}

	if r.JSON.HasMore.Valid() && r.HasMore == false {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	if r.cfg.Request.URL.Query().Has("before_id") {
		next := r.FirstID
		if next == "" {
			return nil, nil
		}
		err = cfg.Apply(option.WithQuery("before_id", next))
		if err != nil {
			return nil, err
		}
	} else {
		next := r.LastID
		if next == "" {
			return nil, nil
		}
		err = cfg.Apply(option.WithQuery("after_id", next))
		if err != nil {
			return nil, err
		}
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *CursorPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &CursorPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type CursorPageAutoPager[T any] struct {
	page *CursorPage[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewCursorPageAutoPager[T any](page *CursorPage[T], err error) *CursorPageAutoPager[T] {
	return &CursorPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *CursorPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Data) == 0 {
		return false
	}
	if r.idx >= len(r.page.Data) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Data) == 0 {
			return false
		}
	}
	r.cur = r.page.Data[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *CursorPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *CursorPageAutoPager[T]) Err() error {
	return r.err
}

func (r *CursorPageAutoPager[T]) Index() int {
	return r.run
}

type OpaqueCursorPage[T any] struct {
	Data     []T    `json:"data"`
	NextPage string `json:"next_page" api:"nullable"`
	HasMore  bool   `json:"has_more"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		NextPage    respjson.Field
		HasMore     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r OpaqueCursorPage[T]) RawJSON() string { return r.JSON.raw }
func (r *OpaqueCursorPage[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *OpaqueCursorPage[T]) GetNextPage() (res *OpaqueCursorPage[T], err error) {
	if len(r.Data) == 0 {
		return nil, nil
	}

	if r.JSON.HasMore.Valid() && r.HasMore == false {
		return nil, nil
	}
	next := r.NextPage
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("page", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *OpaqueCursorPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &OpaqueCursorPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type OpaqueCursorPageAutoPager[T any] struct {
	page *OpaqueCursorPage[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewOpaqueCursorPageAutoPager[T any](page *OpaqueCursorPage[T], err error) *OpaqueCursorPageAutoPager[T] {
	return &OpaqueCursorPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *OpaqueCursorPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Data) == 0 {
		return false
	}
	if r.idx >= len(r.page.Data) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Data) == 0 {
			return false
		}
	}
	r.cur = r.page.Data[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *OpaqueCursorPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *OpaqueCursorPageAutoPager[T]) Err() error {
	return r.err
}

func (r *OpaqueCursorPageAutoPager[T]) Index() int {
	return r.run
}
