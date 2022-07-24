provider "azurerm" {
  features {}
}

module "test_tags" {
  source = "../.."
  tags = {
    "foo" = "bar"
  }
}

module "test_rg" {
  source  = "app.terraform.io/HatBoySoftware/resource_group/azure"
  version = "0.1.2"

  name = "test"
  resource_prefix = "EUS2-HBS-TST"
  tags = module.test_tags.tags
}

data "azurerm_client_config" "current" {}