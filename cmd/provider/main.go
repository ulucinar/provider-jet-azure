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

package main

import (
	"os"
	"path/filepath"

	"gopkg.in/alecthomas/kingpin.v2"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	xpcontroller "github.com/crossplane/crossplane-runtime/pkg/controller"

	"github.com/crossplane/crossplane-runtime/pkg/logging"
	"github.com/crossplane/crossplane-runtime/pkg/ratelimiter"

	tjcontroller "github.com/crossplane/terrajet/pkg/controller"
	"github.com/crossplane/terrajet/pkg/terraform"

	"github.com/crossplane-contrib/provider-jet-azure/apis"
	"github.com/crossplane-contrib/provider-jet-azure/config"

	// genConfig "github.com/crossplane-contrib/provider-jet-azure/cmd/generator/config"
	classicapis "github.com/crossplane-contrib/provider-jet-azure/apis/classic"
	"github.com/crossplane-contrib/provider-jet-azure/internal/clients"
	"github.com/crossplane-contrib/provider-jet-azure/internal/controller"
	classiccontrollers "github.com/crossplane-contrib/provider-jet-azure/pkg/controller"
)

func main() {
	var (
		app              = kingpin.New(filepath.Base(os.Args[0]), "Azure support for Crossplane.").DefaultEnvars()
		debug            = app.Flag("debug", "Run with debug logging.").Short('d').Bool()
		syncPeriod       = app.Flag("sync", "Controller manager sync period such as 300ms, 1.5h, or 2h45m").Short('s').Default("1h").Duration()
		leaderElection   = app.Flag("leader-election", "Use leader election for the controller manager.").Short('l').Default("false").OverrideDefaultFromEnvar("LEADER_ELECTION").Bool()
		terraformVersion = app.Flag("terraform-version", "Terraform version.").Required().Envar("TERRAFORM_VERSION").String()
		providerSource   = app.Flag("terraform-provider-source", "Terraform provider source.").Required().Envar("TERRAFORM_PROVIDER_SOURCE").String()
		providerVersion  = app.Flag("terraform-provider-version", "Terraform provider version.").Required().Envar("TERRAFORM_PROVIDER_VERSION").String()
	)
	kingpin.MustParse(app.Parse(os.Args[1:]))

	zl := zap.New(zap.UseDevMode(*debug))
	log := logging.NewLogrLogger(zl.WithName("provider-jet-azure"))
	if *debug {
		// The controller-runtime runs with a no-op logger by default. It is
		// *very* verbose even at info level, so we only provide it a real
		// logger when we're running in debug mode.
		ctrl.SetLogger(zl)
	}

	log.Debug("Starting", "sync-period", syncPeriod.String())

	cfg, err := ctrl.GetConfig()
	kingpin.FatalIfError(err, "Cannot get API server rest config")

	mgr, err := ctrl.NewManager(cfg, ctrl.Options{
		LeaderElection:   *leaderElection,
		LeaderElectionID: "crossplane-leader-election-provider-jet-azure",
		SyncPeriod:       syncPeriod,
	})
	kingpin.FatalIfError(err, "Cannot create controller manager")

	// genConfig.SetResourceConfigurations()
	ws := terraform.NewWorkspaceStore(log)
	setup := clients.TerraformSetupBuilder(*terraformVersion, *providerSource, *providerVersion)
	rl := ratelimiter.NewGlobal(1)
	kingpin.FatalIfError(apis.AddToScheme(mgr.GetScheme()), "Cannot add Azure APIs to scheme")
	kingpin.FatalIfError(classicapis.AddToScheme(mgr.GetScheme()), "Cannot add classic Azure APIs to scheme")
	kingpin.FatalIfError(classiccontrollers.Setup(mgr, log, rl, *syncPeriod), "Cannot setup classic Azure controllers")
	kingpin.FatalIfError(controller.Setup(mgr, tjcontroller.Options{
		Options: xpcontroller.Options{
			Logger:            log,
			GlobalRateLimiter: rl,
		},
		Provider:       config.GetProvider(),
		WorkspaceStore: ws,
		SetupFn:        setup,
	}),
		"Cannot setup Azure controllers")
	kingpin.FatalIfError(mgr.Start(ctrl.SetupSignalHandler()), "Cannot start controller manager")
}
