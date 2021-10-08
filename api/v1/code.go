package api

type Repo struct {
	Environments map[string]FileFilter
	FileTypes    map[string]string
	Scopes       map[string]string
	Criticality  []CriticalityAnalyzer
}

type CriticalityAnalyzer struct {
	Criticality int
	Branch      string
	Paths       []string
	Comments    []string
	ExternalCmd []string
}
type FileFilter struct {
	Branch string
	Paths  []string
}

type CodeChanges struct {
	Identifier
	LOC, Files, Packages int
	DependencyChanges    []DependencyChange
	Changes              []CodeChange
}

type FileType string

const (
	FileTypeCode          FileType = "code"
	FileTypePackaging     FileType = "packaging"
	FileTypeDocs          FileType = "docs"
	FileTypeBinary        FileType = "binary"
	FileTypeGenerated     FileType = "generated"
	FileTypeDesign        FileType = "design"
	FileTypeUI            FileType = "ui"
	FileTypeIaC           FileType = "iac"
	FileTypeConfiguration FileType = "config"
)

type CodeChange struct {
	File     string
	LOC      int
	FileType string
}

type DependencyChange struct {
	Major, Minor, Patch int
}

type CodePromotion struct {
	Identifier
}

type CodeRevert struct {
	Identifier
}

type ConfigChange struct {
	Identifier
}

type ConfigPromotion struct {
	Identifier
}

type ConfigRevert struct {
	Identifier
}
