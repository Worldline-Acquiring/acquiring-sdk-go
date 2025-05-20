// This file was automatically generated.

package oauth2

import "strings"

var scopesByOperation = map[string]map[string][]string{
	"v1": {
		"processPayment":             {"processing_payment"},
		"getPaymentStatus":           {"processing_payment"},
		"simpleCaptureOfPayment":     {"processing_payment"},
		"reverseAuthorization":       {"processing_payment"},
		"incrementPayment":           {"processing_payment"},
		"createRefund":               {"processing_refund"},
		"processStandaloneRefund":    {"processing_refund"},
		"getRefund":                  {"processing_refund"},
		"captureRefund":              {"processing_refund"},
		"reverseRefundAuthorization": {"processing_refund"},
		"processAccountVerification": {"processing_accountverification"},
		"processBalanceInquiry":      {"processing_balanceinquiry"},
		"technicalReversal":          {"processing_operation_reverse"},
		"requestDccRate":             {"processing_dcc_rate"},
		"ping":                       {"services_ping"},
	},
}

func collectAllScopes() []string {
	seen := map[string]bool{}
	var result []string
	for _, operations := range scopesByOperation {
		for _, scopes := range operations {
			for _, scope := range scopes {
				if _, ok := seen[scope]; !ok {
					result = append(result, scope)
					seen[scope] = true
				}
			}
		}
	}
	return result
}

var allScopes = collectAllScopes()

var allScopesString = strings.Join(allScopes, " ")

// AllScopes returns all available scopes.
func AllScopes() []string {
	return allScopes
}

// ScopesForAPIVersion returns all scopes needed for all operations of the given API version.
func ScopesForAPIVersion(apiVersion string) []string {
	if operations, ok := scopesByOperation[apiVersion]; ok {
		seen := map[string]bool{}
		var result []string
		for _, scopes := range operations {
			for _, scope := range scopes {
				if _, ok := seen[scope]; !ok {
					result = append(result, scope)
					seen[scope] = true
				}
			}
		}
		return result
	}
	return []string{}
}

// ScopesForOperation returns all scopes needed for the given operation of the given API version.
func ScopesForOperation(apiVersion string, operationID string) []string {
	if operations, ok := scopesByOperation[apiVersion]; ok {
		if scopes, ok := operations[operationID]; ok {
			return scopes
		}
	}
	return []string{}
}

// ScopesForOperations returns all scopes needed for the given operations of the given API version.
func ScopesForOperations(apiVersion string, operationIDs ...string) []string {
	if operations, ok := scopesByOperation[apiVersion]; ok {
		seen := map[string]bool{}
		var result []string
		for _, operationID := range operationIDs {
			if scopes, ok := operations[operationID]; ok {
				for _, scope := range scopes {
					if _, ok := seen[scope]; !ok {
						result = append(result, scope)
						seen[scope] = true
					}
				}
			}
		}
		return result
	}
	return []string{}
}

// FilteredScopesForOperation returns all scopes needed for the operations that pass the given filter.
// The first argument to the predicate is the API version, the second is the operation id.
func FilteredScopesForOperation(filter func(string, string) bool) []string {
	seen := map[string]bool{}
	var result []string
	for apiVersion, operations := range scopesByOperation {
		for operationID, scopes := range operations {
			if filter(apiVersion, operationID) {
				for _, scope := range scopes {
					if _, ok := seen[scope]; !ok {
						result = append(result, scope)
						seen[scope] = true
					}
				}
			}
		}
	}
	return result
}
