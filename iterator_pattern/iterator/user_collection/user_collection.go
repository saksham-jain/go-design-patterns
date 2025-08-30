package usercollection

type UserCollection struct {
	Collection []string
}

func NewUserCollection(collection []string) UserCollection {
	return UserCollection{
		Collection: collection,
	}
}
