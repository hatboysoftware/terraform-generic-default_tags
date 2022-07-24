output "resource_group_name" {
  value = module.test_rg.resource_group.name
}

output "subscription_id" {
  value = data.azurerm_client_config.current.subscription_id
}

output "tag_name" {
  value = keys(module.test_rg.resource_group.tags)[0]
}

output "tag_value" {
  value = values(module.test_rg.resource_group.tags)[0]
}