
terraform {
  required_providers {
    cos = {
      source  = "scabarrus.com/terraform/cos"
      versions = ["1.0.0"]
    }
  }
}


# Configure the VMware vSphere Provider
provider "cos" {
 #version = "1.0.0"
  
  apikey = "KBHGG68oooYkBGHWVBhB397z0n"
  bearer = "KBHGG68XYbONooor9mYkBGHWVBhB397z0n"
  endpoint = "192.168.169.129:30061"
  # if you have a self-signed cert
  protocol = "http"
  sslcheck = false
}

data "cos_instance" "instance" { 
    name="Cloud Object Storage-lsx23" 
 
}

resource "cos_bucket" "bucket" {
  name="cloud-object-storage-lsx23-cos-standard-1a2456"
  description="toto"
  instanceid=29

}
