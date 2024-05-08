package requestmodels_postnrel_apigw

type CommentRequest struct {
	PostId          uint   `json:"PostId" validate:"required,number"`
	UserId          string `json:"UserId" validate:"required,number"`
	CommentText     string `json:"CommentText" validate:"required,gte=1,lte=30"`
	ParentCommentId uint   `json:"ParentCommentId" validate:"number"`
}

type CommentEditRequest struct {
	UserId      string `json:"UserId" validate:"required,number"`
	CommentId   uint   `json:"CommentId" validate:"required,number"`
	CommentText string `json:"CommentText" validate:"required,gte=1,lte=30"`
}
