// Copyright 2023 Google Inc. All Rights Reserved.
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
	"flag"
	"fmt"
	"path"
	"syscall"

	"github.com/docker/docker/pkg/system"
	"github.com/golang/glog"
)

const (
	devDirectory = "/dev"

	nvidiaCtlDevice      = "nvidiactl"
	nvidiaUVMDevice      = "nvidia-uvm"
	nvidiaUVMToolsDevice = "nvidia-uvm-tools"
	nvidiaModesetDevice  = "nvidia-modeset"
)

var (
	deviceCount = flag.Int("device-count", 1, "The fake devices count for this node.")
)

func main() {
	flag.Parse()
	glog.Infoln("fake device generator starts")
	if err := system.Mknod(path.Join(devDirectory, nvidiaCtlDevice), syscall.S_IFCHR, int(system.Mkdev(195, 255))); err != nil {
		glog.Error(err)
	}
	if err := system.Mknod(path.Join(devDirectory, nvidiaUVMDevice), syscall.S_IFCHR, int(system.Mkdev(242, 0))); err != nil {
		glog.Error(err)
	}
	if err := system.Mknod(path.Join(devDirectory, nvidiaUVMToolsDevice), syscall.S_IFCHR, int(system.Mkdev(242, 1))); err != nil {
		glog.Error(err)
	}
	if err := system.Mknod(path.Join(devDirectory, nvidiaModesetDevice), syscall.S_IFCHR, int(system.Mkdev(195, 254))); err != nil {
		glog.Error(err)
	}

	for i := 0; i < *deviceCount; i++ {
		if err := system.Mknod(path.Join(devDirectory, fmt.Sprintf("nvidia%d", i)), syscall.S_IFCHR, int(system.Mkdev(195, int64(i)))); err != nil {
			glog.Error(err)
		}
	}
	glog.Infoln("fake device generator completes")
}
