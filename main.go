// Copyright © 2019 Robert Sotomski <sotomski@gmail.com>
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

package main

import (
	"jira-cli/cmd"
)

func main() {

	//g := workflow.New(4)
	//g.AddEdge(0, 1)
	//g.AddEdge(0, 2)
	//g.AddEdge(1, 2)
	//g.AddEdge(2, 0)
	//g.AddEdge(2, 3)
	//g.AddEdge(3, 3)
	//g.BFS(2)
	cmd.Execute()
}
