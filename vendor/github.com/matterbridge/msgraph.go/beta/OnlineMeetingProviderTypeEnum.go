// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// OnlineMeetingProviderType undocumented
type OnlineMeetingProviderType int

const (
	// OnlineMeetingProviderTypeVUnknown undocumented
	OnlineMeetingProviderTypeVUnknown OnlineMeetingProviderType = 0
	// OnlineMeetingProviderTypeVSkypeForBusiness undocumented
	OnlineMeetingProviderTypeVSkypeForBusiness OnlineMeetingProviderType = 1
	// OnlineMeetingProviderTypeVSkypeForConsumer undocumented
	OnlineMeetingProviderTypeVSkypeForConsumer OnlineMeetingProviderType = 2
	// OnlineMeetingProviderTypeVTeamsForBusiness undocumented
	OnlineMeetingProviderTypeVTeamsForBusiness OnlineMeetingProviderType = 3
)

// OnlineMeetingProviderTypePUnknown returns a pointer to OnlineMeetingProviderTypeVUnknown
func OnlineMeetingProviderTypePUnknown() *OnlineMeetingProviderType {
	v := OnlineMeetingProviderTypeVUnknown
	return &v
}

// OnlineMeetingProviderTypePSkypeForBusiness returns a pointer to OnlineMeetingProviderTypeVSkypeForBusiness
func OnlineMeetingProviderTypePSkypeForBusiness() *OnlineMeetingProviderType {
	v := OnlineMeetingProviderTypeVSkypeForBusiness
	return &v
}

// OnlineMeetingProviderTypePSkypeForConsumer returns a pointer to OnlineMeetingProviderTypeVSkypeForConsumer
func OnlineMeetingProviderTypePSkypeForConsumer() *OnlineMeetingProviderType {
	v := OnlineMeetingProviderTypeVSkypeForConsumer
	return &v
}

// OnlineMeetingProviderTypePTeamsForBusiness returns a pointer to OnlineMeetingProviderTypeVTeamsForBusiness
func OnlineMeetingProviderTypePTeamsForBusiness() *OnlineMeetingProviderType {
	v := OnlineMeetingProviderTypeVTeamsForBusiness
	return &v
}