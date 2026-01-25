package dp

import "github.com/pocketbase/pocketbase/core"

const (
	AnnouncementsCollection = "announcements"
)

var _ core.RecordProxy = (*Announcement)(nil)

// Announcement RecordProxy for collection `announcements`.
// Provides type-safe struct access to announcement records.
type Announcement struct {
	core.BaseRecordProxy
}

func (a *Announcement) Title() string {
	return a.GetString("title")
}

func (a *Announcement) SetTitle(title string) {
	a.Set("title", title)
}

func (a *Announcement) BodyText() string {
	return a.GetString("bodytext")
}

func (a *Announcement) SetBodyText(bodyText string) {
	a.Set("bodytext", bodyText)
}

func (a *Announcement) Link() string {
	return a.GetString("link")
}

func (a *Announcement) SetLink(link string) {
	a.Set("link", link)
}

func (a *Announcement) LinkText() string {
	return a.GetString("link_text")
}

func (a *Announcement) SetLinkText(linkText string) {
	a.Set("link_text", linkText)
}

func (a *Announcement) Priority() string {
	return a.GetString("priority")
}

func (a *Announcement) SetPriority(priority string) {
	a.Set("priority", priority)
}

func (a *Announcement) Club() string {
	return a.GetString("club")
}

func (a *Announcement) SetClub(clubID string) {
	a.Set("club", clubID)
}

func (a *Announcement) Team() string {
	return a.GetString("team")
}

func (a *Announcement) SetTeam(teamID string) {
	a.Set("team", teamID)
}

func (a *Announcement) Author() string {
	return a.GetString("author")
}

func (a *Announcement) SetAuthor(authorID string) {
	a.Set("author", authorID)
}
