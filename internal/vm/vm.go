/*
 * Copyright 2020. The Alkaid Authors. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 *
 * Alkaid is a BaaS service based on Hyperledger Fabric.
 *
 */

package vm

type CreateRequest struct {
	ContainerName  string
	ImageName      string
	ImageTag       string
	Environment    []string
	NetworkMode    string
	NetworkAliases []string
	Mounts         map[string]string
	Files          map[string][]byte
	WorkingDir     string
	Command        []string
}
