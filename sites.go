package panda

import "encoding/json"

type FavoriteSiteList struct {
	FavoriteSitesIDs     []string `json:"favoriteSiteIds"`
	AutoFavoritesEnabled bool     `json:"autoFavoritesEnabled"`
}

// GetFavoriteSites get sites with star
func (p Handler) GetFavoriteSites() (list FavoriteSiteList) {
	p.sustainAuth()

	res, err := p.get(BaseURI+"/portal/favorites/list", nil)
	if err != nil {
		return
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&list)
	if err != nil {
		return
	}

	return
}

type Site struct {
	Id string `json:"id"`
}

type DirectSiteJsonResult struct {
	EntityPrefix   string `json:"entityPrefix"`
	SiteCollection []Site `json:"site_collection"`
}

// GetAllSites get all the sites which the authenticated user can see.
func (p Handler) GetAllSites() (list DirectSiteJsonResult) {
	p.sustainAuth()

	res, err := p.get(BaseURI+"/direct/site.json", nil)
	if err != nil {
		return
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&list)
	if err != nil {
		return
	}

	return
}
