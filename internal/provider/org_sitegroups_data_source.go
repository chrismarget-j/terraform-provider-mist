package provider

import (
	"context"
	"fmt"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/Juniper/terraform-provider-mist/internal/datasource_org_sitegroups"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = (*orgSitegroupsDataSource)(nil)

func NewOrgSitegroupsDataSource() datasource.DataSource {
	return &orgSitegroupsDataSource{}
}

type orgSitegroupsDataSource struct {
	client mistapi.ClientInterface
}

func (d *orgSitegroupsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
func (d *orgSitegroupsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_sitegroups"
}

func (d *orgSitegroupsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryWlan + "This data source provides the list of Org Site Groups (sitegroups)." +
			"A site group is a feature that allows users to group multiple sites together based on regions, functions, or other parameters for efficient management of devices. " +
			"Sites can exist in multiple groups simultaneously, and site groups can be used to ensure consistent settings, manage administrator access, and apply specific templates to groups of sites.",
		Attributes: datasource_org_sitegroups.OrgSitegroupsDataSourceSchema(ctx).Attributes,
	}
}

func (d *orgSitegroupsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var ds datasource_org_sitegroups.OrgSitegroupsModel
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

	data, err := d.client.OrgsSitegroups().ListOrgSiteGroups(ctx, orgId, page, limit)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting AP Stats",
			"Could not get AP Stats, unexpected error: "+err.Error(),
		)
		return
	}

	deviceApStat, diags := datasource_org_sitegroups.SdkToTerraform(ctx, data.Data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := resp.State.SetAttribute(ctx, path.Root("org_sitegroups"), deviceApStat); err != nil {
		resp.Diagnostics.Append(err...)
		return
	}
}
