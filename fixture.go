package commandgo

import (
	"log"
	"io/ioutil"
	"os"
)

type Fixture struct {
	HandlesErrors
	Prefix       string
	name         string
	file         *os.File
}

func NewFixture(prefix string) *Fixture {
	fixture := &Fixture{
		HandlesErrors: NewHandlesErrors(),
		Prefix: prefix,
	}

	fixture.build()

	return fixture
}

func (fixture *Fixture) build() {
	var tempfileError error
	fixture.file, tempfileError = ioutil.TempFile("", fixture.Prefix)

	if (tempfileError != nil) {
		fixture.ErrorHandler(tempfileError)
	}

	fixture.name = fixture.file.Name()
}

func (fixture *Fixture) Write(bytes []byte) *Fixture {
	_, writeError := fixture.file.Write(bytes)

	if (writeError != nil) {
		log.Panic(writeError)
	}

	return fixture

}

func (fixture *Fixture) Close() *Fixture {
	err := fixture.file.Close()
	if (err != nil) {
		log.Panic(err)
	}

	return fixture
}

func (fixture *Fixture) Name() string {
	return fixture.name
}
