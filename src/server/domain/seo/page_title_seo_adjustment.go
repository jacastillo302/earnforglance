package domain

// PageTitleSeoAdjustment represents a page title SEO adjustment
type PageTitleSeoAdjustment int

const (
	// PagenameAfterStorename represents pagename comes after storename
	PagenameAfterStorename PageTitleSeoAdjustment = 0

	// StorenameAfterPagename represents storename comes after pagename
	StorenameAfterPagename PageTitleSeoAdjustment = 10
)
