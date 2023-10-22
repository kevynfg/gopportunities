package queries

import "fmt"

func FindAllOpportunitiesQuery(queryLimit int, queryOffset int) string {
	return fmt.Sprintf(`SELECT opportunities.*, GROUP_CONCAT(technologies.name) as technology_names, company.name as company_name FROM opportunities INNER JOIN opportunity_technologies as ot ON opportunities.id = ot.opportunity_id INNER JOIN technologies ON ot.opportunity_technology_id = technologies.id INNER JOIN company ON opportunities.company_id = company.id GROUP BY opportunities.id ORDER BY opportunities.id DESC LIMIT %v OFFSET %v`, queryLimit, queryOffset);
}