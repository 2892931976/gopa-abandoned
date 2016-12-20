/*
Copyright 2016 Medcl (m AT medcl.net)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package pipe

import (. "github.com/medcl/gopa/core/pipeline"
. "github.com/medcl/gopa/core/types"
	"github.com/medcl/gopa/core/tasks"
	log "github.com/cihub/seelog"
)

type End struct {
}

func (this End) Name() string {
	return "end"
}

func (this End) Process(context *Context) (*Context, error) {

	log.Trace("end process")

	task:=context.Get(CONTEXT_CRAWLER_TASK).(*CrawlerTask)

	if(context.IsBreak()){
		task.Message=context.Payload
	}
	task.Url=context.MustGetString(CONTEXT_URL)
	tasks.UpdateTask(task)

	return context, nil
}
