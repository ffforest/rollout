/*
 * Copyright 2023 The KusionStack Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// RandStr returns a random string
func RandStr() string {
	randNum := rand.Intn(999)
	if randNum < 100 {
		randNum += 100
	}
	return fmt.Sprintf("%s%d", time.Now().Format("20060102150405"), randNum)
}
