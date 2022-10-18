package controller

import (
	"fmt"
	"log"
	"net/http"
	"regexp"

	"GoHttpServer/parser"
	"GoHttpServer/validator"
)

const ConfigPath = "./configs/config.yml"

type controller struct {
	ct         []string
	urlPattern string
	validator  ValidatorInt
}

func (c *controller) String() string {
	return fmt.Sprintf("Url %s and content %s", c.urlPattern, c.ct[0])
}

func (c *controller) GetValidator() *ValidatorInt {
	return &c.validator
}

func (c *controller) SetDefault() {
	config, err := parser.ParseConfig(ConfigPath)
	if err != nil {
		panic(err.Error())
	}

	c.urlPattern = config.Auth.Urlvalidreg
	c.ct = config.Auth.Contenttype
}

func (c *controller) UpdateConfig() {
	config, err := parser.ParseConfig(ConfigPath)
	if err != nil {
		log.Println("not updated config", c.ct)
		return
	}

	c.urlPattern = config.Auth.Urlvalidreg
	c.ct = config.Auth.Contenttype
	c.validator = validator.New("boom")
	log.Println("updated config", c.ct)
}

func NewController() *controller {
	var cntr controller
	cntr.SetDefault()
	cntr.validator = validator.New("owl")
	return &cntr
}

func (c *controller) SighupHandler() {
	log.Println("sighup")
	c.UpdateConfig()
}

func (c *controller) ValidateURL(url string) bool {
	if matched, err := regexp.MatchString(c.urlPattern, url); err == nil {
		return matched
	}
	return false
}

func (c *controller) ValidateContentType(content string) bool {
	for _, ct := range c.ct {
		if ct == content {
			return true
		}
	}
	return false
}

func (c *controller) Validate(content, url string) (int, error) {
	if c.ValidateContentType(content) && c.ValidateURL(url) {
		return http.StatusOK, nil
	}

	return http.StatusBadRequest, fmt.Errorf("error in content %s or url %s", content, url)
}
