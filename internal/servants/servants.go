// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package servants

import (
	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/internal/servants/docs"
	"github.com/rocboss/paopao-ce/internal/servants/localoss"
	"github.com/rocboss/paopao-ce/internal/servants/statick"
	"github.com/rocboss/paopao-ce/internal/servants/web"
	"github.com/rocboss/paopao-ce/pkg/cfg"
)

// RegisterWebServants register all the servants to gin.Engine
func RegisterWebServants(e *gin.Engine) {
	docs.RegisterDocs(e)
	statick.RegisterStatick(e)

	cfg.Be("LocalOSS", func() {
		localoss.RouteLocalOSS(e)
	})

	web.RouteWeb(e)
}
