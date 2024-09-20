package obfuscation

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
)

var errInvalidDataType = errors.New("invalid data type - should never happen")

// BodyObfuscator can be used to obfuscate properties in JSON bodies.
type BodyObfuscator struct {
	rules ruleMap
}

func (o BodyObfuscator) obfuscateValue(value interface{}, rule Rule) (string, error) {
	if strVal, ok := value.(string); ok {
		return rule(strVal), nil
	}
	if intVal, ok := value.(int32); ok {
		return rule(strconv.FormatInt(int64(intVal), 10)), nil
	}
	if intVal, ok := value.(int64); ok {
		return rule(strconv.FormatInt(intVal, 10)), nil
	}
	if floatVal, ok := value.(float32); ok {
		return rule(strconv.FormatFloat(float64(floatVal), 'f', -1, 32)), nil
	}
	if floatVal, ok := value.(float64); ok {
		return rule(strconv.FormatFloat(floatVal, 'f', -1, 64)), nil
	}
	if boolVal, ok := value.(bool); ok {
		return rule(strconv.FormatBool(boolVal)), nil
	}
	return "", errInvalidDataType
}

func (o BodyObfuscator) navigateJSON(content interface{}) error {
	if content == nil {
		return nil
	}

	if contentMap, ok := content.(map[string]interface{}); ok {
		for name, obj := range contentMap {
			_, isMap := obj.(map[string]interface{})
			if rule, ok := o.rules[name]; ok && !isMap {
				obfuscatedValue, err := o.obfuscateValue(obj, rule)
				if err != nil {
					return err
				}

				contentMap[name] = obfuscatedValue

			} else {
				err := o.navigateJSON(obj)
				if err != nil {
					return err
				}
			}
		}
	}

	if contentSlice, ok := content.([]interface{}); ok {
		for _, obj := range contentSlice {
			err := o.navigateJSON(obj)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// ObfuscateBody obfuscates the given body as necessary
func (o BodyObfuscator) ObfuscateBody(body string) (string, error) {
	if strings.TrimSpace(body) == "" {
		return body, nil
	}

	var parsedJSON interface{}
	err := json.Unmarshal([]byte(body), &parsedJSON)
	if _, ok := err.(*json.SyntaxError); ok {
		return body, nil
	}
	if err != nil {
		return body, err
	}

	err = o.navigateJSON(parsedJSON)
	if err != nil {
		return body, err
	}

	obfuscatedBody, err := json.MarshalIndent(parsedJSON, "", "    ")
	if err != nil {
		return body, err
	}

	return string(obfuscatedBody), nil
}

// NewBodyObfuscator returns a body obfuscator.
// This will contain some pre-defined obfuscation rules, as well as any provided custom rules.
func NewBodyObfuscator(customRules map[string]Rule) BodyObfuscator {
	rules := ruleMap{
		"address":                 All(),
		"authenticationValue":     KeepingStartCount(4),
		"bin":                     KeepingStartCount(6),
		"cardholderAddress":       All(),
		"cardholderPostalCode":    All(),
		"cardNumber":              KeepingEndCount(4),
		"cardSecurityCode":        All(),
		"city":                    All(),
		"cryptogram":              KeepingStartCount(4),
		"expiryDate":              KeepingEndCount(4),
		"name":                    All(),
		"paymentAccountReference": KeepingStartCount(6),
		"postalCode":              All(),
		"stateCode":               All(),
	}

	for name, rule := range customRules {
		rules[name] = rule
	}

	return BodyObfuscator{rules}
}

var defaultBodyObfuscator = NewBodyObfuscator(ruleMap{})

// DefaultBodyObfuscator returns a default body obfuscator.
// This will be equivalent to calling NewBodyObfuscator with an empty rule map.
func DefaultBodyObfuscator() BodyObfuscator {
	return defaultBodyObfuscator
}
