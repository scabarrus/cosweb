package cos
import (
	"context"
  
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
  )
  
  func resourceBucket() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceBucketCreate,
		ReadContext:   resourceBucketRead,
		UpdateContext: resourceBucketUpdate,
	    DeleteContext: resourceBucketDelete,
	  
	  Schema: map[string]*schema.Schema{
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		 },
		 "description": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		 },
		 "id": &schema.Schema{
			Type:     schema.TypeInt,
			Computed: true,
		 },
		 "instanceid": &schema.Schema{
			Type:     schema.TypeInt,
			Required: true,
		 },
	  },
	}
  }

  func resourceBucketCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
  
	return diags
  }

  func resourceBucketRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	client := m.(CosWeb)
	bucketName:=d.Get("name").(string)
	bucketDescription:=d.Get("description").(string)
	cosInstanceId:=d.Get("instance").(int)
	_,err:=client.CreateBucket(cosInstanceId, bucketName, bucketDescription)

	if err != nil {
		return diag.FromErr(err)
	}
	return diags
  }

  func resourceBucketUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
  
	return diags
  }

  func resourceBucketDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
  
	return diags
  }