package v2client

// InfoService -
type InfoService service

// GetTarget returns the current routing endpoint
//
// it uses the cf cloud foundry client to retrieve it,
// which means it doesn't really need to use this particular client
// to do it... buuut I'm just experimenting, so, eh
func (i *InfoService) GetTarget() (string, error) {
	info, err := i.client.cfc.GetInfo()
	if err != nil {
		return "", err
	}
	return info.RoutingEndpoint, nil
}
