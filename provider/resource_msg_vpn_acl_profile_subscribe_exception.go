package provider

import (
	"context"
	"net/http"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

var _ provider.ResourceType = aclProfileResourceType{}

type aclProfileSubscribeExceptionResourceType struct {
}

func (t aclProfileSubscribeExceptionResourceType) NewResource(ctx context.Context, in provider.Provider) (resource.Resource, diag.Diagnostics) {
	solaceProvider, diags := convertProviderType(in)

	return NewResource[MsgVpnAclProfileSubscribeException](
		aclProfileSubscribeExceptionResource{solaceProvider: solaceProvider}), diags
}

func (t aclProfileSubscribeExceptionResourceType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return MsgVpnAclProfileSubscribeExceptionSchema("msg_vpn_name", "acl_profile_name", "topic_syntax", "subscribe_exception_topic"), nil
}

var _ solaceProviderResource[MsgVpnAclProfile] = aclProfileResource{}

type aclProfileSubscribeExceptionResource struct {
	solaceProvider
}

func (r aclProfileSubscribeExceptionResource) NewData() *MsgVpnAclProfileSubscribeException {
	return &MsgVpnAclProfileSubscribeException{}
}

func (r aclProfileSubscribeExceptionResource) Create(data *MsgVpnAclProfileSubscribeException, diag *diag.Diagnostics) (*http.Response, error) {
	_, httpResponse, err := r.Client.AclProfileApi.
		CreateMsgVpnAclProfileSubscribeException(r.Context, *data.MsgVpnName, *data.AclProfileName).
		Body(*data.ToApi()).
		Execute()
	return httpResponse, err
}

func (r aclProfileSubscribeExceptionResource) Read(data *MsgVpnAclProfileSubscribeException, diag *diag.Diagnostics) (*http.Response, error) {
	apiResponse, httpResponse, err := r.Client.AclProfileApi.
		GetMsgVpnAclProfileSubscribeException(
			r.Context, *data.MsgVpnName, *data.AclProfileName, *data.TopicSyntax, *data.SubscribeExceptionTopic).
		Execute()
	if err == nil && apiResponse != nil && apiResponse.Data != nil {
		data.ToTF(apiResponse.Data)
	}
	return httpResponse, err
}

func (r aclProfileSubscribeExceptionResource) Update(_ *MsgVpnAclProfileSubscribeException, data *MsgVpnAclProfileSubscribeException, diag *diag.Diagnostics) (*http.Response, error) {
	diag.AddError("Update not supported", "Acl Profile Client Connect Exceptions cannot be updated")
	return nil, nil
}

func (r aclProfileSubscribeExceptionResource) Delete(data *MsgVpnAclProfileSubscribeException, diag *diag.Diagnostics) (*http.Response, error) {
	_, httpResponse, err := r.Client.AclProfileApi.
		DeleteMsgVpnAclProfileSubscribeException(
			r.Context, *data.MsgVpnName, *data.AclProfileName, *data.TopicSyntax, *data.SubscribeExceptionTopic).
		Execute()
	return httpResponse, err
}

var msgVpnAclProfileSubscribeExceptionImportRegexp *regexp.Regexp = regexp.MustCompile(
	"^([^\\s\\*\\?\\/]+)\\/([0-9a-zA-Z_\\-]+)\\/(smf|mqtt)\\/([^\\s]+)$")

func (r aclProfileSubscribeExceptionResource) Import(id string, data *MsgVpnAclProfileSubscribeException, diag *diag.Diagnostics) {
	match := msgVpnAclProfileSubscribeExceptionImportRegexp.FindStringSubmatch(id)
	if match != nil {
		data.MsgVpnName = &match[1]
		data.AclProfileName = &match[2]
		data.TopicSyntax = &match[3]
		data.SubscribeExceptionTopic = &match[4]
	} else {
		diag.AddError("Expected <vpn-name>/<acl-profile>/<topic syntax>/<topic name>",
			id+" does not match "+msgVpnAclProfileSubscribeExceptionImportRegexp.String())
		return
	}
}
