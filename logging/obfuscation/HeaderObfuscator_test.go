package obfuscation

import "testing"

func CheckObfuscateHeaderWithMatch(t *testing.T, name, originalValue, expectedObfuscatedValue string) {
	CheckObfuscateHeaderWithMatchWithCustomObfuscator(t, DefaultHeaderObfuscator(), name, originalValue, expectedObfuscatedValue)
}

func CheckObfuscateHeaderWithMatchWithCustomObfuscator(t *testing.T, headerObfuscator HeaderObfuscator, name, originalValue, expectedObfuscatedValue string) {
	obfuscatedValue := headerObfuscator.ObfuscateHeader(name, originalValue)

	if obfuscatedValue != expectedObfuscatedValue {
		t.Fatalf("CheckObfuscateHeaderWithMatch : expected '%s' got '%s'", expectedObfuscatedValue, obfuscatedValue)
	}
}

func CheckObfuscateHeaderWithNoMatch(t *testing.T, name, originalValue string) {
	headerObfuscator := DefaultHeaderObfuscator()
	obfuscatedValue := headerObfuscator.ObfuscateHeader(name, originalValue)

	if obfuscatedValue != originalValue {
		t.Fatalf("CheckObfuscateHeaderWithNoMatch : expected '%s' got '%s'", originalValue, obfuscatedValue)
	}
}

func TestObfuscateHeader(t *testing.T) {
	CheckObfuscateHeaderWithMatch(t, "Authorization", "Basic QWxhZGRpbjpPcGVuU2VzYW1l", "********")
	CheckObfuscateHeaderWithMatch(t, "authorization", "Basic QWxhZGRpbjpPcGVuU2VzYW1l", "********")
	CheckObfuscateHeaderWithMatch(t, "AUTHORIZATION", "Basic QWxhZGRpbjpPcGVuU2VzYW1l", "********")

	CheckObfuscateHeaderWithNoMatch(t, "Content-Type", "application/json")
	CheckObfuscateHeaderWithNoMatch(t, "content-type", "application/json")
	CheckObfuscateHeaderWithNoMatch(t, "CONTENT-TYPE", "application/json")
}

func TestObfuscateCustomHeader(t *testing.T) {
	headerObfuscator := NewHeaderObfuscator(map[string]Rule{
		"content-type": All(),
	})

	CheckObfuscateHeaderWithMatchWithCustomObfuscator(t, headerObfuscator, "Authorization", "Basic QWxhZGRpbjpPcGVuU2VzYW1l", "********")
	CheckObfuscateHeaderWithMatchWithCustomObfuscator(t, headerObfuscator, "authorization", "Basic QWxhZGRpbjpPcGVuU2VzYW1l", "********")
	CheckObfuscateHeaderWithMatchWithCustomObfuscator(t, headerObfuscator, "AUTHORIZATION", "Basic QWxhZGRpbjpPcGVuU2VzYW1l", "********")

	CheckObfuscateHeaderWithMatchWithCustomObfuscator(t, headerObfuscator, "Content-Type", "application/json", "****************")
	CheckObfuscateHeaderWithMatchWithCustomObfuscator(t, headerObfuscator, "content-type", "application/json", "****************")
	CheckObfuscateHeaderWithMatchWithCustomObfuscator(t, headerObfuscator, "CONTENT-TYPE", "application/json", "****************")
}
