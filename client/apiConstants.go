package client

const CD_SUGGEST_NAME_API = "/orchestrator/app/pipeline/suggest/cd/"

const Environment_FOR_ENV_ID = "/orchestrator/env"

const WORKFLOWS_FOR_APP = "/orchestrator/app/app-wf/"

const WORKFLOWS_FOR_ENV = "/orchestrator/env/"

const CD_PIPELINE = "orchestrator/app/cd-pipeline/"

const CD_PIPELINES_ENV = "orchestrator/env/"

const STRATEGIES_FOR_APP = "orchestrator/app/cd-pipeline/strategies/"

const GET_ALL_DEPLOYMENT_TEMPLATE = "orchestrator/chartref/autocomplete/"

const GET_A_DEPLOYMENT_TEMPLATE = "orchestrator/app/env"

const GET_CONFIG_MAP = "orchestrator/config/environment/cm"

const GET_CONFIG_SECRET = "orchestrator/config/environment/cs"

const GET_CONFIG_SECRET_EDIT = "orchestrator/config/environment/cs/edit/"

const CREATE_CD_PIPELINE = "orchestrator/app/cd-pipeline"

const USER_AUTH = "orchestrator/devtron/auth/verify"

const GET_APPS = "orchestrator/app/list/v2"

const CHANGE_TYPE = "orchestrator/app/cd-pipeline/patch/deployment/type"

const TRIGGER_DEPLOY = "orchestrator/app/cd-pipeline/patch/deployment/trigger"

const ADD_KUSTOMIZE_DATA = "orchestrator/app/%v/upload/kustomize/%v" //"/orchestrator/app/{appId}/upload/kustomize/{envId}"

const CI_PIPELINES = "orchestrator/app/ci-pipeline/"

const PATCH_CI_PIPELINES = "orchestrator/app/ci-pipeline/patch"

const PATCH_CD_PIPELINES = "orchestrator/app/cd-pipeline/patch"

const ALL_APP_LIST = "orchestrator/app/allApps"

const APP_LIST_AUTOCOMPLETE = "orchestrator/app/app-listing/autocomplete"

const ADD_ENVIRONMENT = "orchestrator/env"

const MIGRATE_CHART_STORE_APP = "orchestrator/app-store/installed-app/migrate"

const TRIGGER_CHART_STORE_APP = "orchestrator/app-store/installed-app/trigger"
