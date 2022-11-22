package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type IntegrationTestSettings struct {
	PostgresUser, PostgresPassword, PostgresDbName, PostgresHost string
	PostgresPort                                                 int
	MySqlUser, MySqlPassword, MySqlDB                            string
}

func ReadIntegrationTestSettings(contextFile string) IntegrationTestSettings {
	jsonFile, err := os.Open(contextFile)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var integrationTestSettings IntegrationTestSettings

	err = json.Unmarshal(byteValue, &integrationTestSettings)
	if err != nil {
		return integrationTestSettings
	}

	return integrationTestSettings
}
