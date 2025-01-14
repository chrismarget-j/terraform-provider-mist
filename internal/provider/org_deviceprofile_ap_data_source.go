package provider

import (
	"context"
	"fmt"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/Juniper/terraform-provider-mist/internal/datasource_org_deviceprofiles_ap"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = (*orgDeviceprofilesApDataSource)(nil)

func NewOrgDeviceprofilesApDataSource() datasource.DataSource {
	return &orgDeviceprofilesApDataSource{}
}

type orgDeviceprofilesApDataSource struct {
	client mistapi.ClientInterface
}

func (d *orgDeviceprofilesApDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist AP Stats")
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(mistapi.ClientInterface)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *mistapigo.APIClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	d.client = client
}
func (d *orgDeviceprofilesApDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_deviceprofiles_ap"
}

func (d *orgDeviceprofilesApDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryDevices + "This data source provides the list of AP Device Profiles." +
			"An AP Device Profile can be used to define common configuration accross multiple Wireless Access Points" +
			"and is assigned to one or multiple APs as a deviceprofile with the `mist_org_deviceprofile_assign` resource",
		Attributes: datasource_org_deviceprofiles_ap.OrgDeviceprofilesApDataSourceSchema(ctx).Attributes,
	}
}

func (d *orgDeviceprofilesApDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var ds datasource_org_deviceprofiles_ap.OrgDeviceprofilesApModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &ds)...)

	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(ds.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting org_id from ds",
			"Could not get org_id, unexpected error: "+err.Error(),
		)
		return
	}

	var limit *int = models.ToPointer(1000)
	var page *int
	mType := models.ToPointer(models.DeviceTypeEnum("ap"))

	data, err := d.client.OrgsDeviceProfiles().ListOrgDeviceProfiles(ctx, orgId, mType, page, limit)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting AP Stats",
			"Could not get AP Stats, unexpected error: "+err.Error(),
		)
		return
	}

	deviceApStat, diags := datasource_org_deviceprofiles_ap.SdkToTerraform(ctx, data.Data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := resp.State.SetAttribute(ctx, path.Root("deviceprofiles"), deviceApStat); err != nil {
		resp.Diagnostics.Append(err...)
		return
	}
}
