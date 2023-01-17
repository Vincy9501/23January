package unit

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

//func TestHelloTom(t *testing.T) {
//	output := HelloTom()
//	expectOutput := "Tom"
//	if output != expectOutput {
//		t.Errorf("Expected %s do not match actual %s", expectOutput, output)
//	}
//}

func TestHelloTome(t *testing.T) {
	output := HelloTom()
	expectOutput := "Tom"
	assert.Equal(t, expectOutput, output)
}
