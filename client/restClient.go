package client

import (
	"fmt"
	"github.com/devtron-labs/devtron-cli/devtctl/client/models"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
	"os"
	"time"
)

func getServerUrl() string {
	if viper.GetString("server_url") != "" {
		return viper.GetString("server_url")
	}
	if viper.GetString("SERVER_URL") != "" {
		return viper.GetString("SERVER_URL")
	}
	fmt.Println("Please provide server url through devtctl.env or flags ")
	os.Exit(1)
	return ""
}

func getAuthToken() string {
	if viper.GetString("auth_token") != "" {
		return viper.GetString("auth_token")
	}
	if viper.GetString("API_TOKEN") != "" {
		return viper.GetString("API_TOKEN")
	}
	fmt.Println("Please provide auth token through devtctl.env or flags ")
	os.Exit(1)
	return ""
}

// Create a Resty Client
var client = resty.New().
	SetHeader("Accept", "application/json").
	SetRetryCount(3).
	SetRetryWaitTime(500 * time.Millisecond)

func CallGetApi[T any](api string, query map[string]string, result *T) error {

	resp, err := client.R().
		SetQueryParams(query).
		SetResult(result).
		EnableTrace().
		SetHeader("token", getAuthToken()).
		Get(getServerUrl() + api)

	if resp.StatusCode() != 200 {
		err = fmt.Errorf("API %s call failed with reason %s", api, string(resp.Body()))
	}
	return err
}

func CallPostApi[T, R any](api string, request R, response *T) error {

	resp, err := client.R().
		SetBody(request).
		SetResult(response).
		SetHeader("token", getAuthToken()).
		EnableTrace().
		Post(getServerUrl() + api)

	if resp.StatusCode() != 200 {
		err = fmt.Errorf("API %s call failed with reason %s", api, string(resp.Body()))
	}
	return err
}

func CallPostApiWithFiles[R any](api string, fileNameToFilePathMap map[string]string, request R, response *models.Response[string]) error {
	resp, err := client.R().
		SetFiles(fileNameToFilePathMap).
		SetBody(request).
		SetResult(response).
		SetHeader("token", getAuthToken()).
		EnableTrace().
		Post(getServerUrl() + api)

	if resp.StatusCode() != 200 {
		err = fmt.Errorf("API %s call failed with reason %s", api, string(resp.Body()))
	}
	return err
}

func CallPutApi[T, R any](api string, request R, response *T) error {

	resp, err := client.R().
		SetBody(request).
		SetResult(response).
		SetHeader("token", getAuthToken()).
		EnableTrace().
		Put(getServerUrl() + api)

	if resp.StatusCode() != 200 {
		err = fmt.Errorf("API %s call failed with reason %s", api, string(resp.Body()))
	}
	return err
}

func CallDeleteApi[T any](api string, query map[string]string, response *T) error {
	resp, err := client.R().
		SetQueryParams(query).
		SetResult(response).
		SetHeader("token", getAuthToken()).
		EnableTrace().
		Delete(getServerUrl() + api)
	if resp.StatusCode() != 204 {
		err = fmt.Errorf("API %s call failed with reason %s", api, string(resp.Body()))
	}
	return err
}
