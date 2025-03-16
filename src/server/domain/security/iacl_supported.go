package domain

// IAclSupported represents an entity which supports ACL
type IAclSupported interface {
	GetSubjectToAcl() bool
	SetSubjectToAcl(subjectToAcl bool)
}
