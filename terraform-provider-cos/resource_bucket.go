package cos
import (
	"context"
  
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"strconv"
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
		 "instanceid": &schema.Schema{
			Type:     schema.TypeInt,
			Required: true,
		 },
	  },
	}
  }

  func resourceBucketCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] resourceBucketCreate -start")
	// Warning or errors can be collected in a slice type
	//var diags diag.Diagnostics
	client := m.(CosWeb)
	bucketName:=d.Get("name").(string)
	bucketDescription:=d.Get("description").(string)
	cosInstanceID:=d.Get("instanceid").(int)
	bucket,err:=client.CreateBucket(cosInstanceID, bucketName, bucketDescription)
	if err != nil {
		log.Printf("[DEBUG] resourceBucketCreate - error : %s",err)
		return diag.FromErr(err)
	}
	id:=strconv.Itoa(bucket.ID)
	d.SetId(id)
	d.Set("name",bucket.Name)
	d.Set("description",bucket.Description)
	d.Set("instanceid",bucket.CosInstanceGUID)
	log.Printf("[DEBUG] resourceBucketCreate -end")
	return nil
  }

  func resourceBucketRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
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