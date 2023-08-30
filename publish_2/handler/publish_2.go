package handler

import (
	"context"
	pb "publish_2/proto"
)

type Publish2 struct {
	auth   *Authentication
	mapper *Mapper
}

func NewPublish2() *Publish2 {
	return &Publish2{
		auth:   NewAuthentication(),
		mapper: NewMapper(),
	}
}

func (p *Publish2) PublishAction(ctx context.Context, req pb.DouyinPublishActionRequest, rsp pb.DouyinPublishActionResponse) error {

}

func (p *Publish2) PublishList(ctx context.Context, req pb.DouyinPublishListRequest, rsp pb.DouyinPublishListResponse) error {

}
