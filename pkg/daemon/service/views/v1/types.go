//
// Last.Backend LLC CONFIDENTIAL
// __________________
//
// [2014] - [2017] Last.Backend LLC
// All Rights Reserved.
//
// NOTICE:  All information contained herein is, and remains
// the property of Last.Backend LLC and its suppliers,
// if any.  The intellectual and technical concepts contained
// herein are proprietary to Last.Backend LLC
// and its suppliers and may be covered by Russian Federation and Foreign Patents,
// patents in process, and are protected by trade secret or copyright law.
// Dissemination of this information or reproduction of this material
// is strictly forbidden unless prior written permission is obtained
// from Last.Backend LLC.
//

package v1

import (
	"time"
	"github.com/lastbackend/lastbackend/pkg/daemon/pod/views/v1"
)

type Service struct {
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Created     time.Time    `json:"created"`
	Updated     time.Time    `json:"updated"`
	Pods        []v1.PodInfo `json:"pods"`
	Config      Config       `json:"config"`
	Region      string       `json:"region,omitempty"`
}

type Config struct {
	Replicas int    `json:"replicas,omitempty"`
	Memory   int64  `json:"memory,omitempty"`
	Command  string `json:"command,omitempty"`
	Image    string `json:"image,omitempty"`
	Region   string `json:"region,omitempty"`
}

type ServiceList []Service
