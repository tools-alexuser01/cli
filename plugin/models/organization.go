package plugin_models

type Organization struct {
	Guid            string
	Name            string
	QuotaDefinition QuotaFields
	Spaces          []SpaceSummary
	Domains         []DomainFields
	SpaceQuotas     []SpaceQuotaFields
}
