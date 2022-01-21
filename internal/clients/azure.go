package clients

import (
	"context"
	"encoding/json"

	"github.com/crossplane/terrajet/pkg/terraform"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"

	"github.com/crossplane-contrib/provider-jet-azure/apis/v1alpha1"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal Azure credentials as JSON"
	// Azure service principal credentials file JSON keys
	keyAzureSubscriptionID = "subscriptionId"
	keyAzureClientID       = "clientId"
	keyAzureClientSecret   = "clientSecret"
	keyAzureTenantID       = "tenantId"
	// Terraform Provider configuration keys
	keyTerraformFeatures        = "features"
	keySkipProviderRegistration = "skip_provider_registration"
	// Terraform configuration file keys
	keyTerraformSubscriptionID = "subscription_id"
	keyTerraformTenantID       = "tenant_id"
	keyTerraformClientID       = "client_id"
	keyTerraformClientSecret   = "client_secret"
)

// TerraformSetupBuilder returns Terraform setup with provider specific
// configuration like provider credentials used to connect to cloud APIs in the
// expected form of a Terraform provider.
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn { //nolint:gocyclo
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1alpha1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := xpresource.NewProviderConfigUsageTracker(client, &v1alpha1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		data, err := xpresource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		azureCreds := map[string]string{}
		if err := json.Unmarshal(data, &azureCreds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		ps.Configuration = map[string]interface{}{
			keyTerraformFeatures: struct{}{},
			// Terraform AzureRM provider tries to register all resource providers
			// in Azure just in case if the provider of the resource you're
			// trying to create is not registered and the returned error is
			// ambigious. However, this requires service principal to have provider
			// registration permissions which are irrelevant in most contexts.
			// For details, see https://github.com/crossplane-contrib/provider-jet-azure/issues/104
			keySkipProviderRegistration: true,
			keyTerraformSubscriptionID:  azureCreds[keyAzureSubscriptionID],
			keyTerraformTenantID:        azureCreds[keyAzureTenantID],
			keyTerraformClientID:        azureCreds[keyAzureClientID],
			keyTerraformClientSecret:    azureCreds[keyAzureClientSecret],
		}
		return ps, err
	}
}
