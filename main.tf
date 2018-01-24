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
