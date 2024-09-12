package api

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/wundergraph/cosmo/connect-go/wg/cosmo/common"
	platformv1 "github.com/wundergraph/cosmo/connect-go/wg/cosmo/platform/v1"
	"github.com/wundergraph/cosmo/connect-go/wg/cosmo/platform/v1/platformv1connect"
)

func CreateMonograph(ctx context.Context, client platformv1connect.PlatformServiceClient, apiKey string, name string, namespace string, routingUrl string, graphUrl string, subscriptionUrl *string, readme *string, websocketSubprotocol string, subscriptionProtocol string, admissionWebhookUrl string, admissionWebhookSecret string) error {
	request := connect.NewRequest(&platformv1.CreateMonographRequest{
		Name:                   name,
		Namespace:              namespace,
		RoutingUrl:             routingUrl,
		GraphUrl:               graphUrl,
		SubscriptionUrl:        subscriptionUrl,
		Readme:                 readme,
		WebsocketSubprotocol:   resolveWebsocketSubprotocol(websocketSubprotocol),
		SubscriptionProtocol:   resolveSubscriptionProtocol(subscriptionProtocol),
		AdmissionWebhookURL:    admissionWebhookUrl,
		AdmissionWebhookSecret: &admissionWebhookSecret,
	})
	request.Header().Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	_, err := client.CreateMonograph(ctx, request)
	return err
}

func UpdateMonograph(ctx context.Context, client platformv1connect.PlatformServiceClient, apiKey string, name string, namespace string, routingUrl string, graphUrl string, subscriptionUrl *string, readme *string, websocketSubprotocol string, subscriptionProtocol string, admissionWebhookUrl string, admissionWebhookSecret string) error {
	request := connect.NewRequest(&platformv1.UpdateMonographRequest{
		Name:                   name,
		Namespace:              namespace,
		RoutingUrl:             routingUrl,
		GraphUrl:               graphUrl,
		SubscriptionUrl:        subscriptionUrl,
		Readme:                 readme,
		WebsocketSubprotocol:   resolveWebsocketSubprotocol(websocketSubprotocol),
		SubscriptionProtocol:   resolveSubscriptionProtocol(subscriptionProtocol),
		AdmissionWebhookURL:    &admissionWebhookUrl,
		AdmissionWebhookSecret: &admissionWebhookSecret,
	})
	request.Header().Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	_, err := client.UpdateMonograph(ctx, request)
	return err
}

func DeleteMonograph(ctx context.Context, client platformv1connect.PlatformServiceClient, apiKey string, name string, namespace string) error {
	request := connect.NewRequest(&platformv1.DeleteMonographRequest{
		Name:      name,
		Namespace: namespace,
	})
	request.Header().Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	_, err := client.DeleteMonograph(ctx, request)
	return err
}

func GetMonograph(ctx context.Context, client platformv1connect.PlatformServiceClient, apiKey string, name string, namespace string) (*platformv1.FederatedGraph, error) {
	request := connect.NewRequest(&platformv1.GetFederatedGraphByNameRequest{
		Name:      name,
		Namespace: namespace,
	})
	request.Header().Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	response, err := client.GetFederatedGraphByName(ctx, request)
	if err != nil {
		return nil, err
	}

	if response.Msg.GetResponse().Code != common.EnumStatusCode_OK {
		return nil, fmt.Errorf("failed to get monograph: %s", response.Msg.GetResponse().GetDetails())
	}

	return response.Msg.Graph, nil
}
