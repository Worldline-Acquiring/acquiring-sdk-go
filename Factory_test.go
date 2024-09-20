package acquiringsdk

import (
	"reflect"
	"testing"

	"github.com/Worldline-Acquiring/acquiring-sdk-go/authentication/oauth2"
	"github.com/Worldline-Acquiring/acquiring-sdk-go/communicator"
	"github.com/Worldline-Acquiring/acquiring-sdk-go/configuration"
	"github.com/Worldline-Acquiring/acquiring-sdk-go/json"
)

var clientID = "someId"
var clientSecret = "someSecret"
var tokenURI = "https://sso.preprod.acquiring.worldline-solutions.com/auth/realms/acquiring_api/protocol/openid-acquiring/token"

func TestCreateConfiguration(t *testing.T) {
	defaultConfiguration, _ := configuration.DefaultOAuth2Configuration(clientID, clientSecret, tokenURI, "Test")

	conf, err := CreateOAuth2Configuration(clientID, clientSecret, tokenURI, "Test")
	if err != nil {
		t.Fatalf("TestCreateConfiguration: %v", err)
	}

	if conf.APIEndpoint != defaultConfiguration.APIEndpoint {
		t.Fatalf("TestCreateConfiguration: APIEndpoint mismatch %v %v",
			conf.APIEndpoint, defaultConfiguration.APIEndpoint)
	}

	if conf.AuthorizationType != defaultConfiguration.AuthorizationType {
		t.Fatalf("TestCreateConfiguration: AuthorizationType mismatch %v %v",
			conf.AuthorizationType, defaultConfiguration.AuthorizationType)
	}

	if conf.IdleTimeout != defaultConfiguration.IdleTimeout {
		t.Fatalf("TestCreateConfiguration: IdleTimeout mismatch %v %v",
			conf.IdleTimeout, defaultConfiguration.IdleTimeout)
	}

	if conf.ConnectTimeout != defaultConfiguration.ConnectTimeout {
		t.Fatalf("TestCreateConfiguration: ConnectTimeout mismatch %v %v",
			conf.ConnectTimeout, defaultConfiguration.ConnectTimeout)
	}

	if conf.SocketTimeout != defaultConfiguration.SocketTimeout {
		t.Fatalf("TestCreateConfiguration: SocketTimeout mismatch %v %v",
			conf.SocketTimeout, defaultConfiguration.SocketTimeout)
	}

	if conf.MaxConnections != defaultConfiguration.MaxConnections {
		t.Fatalf("TestCreateConfiguration: MaxConnections mismatch %v %v",
			conf.MaxConnections, defaultConfiguration.MaxConnections)
	}

	if conf.GetOAuth2ClientID() != clientID {
		t.Fatalf("TestCreateConfiguration: ClientID mismatch %v %v",
			conf.GetOAuth2ClientID(), clientID)
	}

	if conf.GetOAuth2ClientSecret() != clientSecret {
		t.Fatalf("TestCreateConfiguration: ClientSecret mismatch %v %v",
			conf.GetOAuth2ClientSecret(), clientSecret)
	}
}

func TestCreateCommunicator(t *testing.T) {
	marshaller := json.DefaultMarshaller()

	conf, err := CreateOAuth2Configuration(clientID, clientSecret, tokenURI, "Test")
	if err != nil {
		t.Fatalf("TestCreateCommunicator: %v", err)
	}
	comm, err := CreateCommunicatorFromConfiguration(conf)
	if err != nil {
		t.Fatalf("TestCreateCommunicator: %v", err)
	}
	if comm.Marshaller() != marshaller {
		t.Fatalf("TestCreateCommunicator: marshaller mismatch %v %v",
			comm.Marshaller(), marshaller)
	}

	connection := comm.Connection()
	if _, isDefaultConnection := connection.(*communicator.DefaultConnection); !isDefaultConnection {
		t.Fatalf("TestCreateCommunicator: connection type mismatch %T", connection)
	}

	authenticator := comm.Authenticator()
	if _, isOAuth2Authenticator := authenticator.(*oauth2.Authenticator); !isOAuth2Authenticator {
		t.Fatalf("TestCreateCommunicator: authenticator type mismatch %T", authenticator)
	}

	authClientID := reflect.Indirect(reflect.ValueOf(authenticator)).FieldByName("clientID").String()
	if authClientID != clientID {
		t.Fatalf("TestCreateCommunicator: clientID mismatch %v", authClientID)
	}
	authClientSecret := reflect.Indirect(reflect.ValueOf(authenticator)).FieldByName("clientSecret").String()
	if authClientSecret != clientSecret {
		t.Fatalf("TestCreateCommunicator: clientSecret mismatch %v", authClientSecret)
	}

	metadataProvider := comm.MetadataProvider()
	headers := metadataProvider.MetadataHeaders()
	if len(headers) != 1 {
		t.Fatalf("TestCreateCommunicator: headers len mismatch %v", len(headers))
	}
	if headers[0].Name() != "X-WL-ServerMetaInfo" {
		t.Fatalf("TestCreateCommunicator: header name mismatch %v", headers[0].Name())
	}
}
