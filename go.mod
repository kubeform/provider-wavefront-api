module kubeform.dev/provider-wavefront-api

go 1.15

require (
	github.com/json-iterator/go v1.1.11
	github.com/modern-go/reflect2 v1.0.1
	k8s.io/api v0.21.1
	k8s.io/apimachinery v0.21.1
	k8s.io/client-go v0.21.1
	kmodules.xyz/client-go v0.0.0-20210822203828-5e9cebbf1dfa
	kubeform.dev/apimachinery v0.0.0-20210824104859-ba5604d5a1cc
	sigs.k8s.io/cli-utils v0.25.0
	sigs.k8s.io/controller-runtime v0.9.0
)

replace k8s.io/apimachinery => github.com/kmodules/apimachinery v0.21.0-rc.0.0.20210405112358-ad4c2289ba4c

replace github.com/json-iterator/go => github.com/gomodules/json-iterator v1.1.12-0.20210506053207-2a3ea71074bc
