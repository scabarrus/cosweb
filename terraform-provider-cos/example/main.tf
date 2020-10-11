
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
  load_config_file = "true"
  apiKey = "KBHGG68XYbONIYiNXS71vpilBr9mYkBGHWVBhB397z0n"
  bearer = "KBHGG68XYbONIYiNXS71vpilBr9mYkBGHWVBhB397z0n"
  endpoint = "192.168.169.129:30061"
  # if you have a self-signed cert
  protocol = "http"
  sslCheck = false
}