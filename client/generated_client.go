package client

type RancherCatalogClient struct {
    RancherBaseClient
	
    ApiVersion ApiVersionOperations
    Question QuestionOperations
    Template TemplateOperations
    TemplateVersion TemplateVersionOperations
    Catalog CatalogOperations
    Error ErrorOperations
}

func constructClient() *RancherCatalogClient {
	client := &RancherCatalogClient{
		RancherBaseClient: RancherBaseClient{
			Types: map[string]Schema{},
		},
	}

    
    client.ApiVersion = newApiVersionClient(client)
    client.Question = newQuestionClient(client)
    client.Template = newTemplateClient(client)
    client.TemplateVersion = newTemplateVersionClient(client)
    client.Catalog = newCatalogClient(client)
    client.Error = newErrorClient(client) 


	return client
}

func NewRancherCatalogClient(opts *ClientOpts) (*RancherCatalogClient, error) {
    client := constructClient()
        
    err := setupRancherBaseClient(&client.RancherBaseClient, opts)
    if err != nil {
        return nil, err
    }

    return client, nil
}
