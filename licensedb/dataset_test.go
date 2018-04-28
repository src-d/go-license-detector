package licensedb

import (
	"fmt"
	"os"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/src-d/go-license-detector.v2/licensedb/filer"
)

func TestDataset(t *testing.T) {
	rootFiler, err := filer.FromZIP("dataset.zip")
	assert.Nil(t, err)
	defer rootFiler.Close()
	projects, err := rootFiler.ReadDir("")
	assert.Nil(t, err)
	licenses := map[string]map[string]float32{}
	mutex := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(len(projects))
	for _, project := range projects {
		go func(project filer.File) {
			defer wg.Done()
			myLicenses, _ := Detect(filer.NestFiler(rootFiler, project.Name))
			if len(myLicenses) > 0 {
				mutex.Lock()
				licenses[project.Name] = myLicenses
				mutex.Unlock()
			}
		}(project)
	}
	wg.Wait()
	assert.True(t, len(licenses) >= 893)
	// the rest len(projects) - 902 do not contain any license information
	fmt.Printf("%d %d %d%%\n", len(licenses), 902, (100*len(licenses))/902)
	if os.Getenv("LICENSE_TEST_DEBUG") != "" {
		for _, project := range projects {
			if _, exists := licenses[project.Name]; !exists {
				println(project.Name)
			}
		}
	}
}
