// Copyright 2020 go-azuredevops AUTHORS. All rights reserved.
//
// This example uses the go-azuredevops library to get the Azure DevOps user Id
// by the specified user name and organization name
// 0. Prompts the user for required inputs: organization name, user name, personal access token
// 1. Creates a NewClient() using basic auth and personal access token
// 2. Gets the user Id

package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/mcdafydd/go-azuredevops/azuredevops"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Print("Azure Devops Organization name: ")
	orgName, _ := r.ReadString('\n')
	orgName = strings.TrimSuffix(orgName, "\r\n")

	fmt.Print("Azure Devops Personal Access Token: ")
	bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
	token := string(bytePassword)
	fmt.Print("\n")

	r = bufio.NewReader(os.Stdin)
	fmt.Print("Azure Devops User name: ")
	userName, _ := r.ReadString('\n')
	userName = strings.TrimSuffix(userName, "\r\n")

	tp := azuredevops.BasicAuthTransport{
		Username: "",
		Password: token,
	}

	client, _ := azuredevops.NewClient(tp.Client())

	result, err := client.UserEntitlements.GetUserID(context.Background(), userName, orgName)
	if err != nil {
		fmt.Printf("Error trying to list user entitlements: %+v\n", err)
	}

	if result == nil {
		fmt.Print("User not found")
	} else {
		fmt.Printf("Azure DevOps User ID: %v", *result)
	}
}
