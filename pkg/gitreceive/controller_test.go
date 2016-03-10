package gitreceive

import (
	"testing"

	"github.com/arschles/assert"
	"github.com/deis/builder/pkg"
	"github.com/deis/builder/pkg/gitreceive/git"
	"github.com/deis/builder/pkg/storage"
)

const (
	rawSha   = "c3b4e4ba8b7267226ff02ad07a3a2cca9c9237de"
	bucket   = "git"
	appName  = "myapp"
	slugName = "myslug"
	username = "myuser"
)

func TestCreateBuildHook(t *testing.T) {
	procType := pkg.ProcessType(make(map[string]string))
	sha, err := git.NewSha(rawSha)
	assert.NoErr(t, err)
	endpoint := &storage.Endpoint{URLStr: "s3.amazonaws.com", Secure: false}
	slugBuilderInfo := NewSlugBuilderInfo(endpoint, bucket, appName, slugName, sha)
	hookUsingDockerfile := createBuildHook(slugBuilderInfo, sha, username, appName, procType, true)
	assert.Equal(t, hookUsingDockerfile.Sha, sha.Short(), "git sha")
	assert.Equal(t, hookUsingDockerfile.ReceiveUser, username, "username")
	assert.Equal(t, hookUsingDockerfile.ReceiveRepo, appName, "username")
	assert.Equal(t, hookUsingDockerfile.Image, appName, "image")
	assert.Equal(t, hookUsingDockerfile.Procfile, procType, "procfile")
	assert.Equal(t, hookUsingDockerfile.Dockerfile, "true", "dockerfile field")

	hookNoDockerfile := createBuildHook(slugBuilderInfo, sha, username, appName, procType, false)
	assert.Equal(t, hookNoDockerfile.Sha, sha.Short(), "git sha")
	assert.Equal(t, hookNoDockerfile.ReceiveUser, username, "username")
	assert.Equal(t, hookNoDockerfile.ReceiveRepo, appName, "username")
	assert.Equal(t, hookNoDockerfile.Image, slugBuilderInfo.AbsoluteSlugURL(), "image")
	assert.Equal(t, hookNoDockerfile.Procfile, procType, "procfile")
	assert.Equal(t, hookNoDockerfile.Dockerfile, "", "dockerfile field")

}
