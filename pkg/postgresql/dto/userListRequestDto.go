package dto

type UserFilter struct {
	Name   *string
	Email  *string
	Status *string
}

// convert the dto to struct in order to loop
func ConvertToFilterMap(filter UserFilter) map[string]struct {
	Value   *string
	UseLike bool
} {
	return map[string]struct {
		Value   *string
		UseLike bool
	}{
		"name":   {filter.Name, true},
		"email":  {filter.Email, true},
		"status": {filter.Status, false}, // If its integer we use = instead of LIKE
	}
}
