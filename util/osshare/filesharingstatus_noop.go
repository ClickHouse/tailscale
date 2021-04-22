// Copyright (c) 2021 Tailscale Inc & AUTHORS All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !windows

package osshare

import (
	"tailscale.com/types/logger"
)

func SetFileSharingEnabled(enabled bool, logf logger.Logf) {}
