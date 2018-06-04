// This file was automatically generated by informer-gen

package internalversion

import (
	internalinterfaces "github.com/openshift/origin/pkg/security/generated/informers/internalversion/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// SecurityContextConstraints returns a SecurityContextConstraintsInformer.
	SecurityContextConstraints() SecurityContextConstraintsInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// SecurityContextConstraints returns a SecurityContextConstraintsInformer.
func (v *version) SecurityContextConstraints() SecurityContextConstraintsInformer {
	return &securityContextConstraintsInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
