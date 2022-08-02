module github.com/why19970628/k8s-cheatsheet

go 1.13

require (
	gopkg.in/yaml.v2 v2.4.0 // indirect
	k8s.io/api v0.19.1
	k8s.io/apimachinery v0.19.1
	k8s.io/client-go v11.0.0+incompatible
)

replace k8s.io/client-go => k8s.io/client-go v0.19.1
