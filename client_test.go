package godriblie_test

import (
	"context"
	"flag"
	"net"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/kevin4463-godaddy/godriblie"
)

func TestMain(m *testing.M) {
	flag.Parse()
	javaPath, err := exec.LookPath("java")
	if err != nil {
		panic("cannot execute tests without Java")
	}
	ctx, cancel := context.WithCancel(context.Background())
	cmd := exec.CommandContext(ctx, javaPath, "-Djava.library.path=./DynamoDBLocal_lib", "-jar", "DynamoDBLocal.jar", "-sharedDb")
	cmd.Dir = "/Users/kevin4463/local-dynamodb"
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Start(); err != nil {
		panic("cannot start local dynamodb:" + err.Error())
	}
	for i := 0; i < 10; i++ {
		c, err := net.Dial("tcp", "localhost:8000")
		if err != nil {
			time.Sleep(1 * time.Second)
			continue
		}
		_ = c.Close()
		break
	}
	time.Sleep(1 * time.Second)
	exitCode := m.Run()
	cancel()
	_ = cmd.Wait()
	os.Exit(exitCode)
}

func defaultConfig(t *testing.T) aws.Config {
	t.Helper()
	_ = os.Setenv("AWS_ACCESS_KEY_ID", "DUMMY-ID-EXAMPLE")
	_ = os.Setenv("AWS_SECRET_ACCESS_KEY", "DUMMY-KEY-EXAMPLE")

	return aws.Config{
		Region: "us-west-2",
	}
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

	_, _ = clt.CreateTable(context.Background(), "locks_local")
	data := `"im": { "a": "little-teapot" }`
	err := clt.AcquireLock(context.Background(), "spookyMonster", false, data)
	if err != nil {
		t.Fatal(err)
	}
}
