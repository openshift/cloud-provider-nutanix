package provider

import (
	"fmt"
	"os"
	"time"

	"github.com/nutanix-cloud-native/cloud-provider-nutanix/internal/constants"
)

func GetCCMNamespace() (string, error) {
	ns := os.Getenv(constants.CCMNamespaceKey)
	if ns == "" {
		return "", fmt.Errorf("failed to retrieve CCM namespace. Make sure %s env variable is set", constants.CCMNamespaceKey)
	}
	return ns, nil
}

func NoResyncPeriodFunc() time.Duration {
	return 0
}
