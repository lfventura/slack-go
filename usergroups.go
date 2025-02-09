package slack

import (
	"context"
	"net/url"
	"strings"
)

// UserGroup contains all the information of a user group
type UserGroup struct {
	ID          string         `json:"id"`
	TeamID      string         `json:"team_id"`
	IsUserGroup bool           `json:"is_usergroup"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Handle      string         `json:"handle"`
	IsExternal  bool           `json:"is_external"`
	DateCreate  JSONTime       `json:"date_create"`
	DateUpdate  JSONTime       `json:"date_update"`
	DateDelete  JSONTime       `json:"date_delete"`
	AutoType    string         `json:"auto_type"`
	CreatedBy   string         `json:"created_by"`
	UpdatedBy   string         `json:"updated_by"`
	DeletedBy   string         `json:"deleted_by"`
	Prefs       UserGroupPrefs `json:"prefs"`
	UserCount   int            `json:"user_count"`
	Users       []string       `json:"users"`
}

// UserGroupPrefs contains default channels and groups (private channels)
type UserGroupPrefs struct {
	Channels []string `json:"channels"`
	Groups   []string `json:"groups"`
}

type userGroupResponseFull struct {
	UserGroups []UserGroup `json:"usergroups"`
	UserGroup  UserGroup   `json:"usergroup"`
	Users      []string    `json:"users"`
	SlackResponse
}

func (api *Client) userGroupRequest(ctx context.Context, path string, values url.Values) (*userGroupResponseFull, error) {
	response := &userGroupResponseFull{}
	err := api.postMethod(ctx, path, values, response)
	if err != nil {
		return nil, err
	}

	return response, response.Err()
}

// CreateUserGroup creates a new user group
func (api *Client) CreateUserGroup(userGroup UserGroup) (UserGroup, error) {
	return api.CreateUserGroupContext(context.Background(), userGroup)
}

// CreateUserGroupContext creates a new user group with a custom context
func (api *Client) CreateUserGroupContext(ctx context.Context, userGroup UserGroup) (UserGroup, error) {
	values := url.Values{
		"token": {api.token},
		"name":  {userGroup.Name},
	}

	if userGroup.Handle != "" {
		values["handle"] = []string{userGroup.Handle}
	}

	if userGroup.Description != "" {
		values["description"] = []string{userGroup.Description}
	}

	if len(userGroup.Prefs.Channels) > 0 {
		values["channels"] = []string{strings.Join(userGroup.Prefs.Channels, ",")}
	}

	if userGroup.TeamID != "" {
		values["team_id"] = []string{userGroup.TeamID}
	}

	response, err := api.userGroupRequest(ctx, "usergroups.create", values)
	if err != nil {
		return UserGroup{}, err
	}
	return response.UserGroup, nil
}

// DisableUserGroupOption options for the DisableUserGroup and EnableUserGroup method calls.
type DisableUserGroupOption func(*DisableUserGroupParams)

// DisableUserGroupParams contains arguments for DisableUserGroup and EnableUserGroup method calls.
type DisableUserGroupParams struct {
	IncludeCount bool
	TeamID       string
}

// DisableUserGroupOptionIncludeCount include the count of User Groups (default: false)
func DisableUserGroupOptionIncludeCount(b bool) DisableUserGroupOption {
	return func(params *DisableUserGroupParams) {
		params.IncludeCount = b
	}
}

// DisableUserGroupOptionTeamID include team Id
func DisableUserGroupOptionTeamID(teamID string) DisableUserGroupOption {
	return func(params *DisableUserGroupParams) {
		params.TeamID = teamID
	}
}

// DisableUserGroup disables an existing user group
func (api *Client) DisableUserGroup(userGroup string, options ...DisableUserGroupOption) (UserGroup, error) {
	return api.DisableUserGroupContext(context.Background(), userGroup, options...)
}

// DisableUserGroupContext disables an existing user group with a custom context
func (api *Client) DisableUserGroupContext(ctx context.Context, userGroup string, options ...DisableUserGroupOption) (UserGroup, error) {
	params := DisableUserGroupParams{}

	for _, opt := range options {
		opt(&params)
	}

	values := url.Values{
		"token":     {api.token},
		"usergroup": {userGroup},
	}

	if params.IncludeCount {
		values.Add("include_count", "true")
	}

	if params.TeamID != "" {
		values.Add("team_id", params.TeamID)
	}

	response, err := api.userGroupRequest(ctx, "usergroups.disable", values)
	if err != nil {
		return UserGroup{}, err
	}
	return response.UserGroup, nil
}

// EnableUserGroup enables an existing user group
func (api *Client) EnableUserGroup(userGroup string, options ...DisableUserGroupOption) (UserGroup, error) {
	return api.EnableUserGroupContext(context.Background(), userGroup, options...)
}

// EnableUserGroupContext enables an existing user group with a custom context
func (api *Client) EnableUserGroupContext(ctx context.Context, userGroup string, options ...DisableUserGroupOption) (UserGroup, error) {
	params := DisableUserGroupParams{}

	for _, opt := range options {
		opt(&params)
	}

	values := url.Values{
		"token":     {api.token},
		"usergroup": {userGroup},
	}

	if params.IncludeCount {
		values.Add("include_count", "true")
	}

	if params.TeamID != "" {
		values.Add("team_id", params.TeamID)
	}

	response, err := api.userGroupRequest(ctx, "usergroups.enable", values)
	if err != nil {
		return UserGroup{}, err
	}
	return response.UserGroup, nil
}

// GetUserGroupsOption options for the GetUserGroups method call.
type GetUserGroupsOption func(*GetUserGroupsParams)

// GetUserGroupsOptionIncludeCount include the number of users in each User Group (default: false)
func GetUserGroupsOptionIncludeCount(b bool) GetUserGroupsOption {
	return func(params *GetUserGroupsParams) {
		params.IncludeCount = b
	}
}

// GetUserGroupsOptionIncludeDisabled include disabled User Groups (default: false)
func GetUserGroupsOptionIncludeDisabled(b bool) GetUserGroupsOption {
	return func(params *GetUserGroupsParams) {
		params.IncludeDisabled = b
	}
}

// GetUserGroupsOptionIncludeUsers include the list of users for each User Group (default: false)
func GetUserGroupsOptionIncludeUsers(b bool) GetUserGroupsOption {
	return func(params *GetUserGroupsParams) {
		params.IncludeUsers = b
	}
}

// GetUsersOptionTeamID include team ID
func GetUserGroupsOptionTeamID(teamID string) GetUserGroupsOption {
	return func(params *GetUserGroupsParams) {
		params.TeamID = teamID
	}
}

// GetUserGroupsParams contains arguments for GetUserGroups method call
type GetUserGroupsParams struct {
	IncludeCount    bool
	IncludeDisabled bool
	IncludeUsers    bool
	TeamID          string
}

// GetUserGroups returns a list of user groups for the team
func (api *Client) GetUserGroups(options ...GetUserGroupsOption) ([]UserGroup, error) {
	return api.GetUserGroupsContext(context.Background(), options...)
}

// GetUserGroupsContext returns a list of user groups for the team with a custom context
func (api *Client) GetUserGroupsContext(ctx context.Context, options ...GetUserGroupsOption) ([]UserGroup, error) {
	params := GetUserGroupsParams{}

	for _, opt := range options {
		opt(&params)
	}

	values := url.Values{
		"token": {api.token},
	}
	if params.IncludeCount {
		values.Add("include_count", "true")
	}
	if params.IncludeDisabled {
		values.Add("include_disabled", "true")
	}
	if params.IncludeUsers {
		values.Add("include_users", "true")
	}
	if params.TeamID != "" {
		values.Add("team_id", params.TeamID)
	}

	response, err := api.userGroupRequest(ctx, "usergroups.list", values)
	if err != nil {
		return nil, err
	}
	return response.UserGroups, nil
}

// UpdateUserGroupsOption options for the UpdateUserGroup method call.
type UpdateUserGroupsOption func(*UpdateUserGroupsParams)

// UpdateUserGroupsOptionName change the name of the User Group (default: empty, so it's no-op)
func UpdateUserGroupsOptionName(name string) UpdateUserGroupsOption {
	return func(params *UpdateUserGroupsParams) {
		params.Name = name
	}
}

// UpdateUserGroupsOptionHandle change the handle of the User Group (default: empty, so it's no-op)
func UpdateUserGroupsOptionHandle(handle string) UpdateUserGroupsOption {
	return func(params *UpdateUserGroupsParams) {
		params.Handle = handle
	}
}

// UpdateUserGroupsOptionDescription change the description of the User Group. (default: nil, so it's no-op)
func UpdateUserGroupsOptionDescription(description *string) UpdateUserGroupsOption {
	return func(params *UpdateUserGroupsParams) {
		params.Description = description
	}
}

// UpdateUserGroupsOptionChannels change the default channels of the User Group. (default: unspecified, so it's no-op)
func UpdateUserGroupsOptionChannels(channels []string) UpdateUserGroupsOption {
	return func(params *UpdateUserGroupsParams) {
		params.Channels = &channels
	}
}

// UpdateUserGroupsOptionTeamID specify the team id for the User Group. (default: nil, so it's no-op)
func UpdateUserGroupsOptionTeamID(teamID *string) UpdateUserGroupsOption {
	return func(params *UpdateUserGroupsParams) {
		params.TeamID = teamID
	}
}

// UpdateUserGroupsParams contains arguments for UpdateUserGroup method call
type UpdateUserGroupsParams struct {
	Name        string
	Handle      string
	Description *string
	Channels    *[]string
	TeamID      *string
}

// UpdateUserGroup will update an existing user group
func (api *Client) UpdateUserGroup(userGroupID string, options ...UpdateUserGroupsOption) (UserGroup, error) {
	return api.UpdateUserGroupContext(context.Background(), userGroupID, options...)
}

// UpdateUserGroupContext will update an existing user group with a custom context
func (api *Client) UpdateUserGroupContext(ctx context.Context, userGroupID string, options ...UpdateUserGroupsOption) (UserGroup, error) {
	params := UpdateUserGroupsParams{}

	for _, opt := range options {
		opt(&params)
	}

	values := url.Values{
		"token":     {api.token},
		"usergroup": {userGroupID},
	}

	if params.Name != "" {
		values["name"] = []string{params.Name}
	}

	if params.Handle != "" {
		values["handle"] = []string{params.Handle}
	}

	if params.Description != nil {
		values["description"] = []string{*params.Description}
	}

	if params.Channels != nil {
		values["channels"] = []string{strings.Join(*params.Channels, ",")}
	}

	if params.TeamID != nil {
		values["team_id"] = []string{*params.TeamID}
	}

	response, err := api.userGroupRequest(ctx, "usergroups.update", values)
	if err != nil {
		return UserGroup{}, err
	}
	return response.UserGroup, nil
}

// GetUserGroupMembersOption options for the GetUserGroupMembers method call.
type GetUserGroupMembersOption func(*GetUserGroupMembersParams)

// GetUserGroupMembersParams contains arguments for GetUserGroupMembers method call
type GetUserGroupMembersParams struct {
	IncludeDisabled bool
	TeamID          string
}

// GetUserGroupMembersOptionIncludeDisabled include disabled User Groups (default: false)
func GetUserGroupMembersOptionIncludeDisabled(b bool) GetUserGroupMembersOption {
	return func(params *GetUserGroupMembersParams) {
		params.IncludeDisabled = b
	}
}

// GetUserGroupMembersOptionTeamID include team Id
func GetUserGroupMembersOptionTeamID(teamID string) GetUserGroupMembersOption {
	return func(params *GetUserGroupMembersParams) {
		params.TeamID = teamID
	}
}

// GetUserGroupMembers will retrieve the current list of users in a group
func (api *Client) GetUserGroupMembers(userGroup string, options ...GetUserGroupMembersOption) ([]string, error) {
	return api.GetUserGroupMembersContext(context.Background(), userGroup, options...)
}

// GetUserGroupMembersContext will retrieve the current list of users in a group with a custom context
func (api *Client) GetUserGroupMembersContext(ctx context.Context, userGroup string, options ...GetUserGroupMembersOption) ([]string, error) {
	params := GetUserGroupMembersParams{}

	for _, opt := range options {
		opt(&params)
	}

	values := url.Values{
		"token":     {api.token},
		"usergroup": {userGroup},
	}

	if params.IncludeDisabled {
		values.Add("include_disabled", "true")
	}

	if params.TeamID != "" {
		values.Add("team_id", params.TeamID)
	}

	response, err := api.userGroupRequest(ctx, "usergroups.users.list", values)
	if err != nil {
		return []string{}, err
	}
	return response.Users, nil
}

// UpdateUserGroupMembersOption options for the UpdateUserGroupMembers method call.
type UpdateUserGroupMembersOption func(*UpdateUserGroupMembersParams)

// UpdateUserGroupMembersParams contains arguments for UpdateUserGroupMembers method call
type UpdateUserGroupMembersParams struct {
	IncludeCount bool
	TeamID       string
}

// UpdateUserGroupMembersOptionIncludeCount include the count of User Groups (default: false)
func UpdateUserGroupMembersOptionIncludeCount(b bool) UpdateUserGroupMembersOption {
	return func(params *UpdateUserGroupMembersParams) {
		params.IncludeCount = b
	}
}

// UpdateUserGroupMembersOptionTeamID include team Id
func UpdateUserGroupMembersOptionTeamID(teamID string) UpdateUserGroupMembersOption {
	return func(params *UpdateUserGroupMembersParams) {
		params.TeamID = teamID
	}
}

// UpdateUserGroupMembers will update the members of an existing user group
func (api *Client) UpdateUserGroupMembers(userGroup string, members string, options ...UpdateUserGroupMembersOption) (UserGroup, error) {
	return api.UpdateUserGroupMembersContext(context.Background(), userGroup, members, options...)
}

// UpdateUserGroupMembersContext will update the members of an existing user group with a custom context
func (api *Client) UpdateUserGroupMembersContext(ctx context.Context, userGroup string, members string, options ...UpdateUserGroupMembersOption) (UserGroup, error) {
	params := UpdateUserGroupMembersParams{}

	for _, opt := range options {
		opt(&params)
	}

	values := url.Values{
		"token":     {api.token},
		"usergroup": {userGroup},
		"users":     {members},
	}

	if params.IncludeCount {
		values.Add("include_count", "true")
	}

	if params.TeamID != "" {
		values.Add("team_id", params.TeamID)
	}

	response, err := api.userGroupRequest(ctx, "usergroups.users.update", values)
	if err != nil {
		return UserGroup{}, err
	}
	return response.UserGroup, nil
}
