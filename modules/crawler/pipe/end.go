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

import (
	log "github.com/cihub/seelog"
	"github.com/medcl/gopa/core/errors"
	"github.com/medcl/gopa/core/model"
	. "github.com/medcl/gopa/core/pipeline"
	"github.com/medcl/gopa/core/util"
	"github.com/medcl/gopa/modules/config"
)

const SaveTask JointKey = "save_task"

type SaveTaskJoint struct {
	IsCreate bool
}

func (this SaveTaskJoint) Name() string {
	return string(SaveTask)
}

func (this SaveTaskJoint) Process(context *Context) (*Context, error) {

	log.Trace("end process")
	if context.IsErrorExit() {
		return context, errors.NewWithCode(config.ErrorExitedPipeline, "exited pipeline")
	}

	task := context.Get(CONTEXT_CRAWLER_TASK).(*model.Task)
	task.Status = model.TaskFetchSuccess
	task.Phrase = context.Phrase

	if context.IsBreak() {
		log.Trace("broken pipeline,", context.Payload)
		task.Status = model.TaskFetchFailed
		task.Message = context.Payload
	}

	//update url
	task.Url = context.MustGetString(CONTEXT_URL)
	pageItem := context.Get(CONTEXT_PAGE_ITEM)

	if pageItem != nil {
		task.Page = pageItem.(*model.PageItem)
		meta, b := context.GetMap(CONTEXT_PAGE_METADATA)
		if b {
			task.Page.Metadata = &meta
		}

		text, b := context.GetString(CONTEXT_PAGE_BODY_PLAIN_TEXT)
		if b {
			task.Page.Text = text
		}
	}

	if this.IsCreate {
		log.Trace("create task, url:", task.Url)
		if task.Url == "http://plusx.cn" {
			log.Error(util.ToJson(this.Name(), true))
			log.Error(util.ToJson(context, true))
		}
		model.CreateTask(task)
	} else {
		model.UpdateTask(task)
	}

	return context, nil
}
