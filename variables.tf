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

variable "name" {
  description = "The name of the Log Analytics Storage Insights"
  type        = string
}

variable "resource_group_name" {
  description = "The name of the resource group"
  type        = string
}

variable "workspace_id" {
  description = "The ID of the Log Analytics workspace"
  type        = string
}

variable "storage_account_id" {
  description = "The ID of the storage account"
  type        = string
}

variable "storage_account_key" {
  description = "The key of the storage account"
  type        = string
}

variable "blob_container_names" {
  description = "A list of blob container names"
  type        = list(string)
  default     = null
}

variable "table_names" {
  description = "A list of table names"
  type        = list(string)
  default     = null
}
