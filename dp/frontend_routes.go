package dp

/// This is a temporary solution until I find a better way for frontend routes in backend context.

func RouteEventDetail(id string) string {
    return "/account/event/" + id
}

func RouteAnnouncementDetail(id string) string {
    return "/account/announcements/" + id
}