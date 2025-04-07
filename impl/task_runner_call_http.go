// Copyright 2025 The Serverless Workflow Specification Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package impl

import (
	"fmt"

	"github.com/serverlessworkflow/sdk-go/v3/model"
)

type CallHTTPTaskRunner struct {
	TaskName string
}

func NewCallHttpRunner(taskName string, task *model.CallHTTP) (taskRunner *CallHTTPTaskRunner, err error) {
	if task == nil {
		err = model.NewErrValidation(fmt.Errorf("invalid For task %s", taskName), taskName)
	} else {
		taskRunner = new(CallHTTPTaskRunner)
		taskRunner.TaskName = taskName
	}
	return
}

func (f *CallHTTPTaskRunner) Run(input interface{}, taskSupport TaskSupport) (interface{}, error) {
	return input, nil

}

func (f *CallHTTPTaskRunner) GetTaskName() string {
	return f.TaskName
}
