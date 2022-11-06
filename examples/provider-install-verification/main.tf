terraform {
  required_providers {
    ros = {
      source = "blueserverio/ros"
    }
  }
}

provider "ros" {
    hosturl  = "https://192.168.88.1"
    username = "admin"
    password = ""
    insecure = true
}

resource "ros_system_identity" "MikroTik" {
    name = "UpdatedIdentity-01"
}
