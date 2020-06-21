package config

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockReader struct {
	mock.Mock
}

func (r *MockReader) Read() error {
	args := r.Called()

	return args.Error(0)
}

func (r *MockReader) Decode() (*Config, error) {
	args := r.Called()

	if c, ok := args.Get(0).(*Config); ok {
		return c, args.Error(1)
	}

	return nil, args.Error(1)
}

func TestReadConfig(t *testing.T) {
	assert := assert.New(t)

	mockObj := new(MockReader)
	mockObj.On("Read").Return(nil)
	mockObj.On("Decode").Return(&Config{Host: "localhost", Port: 5000, ChannelSecret: "1234", ChannelToken: "1234"}, nil)

	configService, err := ReadConfig(mockObj)
	config := configService.GetConfig()

	mockObj.AssertCalled(t, "Read")
	mockObj.AssertCalled(t, "Decode")

	assert.NoError(err)
	assert.Equal(config.Host, "localhost")
	assert.Equal(config.Port, 5000)
	assert.Equal(config.ChannelSecret, "1234")
	assert.Equal(config.ChannelToken, "1234")
}

func TestReadConfigReadFail(t *testing.T) {
	assert := assert.New(t)

	mockObj := new(MockReader)
	mockObj.On("Read").Return(fmt.Errorf("Error Read"))
	mockObj.On("Decode").Return(nil, fmt.Errorf("Error Decode"))

	configService, err := ReadConfig(mockObj)
	config := configService.GetConfig()

	mockObj.AssertCalled(t, "Read")
	mockObj.AssertNotCalled(t, "Decode")

	assert.Error(err, "Error Read")
	assert.Nil(config)
}

func TestReadConfigDecodeFail(t *testing.T) {
	assert := assert.New(t)

	mockObj := new(MockReader)
	mockObj.On("Read").Return(nil)
	mockObj.On("Decode").Return(nil, fmt.Errorf("Error Decode"))

	configService, err := ReadConfig(mockObj)
	config := configService.GetConfig()

	mockObj.AssertCalled(t, "Read")
	mockObj.AssertCalled(t, "Decode")

	assert.Error(err, "Error Decode")
	assert.Nil(config)
}

func TestConfigHook(t *testing.T) {
	assert := assert.New(t)

	input, expectedOutput := "5", 5
	output, _ := configHook(reflect.TypeOf(input), reflect.TypeOf(expectedOutput), input)

	assert.Equal(output, expectedOutput)
}

func TestConfigHookNotChanged(t *testing.T) {
	assert := assert.New(t)

	input, expectedOutput := "5", "5"
	output, _ := configHook(reflect.TypeOf(input), reflect.TypeOf(expectedOutput), input)

	assert.Equal(output, "5")
}
