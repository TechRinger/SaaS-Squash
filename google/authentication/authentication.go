package authentication

import (
	"SaaS-Squash/common"
	"context"
	"fmt"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func AuthenticateSheet(credential string) (context.Context, *sheets.Service) {

	ctx := context.Background()
	var client *sheets.Service
	var err error
	if common.AllC2Configs.Google.ShowToken {
		oauth2Token, oauth2Error := google.CredentialsFromJSON(ctx, []byte(credential), "https://www.googleapis.com/auth/spreadsheets")
		client, err = sheets.NewService(ctx, option.WithCredentials(oauth2Token))
		if err != nil {
			common.AllC2Configs.Debug.LogFatalDebug("Sheets - OAuth2 Authentication failed: " + oauth2Error.Error())
		}
		fmt.Printf("Sheets Outh2 AccessToken: %s\n", string(oauth2Token.JSON))
		common.AllC2Configs.Debug.LogDebug("Sheets - OAuth2 Client authentication success")
		return ctx, client
	} else {
		client, err = sheets.NewService(ctx, option.WithCredentialsJSON([]byte(credential)))
		if err != nil {
			common.AllC2Configs.Debug.LogFatalDebug("Sheets - Authentication failed")
		}
		common.AllC2Configs.Debug.LogDebug("Sheets - Client authentication success")
		return ctx, client
	}
}

func AuthenticateDrive(credential string) (context.Context, *drive.Service) {

	ctx := context.Background()
	var client *drive.Service
	var err error

	if common.AllC2Configs.Google.ShowToken {
		oauth2Token, oauth2Error := google.CredentialsFromJSON(ctx, []byte(credential), "https://www.googleapis.com/auth/drive")
		client, err = drive.NewService(ctx, option.WithCredentials(oauth2Token))
		if err != nil {
			common.AllC2Configs.Debug.LogFatalDebug("Drive - OAuth2 Authentication failed: " + oauth2Error.Error())
		}
		fmt.Printf("Drive Outh2 AccessToken: %s\n", string(oauth2Token.JSON))
		common.AllC2Configs.Debug.LogDebug("Drive - OAuth2 Client authentication success")
		return ctx, client
	} else {

		client, err = drive.NewService(ctx, option.WithCredentialsJSON([]byte(credential)))
		if err != nil {
			common.AllC2Configs.Debug.LogFatalDebug("Drive - Authentication failed")
		}
		common.AllC2Configs.Debug.LogDebug("Drive - Client authentication success")
		return ctx, client
	}
}
