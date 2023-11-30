package forms

import (
	"reflect"
	"regexp"
	"strings"
	"sync"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type DefaultValidator struct {
	once     sync.Once
	validate *validator.Validate
}

var _ binding.StructValidator = &DefaultValidator{}

func (v *DefaultValidator) Engine() any {
	v.lazyinit()
	return v.validate
}

func (v *DefaultValidator) ValidateStruct(obj any) error {
	if kindOfData(obj) != reflect.Struct {
		return nil
	}
	v.lazyinit()
	return v.validate.Struct(obj)
}

func kindOfData(data interface{}) reflect.Kind {
	value := reflect.ValueOf(data)
	valueType := value.Kind()

	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}

func (v *DefaultValidator) lazyinit() {
	v.once.Do(func() {
		v.validate = validator.New()
		v.validate.SetTagName("binding")

		// custom validations
		v.validate.RegisterValidation("name", ValidateName)
	})
}

func ValidateName(fl validator.FieldLevel) bool {
	space := regexp.MustCompile(`\s+`)
	name := space.ReplaceAllString(fl.Field().String(), " ")
	name = strings.TrimSpace(name)

	matched, _ := regexp.MatchString(`^[^±!@£$%^&*_+§¡€#¢§¶•ªº«\\/<>?:;'"|=.,0123456789]{3,50}$`, name)
	return matched
}
