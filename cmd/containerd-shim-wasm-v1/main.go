package main

/*
   Copyright The containerd Authors.

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

import (
	"github.com/containerd/containerd/runtime/v2/shim"
	wasm "github.com/denverdino/containerd-wasm"
	"net"
	"os/exec"
)

func main() {
	startLayotto()
	shim.Run("io.containerd.layotto.v2", wasm.New)
}

func startLayotto() {
	conn, err := net.Dial("tcp", "localhost:2045")
	if err == nil {
		conn.Close()
		return
	}
	cmd := exec.Command("layotto", "start", "-c", "/home/docker/config.json")
	cmd.Start()
}
