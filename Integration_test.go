package acquiringsdk

import (
	"crypto/rand"
	"fmt"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/Worldline-Acquiring/acquiring-sdk-go/apiv1/acquirer/merchant/payments"
	"github.com/Worldline-Acquiring/acquiring-sdk-go/apiv1/domain"
	"github.com/Worldline-Acquiring/acquiring-sdk-go/configuration"
)

var envAcquirerID = os.Getenv("acquiring.api.acquirerId")
var envMerchantID = os.Getenv("acquiring.api.merchantId")
var envOAuth2ClientID = os.Getenv("acquiring.api.oauth2.clientId")
var envOAuth2ClientSecret = os.Getenv("acquiring.api.oauth2.clientSecret")
var envOAuthTokenURI = os.Getenv("acquiring.api.oauth2.tokenUri")
var envEndpointURL = os.Getenv("acquiring.api.endpointUrl")
var envProxyURL = os.Getenv("acquiring.api.proxyUrl")

func TestIntegratedProcessPayment(t *testing.T) {
	skipTestIfNeeded(t)

	client, err := getClientIntegration()
	if err != nil {
		t.Fatal(err)
	}
	defer func(client *Client) {
		_ = client.Close()
	}(client)

	paymentsClient := client.V1().Acquirer(envAcquirerID).Merchant(envMerchantID).Payments()

	request := getProcessPaymentIntegrationTestRequest(t)
	response, err := paymentsClient.ProcessPayment(request, nil)
	if err != nil {
		t.Fatal(err)
	}
	assertProcessPaymentResponse(t, request, response)

	paymentID := *response.PaymentID

	query := payments.GetPaymentStatusParams{}
	query.ReturnOperations = NewBool(true)

	status, err := paymentsClient.GetPaymentStatus(paymentID, query, nil)
	if err != nil {
		t.Fatal(err)
	}
	assertPaymentStatusResponse(t, paymentID, status)
}

func TestIntegratedRequestDCCRate(t *testing.T) {
	skipTestIfNeeded(t)

	client, err := getClientIntegration()
	if err != nil {
		t.Fatal(err)
	}
	defer func(client *Client) {
		_ = client.Close()
	}(client)

	request := getDCCRateIntegrationTestRequest(t, 200)
	response, err := client.V1().Acquirer(envAcquirerID).Merchant(envMerchantID).DynamicCurrencyConversion().RequestDccRate(request, nil)
	if err != nil {
		t.Fatal(err)
	}
	assertDccRateResponse(t, request, response)
}

func getProcessPaymentIntegrationTestRequest(t *testing.T) domain.APIPaymentRequest {
	amountData := domain.NewAmountData()
	amountData.Amount = NewInt64(200)
	amountData.CurrencyCode = NewString("GBP")
	amountData.NumberOfDecimals = NewInt32(2)

	plainCardData := domain.NewPlainCardData()
	plainCardData.CardNumber = NewString("4176669999000104")
	plainCardData.CardSecurityCode = NewString("012")
	plainCardData.ExpiryDate = NewString("122031")

	cardPaymentData := domain.NewCardPaymentData()
	cardPaymentData.CardEntryMode = NewString("ECOMMERCE")
	cardPaymentData.CardholderVerificationMethod = NewString("CARD_SECURITY_CODE")
	cardPaymentData.AllowPartialApproval = NewBool(false)
	cardPaymentData.Brand = NewString("VISA")
	cardPaymentData.CaptureImmediately = NewBool(false)
	cardPaymentData.CardData = plainCardData

	references := domain.NewPaymentReferences()
	references.MerchantReference = NewString(fmt.Sprintf("your-order-%s", pseudoUUID(t)))

	request := domain.APIPaymentRequest{}
	request.Amount = amountData
	request.AuthorizationType = NewString("PRE_AUTHORIZATION")
	request.TransactionTimestamp = NewString(FormatDateTime(time.Now()))
	request.CardPaymentData = cardPaymentData
	request.References = references
	request.OperationID = NewString(pseudoUUID(t))

	return request
}

func assertProcessPaymentResponse(t *testing.T, body domain.APIPaymentRequest, response domain.APIPaymentResponse) {
	assertEquals(t, *body.OperationID, response.OperationID)
	assertEquals(t, "0", response.ResponseCode)
	assertEquals(t, "APPROVED", response.ResponseCodeCategory)
	assertNotNil(t, response.ResponseCodeDescription)
	assertEquals(t, "AUTHORIZED", response.Status)
	assertNotNil(t, response.InitialAuthorizationCode)
	assertNotNil(t, response.PaymentID)
	assertNotNil(t, response.TotalAuthorizedAmount)
	assertEquals(t, 200, response.TotalAuthorizedAmount.Amount)
	assertEquals(t, "GBP", response.TotalAuthorizedAmount.CurrencyCode)
	assertEquals(t, 2, response.TotalAuthorizedAmount.NumberOfDecimals)
}

func assertPaymentStatusResponse(t *testing.T, paymentID string, response domain.APIPaymentResource) {
	assertNotNil(t, response.InitialAuthorizationCode)
	assertEquals(t, paymentID, response.PaymentID)
	assertEquals(t, "AUTHORIZED", response.Status)
}

func getDCCRateIntegrationTestRequest(t *testing.T, amount int64) domain.GetDCCRateRequest {
	amountData := domain.NewAmountData()
	amountData.Amount = NewInt64(amount)
	amountData.CurrencyCode = NewString("GBP")
	amountData.NumberOfDecimals = NewInt32(2)

	transactionDataForDcc := domain.NewTransactionDataForDcc()
	transactionDataForDcc.Amount = amountData
	transactionDataForDcc.TransactionType = NewString("PAYMENT")
	transactionDataForDcc.TransactionTimestamp = NewString(FormatDateTime(time.Now()))

	pointOfSaleDataForDcc := domain.NewPointOfSaleDataForDcc()
	pointOfSaleDataForDcc.TerminalID = NewString("12345678")

	cardDataForDcc := domain.NewCardDataForDcc()
	cardDataForDcc.Bin = NewString("41766699")
	cardDataForDcc.Brand = NewString("VISA")

	request := domain.GetDCCRateRequest{}
	request.OperationID = NewString(pseudoUUID(t))
	request.TargetCurrency = NewString("EUR")
	request.CardPaymentData = cardDataForDcc
	request.PointOfSaleData = pointOfSaleDataForDcc
	request.Transaction = transactionDataForDcc

	return request
}

func assertDccRateResponse(t *testing.T, body domain.GetDCCRateRequest, response domain.GetDccRateResponse) {
	assertNotNil(t, response.Proposal)
	assertNotNil(t, response.Proposal.OriginalAmount)
	assertEqualAmounts(t, body.Transaction.Amount, response.Proposal.OriginalAmount)
	assertEquals(t, *body.TargetCurrency, response.Proposal.ResultingAmount.CurrencyCode)
}

func assertEqualAmounts(t *testing.T, expected *domain.AmountData, actual *domain.AmountData) {
	assertEquals(t, *expected.Amount, actual.Amount)
	assertEquals(t, *expected.CurrencyCode, actual.CurrencyCode)
	assertEquals(t, *expected.NumberOfDecimals, actual.NumberOfDecimals)
}

func assertEquals[T comparable](t *testing.T, expected T, actual *T) {
	if actual == nil {
		t.Fatalf("expected '%v' but got nil", expected)
	} else if expected != *actual {
		t.Fatalf("expected '%v' but got '%v'", expected, *actual)
	}
}

func assertNotNil(t *testing.T, actual interface{}) {
	if actual == nil {
		t.Fatalf("unexpected nil")
	}
}

func getConfigurationIntegration() (*configuration.CommunicatorConfiguration, error) {
	conf, _ := configuration.DefaultOAuth2Configuration(envOAuth2ClientID, envOAuth2ClientSecret, envOAuthTokenURI, "Worldline")
	if len(envEndpointURL) == 0 {
		conf.APIEndpoint.Host = "api.preprod.acquiring.worldline-solutions.com"
	} else {
		endpoint, err := url.Parse(envEndpointURL)
		if err != nil {
			return nil, err
		}
		conf.APIEndpoint = *endpoint
	}

	if len(envProxyURL) > 0 {
		proxy, err := url.Parse(envProxyURL)
		if err != nil {
			return nil, err
		}
		conf.Proxy = proxy
	}
	return conf, nil
}

func getClientIntegration() (*Client, error) {
	conf, err := getConfigurationIntegration()
	if err != nil {
		return nil, err
	}
	client, err := CreateClientFromConfiguration(conf)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func pseudoUUID(t *testing.T) string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		t.Fatal("failed to create UUID")
	}
	return fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

func skipTestIfNeeded(t *testing.T) {
	if len(envAcquirerID) == 0 {
		t.Skip("empty env acquiring.api.acquirerId")
	}
	if len(envMerchantID) == 0 {
		t.Skip("empty env acquiring.api.merchantId")
	}
	if len(envOAuth2ClientID) == 0 || len(envOAuth2ClientSecret) == 0 || len(envOAuthTokenURI) == 0 {
		t.Skip("OAuth2 not configured")
	}
}
