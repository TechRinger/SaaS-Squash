package excel

import (
	"SaaS-Squash/common"
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	azidentity "github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

// func (c *Client) Authenticate() (context.Context, *msgraphsdk.GraphServiceClient) {
func (c *ExcelClient) Authenticate() (context.Context, string) {
	ctx := context.Background()

	//retrieve the credential
	cred, err := azidentity.NewClientSecretCredential(common.AllC2Configs.O365.TenantId, common.AllC2Configs.O365.ClientId, common.AllC2Configs.O365.ClientSecret, nil)

	if err != nil {
		common.AllC2Configs.Debug.LogFatalDebug("Excel - Client authentication failed")
	} else {
		common.AllC2Configs.Debug.LogDebug("Excel - Client authentication success")
	}

	tkn, err := cred.GetToken(ctx, policy.TokenRequestOptions{
		Scopes: []string{
			"https://graph.microsoft.com/.default",
		},
	})
	if err != nil {
		common.AllC2Configs.Debug.LogFatalDebug("Excel - Token creation failed" + err.Error())
	} else {
		common.AllC2Configs.Debug.LogDebug("Excel - Token success")
	}
	c.APIKey = tkn.Token
	c.AuthExpire = tkn.ExpiresOn
	//enable to display token
	//common.AllC2Configs.Debug.LogDebug(tkn.Token)

	return ctx, tkn.Token
}
