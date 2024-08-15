package godriblie_test

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
