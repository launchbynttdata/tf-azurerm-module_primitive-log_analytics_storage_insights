// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

resource "azurerm_log_analytics_storage_insights" "main" {
  name                 = var.name
  resource_group_name  = var.resource_group_name
  workspace_id         = var.workspace_id
  storage_account_id   = var.storage_account_id
  storage_account_key  = var.storage_account_key
  blob_container_names = var.blob_container_names
  table_names          = var.table_names
}
