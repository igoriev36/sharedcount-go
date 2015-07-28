package sharedcount

import (
	"github.com/franela/goreq"
)

const HOST_FREE = "free"
const HOST_PLUS = "plus"
const HOST_BUSINESS = "business"

const PATH_URL = "url"
const PATH_QUOTA = "quota"
const PATH_BULK = "bulk"
const PATH_USAGE = "usage"
const PATH_STATUS = "status"
const PATH_DOMAIN_WHITELIST = "domain_whitelist"

var SHAREDCOUNT_API_FORMAT string = "https://%s.sharedcount.com/%s"

type ErrorResponse struct {
	Error string `json:"error"`
	Type  string `json:"type"`
}

type UrlResponse struct {
	StumbleUpon uint32
	Facebook    struct {
		CommentsBoxCount uint32 `json:"commentsbox_count"`
		ClickCount       uint32 `json:"click_count"`
		TotalCount       uint32 `json:"total_count"`
		CommentCount     uint32 `json:"comment_count"`
		LikeCount        uint32 `json:"like_count"`
		ShareCount       uint32 `json:"share_count"`
	}
	GooglePlusOne uint32
	Twitter       uint32
	Pinterest     uint32
	LinkedIn      uint32
	ErrorResponse
}

type QuotaResponse struct {
	QuotaUsed      uint32 `json:"quota_used_today"`
	Plan           string `json:"plan"`
	QuotaRemaining uint32 `json:"quota_remaining_today"`
	QuotaAllocated uint32 `json:"quota_allocated_today"`
	ErrorResponse
}

type StatusResponse struct {
	Status string `json:"status"`
	ErrorResponse
}

type BulkResponse struct {
	BulkId string `json:"bulk_count"`
	Count  uint32 `json:"count"`

	ErrorResponse
}

type SharedCount struct {
	Token string
	Plan  string
}

func New(token string, plan string) *SharedCount {
	return &SharedCount{
		Token: token,
		Plan:  plan,
	}
}

func (api *SharedCount) Quota() (response *QuotaResponse, error error) {
	res, err := goreq.Request()
}

func (api *SharedCount) Url(url string) (response *UrlResponse, error error) {

}

func (api *SharedCount) BulkPost(urls []string) (response *BulkResponse, error error) {

}

func (api *SharedCount) BulkGet(bulkId string) (response *BulkResponse, error error) {

}
