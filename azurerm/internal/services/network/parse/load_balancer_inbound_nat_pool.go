package parse

// NOTE: this file is generated via 'go:generate' - manual changes will be overwritten

import (
	"fmt"

	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/azure"
)

type LoadBalancerInboundNatPoolId struct {
	SubscriptionId     string
	ResourceGroup      string
	LoadBalancerName   string
	InboundNatPoolName string
}

func NewLoadBalancerInboundNatPoolID(subscriptionId, resourceGroup, loadBalancerName, inboundNatPoolName string) LoadBalancerInboundNatPoolId {
	return LoadBalancerInboundNatPoolId{
		SubscriptionId:     subscriptionId,
		ResourceGroup:      resourceGroup,
		LoadBalancerName:   loadBalancerName,
		InboundNatPoolName: inboundNatPoolName,
	}
}

func (id LoadBalancerInboundNatPoolId) ID(_ string) string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Network/loadBalancers/%s/inboundNatPools/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroup, id.LoadBalancerName, id.InboundNatPoolName)
}

// LoadBalancerInboundNatPoolID parses a LoadBalancerInboundNatPool ID into an LoadBalancerInboundNatPoolId struct
func LoadBalancerInboundNatPoolID(input string) (*LoadBalancerInboundNatPoolId, error) {
	id, err := azure.ParseAzureResourceID(input)
	if err != nil {
		return nil, err
	}

	resourceId := LoadBalancerInboundNatPoolId{
		SubscriptionId: id.SubscriptionID,
		ResourceGroup:  id.ResourceGroup,
	}

	if resourceId.LoadBalancerName, err = id.PopSegment("loadBalancers"); err != nil {
		return nil, err
	}
	if resourceId.InboundNatPoolName, err = id.PopSegment("inboundNatPools"); err != nil {
		return nil, err
	}

	if err := id.ValidateNoEmptySegments(input); err != nil {
		return nil, err
	}

	return &resourceId, nil
}
