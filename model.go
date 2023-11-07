package main

type OriginalImage struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type TabletImage struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type MobileImage struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type SkeletonImage struct {
	Images    map[string]OriginalImage `json:"images"`
	MediaType string                   `json:"mediaType"`
}

type ListingModel struct {
	PromoType          string          `json:"promoType"`
	URL                string          `json:"url"`
	Images             []string        `json:"images"`
	SkeletonImages     []SkeletonImage `json:"skeletonImages"`
	BrandingAppearance string          `json:"brandingAppearance"`
	Price              string          `json:"price"`
	HasVideo           bool            `json:"hasVideo"`
	Address            Address         `json:"address"`
}

type Agent struct {
	AgentName  string `json:"agentName"`
	AgentPhoto string `json:"agentPhoto"`
}

type Branding struct {
	AgencyId          int     `json:"agencyId"`
	Agents            []Agent `json:"agents"`
	AgentNames        string  `json:"agentNames"`
	BrandLogo         string  `json:"brandLogo"`
	SkeletonBrandLogo string  `json:"skeletonBrandLogo"`
	BrandName         string  `json:"brandName"`
	BrandColor        string  `json:"brandColor"`
	AgentPhoto        string  `json:"agentPhoto"`
	AgentName         string  `json:"agentName"`
}

type Address struct {
	Street   string  `json:"street"`
	Suburb   string  `json:"suburb"`
	State    string  `json:"state"`
	Postcode string  `json:"postcode"`
	Lat      float64 `json:"lat"`
	Lng      float64 `json:"lng"`
}

type Features struct {
	Beds                  int    `json:"beds"`
	Baths                 int    `json:"baths"`
	Parking               int    `json:"parking"`
	PropertyType          string `json:"propertyType"`
	PropertyTypeFormatted string `json:"propertyTypeFormatted"`
	IsRural               bool   `json:"isRural"`
	LandSize              int    `json:"landSize"`
	LandUnit              string `json:"landUnit"`
	IsRetirement          bool   `json:"isRetirement"`
}

type Inspection struct {
	OpenTime  interface{} `json:"openTime"`
	CloseTime interface{} `json:"closeTime"`
}

type Tags struct {
	TagText      string `json:"tagText"`
	TagClassName string `json:"tagClassName"`
}

type ListingType struct {
	ID                      int          `json:"id"`
	ListingType             string       `json:"listingType"`
	ListingModel            ListingModel `json:"listingModel"`
	Branding                Branding     `json:"branding"`
	Features                Features     `json:"features"`
	Inspection              Inspection   `json:"inspection"`
	Auction                 interface{}  `json:"auction"`
	Tags                    Tags         `json:"tags"`
	DisplaySearchPriceRange interface{}  `json:"displaySearchPriceRange"`
	EnableSingleLineAddress bool         `json:"enableSingleLineAddress"`
}

type ResponseType struct {
	Props struct {
		ListingsMap map[string]ListingType `json:"listingsMap"`
	} `json:"props"`
}
