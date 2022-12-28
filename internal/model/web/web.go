// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"github.com/rocboss/paopao-ce/internal/core"
)

type BaseInfo struct {
	User *core.User
}

type SimpleInfo struct {
	Uid int64
}

func (b *BaseInfo) SetUser(user *core.User) {
	b.User = user
}

func (s *SimpleInfo) SetUserId(id int64) {
	s.Uid = id
}
