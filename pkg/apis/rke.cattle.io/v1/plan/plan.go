package plan

import (
	rkev1 "github.com/rancher/rancher-operator/pkg/apis/rke.cattle.io/v1"
	capi "sigs.k8s.io/cluster-api/api/v1alpha4"
)

type Plan struct {
	Nodes    map[string]*Node         `json:"nodes,omitempty"`
	Machines map[string]*capi.Machine `json:"machines,omitempty"`
	Cluster  *rkev1.RKECluster        `json:"cluster,omitempty"`
}

type Node struct {
	Plan        NodePlan  `json:"plan,omitempty"`
	AppliedPlan *NodePlan `json:"appliedPlan,omitempty"`
	InSync      bool      `json:"inSync,omitempty"`
}

type Secret struct {
	ServerToken string `json:"serverToken,omitempty"`
	AgentToken  string `json:"agentToken,omitempty"`
}

type Instruction struct {
	Name    string   `json:"name,omitempty"`
	Image   string   `json:"image,omitempty"`
	Env     []string `json:"env,omitempty"`
	Args    []string `json:"args,omitempty"`
	Command string   `json:"command,omitempty"`
}

// Name would be `ca.pem`, Path would be `/etc/kubernetes/ssl`, Contents is base64 encoded
type File struct {
	Content string `json:"content,omitempty"`
	Name    string `json:"name,omitempty"`
	Path    string `json:"path,omitempty"`
}

type NodePlan struct {
	Files        []File        `json:"files,omitempty"`
	Instructions []Instruction `json:"instructions,omitempty"`
}
