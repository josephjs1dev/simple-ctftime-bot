package config_test

import (
	"fmt"
	"testing"

	"github.com/josephsalimin/simple-ctftime-bot/internal/config"
	mockconfig "github.com/josephsalimin/simple-ctftime-bot/internal/mock/config"
	"github.com/stretchr/testify/assert"
)

func TestReadConfig(t *testing.T) {
	t.Run("Successfully read config", func(t *testing.T) {
		assert := assert.New(t)

		mockObj := new(mockconfig.MockReader)
		mockObj.On("Read").Return(nil)
		mockObj.On("Decode").Return(&config.Config{Host: "localhost", Port: 5000, ChannelSecret: "1234", ChannelToken: "1234"}, nil)

		config, err := config.ReadConfig(mockObj)

		mockObj.AssertCalled(t, "Read")
		mockObj.AssertCalled(t, "Decode")

		assert.NoError(err)
		assert.Equal(config.Host, "localhost")
		assert.Equal(config.Port, 5000)
		assert.Equal(config.ChannelSecret, "1234")
		assert.Equal(config.ChannelToken, "1234")
	})

	t.Run("Fail to read config", func(t *testing.T) {
		assert := assert.New(t)

		mockObj := new(mockconfig.MockReader)
		mockObj.On("Read").Return(fmt.Errorf("Error Read"))
		mockObj.On("Decode").Return(nil, fmt.Errorf("Error Decode"))

		config, err := config.ReadConfig(mockObj)

		mockObj.AssertCalled(t, "Read")
		mockObj.AssertNotCalled(t, "Decode")

		assert.Error(err, "Error Read")
		assert.Nil(config)
	})

	t.Run("Failed to decode config", func(t *testing.T) {
		assert := assert.New(t)

		mockObj := new(mockconfig.MockReader)
		mockObj.On("Read").Return(nil)
		mockObj.On("Decode").Return(nil, fmt.Errorf("Error Decode"))

		config, err := config.ReadConfig(mockObj)

		mockObj.AssertCalled(t, "Read")
		mockObj.AssertCalled(t, "Decode")

		assert.Error(err, "Error Decode")
		assert.Nil(config)
	})
}

func TestIntegrationReadConfig(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping TestIntegrationReadConfig test")
	}

	t.Run("Successfully read config", func(t *testing.T) {
		assert := assert.New(t)

		config, err := config.ReadConfig(&config.EnvReader{EnvFiles: []string{"../../.env.test"}})

		assert.Nil(err)
		assert.NotNil(config)
	})
}
