package responsemodels_postnrel_apigw

type AddPostResp struct {
	Caption string `json:"caption,omitempty"`
	UserId  string `json:"userid,omitempty"`

	Media string `json:"media,omitempty"`
}

type EditPostResp struct {
	Caption string `json:"caption"`
	PostId  string `json:"postid"`
	UserId  string `json:"userid" `
}
