package godriblie_test

import (
	"context"
	"fmt"
	"net"
	"os"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/kevin4463-godaddy/godriblie"
)

func TestMain(m *testing.M) {
	for i := 0; i < 10; i++ {
		c, err := net.Dial("tcp", "localhost:8000")
		if err != nil {
			time.Sleep(1 * time.Second)
			fmt.Println("retry connection check")
			continue
		}
		_ = c.Close()
		break
	}
	time.Sleep(1 * time.Second)
	exitCode := m.Run()
	os.Exit(exitCode)
}

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
}

func TestClientBasicFlow(t *testing.T) {
	t.Parallel()
	svc := dynamodb.NewFromConfig(defaultConfig(t))

	clt := godriblie.NewLockClient(
		svc,
		"locks_local",
		godriblie.WithOwnerName("local_macos_owner"),
		godriblie.WithPartitionKeyName("key"),
	)

	t.Cleanup(func() {
		t.Log("clean up")
	})

	_, _ = clt.CreateTable(context.Background(),
		"locks_local",
		godriblie.WithProvisionedThroughput(&types.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(5),
			WriteCapacityUnits: aws.Int64(5),
		}),
		godriblie.WithCustomPartitionKeyName("key"))

	data := `"im": { "a": "little-teapot" }`
	err := clt.AcquireLock(context.Background(),
		"spookyMonster",
		false,
		data)
	if err != nil {
		t.Fatal(err)
	}

	ok, lk, err := clt.CheckLock(context.Background(), "spookyMonster")
	if err != nil {
		t.Log(err)
	}

	t.Log(ok, lk)
}
