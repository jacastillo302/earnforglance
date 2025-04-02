package common

const (
	// KeepAlivePath is the request path to the keep-alive URL
	KeepAlivePath = "keepalive/index"

	// AddressAttributeControlName is the name of the custom address attribute control
	// {0} : address attribute id
	AddressAttributeControlName = "address_attribute_%d"

	// RestartTimeout is the default timeout (in milliseconds) before restarting the application
	RestartTimeout = 3000

	// DbBackupsPath is the path to the database backup files
	DbBackupsPath = "db_backups\\"

	// DbBackupFileExtension is the database backup file extension
	DbBackupFileExtension = "bak"

	// HeadCodeFileName is the name of the file with code for the head element
	HeadCodeFileName = "html_code.html"

	// SingleFaviconHeadLink is the head link for the favicon
	// {0} : icon set identifier
	// {1} : icon file name
	SingleFaviconHeadLink = "<link rel=\"shortcut icon\" href=\"/icons/icons_%d/%s\">"

	// FaviconAndAppIconsPath is the path to the favicon and app icons
	// {0} : icon set identifier
	FaviconAndAppIconsPath = "icons/icons_%d"

	// OldFaviconIconName is the name of the old favicon icon for the current store
	// {0} : store identifier
	OldFaviconIconName = "favicon-%d.ico"

	// LocalePatternPath is the path to the localization client-side validation
	// {0} : locale identifier
	LocalePatternPath = "lib_npm/cldr-data/main/%s"

	// LocalePatternArchiveName is the name of the archive with localization of templates
	LocalePatternArchiveName = "main.zip"

	// DefaultLocalePattern is the name of the default pattern locale
	DefaultLocalePattern = "en"

	// DefaultLanguageCulture is the default CultureInfo
	DefaultLanguageCulture = "en-US"

	DefaultLanguageName = "English"

	// LanguagePackMinTranslationProgressToInstall is the minimal progress of language pack translation to download and install
	LanguagePackMinTranslationProgressToInstall = 80

	// LanguagePackProgressAttribute is the name of the generic attribute to store the value of 'LanguagePackProgress'
	LanguagePackProgressAttribute = "LanguagePackProgress"

	// LicenseCheckPath is the path to request the  official site for license compliance check
	// {0} : store URL
	// {1} :  version
	// {2} : admin email
	// {3} : language code
	LicenseCheckPath = "license-check?url=%s&version=%s&email=%s&language=%s"

	// NewsRssPath is the path to request the  official site for news RSS
	// {0} :  version
	// {1} : whether the store is based on localhost
	// {2} : whether advertisements are hidden
	// {3} : store URL
	// {4} : language code
	NewsRssPath = "news-rss?version=%s&localhost=%t&hideAdvertisements=%t&storeUrl=%s&language=%s"

	// InstallationCompletedPath is the path to notify the  official site about successful installation
	// {0} :  version
	// {1} : whether the store is based on localhost
	// {2} : admin email
	// {3} : store URL
	// {4} : language code
	// {5} : culture name
	InstallationCompletedPath = "installation-completed?version=%s&local=%t&email=%s&url=%s&language=%s&culture=%s"

	// SubscribeNewslettersPath is the path to subscribe to the  newsletters
	// {0} : subscriber email
	SubscribeNewslettersPath = "subscribe-newsletters?&email=%s"

	// ExtensionsCategoriesPath is the path to request available categories of marketplace extensions
	// {0} : language code
	ExtensionsCategoriesPath = "extensions-feed?getCategories=1&language=%s"

	// ExtensionsVersionsPath is the path to request available versions of marketplace extensions
	// {0} : language code
	ExtensionsVersionsPath = "extensions-feed?getVersions=1&language=%s"

	// ExtensionsPath is the path to request marketplace extensions
	// {0} : extension category identifier
	// {1} : extension version identifier
	// {2} : extension price identifier
	// {3} : search term
	// {4} : page index
	// {5} : page size
	// {6} : language code
	ExtensionsPath = "extensions-feed?category=%d&version=%d&price=%d&searchTerm=%s&pageIndex=%d&pageSize=%d&language=%s"

	// PdfLtrFontName is the name of the left-to-right PDF font
	PdfLtrFontName = "OpenSans"

	// PdfRtlFontName is the name of the right-to-left PDF font
	PdfRtlFontName = "Vazirmatn"
)
