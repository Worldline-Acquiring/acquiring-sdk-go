package obfuscation

import (
	"bytes"
	"testing"
	"unicode/utf8"
)

func TestObfuscateBodyWithEmptyBody(t *testing.T) {
	emptyString := ""
	obfuscatedBody, err := DefaultBodyObfuscator().ObfuscateBody(emptyString)

	if err != nil {
		t.Fatal(err)
	}

	if obfuscatedBody != emptyString {
		t.Fatalf("TestObfuscateBodyWithEmptyBody : expected '%s' got '%s'", emptyString, obfuscatedBody)
	}
}

func TestObfuscateBodyWithCard(t *testing.T) {
	cardObfuscated := `{
    "amount": {
        "amount": 2345,
        "currencyCode": "CAD"
    },
    "cardPaymentData": {
        "cardData": {
            "cardNumber": "************3456",
            "cardSecurityCode": "***",
            "expiryDate": "**2024"
        }
    }
}`

	cardUnObfuscated := `{
	"amount": {
		"currencyCode": "CAD",
		"amount": 2345
	},
	"cardPaymentData": {
		"cardData": {
			"cardSecurityCode": "123",
			"cardNumber": "1234567890123456",
			"expiryDate": "122024"
		}
	}
}`

	obfuscatedBody, err := DefaultBodyObfuscator().ObfuscateBody(cardUnObfuscated)

	if err != nil {
		t.Fatal(err)
	}

	if cardObfuscated != obfuscatedBody {
		t.Fatalf("TestObfuscateBodyWithCard : expected \n%s\ngot\n%s", cardObfuscated, obfuscatedBody)
	}
}

func TestObfuscateBodyWithCustomCardRule(t *testing.T) {
	cardObfuscated := `{
    "amount": {
        "amount": 2345,
        "currencyCode": "CAD"
    },
    "cardPaymentData": {
        "cardData": {
            "cardNumber": "123456******3456",
            "cardSecurityCode": "***",
            "expiryDate": "**2024"
        }
    }
}`

	cardUnObfuscated := `{
	"amount": {
		"currencyCode": "CAD",
		"amount": 2345
	},
	"cardPaymentData": {
		"cardData": {
			"cardSecurityCode": "123",
			"cardNumber": "1234567890123456",
			"expiryDate": "122024"
		}
	}
}`

	rule := func(value string) string {
		valueLength := utf8.RuneCountInString(value)
		var chars bytes.Buffer
		i := 0
		for _, r := range value {
			if i < 6 || i >= valueLength-4 {
				chars.WriteRune(r)
			} else {
				chars.WriteRune('*')
			}
			i++
		}
		return chars.String()
	}
	bodyObfuscator := NewBodyObfuscator(ruleMap{
		"cardNumber": rule,
	})
	obfuscatedBody, err := bodyObfuscator.ObfuscateBody(cardUnObfuscated)

	if err != nil {
		t.Fatal(err)
	}

	if cardObfuscated != obfuscatedBody {
		t.Fatalf("TestObfuscateBodyWithCustomCardRule : expected \n%s\ngot\n%s", cardObfuscated, obfuscatedBody)
	}
}

func TestObfuscateBodyWithNoMatches(t *testing.T) {
	noObfuscation := `{
	"amount": {
		"currencyCode": "EUR",
		"amount": 1000
	},
	"authorizationType": "PRE_AUTHORIZATION"
}`
	postObfuscation := `{
    "amount": {
        "amount": 1000,
        "currencyCode": "EUR"
    },
    "authorizationType": "PRE_AUTHORIZATION"
}`

	obfuscatedBody, err := DefaultBodyObfuscator().ObfuscateBody(noObfuscation)

	if err != nil {
		t.Fatal(err)
	}

	if postObfuscation != obfuscatedBody {
		t.Fatalf("TestObfuscateBodyWithNoMatches : expected \n%s\ngot\n%s", postObfuscation, obfuscatedBody)
	}
}

func TestObfuscateBodyWithObject(t *testing.T) {
	jsonObfuscated := `[
    {
        "name": "****"
    },
    {
        "name": {}
    }
]`
	jsonUnobfuscated := `[ {
		"name": true
	}, {
		"name": {
	}
} ]`

	obfuscatedBody, err := DefaultBodyObfuscator().ObfuscateBody(jsonUnobfuscated)

	if err != nil {
		t.Fatal(err)
	}

	if jsonObfuscated != obfuscatedBody {
		t.Fatalf("TestObfuscateBodyWithObject : expected \n%s\ngot\n%s", jsonObfuscated, obfuscatedBody)
	}
}
