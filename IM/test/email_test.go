package test

import (
	"IM/helper"
	"testing"
)

func TestEmail(t *testing.T){
	helper.MailSendCode("2960411617@qq.com","你好,世界!")
}
