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
	"github.com/lastbackend/lastbackend/pkg/daemon/container/views/v1"
	"time"
)

type Pod struct {
	// Pod Meta
	Meta PodMeta `json:"meta"`
	// Container spec
	Spec PodSpec `json:"spec"`
	// Pod state
	State PodState `json:"state"`
}

type PodInfo struct {
	// Pod Meta
	Meta PodMeta `json:"meta"`
	// Pod state
	State PodState `json:"state"`
	// Pod containers
	Containers []v1.Container `json:"containers"`
}

type PodState struct {
	// Pod current state
	State string `json:"state"`
	// Pod current status
	Status string `json:"status"`
}

type PodMeta struct {
	// Meta id
	ID string `json:"id"`
	// Meta labels
	Labels map[string]string `json:"lables"`
	// Meta created time
	Created time.Time `json:"created"`
	// Meta updated time
	Updated time.Time `json:"updated"`
}

type PodSpec struct {
	// Provision ID
	ID string `json:"id"`
	// Provision state
	State string `json:"state"`
	// Provision status
	Status string `json:"status"`

	// Containers spec for pod
	Containers []v1.ContainerSpec `json:"containers"`

	// Provision create time
	Created time.Time `json:"created"`
	// Provision update time
	Updated time.Time `json:"updated"`
}
