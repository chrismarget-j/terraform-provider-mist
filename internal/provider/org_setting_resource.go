package provider

import (
	"context"
	"fmt"

	"github.com/Juniper/terraform-provider-mist/internal/resource_org_setting"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource              = &orgSettingResource{}
	_ resource.ResourceWithConfigure = &orgSettingResource{}
)

func NewOrgSettingResource() resource.Resource {
	return &orgSettingResource{}
}

type orgSettingResource struct {
	client mistapi.ClientInterface
}

func (r *orgSettingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist OrgSetting client")
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

	r.client = client
}
func (r *orgSettingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_setting"
}

func (r *orgSettingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategorySite + "This resource manages the Org Settings." +
			"The Org Settings can be used to customize the Org configuration",
		Attributes: resource_org_setting.OrgSettingResourceSchema(ctx).Attributes,
	}
}

func (r *orgSettingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting OrgSetting Create")
	var plan, state resource_org_setting.OrgSettingModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgSetting, diags := resource_org_setting.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Info(ctx, "Starting OrgSetting Create for Org "+plan.OrgId.ValueString())
	orgId, err := uuid.Parse(plan.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting org_id",
			"Could not get org_id, unexpected error: "+err.Error(),
		)
		return
	}
	data, err := r.client.OrgsSetting().UpdateOrgSettings(ctx, orgId, orgSetting)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating orgSetting",
			"Could not create orgSetting, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_org_setting.SdkToTerraform(ctx, &data.Data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}

func (r *orgSettingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_org_setting.OrgSettingModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting OrgSetting Read: orgSetting_id "+state.OrgId.ValueString())
	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting org_id",
			"Could not get org_id, unexpected error: "+err.Error(),
		)
		return
	}
	data, err := r.client.OrgsSetting().GetOrgSettings(ctx, orgId)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting orgSetting",
			"Could not get orgSetting, unexpected error: "+err.Error(),
		)
		return
	}
	state, diags = resource_org_setting.SdkToTerraform(ctx, &data.Data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *orgSettingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_org_setting.OrgSettingModel
	tflog.Info(ctx, "Starting OrgSetting Update")

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	diags = req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgSetting, diags := resource_org_setting.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting OrgSetting Update for OrgSetting "+state.OrgId.ValueString())
	orgId, err := uuid.Parse(plan.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting org_id",
			"Could not get org_id, unexpected error: "+err.Error(),
		)
		return
	}
	data, err := r.client.OrgsSetting().UpdateOrgSettings(ctx, orgId, orgSetting)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating orgSetting",
			"Could not update orgSetting, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_org_setting.SdkToTerraform(ctx, &data.Data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}

func (r *orgSettingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_org_setting.OrgSettingModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting OrgSetting Delete: orgSetting_id "+state.OrgId.ValueString())
	orgSetting, e := resource_org_setting.DeleteTerraformToSdk(ctx)
	diags.Append(e...)
	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting org_id",
			"Could not get org_id, unexpected error: "+err.Error(),
		)
		return
	}
	_, err = r.client.OrgsSetting().UpdateOrgSettings(ctx, orgId, orgSetting)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting orgSetting",
			"Could not delete orgSetting, unexpected error: "+err.Error(),
		)
		return
	}
}
