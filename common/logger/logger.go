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

// description:
// @author renshiwei
// Date: 2022/10/5 19:39

package logger

import "go.uber.org/zap"

func Panic(msg string, fields ...zap.Field) {
	zapLog.Panic(msg, fields...)
}

func DPanic(msg string, fields ...zap.Field) {
	zapLog.DPanic(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	zapLog.Debug(msg, fields...)
}

func Debugf(template string, args ...interface{}) {
	sugarLog.Debugf(template, args...)
}

func Info(msg string, fields ...zap.Field) {
	zapLog.Info(msg, fields...)
}

func Infof(template string, args ...interface{}) {
	sugarLog.Infof(template, args...)
}

func Warn(msg string, fields ...zap.Field) {
	zapLog.Warn(msg, fields...)
}

func Warnf(template string, args ...interface{}) {
	sugarLog.Warnf(template, args...)
}

func Error(msg string, fields ...zap.Field) {
	zapLog.Error(msg, fields...)
}

func Errorf(template string, args ...interface{}) {
	sugarLog.Errorf(template, args...)
}
