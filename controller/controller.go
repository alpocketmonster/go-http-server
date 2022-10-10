package controller

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
)

type controller struct {
	ct         string
	urlPattern string
}

func (c *controller) setDefault() {
	c.ct = "application/vnd.kafka.avro.v2+json"
	c.urlPattern = `^\d{3}-\d(-\d{3}-\d)?\.[a-z0-9-]+\.(db|cdc|cmd|sys|log|tmp)\.[a-z0-9-.]+\.\d+$`
}

func (c *controller) String() string {
	return string(c.ct)
}

func NewController() *controller {
	var cntr controller
	return &cntr
}

func (c *controller) SighupHandler() {

	log.Println("sighup")
}

func (c *controller) ValidateURL(url string) bool {
	if matched, err := regexp.MatchString(c.urlPattern, url); err == nil {
		return matched
	}
	return false
}

func (c *controller) Validate(content, url string) (int, error) {
	c.setDefault()
	if content == c.ct && c.ValidateURL(url) {
		return http.StatusOK, nil
	}

	return http.StatusBadRequest, fmt.Errorf("error in content %s or url %s", content, url)
}
