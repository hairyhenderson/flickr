package people

import (
	"context"
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/hairyhenderson/flickr"
)

type PhotoList struct {
	Page    int `xml:"page,attr"`
	Pages   int `xml:"pages,attr"`
	PerPage int `xml:"perpage,attr"`
	Total   int `xml:"total,attr"`
	Photo   struct {
		Id       string `xml:"id,attr"`
		Owner    string `xml:"owner,attr"`
		Secret   string `xml:"secret,attr"`
		Server   string `xml:"server,attr"`
		Farm     string `xml:"farm,attr"`
		Title    string `xml:"title,attr"`
		IsPublic bool   `xml:"ispublic,attr"`
		IsFriend bool   `xml:"isfriend,attr"`
		IsFamily bool   `xml:"isfamily,attr"`

		// if extras contains "url_o" these are populated
		UrlO    string `xml:"url_o,attr"`
		HeightO int    `xml:"height_o,attr"`
		WidthO  int    `xml:"width_o,attr"`

		Description    string `xml:"description,attr"`
		License        string `xml:"license,attr"`
		DateUpload     string `xml:"date_upload,attr"`
		DateTaken      string `xml:"date_taken,attr"`
		OwnerName      string `xml:"owner_name,attr"`
		IconServer     string `xml:"icon_server,attr"`
		OriginalFormat string `xml:"original_format,attr"`
		LastUpdate     string `xml:"last_udpate,attr"`

		// Geo - these attributes are provided when extras contains "geo"
		Latitude  string `xml:"latitude,attr"`
		Longitude string `xml:"longitude,attr"`
		Accuracy  string `xml:"accuracy,attr"`
		Context   string `xml:"context,attr"`

		// Tags - contains space-separated lists
		Tags        string `xml:"tags,attr"`
		MachineTags string `xml:"machine_tags,attr"`

		// Original Dimensions - these attributes are provided
		// when extras contains "o_dims"
		OWidth  int `xml:"o_width,attr"`
		OHeight int `xml:"o_height,attr"`

		Views     int    `xml:"views,attr"`
		Media     string `xml:"media,attr"`
		PathAlias string `xml:"path_alias,attr"`

		// Square Urls - these attributes are provided when
		// extras contains "url_sq"
		UrlSq    string `xml:"url_sq,attr"`
		HeightSq int    `xml:"height_sq,attr"`
		WidthSq  int    `xml:"width_sq,attr"`

		// Thumbnail Urls - these attributes are provided
		// when extras contains "url_t"
		UrlT    string `xml:"url_t,attr"`
		HeightT int    `xml:"height_t,attr"`
		WidthT  int    `xml:"width_t,attr"`

		// Q Urls - these attributes are provided when
		// extras contains "url_s"
		UrlS    string `xml:"url_s,attr"`
		HeightS int    `xml:"height_s,attr"`
		WidthS  int    `xml:"width_s,attr"`

		// M Urls - these attributes are provided when
		// extras contains "url_m"
		UrlM    string `xml:"url_m,attr"`
		HeightM int    `xml:"height_m,attr"`
		WidthM  int    `xml:"width_m,attr"`

		// N Urls - these attributes are provided when
		// extras contains "url_n"
		UrlN    string `xml:"url_n,attr"`
		HeightN int    `xml:"height_n,attr"`
		WidthN  int    `xml:"width_n,attr"`

		// Z Urls - these attributes are provided when
		// extras contains "url_z"
		UrlZ    string `xml:"url_z,attr"`
		HeightZ int    `xml:"height_z,attr"`
		WidthZ  int    `xml:"width_z,attr"`

		// C Urls - these attributes are provided when
		// extras contains "url_c"
		UrlC    string `xml:"url_c,attr"`
		HeightC int    `xml:"height_c,attr"`
		WidthC  int    `xml:"width_c,attr"`

		// L Urls - these attributes are provided when
		// extras contains "url_l"
		UrlL    string `xml:"url_l,attr"`
		HeightL int    `xml:"height_l,attr"`
		WidthL  int    `xml:"width_l,attr"`
	}
}

type PhotoListResponse struct {
	flickr.BasicResponse
	Photos PhotoList `xml:"photos"`
}

type SafetyLevel int

const (
	NoSafetySpecified SafetyLevel = iota
	Safe
	Moderate
	Restricted
)

type ContentType int

const (
	NoContentTypeSpecified ContentType = iota
	PhotosOnly
	ScreenShotsOnly
	OtherOnly
	PhotosAndScreenshots
	ScreenShotsAndOther
	PhotosAndOther
	All
)

type PrivacyFilterType int

const (
	NoPrivacyFilterSpecified PrivacyFilterType = iota
	Public
	Friends
	Family
	FriendsAndFamily
	Private
)

type GetPhotosOptionalArgs struct {
	SafeSearch    SafetyLevel       // optional, set to NoneSpecified to ignore
	MinUploadDate string            // optional, set to "" to ignore. mysql datetime
	MaxUploadDate string            // optional, set to "" to ignore. mysql datetime
	MinTakenDate  string            // optional, set to "" to ignore. mysql datetime
	MaxTakenDate  string            // optional, set to "" to ignore. mysql datetime
	ContentType   ContentType       // optional, set to NoneSpecified to ignore
	PrivacyFilter PrivacyFilterType // optional, set to NoneSpecified to ignore
	Extras        string            // optional, set to "" to ignore. comma separated string.
	PerPage       int               // 0 to ignore
	Page          int               // 0 to ignore
}

type Person struct {
	ID                           string `xml:"id,attr"`
	NSID                         string `xml:"nsid,attr"`
	IsPro                        bool   `xml:"ispro,attr"`
	IsDeleted                    bool   `xml:"is_deleted,attr"`
	IconServer                   int    `xml:"iconserver,attr"`
	IconFarm                     int    `xml:"iconfarm,attr"`
	PathAlias                    string `xml:"path_alias,attr"`
	HasStats                     bool   `xml:"has_stats,attr"`
	ProBadge                     string `xml:"pro_badge,attr"`
	Expire                       int    `xml:"expire,attr"`
	UploadCount                  int    `xml:"upload_count,attr"`
	UploadLimit                  int    `xml:"upload_limit,attr"`
	UploadLimitStatus            string `xml:"upload_limit_status,attr"`
	IsCognitoUser                bool   `xml:"is_cognito_user,attr"`
	AllRightsReservedPhotosCount int    `xml:"all_rights_reserved_photos_count,attr"`
	HasAdfree                    bool   `xml:"has_adfree,attr"`
	HasFreeStandardShipping      bool   `xml:"has_free_standard_shipping,attr"`
	HasFreeEducationalResources  bool   `xml:"has_free_educational_resources,attr"`

	Username    string      `xml:"username"`
	Realname    string      `xml:"realname"`
	MboxSha1Sum string      `xml:"mbox_sha1sum"`
	Location    string      `xml:"location"`
	Timezone    xmlTimeZone `xml:"timezone"`
	PhotosUrl   string      `xml:"photosurl"`
	ProfileUrl  string      `xml:"profileurl"`
	Photos      struct {
		FirstDate      int     `xml:"firstdate"`
		FirstDateTaken xmlTime `xml:"firstdatetaken"`
		Count          int     `xml:"count"`
	} `xml:"photos"`
}

type xmlTime struct {
	time.Time
}

func (t *xmlTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}
	parse, err := time.Parse("2006-01-02 15:04:05", v)
	if err != nil {
		return err
	}
	*t = xmlTime{parse}
	return nil
}

// xmlTimeZone allows unmarshaling a time.Location from the timezone_id attribute
type xmlTimeZone struct{ time.Location }

func (tz *xmlTimeZone) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	v := struct {
		TimezoneID string `xml:"timezone_id,attr"`
	}{}
	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}

	parsed, err := time.LoadLocation(v.TimezoneID)
	if err != nil {
		return err
	}
	*tz = xmlTimeZone{*parsed}
	return nil
}

type PersonResponse struct {
	flickr.BasicResponse
	Person Person `xml:"person"`
}

type PeopleClient struct {
	hc *http.Client
	fc *flickr.FlickrRequestClient
}

func NewPeopleClient(hc *http.Client, fc *flickr.FlickrRequestClient) *PeopleClient {
	return &PeopleClient{hc: hc, fc: fc}
}

func (pc *PeopleClient) GetInfo(ctx context.Context, userId string) (*Person, error) {
	v := url.Values{}
	v.Set("user_id", userId)

	req, err := pc.fc.NewRequestWithContext(ctx, http.MethodGet, "flickr.people.getInfo", v, nil)
	if err != nil {
		return nil, err
	}

	res, err := pc.hc.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http %s: %v", req.Method, err)
	}
	defer res.Body.Close()

	for k := range res.Header {
		fmt.Printf("%s: %s\n", k, res.Header.Get(k))
	}

	response := PersonResponse{}
	err = flickr.ParseApiResponse(res, &response)
	if err != nil {
		return nil, fmt.Errorf("parse api response: %w", err)
	}

	return &response.Person, nil
}

func (pc *PeopleClient) GetPhotos(ctx context.Context,
	userId string, opts GetPhotosOptionalArgs) (*PhotoList, error) {

	v := url.Values{}

	v.Set("user_id", userId)
	if opts.SafeSearch != NoSafetySpecified {
		v.Set("safe_search", strconv.Itoa(int(opts.SafeSearch)))
	}
	if opts.MinUploadDate != "" {
		v.Set("min_upload_date", opts.MinUploadDate)
	}
	if opts.MaxUploadDate != "" {
		v.Set("min_upload_date", opts.MaxUploadDate)
	}
	if opts.MinTakenDate != "" {
		v.Set("min_taken_date", opts.MinTakenDate)
	}
	if opts.MaxTakenDate != "" {
		v.Set("max_taken_date", opts.MaxTakenDate)
	}
	if opts.ContentType != NoContentTypeSpecified {
		v.Set("content_type", strconv.Itoa(int(opts.ContentType)))
	}
	if opts.PrivacyFilter != NoPrivacyFilterSpecified {
		v.Set("privacy_filter", strconv.Itoa(int(opts.PrivacyFilter)))
	}
	if opts.PerPage != 0 {
		v.Set("per_page", strconv.Itoa(opts.PerPage))
	}
	if opts.Page != 0 {
		v.Set("page", strconv.Itoa(opts.Page))
	}
	if opts.Extras != "" {
		v.Set("extras", opts.Extras)
	}

	req, err := pc.fc.NewRequestWithContext(ctx, http.MethodGet, "flickr.people.getPhotos", v, nil)
	if err != nil {
		return nil, err
	}

	res, err := pc.hc.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http %s: %v", req.Method, err)
	}
	defer res.Body.Close()

	response := PhotoListResponse{}
	err = flickr.ParseApiResponse(res, &response)
	if err != nil {
		return nil, fmt.Errorf("parse api response: %w", err)
	}

	return &response.Photos, err
}
