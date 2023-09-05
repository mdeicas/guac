// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (a *Artifact) Occurrences(ctx context.Context) (result []*Occurrence, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = a.NamedOccurrences(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = a.Edges.OccurrencesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = a.QueryOccurrences().All(ctx)
	}
	return result, err
}

func (a *Artifact) Sbom(ctx context.Context) (result []*BillOfMaterials, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = a.NamedSbom(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = a.Edges.SbomOrErr()
	}
	if IsNotLoaded(err) {
		result, err = a.QuerySbom().All(ctx)
	}
	return result, err
}

func (a *Artifact) Attestations(ctx context.Context) (result []*SLSAAttestation, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = a.NamedAttestations(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = a.Edges.AttestationsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = a.QueryAttestations().All(ctx)
	}
	return result, err
}

func (a *Artifact) Same(ctx context.Context) (result []*HashEqual, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = a.NamedSame(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = a.Edges.SameOrErr()
	}
	if IsNotLoaded(err) {
		result, err = a.QuerySame().All(ctx)
	}
	return result, err
}

func (bom *BillOfMaterials) Package(ctx context.Context) (*PackageVersion, error) {
	result, err := bom.Edges.PackageOrErr()
	if IsNotLoaded(err) {
		result, err = bom.QueryPackage().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (bom *BillOfMaterials) Artifact(ctx context.Context) (*Artifact, error) {
	result, err := bom.Edges.ArtifactOrErr()
	if IsNotLoaded(err) {
		result, err = bom.QueryArtifact().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (b *Builder) SlsaAttestations(ctx context.Context) (result []*SLSAAttestation, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = b.NamedSlsaAttestations(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = b.Edges.SlsaAttestationsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = b.QuerySlsaAttestations().All(ctx)
	}
	return result, err
}

func (c *Certification) Source(ctx context.Context) (*SourceName, error) {
	result, err := c.Edges.SourceOrErr()
	if IsNotLoaded(err) {
		result, err = c.QuerySource().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (c *Certification) PackageVersion(ctx context.Context) (*PackageVersion, error) {
	result, err := c.Edges.PackageVersionOrErr()
	if IsNotLoaded(err) {
		result, err = c.QueryPackageVersion().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (c *Certification) AllVersions(ctx context.Context) (*PackageName, error) {
	result, err := c.Edges.AllVersionsOrErr()
	if IsNotLoaded(err) {
		result, err = c.QueryAllVersions().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (c *Certification) Artifact(ctx context.Context) (*Artifact, error) {
	result, err := c.Edges.ArtifactOrErr()
	if IsNotLoaded(err) {
		result, err = c.QueryArtifact().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (cs *CertifyScorecard) Scorecard(ctx context.Context) (*Scorecard, error) {
	result, err := cs.Edges.ScorecardOrErr()
	if IsNotLoaded(err) {
		result, err = cs.QueryScorecard().Only(ctx)
	}
	return result, err
}

func (cs *CertifyScorecard) Source(ctx context.Context) (*SourceName, error) {
	result, err := cs.Edges.SourceOrErr()
	if IsNotLoaded(err) {
		result, err = cs.QuerySource().Only(ctx)
	}
	return result, err
}

func (cv *CertifyVex) Package(ctx context.Context) (*PackageVersion, error) {
	result, err := cv.Edges.PackageOrErr()
	if IsNotLoaded(err) {
		result, err = cv.QueryPackage().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (cv *CertifyVex) Artifact(ctx context.Context) (*Artifact, error) {
	result, err := cv.Edges.ArtifactOrErr()
	if IsNotLoaded(err) {
		result, err = cv.QueryArtifact().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (cv *CertifyVex) Vulnerability(ctx context.Context) (*VulnerabilityType, error) {
	result, err := cv.Edges.VulnerabilityOrErr()
	if IsNotLoaded(err) {
		result, err = cv.QueryVulnerability().Only(ctx)
	}
	return result, err
}

func (cv *CertifyVuln) Vulnerability(ctx context.Context) (*VulnerabilityID, error) {
	result, err := cv.Edges.VulnerabilityOrErr()
	if IsNotLoaded(err) {
		result, err = cv.QueryVulnerability().Only(ctx)
	}
	return result, err
}

func (cv *CertifyVuln) Package(ctx context.Context) (*PackageVersion, error) {
	result, err := cv.Edges.PackageOrErr()
	if IsNotLoaded(err) {
		result, err = cv.QueryPackage().Only(ctx)
	}
	return result, err
}

func (d *Dependency) Package(ctx context.Context) (*PackageVersion, error) {
	result, err := d.Edges.PackageOrErr()
	if IsNotLoaded(err) {
		result, err = d.QueryPackage().Only(ctx)
	}
	return result, err
}

func (d *Dependency) DependentPackage(ctx context.Context) (*PackageName, error) {
	result, err := d.Edges.DependentPackageOrErr()
	if IsNotLoaded(err) {
		result, err = d.QueryDependentPackage().Only(ctx)
	}
	return result, err
}

func (hsa *HasSourceAt) PackageVersion(ctx context.Context) (*PackageVersion, error) {
	result, err := hsa.Edges.PackageVersionOrErr()
	if IsNotLoaded(err) {
		result, err = hsa.QueryPackageVersion().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (hsa *HasSourceAt) AllVersions(ctx context.Context) (*PackageName, error) {
	result, err := hsa.Edges.AllVersionsOrErr()
	if IsNotLoaded(err) {
		result, err = hsa.QueryAllVersions().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (hsa *HasSourceAt) Source(ctx context.Context) (*SourceName, error) {
	result, err := hsa.Edges.SourceOrErr()
	if IsNotLoaded(err) {
		result, err = hsa.QuerySource().Only(ctx)
	}
	return result, err
}

func (he *HashEqual) Artifacts(ctx context.Context) (result []*Artifact, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = he.NamedArtifacts(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = he.Edges.ArtifactsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = he.QueryArtifacts().All(ctx)
	}
	return result, err
}

func (iv *IsVulnerability) Osv(ctx context.Context) (*VulnerabilityType, error) {
	result, err := iv.Edges.OsvOrErr()
	if IsNotLoaded(err) {
		result, err = iv.QueryOsv().Only(ctx)
	}
	return result, err
}

func (iv *IsVulnerability) Vulnerability(ctx context.Context) (*VulnerabilityType, error) {
	result, err := iv.Edges.VulnerabilityOrErr()
	if IsNotLoaded(err) {
		result, err = iv.QueryVulnerability().Only(ctx)
	}
	return result, err
}

func (o *Occurrence) Artifact(ctx context.Context) (*Artifact, error) {
	result, err := o.Edges.ArtifactOrErr()
	if IsNotLoaded(err) {
		result, err = o.QueryArtifact().Only(ctx)
	}
	return result, err
}

func (o *Occurrence) Package(ctx context.Context) (*PackageVersion, error) {
	result, err := o.Edges.PackageOrErr()
	if IsNotLoaded(err) {
		result, err = o.QueryPackage().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (o *Occurrence) Source(ctx context.Context) (*SourceName, error) {
	result, err := o.Edges.SourceOrErr()
	if IsNotLoaded(err) {
		result, err = o.QuerySource().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (pn *PackageName) Namespace(ctx context.Context) (*PackageNamespace, error) {
	result, err := pn.Edges.NamespaceOrErr()
	if IsNotLoaded(err) {
		result, err = pn.QueryNamespace().Only(ctx)
	}
	return result, err
}

func (pn *PackageName) Versions(ctx context.Context) (result []*PackageVersion, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = pn.NamedVersions(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = pn.Edges.VersionsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = pn.QueryVersions().All(ctx)
	}
	return result, err
}

func (pn *PackageNamespace) Package(ctx context.Context) (*PackageType, error) {
	result, err := pn.Edges.PackageOrErr()
	if IsNotLoaded(err) {
		result, err = pn.QueryPackage().Only(ctx)
	}
	return result, err
}

func (pn *PackageNamespace) Names(ctx context.Context) (result []*PackageName, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = pn.NamedNames(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = pn.Edges.NamesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = pn.QueryNames().All(ctx)
	}
	return result, err
}

func (pt *PackageType) Namespaces(ctx context.Context) (result []*PackageNamespace, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = pt.NamedNamespaces(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = pt.Edges.NamespacesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = pt.QueryNamespaces().All(ctx)
	}
	return result, err
}

func (pv *PackageVersion) Name(ctx context.Context) (*PackageName, error) {
	result, err := pv.Edges.NameOrErr()
	if IsNotLoaded(err) {
		result, err = pv.QueryName().Only(ctx)
	}
	return result, err
}

func (pv *PackageVersion) Occurrences(ctx context.Context) (result []*Occurrence, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = pv.NamedOccurrences(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = pv.Edges.OccurrencesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = pv.QueryOccurrences().All(ctx)
	}
	return result, err
}

func (pv *PackageVersion) Sbom(ctx context.Context) (result []*BillOfMaterials, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = pv.NamedSbom(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = pv.Edges.SbomOrErr()
	}
	if IsNotLoaded(err) {
		result, err = pv.QuerySbom().All(ctx)
	}
	return result, err
}

func (pv *PackageVersion) EqualPackages(ctx context.Context) (result []*PkgEqual, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = pv.NamedEqualPackages(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = pv.Edges.EqualPackagesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = pv.QueryEqualPackages().All(ctx)
	}
	return result, err
}

func (pe *PkgEqual) Packages(ctx context.Context) (result []*PackageVersion, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = pe.NamedPackages(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = pe.Edges.PackagesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = pe.QueryPackages().All(ctx)
	}
	return result, err
}

func (sa *SLSAAttestation) BuiltFrom(ctx context.Context) (result []*Artifact, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = sa.NamedBuiltFrom(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = sa.Edges.BuiltFromOrErr()
	}
	if IsNotLoaded(err) {
		result, err = sa.QueryBuiltFrom().All(ctx)
	}
	return result, err
}

func (sa *SLSAAttestation) BuiltBy(ctx context.Context) (*Builder, error) {
	result, err := sa.Edges.BuiltByOrErr()
	if IsNotLoaded(err) {
		result, err = sa.QueryBuiltBy().Only(ctx)
	}
	return result, err
}

func (sa *SLSAAttestation) Subject(ctx context.Context) (*Artifact, error) {
	result, err := sa.Edges.SubjectOrErr()
	if IsNotLoaded(err) {
		result, err = sa.QuerySubject().Only(ctx)
	}
	return result, err
}

func (s *Scorecard) Certifications(ctx context.Context) (result []*CertifyScorecard, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = s.NamedCertifications(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = s.Edges.CertificationsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = s.QueryCertifications().All(ctx)
	}
	return result, err
}

func (sn *SourceName) Namespace(ctx context.Context) (*SourceNamespace, error) {
	result, err := sn.Edges.NamespaceOrErr()
	if IsNotLoaded(err) {
		result, err = sn.QueryNamespace().Only(ctx)
	}
	return result, err
}

func (sn *SourceName) Occurrences(ctx context.Context) (result []*Occurrence, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = sn.NamedOccurrences(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = sn.Edges.OccurrencesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = sn.QueryOccurrences().All(ctx)
	}
	return result, err
}

func (sn *SourceNamespace) SourceType(ctx context.Context) (*SourceType, error) {
	result, err := sn.Edges.SourceTypeOrErr()
	if IsNotLoaded(err) {
		result, err = sn.QuerySourceType().Only(ctx)
	}
	return result, err
}

func (sn *SourceNamespace) Names(ctx context.Context) (result []*SourceName, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = sn.NamedNames(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = sn.Edges.NamesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = sn.QueryNames().All(ctx)
	}
	return result, err
}

func (st *SourceType) Namespaces(ctx context.Context) (result []*SourceNamespace, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = st.NamedNamespaces(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = st.Edges.NamespacesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = st.QueryNamespaces().All(ctx)
	}
	return result, err
}

func (ve *VulnEqual) VulnerabilityIds(ctx context.Context) (result []*VulnerabilityID, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = ve.NamedVulnerabilityIds(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = ve.Edges.VulnerabilityIdsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = ve.QueryVulnerabilityIds().All(ctx)
	}
	return result, err
}

func (vi *VulnerabilityID) Type(ctx context.Context) (*VulnerabilityType, error) {
	result, err := vi.Edges.TypeOrErr()
	if IsNotLoaded(err) {
		result, err = vi.QueryType().Only(ctx)
	}
	return result, err
}

func (vi *VulnerabilityID) VulnEquals(ctx context.Context) (result []*VulnEqual, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = vi.NamedVulnEquals(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = vi.Edges.VulnEqualsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = vi.QueryVulnEquals().All(ctx)
	}
	return result, err
}

func (vt *VulnerabilityType) VulnerabilityIds(ctx context.Context) (result []*VulnerabilityID, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = vt.NamedVulnerabilityIds(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = vt.Edges.VulnerabilityIdsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = vt.QueryVulnerabilityIds().All(ctx)
	}
	return result, err
}
