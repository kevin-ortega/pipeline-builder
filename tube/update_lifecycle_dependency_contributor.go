/*
 * Copyright 2018-2020 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package tube

type UpdateLifecycleDependencyContributor struct {
	Descriptor Descriptor
	Salt       string
}

func (UpdateLifecycleDependencyContributor) Group() string {
	return "builder-dependencies"
}

func (u UpdateLifecycleDependencyContributor) Job() Job {
	b := NewBuildCommonResource()
	l := NewLifecycleResource()
	s := NewSourceResource(u.Descriptor, u.Salt)

	return Job{
		Name:   "update-lifecycle",
		Public: true,
		Plan: []map[string]interface{}{
			{
				"in_parallel": []map[string]interface{}{
					{
						"get":      "build-common",
						"resource": b.Name,
					},
					{
						"get":     l.Name,
						"trigger": true,
						"params": map[string]interface{}{
							"globs": []string{""},
						},
					},
					{
						"get":      "source",
						"resource": s.Name,
					},
				},
			},
			{
				"task": "update-lifecycle-dependency",
				"file": "build-common/update-lifecycle-dependency.yml",
			},
			{
				"put": s.Name,
				"params": map[string]interface{}{
					"repository": "source",
					"rebase":     true,
				},
			},
		},
	}

}

func (u UpdateLifecycleDependencyContributor) Resources() []Resource {
	return []Resource{
		NewBuildCommonResource(),
		NewLifecycleResource(),
		NewSourceResource(u.Descriptor, u.Salt),
	}
}
