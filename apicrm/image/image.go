package image

func (svc *APIImageService) Get(imageID string) ([]byte, error) {
	response, err := svc.client.CRMHandlerImage(imageID)

	return response, err
}
