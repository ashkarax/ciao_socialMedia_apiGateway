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

// type PostData struct {
// 	UserId            uint   `json:"UserId"`
// 	UserName          string `json:"UserName"`
// 	UserProfileImgURL string `json:"UserProfileImgURL"`

// 	PostId  uint   `json:"PostId"`
// 	Caption string `json:"Caption"`

// 	PostAge       string   `json:"PostAge"`
// 	MediaUrl      []string `json:"MediaUrl"`
// 	LikeStatus    bool     `json:"LikeStatus" `
// 	LikesCount    string   `json:"LikesCount"`
// 	CommentsCount string   `json:"CommentsCount"`
// }
