package constants

import "fwds/pkg/mq"

var (
	MemberCancellation = mq.NewBusiness("用户注销", "member_cancellation", "member_cancellation", "member_cancellation")
)
