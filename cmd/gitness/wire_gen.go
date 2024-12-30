// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"context"

	aiagent2 "github.com/harness/gitness/app/api/controller/aiagent"
	capabilities2 "github.com/harness/gitness/app/api/controller/capabilities"
	check2 "github.com/harness/gitness/app/api/controller/check"
	connector2 "github.com/harness/gitness/app/api/controller/connector"
	"github.com/harness/gitness/app/api/controller/execution"
	"github.com/harness/gitness/app/api/controller/githook"
	gitspace2 "github.com/harness/gitness/app/api/controller/gitspace"
	infraprovider3 "github.com/harness/gitness/app/api/controller/infraprovider"
	keywordsearch2 "github.com/harness/gitness/app/api/controller/keywordsearch"
	"github.com/harness/gitness/app/api/controller/limiter"
	logs2 "github.com/harness/gitness/app/api/controller/logs"
	migrate2 "github.com/harness/gitness/app/api/controller/migrate"
	"github.com/harness/gitness/app/api/controller/pipeline"
	"github.com/harness/gitness/app/api/controller/plugin"
	"github.com/harness/gitness/app/api/controller/principal"
	pullreq2 "github.com/harness/gitness/app/api/controller/pullreq"
	"github.com/harness/gitness/app/api/controller/repo"
	"github.com/harness/gitness/app/api/controller/reposettings"
	secret2 "github.com/harness/gitness/app/api/controller/secret"
	"github.com/harness/gitness/app/api/controller/service"
	"github.com/harness/gitness/app/api/controller/serviceaccount"
	"github.com/harness/gitness/app/api/controller/space"
	"github.com/harness/gitness/app/api/controller/system"
	"github.com/harness/gitness/app/api/controller/template"
	"github.com/harness/gitness/app/api/controller/trigger"
	"github.com/harness/gitness/app/api/controller/upload"
	"github.com/harness/gitness/app/api/controller/user"
	usergroup2 "github.com/harness/gitness/app/api/controller/usergroup"
	webhook2 "github.com/harness/gitness/app/api/controller/webhook"
	"github.com/harness/gitness/app/api/openapi"
	"github.com/harness/gitness/app/auth/authn"
	"github.com/harness/gitness/app/auth/authz"
	"github.com/harness/gitness/app/bootstrap"
	"github.com/harness/gitness/app/connector"
	events7 "github.com/harness/gitness/app/events/git"
	events3 "github.com/harness/gitness/app/events/gitspace"
	events4 "github.com/harness/gitness/app/events/gitspaceinfra"
	events5 "github.com/harness/gitness/app/events/pipeline"
	events6 "github.com/harness/gitness/app/events/pullreq"
	events2 "github.com/harness/gitness/app/events/repo"
	"github.com/harness/gitness/app/gitspace/infrastructure"
	"github.com/harness/gitness/app/gitspace/logutil"
	"github.com/harness/gitness/app/gitspace/orchestrator"
	"github.com/harness/gitness/app/gitspace/orchestrator/container"
	"github.com/harness/gitness/app/gitspace/orchestrator/ide"
	"github.com/harness/gitness/app/gitspace/orchestrator/runarg"
	"github.com/harness/gitness/app/gitspace/platformconnector"
	"github.com/harness/gitness/app/gitspace/scm"
	"github.com/harness/gitness/app/gitspace/secret"
	"github.com/harness/gitness/app/pipeline/canceler"
	"github.com/harness/gitness/app/pipeline/commit"
	"github.com/harness/gitness/app/pipeline/converter"
	"github.com/harness/gitness/app/pipeline/file"
	"github.com/harness/gitness/app/pipeline/manager"
	"github.com/harness/gitness/app/pipeline/resolver"
	"github.com/harness/gitness/app/pipeline/runner"
	"github.com/harness/gitness/app/pipeline/scheduler"
	"github.com/harness/gitness/app/pipeline/triggerer"
	router2 "github.com/harness/gitness/app/router"
	server2 "github.com/harness/gitness/app/server"
	"github.com/harness/gitness/app/services"
	"github.com/harness/gitness/app/services/aiagent"
	"github.com/harness/gitness/app/services/capabilities"
	"github.com/harness/gitness/app/services/cleanup"
	"github.com/harness/gitness/app/services/codecomments"
	"github.com/harness/gitness/app/services/codeowners"
	"github.com/harness/gitness/app/services/exporter"
	"github.com/harness/gitness/app/services/gitspace"
	"github.com/harness/gitness/app/services/gitspaceevent"
	"github.com/harness/gitness/app/services/gitspaceinfraevent"
	"github.com/harness/gitness/app/services/importer"
	infraprovider2 "github.com/harness/gitness/app/services/infraprovider"
	"github.com/harness/gitness/app/services/instrument"
	"github.com/harness/gitness/app/services/keywordsearch"
	"github.com/harness/gitness/app/services/label"
	"github.com/harness/gitness/app/services/locker"
	"github.com/harness/gitness/app/services/messaging"
	"github.com/harness/gitness/app/services/metric"
	"github.com/harness/gitness/app/services/migrate"
	"github.com/harness/gitness/app/services/notification"
	"github.com/harness/gitness/app/services/notification/mailer"
	"github.com/harness/gitness/app/services/protection"
	"github.com/harness/gitness/app/services/publicaccess"
	"github.com/harness/gitness/app/services/publickey"
	"github.com/harness/gitness/app/services/pullreq"
	"github.com/harness/gitness/app/services/refcache"
	repo2 "github.com/harness/gitness/app/services/repo"
	"github.com/harness/gitness/app/services/rules"
	secret3 "github.com/harness/gitness/app/services/secret"
	"github.com/harness/gitness/app/services/settings"
	system2 "github.com/harness/gitness/app/services/system"
	trigger2 "github.com/harness/gitness/app/services/trigger"
	"github.com/harness/gitness/app/services/usage"
	"github.com/harness/gitness/app/services/usergroup"
	"github.com/harness/gitness/app/services/webhook"
	"github.com/harness/gitness/app/sse"
	"github.com/harness/gitness/app/store"
	"github.com/harness/gitness/app/store/cache"
	"github.com/harness/gitness/app/store/database"
	"github.com/harness/gitness/app/store/logs"
	"github.com/harness/gitness/app/url"
	"github.com/harness/gitness/audit"
	"github.com/harness/gitness/blob"
	"github.com/harness/gitness/cli/operations/server"
	"github.com/harness/gitness/encrypt"
	"github.com/harness/gitness/events"
	"github.com/harness/gitness/git"
	"github.com/harness/gitness/git/api"
	"github.com/harness/gitness/git/storage"
	"github.com/harness/gitness/infraprovider"
	"github.com/harness/gitness/job"
	"github.com/harness/gitness/livelog"
	"github.com/harness/gitness/lock"
	"github.com/harness/gitness/pubsub"
	api2 "github.com/harness/gitness/registry/app/api"
	"github.com/harness/gitness/registry/app/api/router"
	"github.com/harness/gitness/registry/app/pkg"
	"github.com/harness/gitness/registry/app/pkg/docker"
	database2 "github.com/harness/gitness/registry/app/store/database"
	"github.com/harness/gitness/registry/gc"
	"github.com/harness/gitness/ssh"
	"github.com/harness/gitness/store/database/dbtx"
	"github.com/harness/gitness/types"
	"github.com/harness/gitness/types/check"

	_ "github.com/lib/pq"

	_ "github.com/mattn/go-sqlite3"
)

// Injectors from wire.go:

func initSystem(ctx context.Context, config *types.Config) (*server.System, error) {
	databaseConfig := server.ProvideDatabaseConfig(config)
	db, err := database.ProvideDatabase(ctx, databaseConfig)
	if err != nil {
		return nil, err
	}
	accessorTx := dbtx.ProvideAccessorTx(db)
	transactor := dbtx.ProvideTransactor(accessorTx)
	principalUID := check.ProvidePrincipalUIDCheck()
	spacePathTransformation := store.ProvidePathTransformation()
	spacePathStore := database.ProvideSpacePathStore(db, spacePathTransformation)
	spacePathCache := cache.ProvidePathCache(spacePathStore, spacePathTransformation)
	spaceStore := database.ProvideSpaceStore(db, spacePathCache, spacePathStore)
	principalInfoView := database.ProvidePrincipalInfoView(db)
	principalInfoCache := cache.ProvidePrincipalInfoCache(principalInfoView)
	membershipStore := database.ProvideMembershipStore(db, principalInfoCache, spacePathStore, spaceStore)
	permissionCache := authz.ProvidePermissionCache(spaceStore, membershipStore)
	publicAccessStore := database.ProvidePublicAccessStore(db)
	spaceCache := refcache.ProvideSpaceCache(spacePathStore, spaceStore, spacePathTransformation)
	repoStore := database.ProvideRepoStore(db, spacePathCache, spacePathStore, spaceStore)
	repoFinder := refcache.ProvideRepoFinder(repoStore, spaceCache)
	publicaccessService := publicaccess.ProvidePublicAccess(config, publicAccessStore, spaceCache, repoFinder)
	authorizer := authz.ProvideAuthorizer(permissionCache, spaceStore, publicaccessService)
	principalUIDTransformation := store.ProvidePrincipalUIDTransformation()
	principalStore := database.ProvidePrincipalStore(db, principalUIDTransformation)
	tokenStore := database.ProvideTokenStore(db)
	publicKeyStore := database.ProvidePublicKeyStore(db)
	controller := user.ProvideController(transactor, principalUID, authorizer, principalStore, tokenStore, membershipStore, publicKeyStore)
	serviceController := service.NewController(principalUID, authorizer, principalStore)
	bootstrapBootstrap := bootstrap.ProvideBootstrap(config, controller, serviceController)
	authenticator := authn.ProvideAuthenticator(config, principalStore, tokenStore)
	provider, err := url.ProvideURLProvider(config)
	if err != nil {
		return nil, err
	}
	pipelineStore := database.ProvidePipelineStore(db)
	executionStore := database.ProvideExecutionStore(db)
	ruleStore := database.ProvideRuleStore(db, principalInfoCache)
	checkStore := database.ProvideCheckStore(db, principalInfoCache)
	pullReqStore := database.ProvidePullReqStore(db, principalInfoCache)
	settingsStore := database.ProvideSettingsStore(db)
	settingsService := settings.ProvideService(settingsStore)
	protectionManager, err := protection.ProvideManager(ruleStore)
	if err != nil {
		return nil, err
	}
	typesConfig := server.ProvideGitConfig(config)
	universalClient, err := server.ProvideRedis(config)
	if err != nil {
		return nil, err
	}
	cacheCache, err := api.ProvideLastCommitCache(typesConfig, universalClient)
	if err != nil {
		return nil, err
	}
	clientFactory := githook.ProvideFactory()
	apiGit, err := git.ProvideGITAdapter(typesConfig, cacheCache, clientFactory)
	if err != nil {
		return nil, err
	}
	storageStore := storage.ProvideLocalStore()
	gitInterface, err := git.ProvideService(typesConfig, apiGit, clientFactory, storageStore)
	if err != nil {
		return nil, err
	}
	triggerStore := database.ProvideTriggerStore(db)
	encrypter, err := encrypt.ProvideEncrypter(config)
	if err != nil {
		return nil, err
	}
	jobStore := database.ProvideJobStore(db)
	pubsubConfig := server.ProvidePubsubConfig(config)
	pubSub := pubsub.ProvidePubSub(pubsubConfig, universalClient)
	executor := job.ProvideExecutor(jobStore, pubSub)
	lockConfig := server.ProvideLockConfig(config)
	mutexManager := lock.ProvideMutexManager(lockConfig, universalClient)
	jobConfig := server.ProvideJobsConfig(config)
	jobScheduler, err := job.ProvideScheduler(jobStore, executor, mutexManager, pubSub, jobConfig)
	if err != nil {
		return nil, err
	}
	streamer := sse.ProvideEventsStreaming(pubSub)
	localIndexSearcher := keywordsearch.ProvideLocalIndexSearcher()
	indexer := keywordsearch.ProvideIndexer(localIndexSearcher)
	auditService := audit.ProvideAuditService()
	repository, err := importer.ProvideRepoImporter(config, provider, gitInterface, transactor, repoStore, pipelineStore, triggerStore, encrypter, jobScheduler, executor, streamer, indexer, publicaccessService, auditService)
	if err != nil {
		return nil, err
	}
	codeownersConfig := server.ProvideCodeOwnerConfig(config)
	usergroupResolver := usergroup.ProvideUserGroupResolver()
	codeownersService := codeowners.ProvideCodeOwners(gitInterface, repoStore, codeownersConfig, principalStore, usergroupResolver)
	eventsConfig := server.ProvideEventsConfig(config)
	eventsSystem, err := events.ProvideSystem(eventsConfig, universalClient)
	if err != nil {
		return nil, err
	}
	reporter, err := events2.ProvideReporter(eventsSystem)
	if err != nil {
		return nil, err
	}
	resourceLimiter, err := limiter.ProvideLimiter()
	if err != nil {
		return nil, err
	}
	lockerLocker := locker.ProvideLocker(mutexManager)
	repoIdentifier := check.ProvideRepoIdentifierCheck()
	repoCheck := repo.ProvideRepoCheck()
	labelStore := database.ProvideLabelStore(db)
	labelValueStore := database.ProvideLabelValueStore(db)
	pullReqLabelAssignmentStore := database.ProvidePullReqLabelStore(db)
	labelService := label.ProvideLabel(transactor, spaceStore, labelStore, labelValueStore, pullReqLabelAssignmentStore)
	instrumentService := instrument.ProvideService()
	userGroupStore := database.ProvideUserGroupStore(db)
	searchService := usergroup.ProvideSearchService()
	rulesService := rules.ProvideService(transactor, ruleStore, repoStore, spaceStore, protectionManager, auditService, instrumentService, principalInfoCache, userGroupStore, searchService, streamer)
	repoController := repo.ProvideController(config, transactor, provider, authorizer, repoStore, spaceStore, pipelineStore, principalStore, executionStore, ruleStore, checkStore, pullReqStore, settingsService, principalInfoCache, protectionManager, gitInterface, spaceCache, repoFinder, repository, codeownersService, reporter, indexer, resourceLimiter, lockerLocker, auditService, mutexManager, repoIdentifier, repoCheck, publicaccessService, labelService, instrumentService, userGroupStore, searchService, rulesService, streamer)
	reposettingsController := reposettings.ProvideController(authorizer, repoFinder, settingsService, auditService)
	stageStore := database.ProvideStageStore(db)
	schedulerScheduler, err := scheduler.ProvideScheduler(stageStore, mutexManager)
	if err != nil {
		return nil, err
	}
	stepStore := database.ProvideStepStore(db)
	cancelerCanceler := canceler.ProvideCanceler(executionStore, streamer, repoStore, schedulerScheduler, stageStore, stepStore)
	commitService := commit.ProvideService(gitInterface)
	fileService := file.ProvideService(gitInterface)
	converterService := converter.ProvideService(fileService, publicaccessService)
	templateStore := database.ProvideTemplateStore(db)
	pluginStore := database.ProvidePluginStore(db)
	triggererTriggerer := triggerer.ProvideTriggerer(executionStore, checkStore, stageStore, transactor, pipelineStore, fileService, converterService, schedulerScheduler, repoStore, provider, templateStore, pluginStore, publicaccessService)
	executionController := execution.ProvideController(transactor, authorizer, executionStore, checkStore, cancelerCanceler, commitService, triggererTriggerer, stageStore, pipelineStore, repoFinder)
	logStore := logs.ProvideLogStore(db, config)
	logStream := livelog.ProvideLogStream()
	logsController := logs2.ProvideController(authorizer, executionStore, pipelineStore, stageStore, stepStore, logStore, logStream, repoFinder)
	spaceIdentifier := check.ProvideSpaceIdentifierCheck()
	secretStore := database.ProvideSecretStore(db)
	connectorStore := database.ProvideConnectorStore(db, secretStore)
	repoGitInfoView := database.ProvideRepoGitInfoView(db)
	repoGitInfoCache := cache.ProvideRepoGitInfoCache(repoGitInfoView)
	listService := pullreq.ProvideListService(transactor, gitInterface, authorizer, spaceStore, repoStore, repoGitInfoCache, pullReqStore, checkStore, labelService, protectionManager)
	exporterRepository, err := exporter.ProvideSpaceExporter(provider, gitInterface, repoStore, jobScheduler, executor, encrypter, streamer)
	if err != nil {
		return nil, err
	}
	infraProviderResourceView := database.ProvideInfraProviderResourceView(db)
	infraProviderResourceCache := cache.ProvideInfraProviderResourceCache(infraProviderResourceView)
	gitspaceConfigStore := database.ProvideGitspaceConfigStore(db, principalInfoCache, infraProviderResourceCache)
	gitspaceInstanceStore := database.ProvideGitspaceInstanceStore(db)
	eventsReporter, err := events3.ProvideReporter(eventsSystem)
	if err != nil {
		return nil, err
	}
	gitspaceEventStore := database.ProvideGitspaceEventStore(db)
	infraProviderResourceStore := database.ProvideInfraProviderResourceStore(db)
	infraProviderConfigStore := database.ProvideInfraProviderConfigStore(db)
	infraProviderTemplateStore := database.ProvideInfraProviderTemplateStore(db)
	dockerConfig, err := server.ProvideDockerConfig(config)
	if err != nil {
		return nil, err
	}
	dockerClientFactory := infraprovider.ProvideDockerClientFactory(dockerConfig)
	reporter2, err := events4.ProvideReporter(eventsSystem)
	if err != nil {
		return nil, err
	}
	dockerProvider := infraprovider.ProvideDockerProvider(dockerConfig, dockerClientFactory, reporter2)
	factory := infraprovider.ProvideFactory(dockerProvider)
	infraproviderService := infraprovider2.ProvideInfraProvider(transactor, infraProviderResourceStore, infraProviderConfigStore, infraProviderTemplateStore, factory, spaceStore)
	gitnessSCM := scm.ProvideGitnessSCM(repoStore, repoFinder, gitInterface, tokenStore, principalStore, provider)
	genericSCM := scm.ProvideGenericSCM()
	scmFactory := scm.ProvideFactory(gitnessSCM, genericSCM)
	scmSCM := scm.ProvideSCM(scmFactory)
	platformConnector := platformconnector.ProvideGitnessPlatformConnector()
	infraProvisionedStore := database.ProvideInfraProvisionedStore(db)
	infrastructureConfig := server.ProvideGitspaceInfraProvisionerConfig(config)
	infraProvisioner := infrastructure.ProvideInfraProvisionerService(infraProviderConfigStore, infraProviderResourceStore, factory, infraProviderTemplateStore, infraProvisionedStore, infrastructureConfig)
	statefulLogger := logutil.ProvideStatefulLogger(logStream)
	runargResolver, err := runarg.ProvideResolver()
	if err != nil {
		return nil, err
	}
	runargProvider, err := runarg.ProvideStaticProvider(runargResolver)
	if err != nil {
		return nil, err
	}
	containerOrchestrator := container.ProvideEmbeddedDockerOrchestrator(dockerClientFactory, statefulLogger, runargProvider)
	orchestratorConfig := server.ProvideGitspaceOrchestratorConfig(config)
	vsCodeConfig := server.ProvideIDEVSCodeConfig(config)
	vsCode := ide.ProvideVSCodeService(vsCodeConfig)
	vsCodeWebConfig := server.ProvideIDEVSCodeWebConfig(config)
	vsCodeWeb := ide.ProvideVSCodeWebService(vsCodeWebConfig)
	intellijConfig := server.ProvideIDEIntellijConfig(config)
	intellij := ide.ProvideIntellijService(intellijConfig)
	ideFactory := ide.ProvideIDEFactory(vsCode, vsCodeWeb, intellij)
	passwordResolver := secret.ProvidePasswordResolver()
	resolverFactory := secret.ProvideResolverFactory(passwordResolver)
	orchestratorOrchestrator := orchestrator.ProvideOrchestrator(scmSCM, platformConnector, infraProvisioner, containerOrchestrator, eventsReporter, orchestratorConfig, ideFactory, resolverFactory)
	gitspaceService := gitspace.ProvideGitspace(transactor, gitspaceConfigStore, gitspaceInstanceStore, eventsReporter, gitspaceEventStore, spaceStore, infraproviderService, orchestratorOrchestrator, scmSCM, config)
	usageMetricStore := database.ProvideUsageMetricStore(db)
	spaceController := space.ProvideController(config, transactor, provider, streamer, spaceIdentifier, authorizer, spacePathStore, pipelineStore, secretStore, connectorStore, templateStore, spaceStore, repoStore, principalStore, repoController, membershipStore, listService, spaceCache, repository, exporterRepository, resourceLimiter, publicaccessService, auditService, gitspaceService, labelService, instrumentService, executionStore, rulesService, usageMetricStore)
	reporter3, err := events5.ProvideReporter(eventsSystem)
	if err != nil {
		return nil, err
	}
	pipelineController := pipeline.ProvideController(triggerStore, authorizer, pipelineStore, reporter3, repoFinder)
	secretController := secret2.ProvideController(encrypter, secretStore, authorizer, spaceStore)
	triggerController := trigger.ProvideController(authorizer, triggerStore, pipelineStore, repoFinder)
	scmService := connector.ProvideSCMConnectorHandler(secretStore)
	connectorService := connector.ProvideConnectorHandler(secretStore, scmService)
	connectorController := connector2.ProvideController(connectorStore, connectorService, authorizer, spaceCache)
	templateController := template.ProvideController(templateStore, authorizer, spaceStore)
	pluginController := plugin.ProvideController(pluginStore)
	pullReqActivityStore := database.ProvidePullReqActivityStore(db, principalInfoCache)
	codeCommentView := database.ProvideCodeCommentView(db)
	pullReqReviewStore := database.ProvidePullReqReviewStore(db)
	pullReqReviewerStore := database.ProvidePullReqReviewerStore(db, principalInfoCache)
	userGroupReviewersStore := database.ProvideUserGroupReviewerStore(db, principalInfoCache, userGroupStore)
	pullReqFileViewStore := database.ProvidePullReqFileViewStore(db)
	reporter4, err := events6.ProvideReporter(eventsSystem)
	if err != nil {
		return nil, err
	}
	migrator := codecomments.ProvideMigrator(gitInterface)
	readerFactory, err := events7.ProvideReaderFactory(eventsSystem)
	if err != nil {
		return nil, err
	}
	eventsReaderFactory, err := events6.ProvideReaderFactory(eventsSystem)
	if err != nil {
		return nil, err
	}
	pullreqService, err := pullreq.ProvideService(ctx, config, readerFactory, eventsReaderFactory, reporter4, gitInterface, repoGitInfoCache, repoStore, pullReqStore, pullReqActivityStore, principalInfoCache, codeCommentView, migrator, pullReqFileViewStore, pubSub, provider, streamer)
	if err != nil {
		return nil, err
	}
	pullReq := migrate.ProvidePullReqImporter(provider, gitInterface, principalStore, spaceStore, repoStore, pullReqStore, pullReqActivityStore, labelStore, labelValueStore, pullReqLabelAssignmentStore, transactor, mutexManager)
	pullreqController := pullreq2.ProvideController(transactor, provider, authorizer, auditService, pullReqStore, pullReqActivityStore, codeCommentView, pullReqReviewStore, pullReqReviewerStore, repoStore, principalStore, userGroupStore, userGroupReviewersStore, principalInfoCache, pullReqFileViewStore, membershipStore, checkStore, gitInterface, repoFinder, reporter4, migrator, pullreqService, listService, protectionManager, streamer, codeownersService, lockerLocker, pullReq, labelService, instrumentService, searchService)
	webhookConfig := server.ProvideWebhookConfig(config)
	webhookStore := database.ProvideWebhookStore(db)
	webhookExecutionStore := database.ProvideWebhookExecutionStore(db)
	urlProvider := webhook.ProvideURLProvider(ctx)
	webhookService, err := webhook.ProvideService(ctx, webhookConfig, transactor, readerFactory, eventsReaderFactory, webhookStore, webhookExecutionStore, spaceStore, repoStore, pullReqStore, pullReqActivityStore, provider, principalStore, gitInterface, encrypter, labelStore, urlProvider, labelValueStore, streamer)
	if err != nil {
		return nil, err
	}
	preprocessor := webhook2.ProvidePreprocessor()
	webhookController := webhook2.ProvideController(authorizer, spaceCache, repoFinder, webhookService, encrypter, preprocessor)
	reporter5, err := events7.ProvideReporter(eventsSystem)
	if err != nil {
		return nil, err
	}
	preReceiveExtender, err := githook.ProvidePreReceiveExtender()
	if err != nil {
		return nil, err
	}
	updateExtender, err := githook.ProvideUpdateExtender()
	if err != nil {
		return nil, err
	}
	postReceiveExtender, err := githook.ProvidePostReceiveExtender()
	if err != nil {
		return nil, err
	}
	githookController := githook.ProvideController(authorizer, principalStore, repoStore, reporter5, reporter, gitInterface, pullReqStore, provider, protectionManager, clientFactory, resourceLimiter, settingsService, preReceiveExtender, updateExtender, postReceiveExtender, streamer)
	serviceaccountController := serviceaccount.NewController(principalUID, authorizer, principalStore, spaceStore, repoStore, tokenStore)
	principalController := principal.ProvideController(principalStore, authorizer)
	usergroupController := usergroup2.ProvideController(userGroupStore, spaceStore, authorizer, searchService)
	v := check2.ProvideCheckSanitizers()
	checkController := check2.ProvideController(transactor, authorizer, spaceStore, checkStore, spaceCache, repoFinder, gitInterface, v, streamer)
	systemController := system.NewController(principalStore, config)
	blobConfig, err := server.ProvideBlobStoreConfig(config)
	if err != nil {
		return nil, err
	}
	blobStore, err := blob.ProvideStore(ctx, blobConfig)
	if err != nil {
		return nil, err
	}
	uploadController := upload.ProvideController(authorizer, repoFinder, blobStore)
	searcher := keywordsearch.ProvideSearcher(localIndexSearcher)
	keywordsearchController := keywordsearch2.ProvideController(authorizer, searcher, repoController, spaceController)
	infraproviderController := infraprovider3.ProvideController(authorizer, spaceStore, infraproviderService)
	limiterGitspace := limiter.ProvideGitspaceLimiter()
	gitspaceController := gitspace2.ProvideController(transactor, authorizer, infraproviderService, gitspaceConfigStore, gitspaceInstanceStore, spaceStore, gitspaceEventStore, statefulLogger, scmSCM, gitspaceService, limiterGitspace, repoFinder)
	rule := migrate.ProvideRuleImporter(ruleStore, transactor, principalStore)
	migrateWebhook := migrate.ProvideWebhookImporter(webhookConfig, transactor, webhookStore)
	migrateLabel := migrate.ProvideLabelImporter(transactor, labelStore, labelValueStore, spaceStore)
	migrateController := migrate2.ProvideController(authorizer, publicaccessService, gitInterface, provider, pullReq, rule, migrateWebhook, migrateLabel, resourceLimiter, auditService, repoIdentifier, transactor, spaceStore, repoStore, spaceCache, repoFinder)
	registry, err := capabilities.ProvideCapabilities()
	if err != nil {
		return nil, err
	}
	capabilitiesController := capabilities2.ProvideController(registry)
	intelligence, err := aiagent.ProvideAiAgent(authorizer, registry, capabilitiesController)
	if err != nil {
		return nil, err
	}
	slack, err := messaging.ProvideSlack(repoStore, gitInterface)
	if err != nil {
		return nil, err
	}
	aiagentController := aiagent2.ProvideController(authorizer, intelligence, repoFinder, pipelineStore, executionStore, gitInterface, provider, slack)
	openapiService := openapi.ProvideOpenAPIService()
	storageDriver, err := api2.BlobStorageProvider(config)
	if err != nil {
		return nil, err
	}
	storageDeleter := gc.StorageDeleterProvider(storageDriver)
	mediaTypesRepository := database2.ProvideMediaTypeDao(db)
	blobRepository := database2.ProvideBlobDao(db, mediaTypesRepository)
	storageService := docker.StorageServiceProvider(config, storageDriver)
	gcService := gc.ServiceProvider()
	app := docker.NewApp(ctx, storageDeleter, blobRepository, spaceStore, config, storageService, gcService)
	registryRepository := database2.ProvideRepoDao(db, mediaTypesRepository)
	manifestRepository := database2.ProvideManifestDao(db, mediaTypesRepository)
	manifestReferenceRepository := database2.ProvideManifestRefDao(db)
	tagRepository := database2.ProvideTagDao(db)
	imageRepository := database2.ProvideImageDao(db)
	artifactRepository := database2.ProvideArtifactDao(db)
	layerRepository := database2.ProvideLayerDao(db, mediaTypesRepository)
	eventReporter := docker.ProvideReporter()
	ociImageIndexMappingRepository := database2.ProvideOCIImageIndexMappingDao(db)
	manifestService := docker.ManifestServiceProvider(registryRepository, manifestRepository, blobRepository, mediaTypesRepository, manifestReferenceRepository, tagRepository, imageRepository, artifactRepository, layerRepository, gcService, transactor, eventReporter, spacePathStore, ociImageIndexMappingRepository)
	registryBlobRepository := database2.ProvideRegistryBlobDao(db)
	bandwidthStatRepository := database2.ProvideBandwidthStatDao(db)
	downloadStatRepository := database2.ProvideDownloadStatDao(db)
	localRegistry := docker.LocalRegistryProvider(app, manifestService, blobRepository, registryRepository, manifestRepository, registryBlobRepository, mediaTypesRepository, tagRepository, imageRepository, artifactRepository, bandwidthStatRepository, downloadStatRepository, gcService, transactor)
	upstreamProxyConfigRepository := database2.ProvideUpstreamDao(db, registryRepository, spacePathStore)
	secretService := secret3.ProvideSecretService(secretStore, encrypter, spacePathStore)
	proxyController := docker.ProvideProxyController(localRegistry, manifestService, secretService, spacePathStore)
	remoteRegistry := docker.RemoteRegistryProvider(localRegistry, app, upstreamProxyConfigRepository, spacePathStore, secretService, proxyController)
	coreController := pkg.CoreControllerProvider(registryRepository)
	dbStore := docker.DBStoreProvider(blobRepository, imageRepository, artifactRepository, bandwidthStatRepository, downloadStatRepository)
	dockerController := docker.ControllerProvider(localRegistry, remoteRegistry, coreController, spaceStore, authorizer, dbStore)
	handler := api2.NewHandlerProvider(dockerController, spaceStore, tokenStore, controller, authenticator, provider, authorizer, config)
	registryOCIHandler := router.OCIHandlerProvider(handler)
	cleanupPolicyRepository := database2.ProvideCleanupPolicyDao(db, transactor)
	apiHandler := router.APIHandlerProvider(registryRepository, upstreamProxyConfigRepository, tagRepository, manifestRepository, cleanupPolicyRepository, imageRepository, storageDriver, spaceStore, transactor, authenticator, provider, authorizer, auditService, spacePathStore)
	appRouter := router.AppRouterProvider(registryOCIHandler, apiHandler)
	sender := usage.ProvideMediator(ctx, config, spaceStore, usageMetricStore)
	routerRouter := router2.ProvideRouter(ctx, config, authenticator, repoController, reposettingsController, executionController, logsController, spaceController, pipelineController, secretController, triggerController, connectorController, templateController, pluginController, pullreqController, webhookController, githookController, gitInterface, serviceaccountController, controller, principalController, usergroupController, checkController, systemController, uploadController, keywordsearchController, infraproviderController, gitspaceController, migrateController, aiagentController, capabilitiesController, provider, openapiService, appRouter, sender)
	serverServer := server2.ProvideServer(config, routerRouter)
	publickeyService := publickey.ProvidePublicKey(publicKeyStore, principalInfoCache)
	sshServer := ssh.ProvideServer(config, publickeyService, repoController)
	executionManager := manager.ProvideExecutionManager(config, executionStore, pipelineStore, provider, streamer, fileService, converterService, logStore, logStream, checkStore, repoStore, schedulerScheduler, secretStore, stageStore, stepStore, principalStore, publicaccessService, reporter3)
	client := manager.ProvideExecutionClient(executionManager, provider, config)
	resolverManager := resolver.ProvideResolver(config, pluginStore, templateStore, executionStore, repoStore)
	runtimeRunner, err := runner.ProvideExecutionRunner(config, client, resolverManager)
	if err != nil {
		return nil, err
	}
	poller := runner.ProvideExecutionPoller(runtimeRunner, client)
	triggerConfig := server.ProvideTriggerConfig(config)
	triggerService, err := trigger2.ProvideService(ctx, triggerConfig, triggerStore, commitService, pullReqStore, repoStore, pipelineStore, triggererTriggerer, readerFactory, eventsReaderFactory)
	if err != nil {
		return nil, err
	}
	systemService := system2.ProvideService(settingsService)
	collector, err := metric.ProvideCollector(config, principalStore, repoStore, pipelineStore, executionStore, jobScheduler, executor, gitspaceConfigStore, systemService, registryRepository, artifactRepository)
	if err != nil {
		return nil, err
	}
	sizeCalculator, err := repo2.ProvideCalculator(config, gitInterface, repoStore, jobScheduler, executor)
	if err != nil {
		return nil, err
	}
	readerFactory2, err := events2.ProvideReaderFactory(eventsSystem)
	if err != nil {
		return nil, err
	}
	repoService, err := repo2.ProvideService(ctx, config, reporter, readerFactory2, repoStore, provider, gitInterface, lockerLocker)
	if err != nil {
		return nil, err
	}
	cleanupConfig := server.ProvideCleanupConfig(config)
	cleanupService, err := cleanup.ProvideService(cleanupConfig, jobScheduler, executor, webhookExecutionStore, tokenStore, repoStore, repoController)
	if err != nil {
		return nil, err
	}
	mailerMailer := mailer.ProvideMailClient(config)
	notificationClient := notification.ProvideMailClient(mailerMailer)
	notificationConfig := server.ProvideNotificationConfig(config)
	notificationService, err := notification.ProvideNotificationService(ctx, notificationClient, notificationConfig, eventsReaderFactory, pullReqStore, repoStore, principalInfoView, principalInfoCache, pullReqReviewerStore, pullReqActivityStore, spacePathStore, provider)
	if err != nil {
		return nil, err
	}
	keywordsearchConfig := server.ProvideKeywordSearchConfig(config)
	keywordsearchService, err := keywordsearch.ProvideService(ctx, keywordsearchConfig, readerFactory, readerFactory2, repoStore, indexer)
	if err != nil {
		return nil, err
	}
	gitspaceeventConfig := server.ProvideGitspaceEventConfig(config)
	readerFactory3, err := events3.ProvideReaderFactory(eventsSystem)
	if err != nil {
		return nil, err
	}
	gitspaceeventService, err := gitspaceevent.ProvideService(ctx, gitspaceeventConfig, readerFactory3, gitspaceEventStore)
	if err != nil {
		return nil, err
	}
	readerFactory4, err := events4.ProvideReaderFactory(eventsSystem)
	if err != nil {
		return nil, err
	}
	gitspaceinfraeventService, err := gitspaceinfraevent.ProvideService(ctx, gitspaceeventConfig, readerFactory4, orchestratorOrchestrator, gitspaceService, eventsReporter)
	if err != nil {
		return nil, err
	}
	gitspaceServices := services.ProvideGitspaceServices(gitspaceeventService, infraproviderService, gitspaceService, gitspaceinfraeventService)
	consumer, err := instrument.ProvideGitConsumer(ctx, config, readerFactory, repoStore, principalInfoCache, instrumentService)
	if err != nil {
		return nil, err
	}
	repositoryCount, err := instrument.ProvideRepositoryCount(ctx, config, instrumentService, repoStore, jobScheduler, executor)
	if err != nil {
		return nil, err
	}
	servicesServices := services.ProvideServices(webhookService, pullreqService, triggerService, jobScheduler, collector, sizeCalculator, repoService, cleanupService, notificationService, keywordsearchService, gitspaceServices, instrumentService, consumer, repositoryCount)
	serverSystem := server.NewSystem(bootstrapBootstrap, serverServer, sshServer, poller, resolverManager, servicesServices)
	return serverSystem, nil
}
