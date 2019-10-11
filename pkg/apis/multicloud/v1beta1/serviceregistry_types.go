// Package v1beta1 of apis contain the API type definition for the components
// IBM Confidential
// OCO Source Materials
// 5737-E67
// (C) Copyright IBM Corporation 2019 All Rights Reserved
// The source code for this program is not published or otherwise divested of its trade secrets, irrespective of what has been deposited with the U.S. Copyright Office.
package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.ibm.com/IBMPrivateCloud/ibm-klusterlet-operator/pkg/image"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ServiceRegistrySpec defines the desired state of ServiceRegistry
// +k8s:openapi-gen=true
type ServiceRegistrySpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	ClusterName      string               `json:"clusterName"`
	ClusterNamespace string               `json:"clusterNamespace"`
	FullNameOverride string               `json:"fullnameOverride"`
	ServiceRegistry  ServiceRegistryImage `json:"serviceRegistry"`
	CoreDNS          CoreDNSImage         `json:"coredns"`
	ImagePullSecret  string               `json:"imagePullSecret,omitempty"`

	ConnectionManager                  string `json:"connectionManager"`
	DNSSuffix                          string `json:"dnsSuffix,omitempty"`
	Plugins                            string `json:"plugins,omitempty"`
	IstioIngressGateway                string `json:"istioIngressGateway,omitempty"`
	IstioServiceEntryRegistryNamespace string `json:"istioServiceEntryRegistryNamespace,omitempty"`
}

// ServiceRegistryImage defines service registry configuration in the spec
// +k8s:openapi-gen=true
type ServiceRegistryImage struct {
	Image image.Image `json:"image"`
}

// CoreDNSImage defines CoreDNSImage configuration in the spec
// +k8s:openapi-gen=true
type CoreDNSImage struct {
	Image image.Image `json:"image,omitempty"`
}

// ServiceRegistryStatus defines the observed state of ServiceRegistry
// +k8s:openapi-gen=true
type ServiceRegistryStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceRegistry is the Schema for the serviceregistries API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type ServiceRegistry struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServiceRegistrySpec   `json:"spec,omitempty"`
	Status ServiceRegistryStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceRegistryList contains a list of ServiceRegistry
type ServiceRegistryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceRegistry `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ServiceRegistry{}, &ServiceRegistryList{})
}
