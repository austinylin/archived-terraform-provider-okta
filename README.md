# terraform-provider-okta
_A very unstable, early pass at a terraform provider for Okta._

Currently the only thing this provider can do is manage Okta groups, but I'll be expanding the functionality over time in parallel with expanding the Okta API Client ([austinylin/go-okta][]).

## Install
```
go get github.com/austinylin/terraform-provider-okta
```

## Configuration
```
provider "okta" {
  base_url  = "https://{domain.okta[preview].com}/api/v1/"
  api_token = "{apiToken}"
}
```

## Make Something Cool

### Groups
You can manage Okta groups using Terraform as shown below. 

```
resource "okta_group" "tf_users" {
  name        = "Terraform Users"
  description = "A group for cool people using Terraform."
}
```

This will create a group with the given name, and a description of "[Managed by TF] + {description}". The prefix will move to a config variable at somepoint.