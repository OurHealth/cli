package helpers

import (
	"encoding/json"
	"fmt"

	"github.com/catalyzeio/catalyze/httpclient"
	"github.com/catalyzeio/catalyze/models"
)

// ListEnvironmentInvites lists all invites for the associated environment.
func ListEnvironmentInvites(settings *models.Settings) *[]models.Invite {
	resp := httpclient.Get(fmt.Sprintf("%s/v1/environments/%s/invites", settings.PaasHost, settings.EnvironmentID), true, settings)
	var invites []models.Invite
	json.Unmarshal(resp, &invites)
	return &invites
}

// CreateInvite invites a user by email to the associated environment. This user
// does not need to have a Dashboard account to send them an invite, but
// requires a Dashboard account to accept it.
func CreateInvite(email string, settings *models.Settings) *models.Invite {
	i := models.Invite{
		Email: email,
	}
	b, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	resp := httpclient.Post(b, fmt.Sprintf("%s/v1/environments/%s/invites", settings.PaasHost, settings.EnvironmentID), true, settings)
	var invite models.Invite
	json.Unmarshal(resp, &invite)
	return &invite
}

// DeleteInvite deletes a pending invite. If an invite has already been accepted
// it cannot be deleted. Instead use the `catalyze users rm` command to revoke
// their access. This DeleteInvite method would be used if you typed the email
// incorrectly and wanted to revoke the invitation.
func DeleteInvite(inviteID string, settings *models.Settings) {
	httpclient.Delete(fmt.Sprintf("%s/v1/environments/%s/invites/%s", settings.PaasHost, settings.EnvironmentID, inviteID), true, settings)
}
