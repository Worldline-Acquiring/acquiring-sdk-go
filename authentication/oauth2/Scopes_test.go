package oauth2

import (
	"strings"
	"testing"
)

func TestAllScopes(t *testing.T) {
	scopes := AllScopes()
	assertContains(t, "TestAllScopes", scopes, "processing_payment")
	assertContains(t, "TestAllScopes", scopes, "processing_dcc_rate")
	assertContains(t, "TestAllScopes", scopes, "services_ping")

	allScopesString := strings.Join(scopes, " ")
	if len(allScopesString) > 260 {
		t.Fatalf("TestAllScopes : %s is too long", allScopesString)
	}
}

func TestScopesForV1(t *testing.T) {
	scopes := ScopesForAPIVersion("v1")
	assertContains(t, "TestScopesForV1", scopes, "processing_payment")
	assertContains(t, "TestScopesForV1", scopes, "processing_dcc_rate")
	assertContains(t, "TestScopesForV1", scopes, "services_ping")
}

func TestScopesForUnknownApiVersion(t *testing.T) {
	scopes := ScopesForAPIVersion("v-1")
	if len(scopes) != 0 {
		t.Fatalf("TestScopesForUnknownApiVersion : %v is not empty", scopes)
	}
}

func TestScopesForV1ProcessPayment(t *testing.T) {
	scopes := ScopesForOperation("v1", "processPayment")
	assertContains(t, "TestScopesForV1ProcessPayment", scopes, "processing_payment")
}

func TestScopesForV1RequestDccRate(t *testing.T) {
	scopes := ScopesForOperation("v1", "requestDccRate")
	assertContains(t, "TestScopesForV1RequestDccRate", scopes, "processing_dcc_rate")
}

func TestScopesForUnknownOperation(t *testing.T) {
	scopes := ScopesForOperation("v1", "unknown")
	if len(scopes) != 0 {
		t.Fatalf("TestScopesForUnknownOperation : %v is not empty", scopes)
	}
}

func TestScopesForOperationOfUnknownApiVersion(t *testing.T) {
	scopes := ScopesForOperation("v-1", "processPayment")
	if len(scopes) != 0 {
		t.Fatalf("TestScopesForOperationOfUnknownApiVersion : %v is not empty", scopes)
	}
}

func TestScopesForV1Operations(t *testing.T) {
	scopes := ScopesForOperations("v1", "processPayment", "requestDccRate", "unknown")
	assertContains(t, "TestScopesForV1Operations", scopes, "processing_payment")
	assertContains(t, "TestScopesForV1Operations", scopes, "processing_dcc_rate")
	assertNotContains(t, "TestScopesForV1Operations", scopes, "services_ping")
}

func TestScopesForOperationsOfUnknownApiVersion(t *testing.T) {
	scopes := ScopesForOperations("v-1", "processPayment", "requestDccRate")
	if len(scopes) != 0 {
		t.Fatalf("TestScopesForOperationsOfUnknownApiVersion : %v is not empty", scopes)
	}
}

func TestFilteredScopes(t *testing.T) {
	operationIDs := map[string]bool{
		"processPayment": true,
		"requestDccRate": true,
		"unknown":        true,
	}
	scopes := FilteredScopesForOperation(func(apiVersion string, operationID string) bool {
		if _, ok := operationIDs[operationID]; ok {
			return apiVersion == "v1"
		}
		return false
	})
	assertContains(t, "TestScopesForV1Operations", scopes, "processing_payment")
	assertContains(t, "TestScopesForV1Operations", scopes, "processing_dcc_rate")
	assertNotContains(t, "TestScopesForV1Operations", scopes, "services_ping")
}

func assertContains(t *testing.T, testName string, slice []string, value string) {
	t.Helper()

	for _, s := range slice {
		if value == s {
			return
		}
	}
	t.Fatalf("%s : %s does not contain %s", testName, slice, value)
}

func assertNotContains(t *testing.T, testName string, slice []string, value string) {
	t.Helper()

	for _, s := range slice {
		if value == s {
			t.Fatalf("%s : %s contains %s", testName, slice, value)
			return
		}
	}
}
