package provider

import (
	"net/http"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func NewMsgVpnClientProfileResource() resource.Resource {
	return &solaceResource[MsgVpnClientProfile]{spr: &clientProfileResource{}}
}

var _ solaceProviderResource[MsgVpnClientProfile] = &clientProfileResource{}

type clientProfileResource struct {
	*solaceProvider
}

func (t clientProfileResource) Name() string {
	return "clientprofile"
}

func (r clientProfileResource) Schema() schema.Schema {
	return MsgVpnClientProfileResourceSchema("msg_vpn_name", "client_profile_name")
}

func (r *clientProfileResource) SetProvider(provider *solaceProvider) {
	r.solaceProvider = provider
}

func (r clientProfileResource) NewData() *MsgVpnClientProfile {
	return &MsgVpnClientProfile{}
}

func (r clientProfileResource) Create(data *MsgVpnClientProfile, diag *diag.Diagnostics) (*http.Response, error) {
	apiReq := r.Client.ClientProfileApi.CreateMsgVpnClientProfile(r.Context, *data.MsgVpnName).Body(*data.ToApi())
	_, httpResponse, err := apiReq.Execute()
	return httpResponse, err
}

func (r clientProfileResource) Read(data *MsgVpnClientProfile, diag *diag.Diagnostics) (*http.Response, error) {
	apiReq := r.Client.ClientProfileApi.GetMsgVpnClientProfile(r.Context, *data.MsgVpnName, *data.ClientProfileName)
	apiResponse, httpResponse, err := apiReq.Execute()
	if err == nil && apiResponse != nil && apiResponse.Data != nil {
		data.ToTF(apiResponse.Data)
	}
	return httpResponse, err
}

func (r clientProfileResource) Update(_ *MsgVpnClientProfile, data *MsgVpnClientProfile, diag *diag.Diagnostics) (*http.Response, error) {
	apiReq := r.Client.ClientProfileApi.UpdateMsgVpnClientProfile(r.Context, *data.MsgVpnName, *data.ClientProfileName).Body(*data.ToApi())
	_, httpResponse, err := apiReq.Execute()
	return httpResponse, err
}

func (r clientProfileResource) Delete(data *MsgVpnClientProfile, diag *diag.Diagnostics) (*http.Response, error) {
	apiReq := r.Client.ClientProfileApi.DeleteMsgVpnClientProfile(r.Context, *data.MsgVpnName, *data.ClientProfileName)
	_, httpResponse, err := apiReq.Execute()
	return httpResponse, err
}

var msgVpnClientProfileImportRegexp *regexp.Regexp = regexp.MustCompile(
	"^([^\\s\\*\\?\\/]+)\\/([0-9a-zA-Z_\\-]+)$")

func (r clientProfileResource) Import(id string, data *MsgVpnClientProfile, diag *diag.Diagnostics) {
	match := msgVpnClientProfileImportRegexp.FindStringSubmatch(id)
	if match != nil {
		data.MsgVpnName = &match[1]
		data.ClientProfileName = &match[2]
	} else {
		diag.AddError("Expected <vpn-name>/<client-profile>", id+" does not match "+msgVpnClientProfileImportRegexp.String())
		return
	}
}
