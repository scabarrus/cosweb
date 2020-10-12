package cos

import (
    "context"
    "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "log"
    //"fmt"
)
//Provider instanciate the provider
//For cosWeb need to provide an apiKey, a bearer token, endpoint, protocol and if ssl certificate need to be checked
//it return the provider sechema with parameter expecting, datasource and resource methods
func Provider() *schema.Provider {
    return &schema.Provider{
      Schema: map[string]*schema.Schema{
        "apikey": &schema.Schema{
          Type:        schema.TypeString,
          //Optional:    true,
          Required: true,
          DefaultFunc: schema.EnvDefaultFunc("COS_APIKEY", nil),
        },
        "bearer": &schema.Schema{
          Type:        schema.TypeString,
          Required:    true,
          Sensitive:   true,
          DefaultFunc: schema.EnvDefaultFunc("COS_BEARER", nil),
        },
        "endpoint": &schema.Schema{
            Type:     schema.TypeString,
            Required: true,
        },
        "protocol": &schema.Schema{
            Type:     schema.TypeString,
            Required: true,
        },
        "sslcheck": &schema.Schema{
            Type:     schema.TypeBool,
            Optional: true,
        },
      },
      ResourcesMap: map[string]*schema.Resource{
          "cos_bucket": resourceBucket(),
      },
      DataSourcesMap: map[string]*schema.Resource{
        "cos_instance": dataSourceCosInstance(),
      },
      ConfigureContextFunc: configureProvider,
    }
  }

//configureProvider instanciate the client use by this provider
//the client is CosWeb struct which provide a set of method for managing cos instance and bucket across the CosWeb microservice
//It returns the client CosWeb or an error
func configureProvider(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
    // Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
    apiKey:=d.Get("apikey").(string)
    bearer:=d.Get("bearer").(string)
    endpoint:=d.Get("endpoint").(string)
    protocol:=d.Get("protocol").(string)
    sslCheck:=d.Get("sslcheck").(bool)
    client :=CosWeb{apiKey,bearer,endpoint,protocol,sslCheck}
    log.Printf("[INFO] Initializing CosWweb Client")
    log.Println("client : ",client)
    return client,diags
}
