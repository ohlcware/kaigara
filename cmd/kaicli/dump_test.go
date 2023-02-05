package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ohlcware/kaigara/pkg/storage"
	"github.com/ohlcware/kaigara/utils/testenv"
)

var testdataPath = "../testdata/testenv.yml"

func TestKaidumpListAppNames(t *testing.T) {
	ss := testenv.GetTestStorage(testdataPath, conf)

	b := kaidumpRun(ss)
	assert.NotNil(t, b)

	appNames := strings.Split(conf.AppNames, ",")
	scopes := strings.Split(conf.Scopes, ",")
	if err := storage.CleanAll(ss, appNames, scopes); err != nil {
		panic(err)
	}
}
