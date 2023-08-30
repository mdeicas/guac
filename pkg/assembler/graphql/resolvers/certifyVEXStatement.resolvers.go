package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"strings"

	"github.com/guacsec/guac/pkg/assembler/graphql/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// IngestVEXStatement is the resolver for the ingestVEXStatement field.
func (r *mutationResolver) IngestVEXStatement(ctx context.Context, subject model.PackageOrArtifactInput, vulnerability model.VulnerabilityInputSpec, vexStatement model.VexStatementInputSpec) (string, error) {
	if strings.ToLower(vulnerability.Type) == "novuln" {
		return "", gqlerror.Errorf("novuln type cannot be used for VEX")
	}
	// vulnerability input (type and vulnerability ID) will be enforced to be lowercase
	ingestedVEXStatement, err := r.Backend.IngestVEXStatement(ctx, subject,
		model.VulnerabilityInputSpec{Type: strings.ToLower(vulnerability.Type), VulnerabilityID: strings.ToLower(vulnerability.VulnerabilityID)},
		vexStatement)
	if err == nil {
		return "", err
	}
	return ingestedVEXStatement.ID, err
}

// IngestVEXStatements is the resolver for the ingestVEXStatements field.
func (r *mutationResolver) IngestVEXStatements(ctx context.Context, subjects model.PackageOrArtifactInputs, vulnerabilities []*model.VulnerabilityInputSpec, vexStatements []*model.VexStatementInputSpec) ([]string, error) {
	// vulnerability input (type and vulnerability ID) will be enforced to be lowercase
	valuesDefined := 0
	if len(subjects.Packages) > 0 {
		if len(subjects.Packages) != len(vexStatements) {
			return []string{}, gqlerror.Errorf("uneven packages and vexStatements for ingestion")
		}
		valuesDefined = valuesDefined + 1
	}
	if len(subjects.Artifacts) > 0 {
		if len(subjects.Artifacts) != len(vexStatements) {
			return []string{}, gqlerror.Errorf("uneven artifact and vexStatements for ingestion")
		}
		valuesDefined = valuesDefined + 1
	}
	if len(vulnerabilities) != len(vexStatements) {
		return []string{}, gqlerror.Errorf("uneven vulnerabilities and vexStatements for ingestion")
	}
	if valuesDefined != 1 {
		return []string{}, gqlerror.Errorf("must specify at most packages or artifacts for %v", "IngestVEXStatements")
	}

	var lowercaseVulnInputList []*model.VulnerabilityInputSpec
	for _, v := range vulnerabilities {
		if strings.ToLower(v.Type) == "novuln" {
			return []string{}, gqlerror.Errorf("novuln type cannot be used for VEX")
		}
		lowercaseVulnInput := model.VulnerabilityInputSpec{
			Type:            strings.ToLower(v.Type),
			VulnerabilityID: strings.ToLower(v.VulnerabilityID),
		}
		lowercaseVulnInputList = append(lowercaseVulnInputList, &lowercaseVulnInput)
	}
	return r.Backend.IngestVEXStatements(ctx, subjects, lowercaseVulnInputList, vexStatements)
}

// CertifyVEXStatement is the resolver for the CertifyVEXStatement field.
func (r *queryResolver) CertifyVEXStatement(ctx context.Context, certifyVEXStatementSpec model.CertifyVEXStatementSpec) ([]*model.CertifyVEXStatement, error) {
	// vulnerability input (type and vulnerability ID) will be enforced to be lowercase

	if certifyVEXStatementSpec.Vulnerability != nil {
		var typeLowerCase *string = nil
		var vulnIDLowerCase *string = nil
		if certifyVEXStatementSpec.Vulnerability.Type != nil {
			lower := strings.ToLower(*certifyVEXStatementSpec.Vulnerability.Type)
			typeLowerCase = &lower
		}
		if certifyVEXStatementSpec.Vulnerability.VulnerabilityID != nil {
			lower := strings.ToLower(*certifyVEXStatementSpec.Vulnerability.VulnerabilityID)
			vulnIDLowerCase = &lower
		}

		lowercaseVulnFilter := &model.VulnerabilitySpec{
			ID:              certifyVEXStatementSpec.Vulnerability.ID,
			Type:            typeLowerCase,
			VulnerabilityID: vulnIDLowerCase,
			NoVuln:          certifyVEXStatementSpec.Vulnerability.NoVuln,
		}

		lowercaseCertifyVexFilter := &model.CertifyVEXStatementSpec{
			ID:               certifyVEXStatementSpec.ID,
			Subject:          certifyVEXStatementSpec.Subject,
			Vulnerability:    lowercaseVulnFilter,
			Status:           certifyVEXStatementSpec.Status,
			VexJustification: certifyVEXStatementSpec.VexJustification,
			Statement:        certifyVEXStatementSpec.Statement,
			StatusNotes:      certifyVEXStatementSpec.StatusNotes,
			KnownSince:       certifyVEXStatementSpec.KnownSince,
			Origin:           certifyVEXStatementSpec.Origin,
			Collector:        certifyVEXStatementSpec.Collector,
		}
		return r.Backend.CertifyVEXStatement(ctx, lowercaseCertifyVexFilter)
	} else {
		return r.Backend.CertifyVEXStatement(ctx, &certifyVEXStatementSpec)
	}
}
