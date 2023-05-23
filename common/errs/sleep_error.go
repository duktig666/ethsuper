//  Copyright © 2022-2023. duktig666 Limited.
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package errs

import (
	"fmt"
	"time"
)

// SleepError With time. Duration's error
// The most intuitive usage is: the length of time that can be specified when encountering Sleep
type SleepError struct {
	Err   error         `json:"-"`
	Msg   string        `json:"msg"`
	Sleep time.Duration `json:"sleep"`
}

func (e *SleepError) Error() string {
	return e.Err.Error()
}

func NewSleepError(msg string, sleep time.Duration) error {
	return &SleepError{
		Err:   fmt.Errorf(msg),
		Sleep: sleep,
		Msg:   msg,
	}
}
