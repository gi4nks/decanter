package main

import (
	"encoding/json"
	"io/ioutil"
)

type Configuration struct {
	RepositoryDirectory string
	RepositoryUrl       string
	RestUrl             string
	DebugMode           bool
}

type Settings struct {
	configs Configuration
}

func (sts *Settings) LoadSettings() {
	folder := executableFolder()

	file, err := ioutil.ReadFile(folder + "/conf.json")

	if err != nil {
		sts.configs = Configuration{}
		sts.configs.RepositoryDirectory = folder + "/" + ConstRepositoryDirectory
		sts.configs.RepositoryUrl = ConstRepositoryUrl
		sts.configs.RestUrl = ConstRestUrl
		sts.configs.DebugMode = ConstDebugMode

	} else {
		json.Unmarshal(file, &sts.configs)

		parrot.Debug("folder: " + folder)
		parrot.Debug("file: " + asJson(sts.configs))

	}
}

func (sts Settings) RepositoryDirectory() string {
	return sts.configs.RepositoryDirectory
}

func (sts Settings) RepositoryUrl() string {
	return sts.configs.RepositoryUrl
}

func (sts Settings) RestUrl() string {
	return sts.configs.RestUrl
}

func (sts Settings) DebugMode() bool {
	return sts.configs.DebugMode
}

func (sts Settings) String() string {
	b, err := json.Marshal(sts.configs)
	if err != nil {
		parrot.Error("Warning", err)
		return "{}"
	}
	return string(b)
}
