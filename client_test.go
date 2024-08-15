package godriblie_test

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/kevin4463-godaddy/godriblie"
	"github.com/kevin4463-godaddy/godriblie/internal/mocks"
	"github.com/kevin4463-godaddy/godriblie/internal/utils"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"log"
	"testing"
	"time"
)

type client_UnitTestSuite struct {
	suite.Suite

	client       *godriblie.DribbleClient
	mockDynamoDb *mocks.DynamoDbProvider
	mockContext  context.Context

	defaultTableName      string
	defaultLockOwner      string
	defaultPartitionKey   string
	deleteOnReleaseOption bool
	isReleasedOption      bool
	dataOption            []byte
	lockOutputValue       *utils.LockDto
}

func Test_client_UnitTestSuite(t *testing.T) {
	suite.Run(t, &client_UnitTestSuite{})
}

func (suite *client_UnitTestSuite) SetupTest() {
	suite.defaultTableName = "mock_lock_table"
	suite.defaultPartitionKey = "mock_partition_key"
	suite.defaultLockOwner = "mock_lock_owner"

	suite.deleteOnReleaseOption = false
	suite.isReleasedOption = false
}

func (suite *client_UnitTestSuite) setupMocks() {
	suite.mockDynamoDb = mocks.NewDynamoDbProvider(suite.T())
	suite.mockContext = context.TODO()

	suite.lockOutputValue = &utils.LockDto{
		PartitionKey:    suite.defaultPartitionKey,
		Owner:           suite.defaultLockOwner,
		Timestamp:       time.Now().Unix(),
		ExpTime:         5 * 60 * 60 * 24,
		DeleteOnRelease: suite.deleteOnReleaseOption,
		IsReleased:      suite.isReleasedOption,
		Data:            suite.dataOption,
	}

	suite.client = godriblie.NewLockClient(suite.mockDynamoDb,
		suite.defaultTableName,
		godriblie.WithOwnerName(suite.defaultLockOwner))
}

func (suite *client_UnitTestSuite) Test_AcquireLock_Success() {
	suite.setupMocks()
	suite.mockDynamoDb.EXPECT().
		Query(suite.mockContext, mock.Anything).
		Return(&dynamodb.QueryOutput{Count: 1, Items: suite.getAttributeValueMap(suite.defaultLockOwner)}, nil).
		Once()
	suite.mockDynamoDb.EXPECT().
		PutItem(suite.mockContext, mock.Anything).
		Return(&dynamodb.PutItemOutput{Attributes: suite.getAttributeValueMap(suite.defaultLockOwner)[0]}, nil).
		Once()

	lock, err := suite.client.AcquireLock(suite.mockContext, suite.defaultPartitionKey)

	log.Printf("lock acquired: %+v\n", lock)

	suite.Nil(err)
	suite.Equal(suite.defaultPartitionKey, lock.PartitionKey)
}

func (suite *client_UnitTestSuite) Test_AcquireLock_NotOwner_Fail() {
	suite.setupMocks()
	newOwner := "some_different_owner_name"
	suite.mockDynamoDb.EXPECT().
		Query(suite.mockContext, mock.Anything).
		Return(&dynamodb.QueryOutput{Count: 1, Items: suite.getAttributeValueMap(newOwner)}, nil).
		Once()

	lock, err := suite.client.AcquireLock(suite.mockContext, suite.defaultPartitionKey)

	suite.Equal(err.Error(), "lock is in use by another and not released")
	suite.Equal((*utils.LockDto)(nil), lock)
}

func (suite *client_UnitTestSuite) getAttributeValueMap(ownerName string) []map[string]types.AttributeValue {
	return []map[string]types.AttributeValue{
		{
			"key":             &types.AttributeValueMemberS{Value: suite.defaultPartitionKey},
			"owner":           &types.AttributeValueMemberS{Value: ownerName},
			"timestamp":       &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", suite.lockOutputValue.Timestamp)},
			"expTime":         &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", suite.lockOutputValue.ExpTime)},
			"isReleased":      &types.AttributeValueMemberBOOL{Value: suite.isReleasedOption},
			"deleteOnRelease": &types.AttributeValueMemberBOOL{Value: suite.deleteOnReleaseOption},
		},
	}
}
