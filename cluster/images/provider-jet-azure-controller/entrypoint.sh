#!/usr/bin/env bash

terraform-provider-azurerm -debuggable &
crossplane-jet-azure-provider "$@"
