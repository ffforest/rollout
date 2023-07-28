package builder

import (
	"fmt"

	"github.com/google/uuid"
)

const (
	DefaultNamespace = "default"
	DefaultName      = "test"
	DefaultAppName   = "app"
	DefaultCluster   = "cluster"
)

// builder is a base builder for resource
type builder struct {
	namespace  string
	name       string
	namePrefix string
	appName    string
	cluster    string
}

// complete sets default values
func (b *builder) complete() {
	b.buildName()

	if b.namespace == "" {
		b.namespace = DefaultNamespace
	}

	if b.appName == "" {
		b.appName = DefaultAppName
	}
	if b.cluster == "" {
		b.cluster = DefaultCluster
	}
}

// buildName generates the name of resource
func (b *builder) buildName() *builder {
	if b.name != "" {
		return b
	}
	if b.namePrefix != "" {
		b.name = randomName(b.namePrefix)
		return b
	}
	b.name = randomName(DefaultName)
	return b
}

// randomName generates a random name with prefix
func randomName(prefix string) string {
	return fmt.Sprintf("%s-%s", prefix, uuid.New().String())
}
