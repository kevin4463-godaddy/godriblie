package godriblie_test

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/kevin4463-godaddy/godriblie"
	"github.com/kevin4463-godaddy/godriblie/internal/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type client_UnitTestSuite struct {
	suite.Suite

	client           *godriblie.DribbleClient
	mockDynamoDb     *mocks.DynamoDbProvider
	mockContext      context.Context
	defaultTableName string
}

func Test_client_UnitTestSuite(t *testing.T) {
	suite.Run(t, &client_UnitTestSuite{})
}

func (suite *client_UnitTestSuite) SetupTest() {
	suite.defaultTableName = "mock_lock_table"
}

func (suite *client_UnitTestSuite) setupMocks() {
	suite.mockDynamoDb = mocks.NewDynamoDbProvider(suite.T())
	suite.mockContext = context.TODO()
	suite.client = godriblie.NewLockClient(suite.mockDynamoDb, suite.defaultTableName,
		godriblie.WithOwnerName("mock_owner"))
}

func (suite *client_UnitTestSuite) Test_AcquireLock_Success() {
	suite.setupMocks()

	suite.mockDynamoDb.EXPECT().
		PutItem(suite.mockContext, mock.Anything).
		Return(&dynamodb.PutItemOutput{
			Attributes: map[string]types.AttributeValue{
				"key": &types.AttributeValueMemberS{Value: "test_key_value"},
			},
		}, nil)

	pkey := "test_key_value"

	o, err := suite.client.AcquireLock(suite.mockContext, pkey, godriblie.WithDeleteLockOnRelease())
	suite.Nil(err)
	suite.Equal(pkey, o.PartitionKey)
}

/*
func defaultConfig(t *testing.T) aws.Config {
	t.Helper()
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-west-2"),
		config.WithEndpointResolver(aws.EndpointResolverFunc(
			func(service, region string) (aws.Endpoint, error) {
				return aws.Endpoint{URL: "http://localhost:8000"}, nil
			})),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID:     "DUMMY-ID-EXAMPLE",
				SecretAccessKey: "DUMMY-KEY-EXAMPLE",
				SessionToken:    "dummy",
				Source:          "Hard-coded credentials; values are irrelevant for local db",
			},
		}))
	if err != nil {
		t.Fatal(err)
	}

	return cfg
}*/
