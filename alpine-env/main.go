package main

import (
	"context"
	"fmt"

	"dagger.io/dagger"
)

func main() {
	ctx := context.Background()
	client, err := dagger.Connect(ctx)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	container := client.Container().From("alpine:3.16.2")

	container = container.WithEnvVariable("FOO", "bar")

	out, err := container.WithExec([]string{"sh", "-c", "echo $FOO"}).Stdout(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)

	out, err = container.WithExec([]string{"sh", "-c", "uname -a"}).Stdout(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println(out)

}
