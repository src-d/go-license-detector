package licensedb

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDataset(t *testing.T) {
	zipfile, err := zip.OpenReader("dataset.zip")
	assert.Nil(t, err)
	defer zipfile.Close()
	projects := map[string][]*zip.File{}
	for _, f := range zipfile.File {
		path := strings.Split(f.Name, "/")
		if path[1] != "" {
			files := projects[path[0]]
			if files == nil {
				files = []*zip.File{}
			}
			files = append(files, f)
			projects[path[0]] = files
		}
	}
	licenses := map[string]map[string]float32{}
	mutex := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(len(projects))
	for project, files := range projects {
		go func(project string, files []*zip.File) {
			defer wg.Done()
			myFilesList := make([]string, 0, len(files))
			myFilesMap := map[string]*zip.File{}
			for _, f := range files {
				name := f.Name[strings.Index(f.Name, "/")+1:]
				myFilesList = append(myFilesList, name)
				myFilesMap[name] = f
			}
			myLicenses, _ := InvestigateFilesLicenses(myFilesList, func(name string) (string, error) {
				reader, err := myFilesMap[name].Open()
				if err != nil {
					return "", err
				}
				defer reader.Close()
				bytes, err := ioutil.ReadAll(reader)
				if err != nil {
					return "", err
				}
				return string(bytes), nil
			})
			if len(myLicenses) > 0 {
				mutex.Lock()
				licenses[project] = myLicenses
				mutex.Unlock()
			}
		}(project, files)
	}
	wg.Wait()
	assert.True(t, len(licenses) >= 817)
	// the rest len(projects) - 902 do not contain any license information
	fmt.Printf("%d %d %d%%\n", len(licenses), 902, (100 * len(licenses)) / 902)
}
