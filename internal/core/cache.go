package core

const (
	IdxActNop IdxAct = iota + 1
	IdxActCreatePost
	IdxActUpdatePost
	IdxActDeletePost
	IdxActStickPost
	IdxActVisiblePost
)

type IdxAct uint8

type IndexAction struct {
	Act IdxAct
}

func (a IdxAct) String() string {
	switch a {
	case IdxActNop:
		return "no operator"
	case IdxActCreatePost:
		return "create post"
	case IdxActUpdatePost:
		return "update post"
	case IdxActDeletePost:
		return "delete post"
	case IdxActStickPost:
		return "stick post"
	case IdxActVisiblePost:
		return "visible post"
	default:
		return "unknow action"
	}
}
