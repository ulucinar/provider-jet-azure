/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package zone

import (
	"context"
	"time"

	dnsapi "github.com/Azure/azure-sdk-for-go/services/dns/mgmt/2018-05-01/dns"
	"github.com/pkg/errors"
	"k8s.io/client-go/util/workqueue"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/event"
	"github.com/crossplane/crossplane-runtime/pkg/logging"
	"github.com/crossplane/crossplane-runtime/pkg/ratelimiter"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/resource"

	"github.com/crossplane-contrib/provider-jet-azure/apis/classic/dns/v1alpha1"
	azureclients "github.com/crossplane-contrib/provider-jet-azure/pkg/clients"
	"github.com/crossplane-contrib/provider-jet-azure/pkg/clients/dns"
)

// Error strings.
const (
	errNotDNSZone    = "managed resource is not an DNS Zone"
	errCreateDNSZone = "cannot create DNS Zone"
	errUpdateDNSZone = "cannot update DNS Zone"
	errGetDNSZone    = "cannot get DNS Zone"
	errDeleteDNSZone = "cannot delete DNS Zone"
)

// Setup adds a controller that reconciles DNS Zones.
func Setup(mgr ctrl.Manager, l logging.Logger, rl workqueue.RateLimiter, poll time.Duration) error {
	name := managed.ControllerName(v1alpha1.ZoneGroupKind)

	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(controller.Options{
			RateLimiter: ratelimiter.NewDefaultManagedRateLimiter(rl),
		}).
		For(&v1alpha1.Zone{}).
		Complete(managed.NewReconciler(mgr,
			resource.ManagedKind(v1alpha1.ZoneGroupVersionKind),
			managed.WithConnectionPublishers(),
			managed.WithExternalConnecter(&connecter{client: mgr.GetClient()}),
			managed.WithReferenceResolver(managed.NewAPISimpleReferenceResolver(mgr.GetClient())),
			managed.WithPollInterval(poll),
			managed.WithLogger(l.WithValues("controller", name)),
			managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name)))))
}

type connecter struct {
	client client.Client
}

func (c *connecter) Connect(ctx context.Context, mg resource.Managed) (managed.ExternalClient, error) {
	creds, auth, err := azureclients.GetAuthInfo(ctx, c.client, mg)
	if err != nil {
		return nil, err
	}
	cl := dnsapi.NewZonesClient(creds[azureclients.CredentialsKeySubscriptionID])
	cl.Authorizer = auth
	return &external{
		client: dns.NewZoneClient(cl),
	}, nil
}

type external struct {
	client dns.ZoneAPI
}

func (e *external) Observe(ctx context.Context, mg resource.Managed) (managed.ExternalObservation, error) {
	z, ok := mg.(*v1alpha1.Zone)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errNotDNSZone)
	}

	az, err := e.client.Get(ctx, z)
	if azureclients.IsNotFound(err) {
		return managed.ExternalObservation{ResourceExists: false}, nil
	}
	if err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, errGetDNSZone)
	}

	dns.UpdateZoneStatusFromAzure(z, az)

	z.SetConditions(xpv1.Available())

	o := managed.ExternalObservation{
		ResourceExists:   true,
		ResourceUpToDate: dns.ZoneIsUpToDate(z, az),
	}

	return o, nil
}

func (e *external) Create(ctx context.Context, mg resource.Managed) (managed.ExternalCreation, error) {
	z, ok := mg.(*v1alpha1.Zone)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errNotDNSZone)
	}

	return managed.ExternalCreation{}, errors.Wrap(e.client.CreateOrUpdate(ctx, z), errCreateDNSZone)
}

func (e *external) Update(ctx context.Context, mg resource.Managed) (managed.ExternalUpdate, error) {
	z, ok := mg.(*v1alpha1.Zone)
	if !ok {
		return managed.ExternalUpdate{}, errors.New(errNotDNSZone)
	}

	az, err := e.client.Get(ctx, z)
	if err != nil {
		return managed.ExternalUpdate{}, errors.Wrap(err, errGetDNSZone)
	}

	dns.UpdateZoneStatusFromAzure(z, az)

	return managed.ExternalUpdate{}, errors.Wrap(e.client.CreateOrUpdate(ctx, z), errUpdateDNSZone)
}

func (e *external) Delete(ctx context.Context, mg resource.Managed) error {
	z, ok := mg.(*v1alpha1.Zone)
	if !ok {
		return errors.New(errNotDNSZone)
	}

	az, err := e.client.Get(ctx, z)
	if err != nil {
		return errors.Wrap(err, errGetDNSZone)
	}

	dns.UpdateZoneStatusFromAzure(z, az)

	err = e.client.Delete(ctx, z)

	return errors.Wrap(resource.Ignore(azureclients.IsNotFound, err), errDeleteDNSZone)
}
