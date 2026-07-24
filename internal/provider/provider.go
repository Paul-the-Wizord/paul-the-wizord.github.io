package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// New returns a factory function for creating a new HelloProvider.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &HelloProvider{version: version}
	}
}

// HelloProvider is the root provider implementation.
type HelloProvider struct {
	version string
}

// HelloProviderModel describes the provider configuration schema.
type HelloProviderModel struct{}

// Metadata returns the provider type name.
func (p *HelloProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "hello"
	resp.Version = p.version
}

// Schema defines the provider configuration schema.
func (p *HelloProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{}
}

// Configure is a no-op for this provider.
func (p *HelloProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config HelloProviderModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Resources returns the resources offered by this provider.
func (p *HelloProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewHelloWorldResource,
	}
}

// DataSources returns no data sources for this provider.
func (p *HelloProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return nil
}
