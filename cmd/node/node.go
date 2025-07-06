/*
Copyright 2024 YANDEX LLC

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

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/yandex-cloud/yc-csi-driver/cmd"
	"github.com/yandex-cloud/yc-csi-driver/pkg/server/node"
	"github.com/yandex-cloud/yc-csi-driver/pkg/version"

	"k8s.io/klog"
)

func main() {
	var showVersion bool

	opts := node.DefaultOptions()

	flag.BoolVar(&showVersion, "version", false, "show version")
	flag.StringVar(&opts.Endpoint, "endpoint", opts.Endpoint, "CSI Endpoint")
	flag.StringVar(&opts.DriverName, "driver-name", opts.DriverName, "driver name to identify as for the CSI")
	flag.IntVar(&opts.MaxVolumesPerNode, "max-volumes", opts.MaxVolumesPerNode, "the maximum number of volumes attachable to a node")
	klog.InitFlags(nil)

	flag.Parse()

	if showVersion {
		fmt.Println(version.Version())
		os.Exit(0)
	}

	srv, err := node.New(opts)
	if err != nil {
		klog.Fatal(err.Error())
		return
	}

	cmd.Execute(srv)
}
