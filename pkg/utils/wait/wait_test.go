/*
 * Tencent is pleased to support the open source community by making TKEStack available.
 *
 * Copyright (C) 2012-2019 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

package wait

import (
	"context"
	"github.com/sirupsen/logrus"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunUntil(t *testing.T) {
	called := false
	ctx, cancel := context.WithCancel(context.Background())
	_ = RunUntil(ctx, logrus.New(), 0, func() error {
		called = true
		cancel()
		return nil
	})
	require.True(t, called)
}
