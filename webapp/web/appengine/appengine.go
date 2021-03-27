package appengine

import (
	"os"
	"strings"
	"time"
)

var ProjectID = func() string {
	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	if projectID == "" {
		return "local"
	}
	return projectID
}()

var VersionID = func() string {
	versionID := os.Getenv("GAE_VERSION")
	if versionID == "" {
		return time.Now().Format("20060102t150405")
	}
	return versionID
}()

func IsDevAppServer() bool {
	return ProjectID == "local"
}

// IsProd ...
func IsProd() bool {
	return strings.HasSuffix(ProjectID, "-prod")
}
