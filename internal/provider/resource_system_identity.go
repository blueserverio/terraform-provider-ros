package ros

import (
	"context"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemIdentity() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Setting the System's Identity provides a unique identifying name for when the system identifies itself to other routers in the network and when accessing services such as DHCP, Neighbour Discovery, and default wireless SSID. The default system Identity is set to 'MikroTik'.		.",

		CreateContext: resourceSystemIdentityCreate,
		ReadContext:   resourceSystemIdentityRead,
		UpdateContext: resourceSystemIdentityUpdate,
		DeleteContext: resourceSystemIdentityDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Description: 	"The display name for the device.",
				Type:        	schema.TypeString,
				Required:		true,
			},
		},
	}
}

func resourceSystemIdentityCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	//client := meta.(*)
	
	idFromAPI := "my-id"
	d.SetId(idFromAPI)

	tflog.Trace(ctx, "created a system identity resource")

	return diag.Errorf("not implemented")
}

func resourceSystemIdentityRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}

func resourceSystemIdentityUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}

func resourceSystemIdentityDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}
