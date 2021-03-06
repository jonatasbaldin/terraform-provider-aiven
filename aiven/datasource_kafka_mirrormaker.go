package aiven

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceKafkaMirrormaker() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourceServiceRead,
		Schema:      resourceSchemaAsDatasourceSchema(aivenKafkaMirrormakerSchema(), "project", "service_name"),
	}
}
