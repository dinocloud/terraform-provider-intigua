provider "intigua" {
  "host"     = "localhost"
  "username" = "admin"
  "api_key"  = "a"
}

resource "intigua_endpoint" "test" {
  "os"      = "test"
  "ip"      = "1.1.1.1"
  "dnsname" = "hola.com"
  "name"    = "test"
}

resource "intigua_connector" "test" {
  "endpoint_connector_url" = "www.url.com"
}

resource "intigua_credential" "test" {
  "accountname" = "test"
  "username"    = "1.1.1.1"
  "password"    = "hola.com"
  "type"        = "test"
}

resource "intigua_credential_association" "test" {
  "endpoint_id"   = "www.url.com"
  "credential_id" = "intigua_credential"
}

resource "intigua_management_service_registration" "test" {
  "name"         = "test"
  "version"      = "1.1.1.1"
  "configpackage" = "hola.com"
  "endpoint_id"  = "test"
}
