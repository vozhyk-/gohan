// Copyright (C) 2015 NTT Innovation Institute, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package framework

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/cloudwan/gohan/extension/framework/buflog"
	"github.com/cloudwan/gohan/extension/framework/runner"
	"github.com/codegangsta/cli"
	logging "github.com/op/go-logging"
)

var log = logging.MustGetLogger("extest")

// RunTests runs extension tests when invoked from Gohan CLI
func RunTests(c *cli.Context) {
	buflog.SetUpDefaultLogging()

	testFiles := getTestFiles(c.Args())
	summary := map[string]error{}
	for _, testFile := range testFiles {
		testRunner := runner.NewTestRunner(testFile, c.Bool("verbose"))
		errors := testRunner.Run()
		if err, ok := errors[runner.GeneralError]; ok {
			summary[testFile] = fmt.Errorf("%s", err.Error())
			log.Error(fmt.Sprintf("Error: %s", err.Error()))
			continue
		}

		failed := 0
		for _, err := range errors {
			if err != nil {
				failed++
			}
		}
		summary[testFile] = nil
		if failed > 0 {
			summary[testFile] = fmt.Errorf("%d/%d tests failed", failed, len(errors))
		}
	}

	returnCode := 0
	log.Info("Run %d test files:", len(summary))
	for testFile, err := range summary {
		if err != nil {
			returnCode = 1
			log.Error(fmt.Sprintf("Failure in %s: %s", testFile, err.Error()))
		} else {
			log.Notice("OK %s ", testFile)
		}
	}
	os.Exit(returnCode)
}

func getTestFiles(args cli.Args) []string {
	paths := args
	if len(paths) == 0 {
		paths = append(paths, ".")
	}

	pattern := regexp.MustCompile(`^test_.*\.js$`)
	seen := map[string]bool{}
	testFiles := []string{}
	for _, path := range paths {
		filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
			if err != nil {
				log.Error(fmt.Sprintf("Failed to process '%s': %s", filePath, err.Error()))
				return nil
			}

			if info.IsDir() {
				return nil
			}

			if !pattern.MatchString(info.Name()) {
				return nil
			}

			filePath = filepath.Clean(filePath)

			if !seen[filePath] {
				testFiles = append(testFiles, filePath)
				seen[filePath] = true
			}

			return nil
		})
	}

	return testFiles
}
