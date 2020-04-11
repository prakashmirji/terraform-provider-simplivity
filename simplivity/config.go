package simplivity

import (
	"errors"
	"fmt"

	"github.com/HewlettPackard/simplivity-go/ovc"
)

// Config struct to store creds
type Config struct {
	Username       string
	Password       string
	OvcHostAddress string
	CertPath       string

	ovcClient *ovc.Client
}

var ErrConfigNotInitialized = errors.New("config not initialized")

func (c *Config) createClient() error {
	if c == nil {
		return ErrConfigNotInitialized
	}

	client, err := ovc.NewClient(c.Username, c.Password, c.OvcHostAddress, "")

	if err != nil {
		fmt.Println(err)
	}

	c.ovcClient = client

	return nil

}
