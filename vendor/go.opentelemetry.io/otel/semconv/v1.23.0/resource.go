// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated from semantic convention specification. DO NOT EDIT.

package semconv // import "go.opentelemetry.io/otel/semconv/v1.23.0"

import "go.opentelemetry.io/otel/attribute"

// The Android platform on which the Android application is running.
const (
	// AndroidOSAPILevelKey is the attribute Key conforming to the
	// "android.os.api_level" semantic conventions. It represents the uniquely
	// identifies the framework API revision offered by a version
	// (`os.version`) of the android operating system. More information can be
	// found
	// [here](https://developer.android.com/guide/topics/manifest/uses-sdk-element#APILevels).
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '33', '32'
	AndroidOSAPILevelKey = attribute.Key("android.os.api_level")
)

// AndroidOSAPILevel returns an attribute KeyValue conforming to the
// "android.os.api_level" semantic conventions. It represents the uniquely
// identifies the framework API revision offered by a version (`os.version`) of
// the android operating system. More information can be found
// [here](https://developer.android.com/guide/topics/manifest/uses-sdk-element#APILevels).
func AndroidOSAPILevel(val string) attribute.KeyValue {
	return AndroidOSAPILevelKey.String(val)
}

// The web browser in which the application represented by the resource is
// running. The `browser.*` attributes MUST be used only for resources that
// represent applications running in a web browser (regardless of whether
// running on a mobile or desktop device).
const (
	// BrowserBrandsKey is the attribute Key conforming to the "browser.brands"
	// semantic conventions. It represents the array of brand name and version
	// separated by a space
	//
	// Type: string[]
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: ' Not A;Brand 99', 'Chromium 99', 'Chrome 99'
	// Note: This value is intended to be taken from the [UA client hints
	// API](https://wicg.github.io/ua-client-hints/#interface)
	// (`navigator.userAgentData.brands`).
	BrowserBrandsKey = attribute.Key("browser.brands")

	// BrowserLanguageKey is the attribute Key conforming to the
	// "browser.language" semantic conventions. It represents the preferred
	// language of the user using the browser
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'en', 'en-US', 'fr', 'fr-FR'
	// Note: This value is intended to be taken from the Navigator API
	// `navigator.language`.
	BrowserLanguageKey = attribute.Key("browser.language")

	// BrowserMobileKey is the attribute Key conforming to the "browser.mobile"
	// semantic conventions. It represents a boolean that is true if the
	// browser is running on a mobile device
	//
	// Type: boolean
	// RequirementLevel: Optional
	// Stability: experimental
	// Note: This value is intended to be taken from the [UA client hints
	// API](https://wicg.github.io/ua-client-hints/#interface)
	// (`navigator.userAgentData.mobile`). If unavailable, this attribute
	// SHOULD be left unset.
	BrowserMobileKey = attribute.Key("browser.mobile")

	// BrowserPlatformKey is the attribute Key conforming to the
	// "browser.platform" semantic conventions. It represents the platform on
	// which the browser is running
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'Windows', 'macOS', 'Android'
	// Note: This value is intended to be taken from the [UA client hints
	// API](https://wicg.github.io/ua-client-hints/#interface)
	// (`navigator.userAgentData.platform`). If unavailable, the legacy
	// `navigator.platform` API SHOULD NOT be used instead and this attribute
	// SHOULD be left unset in order for the values to be consistent.
	// The list of possible values is defined in the [W3C User-Agent Client
	// Hints
	// specification](https://wicg.github.io/ua-client-hints/#sec-ch-ua-platform).
	// Note that some (but not all) of these values can overlap with values in
	// the [`os.type` and `os.name` attributes](./os.md). However, for
	// consistency, the values in the `browser.platform` attribute should
	// capture the exact value that the user agent provides.
	BrowserPlatformKey = attribute.Key("browser.platform")
)

// BrowserBrands returns an attribute KeyValue conforming to the
// "browser.brands" semantic conventions. It represents the array of brand name
// and version separated by a space
func BrowserBrands(val ...string) attribute.KeyValue {
	return BrowserBrandsKey.StringSlice(val)
}

// BrowserLanguage returns an attribute KeyValue conforming to the
// "browser.language" semantic conventions. It represents the preferred
// language of the user using the browser
func BrowserLanguage(val string) attribute.KeyValue {
	return BrowserLanguageKey.String(val)
}

// BrowserMobile returns an attribute KeyValue conforming to the
// "browser.mobile" semantic conventions. It represents a boolean that is true
// if the browser is running on a mobile device
func BrowserMobile(val bool) attribute.KeyValue {
	return BrowserMobileKey.Bool(val)
}

// BrowserPlatform returns an attribute KeyValue conforming to the
// "browser.platform" semantic conventions. It represents the platform on which
// the browser is running
func BrowserPlatform(val string) attribute.KeyValue {
	return BrowserPlatformKey.String(val)
}

// Resources used by AWS Elastic Container Service (ECS).
const (
	// AWSECSClusterARNKey is the attribute Key conforming to the
	// "aws.ecs.cluster.arn" semantic conventions. It represents the ARN of an
	// [ECS
	// cluster](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/clusters.html).
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'arn:aws:ecs:us-west-2:123456789123:cluster/my-cluster'
	AWSECSClusterARNKey = attribute.Key("aws.ecs.cluster.arn")

	// AWSECSContainerARNKey is the attribute Key conforming to the
	// "aws.ecs.container.arn" semantic conventions. It represents the Amazon
	// Resource Name (ARN) of an [ECS container
	// instance](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ECS_instances.html).
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples:
	// 'arn:aws:ecs:us-west-1:123456789123:container/32624152-9086-4f0e-acae-1a75b14fe4d9'
	AWSECSContainerARNKey = attribute.Key("aws.ecs.container.arn")

	// AWSECSLaunchtypeKey is the attribute Key conforming to the
	// "aws.ecs.launchtype" semantic conventions. It represents the [launch
	// type](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/launch_types.html)
	// for an ECS task.
	//
	// Type: Enum
	// RequirementLevel: Optional
	// Stability: experimental
	AWSECSLaunchtypeKey = attribute.Key("aws.ecs.launchtype")

	// AWSECSTaskARNKey is the attribute Key conforming to the
	// "aws.ecs.task.arn" semantic conventions. It represents the ARN of an
	// [ECS task
	// definition](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definitions.html).
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples:
	// 'arn:aws:ecs:us-west-1:123456789123:task/10838bed-421f-43ef-870a-f43feacbbb5b'
	AWSECSTaskARNKey = attribute.Key("aws.ecs.task.arn")

	// AWSECSTaskFamilyKey is the attribute Key conforming to the
	// "aws.ecs.task.family" semantic conventions. It represents the task
	// definition family this task definition is a member of.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'opentelemetry-family'
	AWSECSTaskFamilyKey = attribute.Key("aws.ecs.task.family")

	// AWSECSTaskRevisionKey is the attribute Key conforming to the
	// "aws.ecs.task.revision" semantic conventions. It represents the revision
	// for this task definition.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '8', '26'
	AWSECSTaskRevisionKey = attribute.Key("aws.ecs.task.revision")
)

var (
	// ec2
	AWSECSLaunchtypeEC2 = AWSECSLaunchtypeKey.String("ec2")
	// fargate
	AWSECSLaunchtypeFargate = AWSECSLaunchtypeKey.String("fargate")
)

// AWSECSClusterARN returns an attribute KeyValue conforming to the
// "aws.ecs.cluster.arn" semantic conventions. It represents the ARN of an [ECS
// cluster](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/clusters.html).
func AWSECSClusterARN(val string) attribute.KeyValue {
	return AWSECSClusterARNKey.String(val)
}

// AWSECSContainerARN returns an attribute KeyValue conforming to the
// "aws.ecs.container.arn" semantic conventions. It represents the Amazon
// Resource Name (ARN) of an [ECS container
// instance](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ECS_instances.html).
func AWSECSContainerARN(val string) attribute.KeyValue {
	return AWSECSContainerARNKey.String(val)
}

// AWSECSTaskARN returns an attribute KeyValue conforming to the
// "aws.ecs.task.arn" semantic conventions. It represents the ARN of an [ECS
// task
// definition](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definitions.html).
func AWSECSTaskARN(val string) attribute.KeyValue {
	return AWSECSTaskARNKey.String(val)
}

// AWSECSTaskFamily returns an attribute KeyValue conforming to the
// "aws.ecs.task.family" semantic conventions. It represents the task
// definition family this task definition is a member of.
func AWSECSTaskFamily(val string) attribute.KeyValue {
	return AWSECSTaskFamilyKey.String(val)
}

// AWSECSTaskRevision returns an attribute KeyValue conforming to the
// "aws.ecs.task.revision" semantic conventions. It represents the revision for
// this task definition.
func AWSECSTaskRevision(val string) attribute.KeyValue {
	return AWSECSTaskRevisionKey.String(val)
}

// Resources used by AWS Elastic Kubernetes Service (EKS).
const (
	// AWSEKSClusterARNKey is the attribute Key conforming to the
	// "aws.eks.cluster.arn" semantic conventions. It represents the ARN of an
	// EKS cluster.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'arn:aws:ecs:us-west-2:123456789123:cluster/my-cluster'
	AWSEKSClusterARNKey = attribute.Key("aws.eks.cluster.arn")
)

// AWSEKSClusterARN returns an attribute KeyValue conforming to the
// "aws.eks.cluster.arn" semantic conventions. It represents the ARN of an EKS
// cluster.
func AWSEKSClusterARN(val string) attribute.KeyValue {
	return AWSEKSClusterARNKey.String(val)
}

// Resources specific to Amazon Web Services.
const (
	// AWSLogGroupARNsKey is the attribute Key conforming to the
	// "aws.log.group.arns" semantic conventions. It represents the Amazon
	// Resource Name(s) (ARN) of the AWS log group(s).
	//
	// Type: string[]
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples:
	// 'arn:aws:logs:us-west-1:123456789012:log-group:/aws/my/group:*'
	// Note: See the [log group ARN format
	// documentation](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/iam-access-control-overview-cwl.html#CWL_ARN_Format).
	AWSLogGroupARNsKey = attribute.Key("aws.log.group.arns")

	// AWSLogGroupNamesKey is the attribute Key conforming to the
	// "aws.log.group.names" semantic conventions. It represents the name(s) of
	// the AWS log group(s) an application is writing to.
	//
	// Type: string[]
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '/aws/lambda/my-function', 'opentelemetry-service'
	// Note: Multiple log groups must be supported for cases like
	// multi-container applications, where a single application has sidecar
	// containers, and each write to their own log group.
	AWSLogGroupNamesKey = attribute.Key("aws.log.group.names")

	// AWSLogStreamARNsKey is the attribute Key conforming to the
	// "aws.log.stream.arns" semantic conventions. It represents the ARN(s) of
	// the AWS log stream(s).
	//
	// Type: string[]
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples:
	// 'arn:aws:logs:us-west-1:123456789012:log-group:/aws/my/group:log-stream:logs/main/10838bed-421f-43ef-870a-f43feacbbb5b'
	// Note: See the [log stream ARN format
	// documentation](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/iam-access-control-overview-cwl.html#CWL_ARN_Format).
	// One log group can contain several log streams, so these ARNs necessarily
	// identify both a log group and a log stream.
	AWSLogStreamARNsKey = attribute.Key("aws.log.stream.arns")

	// AWSLogStreamNamesKey is the attribute Key conforming to the
	// "aws.log.stream.names" semantic conventions. It represents the name(s)
	// of the AWS log stream(s) an application is writing to.
	//
	// Type: string[]
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'logs/main/10838bed-421f-43ef-870a-f43feacbbb5b'
	AWSLogStreamNamesKey = attribute.Key("aws.log.stream.names")
)

// AWSLogGroupARNs returns an attribute KeyValue conforming to the
// "aws.log.group.arns" semantic conventions. It represents the Amazon Resource
// Name(s) (ARN) of the AWS log group(s).
func AWSLogGroupARNs(val ...string) attribute.KeyValue {
	return AWSLogGroupARNsKey.StringSlice(val)
}

// AWSLogGroupNames returns an attribute KeyValue conforming to the
// "aws.log.group.names" semantic conventions. It represents the name(s) of the
// AWS log group(s) an application is writing to.
func AWSLogGroupNames(val ...string) attribute.KeyValue {
	return AWSLogGroupNamesKey.StringSlice(val)
}

// AWSLogStreamARNs returns an attribute KeyValue conforming to the
// "aws.log.stream.arns" semantic conventions. It represents the ARN(s) of the
// AWS log stream(s).
func AWSLogStreamARNs(val ...string) attribute.KeyValue {
	return AWSLogStreamARNsKey.StringSlice(val)
}

// AWSLogStreamNames returns an attribute KeyValue conforming to the
// "aws.log.stream.names" semantic conventions. It represents the name(s) of
// the AWS log stream(s) an application is writing to.
func AWSLogStreamNames(val ...string) attribute.KeyValue {
	return AWSLogStreamNamesKey.StringSlice(val)
}

// Resource used by Google Cloud Run.
const (
	// GCPCloudRunJobExecutionKey is the attribute Key conforming to the
	// "gcp.cloud_run.job.execution" semantic conventions. It represents the
	// name of the Cloud Run
	// [execution](https://cloud.google.com/run/docs/managing/job-executions)
	// being run for the Job, as set by the
	// [`CLOUD_RUN_EXECUTION`](https://cloud.google.com/run/docs/container-contract#jobs-env-vars)
	// environment variable.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'job-name-xxxx', 'sample-job-mdw84'
	GCPCloudRunJobExecutionKey = attribute.Key("gcp.cloud_run.job.execution")

	// GCPCloudRunJobTaskIndexKey is the attribute Key conforming to the
	// "gcp.cloud_run.job.task_index" semantic conventions. It represents the
	// index for a task within an execution as provided by the
	// [`CLOUD_RUN_TASK_INDEX`](https://cloud.google.com/run/docs/container-contract#jobs-env-vars)
	// environment variable.
	//
	// Type: int
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 0, 1
	GCPCloudRunJobTaskIndexKey = attribute.Key("gcp.cloud_run.job.task_index")
)

// GCPCloudRunJobExecution returns an attribute KeyValue conforming to the
// "gcp.cloud_run.job.execution" semantic conventions. It represents the name
// of the Cloud Run
// [execution](https://cloud.google.com/run/docs/managing/job-executions) being
// run for the Job, as set by the
// [`CLOUD_RUN_EXECUTION`](https://cloud.google.com/run/docs/container-contract#jobs-env-vars)
// environment variable.
func GCPCloudRunJobExecution(val string) attribute.KeyValue {
	return GCPCloudRunJobExecutionKey.String(val)
}

// GCPCloudRunJobTaskIndex returns an attribute KeyValue conforming to the
// "gcp.cloud_run.job.task_index" semantic conventions. It represents the index
// for a task within an execution as provided by the
// [`CLOUD_RUN_TASK_INDEX`](https://cloud.google.com/run/docs/container-contract#jobs-env-vars)
// environment variable.
func GCPCloudRunJobTaskIndex(val int) attribute.KeyValue {
	return GCPCloudRunJobTaskIndexKey.Int(val)
}

// Resources used by Google Compute Engine (GCE).
const (
	// GCPGceInstanceHostnameKey is the attribute Key conforming to the
	// "gcp.gce.instance.hostname" semantic conventions. It represents the
	// hostname of a GCE instance. This is the full value of the default or
	// [custom
	// hostname](https://cloud.google.com/compute/docs/instances/custom-hostname-vm).
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'my-host1234.example.com',
	// 'sample-vm.us-west1-b.c.my-project.internal'
	GCPGceInstanceHostnameKey = attribute.Key("gcp.gce.instance.hostname")

	// GCPGceInstanceNameKey is the attribute Key conforming to the
	// "gcp.gce.instance.name" semantic conventions. It represents the instance
	// name of a GCE instance. This is the value provided by `host.name`, the
	// visible name of the instance in the Cloud Console UI, and the prefix for
	// the default hostname of the instance as defined by the [default internal
	// DNS
	// name](https://cloud.google.com/compute/docs/internal-dns#instance-fully-qualified-domain-names).
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'instance-1', 'my-vm-name'
	GCPGceInstanceNameKey = attribute.Key("gcp.gce.instance.name")
)

// GCPGceInstanceHostname returns an attribute KeyValue conforming to the
// "gcp.gce.instance.hostname" semantic conventions. It represents the hostname
// of a GCE instance. This is the full value of the default or [custom
// hostname](https://cloud.google.com/compute/docs/instances/custom-hostname-vm).
func GCPGceInstanceHostname(val string) attribute.KeyValue {
	return GCPGceInstanceHostnameKey.String(val)
}

// GCPGceInstanceName returns an attribute KeyValue conforming to the
// "gcp.gce.instance.name" semantic conventions. It represents the instance
// name of a GCE instance. This is the value provided by `host.name`, the
// visible name of the instance in the Cloud Console UI, and the prefix for the
// default hostname of the instance as defined by the [default internal DNS
// name](https://cloud.google.com/compute/docs/internal-dns#instance-fully-qualified-domain-names).
func GCPGceInstanceName(val string) attribute.KeyValue {
	return GCPGceInstanceNameKey.String(val)
}

// Heroku dyno metadata
const (
	// HerokuAppIDKey is the attribute Key conforming to the "heroku.app.id"
	// semantic conventions. It represents the unique identifier for the
	// application
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '2daa2797-e42b-4624-9322-ec3f968df4da'
	HerokuAppIDKey = attribute.Key("heroku.app.id")

	// HerokuReleaseCommitKey is the attribute Key conforming to the
	// "heroku.release.commit" semantic conventions. It represents the commit
	// hash for the current release
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'e6134959463efd8966b20e75b913cafe3f5ec'
	HerokuReleaseCommitKey = attribute.Key("heroku.release.commit")

	// HerokuReleaseCreationTimestampKey is the attribute Key conforming to the
	// "heroku.release.creation_timestamp" semantic conventions. It represents
	// the time and date the release was created
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '2022-10-23T18:00:42Z'
	HerokuReleaseCreationTimestampKey = attribute.Key("heroku.release.creation_timestamp")
)

// HerokuAppID returns an attribute KeyValue conforming to the
// "heroku.app.id" semantic conventions. It represents the unique identifier
// for the application
func HerokuAppID(val string) attribute.KeyValue {
	return HerokuAppIDKey.String(val)
}

// HerokuReleaseCommit returns an attribute KeyValue conforming to the
// "heroku.release.commit" semantic conventions. It represents the commit hash
// for the current release
func HerokuReleaseCommit(val string) attribute.KeyValue {
	return HerokuReleaseCommitKey.String(val)
}

// HerokuReleaseCreationTimestamp returns an attribute KeyValue conforming
// to the "heroku.release.creation_timestamp" semantic conventions. It
// represents the time and date the release was created
func HerokuReleaseCreationTimestamp(val string) attribute.KeyValue {
	return HerokuReleaseCreationTimestampKey.String(val)
}

// The software deployment.
const (
	// DeploymentEnvironmentKey is the attribute Key conforming to the
	// "deployment.environment" semantic conventions. It represents the name of
	// the [deployment
	// environment](https://wikipedia.org/wiki/Deployment_environment) (aka
	// deployment tier).
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'staging', 'production'
	DeploymentEnvironmentKey = attribute.Key("deployment.environment")
)

// DeploymentEnvironment returns an attribute KeyValue conforming to the
// "deployment.environment" semantic conventions. It represents the name of the
// [deployment environment](https://wikipedia.org/wiki/Deployment_environment)
// (aka deployment tier).
func DeploymentEnvironment(val string) attribute.KeyValue {
	return DeploymentEnvironmentKey.String(val)
}

// The device on which the process represented by this resource is running.
const (
	// DeviceIDKey is the attribute Key conforming to the "device.id" semantic
	// conventions. It represents a unique identifier representing the device
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '2ab2916d-a51f-4ac8-80ee-45ac31a28092'
	// Note: The device identifier MUST only be defined using the values
	// outlined below. This value is not an advertising identifier and MUST NOT
	// be used as such. On iOS (Swift or Objective-C), this value MUST be equal
	// to the [vendor
	// identifier](https://developer.apple.com/documentation/uikit/uidevice/1620059-identifierforvendor).
	// On Android (Java or Kotlin), this value MUST be equal to the Firebase
	// Installation ID or a globally unique UUID which is persisted across
	// sessions in your application. More information can be found
	// [here](https://developer.android.com/training/articles/user-data-ids) on
	// best practices and exact implementation details. Caution should be taken
	// when storing personal data or anything which can identify a user. GDPR
	// and data protection laws may apply, ensure you do your own due
	// diligence.
	DeviceIDKey = attribute.Key("device.id")

	// DeviceManufacturerKey is the attribute Key conforming to the
	// "device.manufacturer" semantic conventions. It represents the name of
	// the device manufacturer
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'Apple', 'Samsung'
	// Note: The Android OS provides this field via
	// [Build](https://developer.android.com/reference/android/os/Build#MANUFACTURER).
	// iOS apps SHOULD hardcode the value `Apple`.
	DeviceManufacturerKey = attribute.Key("device.manufacturer")

	// DeviceModelIdentifierKey is the attribute Key conforming to the
	// "device.model.identifier" semantic conventions. It represents the model
	// identifier for the device
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'iPhone3,4', 'SM-G920F'
	// Note: It's recommended this value represents a machine readable version
	// of the model identifier rather than the market or consumer-friendly name
	// of the device.
	DeviceModelIdentifierKey = attribute.Key("device.model.identifier")

	// DeviceModelNameKey is the attribute Key conforming to the
	// "device.model.name" semantic conventions. It represents the marketing
	// name for the device model
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'iPhone 6s Plus', 'Samsung Galaxy S6'
	// Note: It's recommended this value represents a human readable version of
	// the device model rather than a machine readable alternative.
	DeviceModelNameKey = attribute.Key("device.model.name")
)

// DeviceID returns an attribute KeyValue conforming to the "device.id"
// semantic conventions. It represents a unique identifier representing the
// device
func DeviceID(val string) attribute.KeyValue {
	return DeviceIDKey.String(val)
}

// DeviceManufacturer returns an attribute KeyValue conforming to the
// "device.manufacturer" semantic conventions. It represents the name of the
// device manufacturer
func DeviceManufacturer(val string) attribute.KeyValue {
	return DeviceManufacturerKey.String(val)
}

// DeviceModelIdentifier returns an attribute KeyValue conforming to the
// "device.model.identifier" semantic conventions. It represents the model
// identifier for the device
func DeviceModelIdentifier(val string) attribute.KeyValue {
	return DeviceModelIdentifierKey.String(val)
}

// DeviceModelName returns an attribute KeyValue conforming to the
// "device.model.name" semantic conventions. It represents the marketing name
// for the device model
func DeviceModelName(val string) attribute.KeyValue {
	return DeviceModelNameKey.String(val)
}

// A serverless instance.
const (
	// FaaSInstanceKey is the attribute Key conforming to the "faas.instance"
	// semantic conventions. It represents the execution environment ID as a
	// string, that will be potentially reused for other invocations to the
	// same function/function version.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '2021/06/28/[$LATEST]2f399eb14537447da05ab2a2e39309de'
	// Note: * **AWS Lambda:** Use the (full) log stream name.
	FaaSInstanceKey = attribute.Key("faas.instance")

	// FaaSMaxMemoryKey is the attribute Key conforming to the
	// "faas.max_memory" semantic conventions. It represents the amount of
	// memory available to the serverless function converted to Bytes.
	//
	// Type: int
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 134217728
	// Note: It's recommended to set this attribute since e.g. too little
	// memory can easily stop a Java AWS Lambda function from working
	// correctly. On AWS Lambda, the environment variable
	// `AWS_LAMBDA_FUNCTION_MEMORY_SIZE` provides this information (which must
	// be multiplied by 1,048,576).
	FaaSMaxMemoryKey = attribute.Key("faas.max_memory")

	// FaaSNameKey is the attribute Key conforming to the "faas.name" semantic
	// conventions. It represents the name of the single function that this
	// runtime instance executes.
	//
	// Type: string
	// RequirementLevel: Required
	// Stability: experimental
	// Examples: 'my-function', 'myazurefunctionapp/some-function-name'
	// Note: This is the name of the function as configured/deployed on the
	// FaaS
	// platform and is usually different from the name of the callback
	// function (which may be stored in the
	// [`code.namespace`/`code.function`](/docs/general/attributes.md#source-code-attributes)
	// span attributes).
	//
	// For some cloud providers, the above definition is ambiguous. The
	// following
	// definition of function name MUST be used for this attribute
	// (and consequently the span name) for the listed cloud
	// providers/products:
	//
	// * **Azure:**  The full name `<FUNCAPP>/<FUNC>`, i.e., function app name
	//   followed by a forward slash followed by the function name (this form
	//   can also be seen in the resource JSON for the function).
	//   This means that a span attribute MUST be used, as an Azure function
	//   app can host multiple functions that would usually share
	//   a TracerProvider (see also the `cloud.resource_id` attribute).
	FaaSNameKey = attribute.Key("faas.name")

	// FaaSVersionKey is the attribute Key conforming to the "faas.version"
	// semantic conventions. It represents the immutable version of the
	// function being executed.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '26', 'pinkfroid-00002'
	// Note: Depending on the cloud provider and platform, use:
	//
	// * **AWS Lambda:** The [function
	// version](https://docs.aws.amazon.com/lambda/latest/dg/configuration-versions.html)
	//   (an integer represented as a decimal string).
	// * **Google Cloud Run (Services):** The
	// [revision](https://cloud.google.com/run/docs/managing/revisions)
	//   (i.e., the function name plus the revision suffix).
	// * **Google Cloud Functions:** The value of the
	//   [`K_REVISION` environment
	// variable](https://cloud.google.com/functions/docs/env-var#runtime_environment_variables_set_automatically).
	// * **Azure Functions:** Not applicable. Do not set this attribute.
	FaaSVersionKey = attribute.Key("faas.version")
)

// FaaSInstance returns an attribute KeyValue conforming to the
// "faas.instance" semantic conventions. It represents the execution
// environment ID as a string, that will be potentially reused for other
// invocations to the same function/function version.
func FaaSInstance(val string) attribute.KeyValue {
	return FaaSInstanceKey.String(val)
}

// FaaSMaxMemory returns an attribute KeyValue conforming to the
// "faas.max_memory" semantic conventions. It represents the amount of memory
// available to the serverless function converted to Bytes.
func FaaSMaxMemory(val int) attribute.KeyValue {
	return FaaSMaxMemoryKey.Int(val)
}

// FaaSName returns an attribute KeyValue conforming to the "faas.name"
// semantic conventions. It represents the name of the single function that
// this runtime instance executes.
func FaaSName(val string) attribute.KeyValue {
	return FaaSNameKey.String(val)
}

// FaaSVersion returns an attribute KeyValue conforming to the
// "faas.version" semantic conventions. It represents the immutable version of
// the function being executed.
func FaaSVersion(val string) attribute.KeyValue {
	return FaaSVersionKey.String(val)
}

// A host is defined as a computing instance. For example, physical servers,
// virtual machines, switches or disk array.
const (
	// HostArchKey is the attribute Key conforming to the "host.arch" semantic
	// conventions. It represents the CPU architecture the host system is
	// running on.
	//
	// Type: Enum
	// RequirementLevel: Optional
	// Stability: experimental
	HostArchKey = attribute.Key("host.arch")

	// HostIDKey is the attribute Key conforming to the "host.id" semantic
	// conventions. It represents the unique host ID. For Cloud, this must be
	// the instance_id assigned by the cloud provider. For non-containerized
	// systems, this should be the `machine-id`. See the table below for the
	// sources to use to determine the `machine-id` based on operating system.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'fdbf79e8af94cb7f9e8df36789187052'
	HostIDKey = attribute.Key("host.id")

	// HostImageIDKey is the attribute Key conforming to the "host.image.id"
	// semantic conventions. It represents the vM image ID or host OS image ID.
	// For Cloud, this value is from the provider.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'ami-07b06b442921831e5'
	HostImageIDKey = attribute.Key("host.image.id")

	// HostImageNameKey is the attribute Key conforming to the
	// "host.image.name" semantic conventions. It represents the name of the VM
	// image or OS install the host was instantiated from.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'infra-ami-eks-worker-node-7d4ec78312', 'CentOS-8-x86_64-1905'
	HostImageNameKey = attribute.Key("host.image.name")

	// HostImageVersionKey is the attribute Key conforming to the
	// "host.image.version" semantic conventions. It represents the version
	// string of the VM image or host OS as defined in [Version
	// Attributes](README.md#version-attributes).
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '0.1'
	HostImageVersionKey = attribute.Key("host.image.version")

	// HostIPKey is the attribute Key conforming to the "host.ip" semantic
	// conventions. It represents the available IP addresses of the host,
	// excluding loopback interfaces.
	//
	// Type: string[]
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '192.168.1.140', 'fe80::abc2:4a28:737a:609e'
	// Note: IPv4 Addresses MUST be specified in dotted-quad notation. IPv6
	// addresses MUST be specified in the [RFC
	// 5952](https://www.rfc-editor.org/rfc/rfc5952.html) format.
	HostIPKey = attribute.Key("host.ip")

	// HostMacKey is the attribute Key conforming to the "host.mac" semantic
	// conventions. It represents the available MAC addresses of the host,
	// excluding loopback interfaces.
	//
	// Type: string[]
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'AC-DE-48-23-45-67', 'AC-DE-48-23-45-67-01-9F'
	// Note: MAC Addresses MUST be represented in [IEEE RA hexadecimal
	// form](https://standards.ieee.org/wp-content/uploads/import/documents/tutorials/eui.pdf):
	// as hyphen-separated octets in uppercase hexadecimal form from most to
	// least significant.
	HostMacKey = attribute.Key("host.mac")

	// HostNameKey is the attribute Key conforming to the "host.name" semantic
	// conventions. It represents the name of the host. On Unix systems, it may
	// contain what the hostname command returns, or the fully qualified
	// hostname, or another name specified by the user.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'opentelemetry-test'
	HostNameKey = attribute.Key("host.name")

	// HostTypeKey is the attribute Key conforming to the "host.type" semantic
	// conventions. It represents the type of host. For Cloud, this must be the
	// machine type.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'n1-standard-1'
	HostTypeKey = attribute.Key("host.type")
)

var (
	// AMD64
	HostArchAMD64 = HostArchKey.String("amd64")
	// ARM32
	HostArchARM32 = HostArchKey.String("arm32")
	// ARM64
	HostArchARM64 = HostArchKey.String("arm64")
	// Itanium
	HostArchIA64 = HostArchKey.String("ia64")
	// 32-bit PowerPC
	HostArchPPC32 = HostArchKey.String("ppc32")
	// 64-bit PowerPC
	HostArchPPC64 = HostArchKey.String("ppc64")
	// IBM z/Architecture
	HostArchS390x = HostArchKey.String("s390x")
	// 32-bit x86
	HostArchX86 = HostArchKey.String("x86")
)

// HostID returns an attribute KeyValue conforming to the "host.id" semantic
// conventions. It represents the unique host ID. For Cloud, this must be the
// instance_id assigned by the cloud provider. For non-containerized systems,
// this should be the `machine-id`. See the table below for the sources to use
// to determine the `machine-id` based on operating system.
func HostID(val string) attribute.KeyValue {
	return HostIDKey.String(val)
}

// HostImageID returns an attribute KeyValue conforming to the
// "host.image.id" semantic conventions. It represents the vM image ID or host
// OS image ID. For Cloud, this value is from the provider.
func HostImageID(val string) attribute.KeyValue {
	return HostImageIDKey.String(val)
}

// HostImageName returns an attribute KeyValue conforming to the
// "host.image.name" semantic conventions. It represents the name of the VM
// image or OS install the host was instantiated from.
func HostImageName(val string) attribute.KeyValue {
	return HostImageNameKey.String(val)
}

// HostImageVersion returns an attribute KeyValue conforming to the
// "host.image.version" semantic conventions. It represents the version string
// of the VM image or host OS as defined in [Version
// Attributes](README.md#version-attributes).
func HostImageVersion(val string) attribute.KeyValue {
	return HostImageVersionKey.String(val)
}

// HostIP returns an attribute KeyValue conforming to the "host.ip" semantic
// conventions. It represents the available IP addresses of the host, excluding
// loopback interfaces.
func HostIP(val ...string) attribute.KeyValue {
	return HostIPKey.StringSlice(val)
}

// HostMac returns an attribute KeyValue conforming to the "host.mac"
// semantic conventions. It represents the available MAC addresses of the host,
// excluding loopback interfaces.
func HostMac(val ...string) attribute.KeyValue {
	return HostMacKey.StringSlice(val)
}

// HostName returns an attribute KeyValue conforming to the "host.name"
// semantic conventions. It represents the name of the host. On Unix systems,
// it may contain what the hostname command returns, or the fully qualified
// hostname, or another name specified by the user.
func HostName(val string) attribute.KeyValue {
	return HostNameKey.String(val)
}

// HostType returns an attribute KeyValue conforming to the "host.type"
// semantic conventions. It represents the type of host. For Cloud, this must
// be the machine type.
func HostType(val string) attribute.KeyValue {
	return HostTypeKey.String(val)
}

// A host's CPU information
const (
	// HostCPUCacheL2SizeKey is the attribute Key conforming to the
	// "host.cpu.cache.l2.size" semantic conventions. It represents the amount
	// of level 2 memory cache available to the processor (in Bytes).
	//
	// Type: int
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 12288000
	HostCPUCacheL2SizeKey = attribute.Key("host.cpu.cache.l2.size")

	// HostCPUFamilyKey is the attribute Key conforming to the
	// "host.cpu.family" semantic conventions. It represents the numeric value
	// specifying the family or generation of the CPU.
	//
	// Type: int
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 6
	HostCPUFamilyKey = attribute.Key("host.cpu.family")

	// HostCPUModelIDKey is the attribute Key conforming to the
	// "host.cpu.model.id" semantic conventions. It represents the model
	// identifier. It provides more granular information about the CPU,
	// distinguishing it from other CPUs within the same family.
	//
	// Type: int
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 6
	HostCPUModelIDKey = attribute.Key("host.cpu.model.id")

	// HostCPUModelNameKey is the attribute Key conforming to the
	// "host.cpu.model.name" semantic conventions. It represents the model
	// designation of the processor.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '11th Gen Intel(R) Core(TM) i7-1185G7 @ 3.00GHz'
	HostCPUModelNameKey = attribute.Key("host.cpu.model.name")

	// HostCPUSteppingKey is the attribute Key conforming to the
	// "host.cpu.stepping" semantic conventions. It represents the stepping or
	// core revisions.
	//
	// Type: int
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 1
	HostCPUSteppingKey = attribute.Key("host.cpu.stepping")

	// HostCPUVendorIDKey is the attribute Key conforming to the
	// "host.cpu.vendor.id" semantic conventions. It represents the processor
	// manufacturer identifier. A maximum 12-character string.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'GenuineIntel'
	// Note: [CPUID](https://wiki.osdev.org/CPUID) command returns the vendor
	// ID string in EBX, EDX and ECX registers. Writing these to memory in this
	// order results in a 12-character string.
	HostCPUVendorIDKey = attribute.Key("host.cpu.vendor.id")
)

// HostCPUCacheL2Size returns an attribute KeyValue conforming to the
// "host.cpu.cache.l2.size" semantic conventions. It represents the amount of
// level 2 memory cache available to the processor (in Bytes).
func HostCPUCacheL2Size(val int) attribute.KeyValue {
	return HostCPUCacheL2SizeKey.Int(val)
}

// HostCPUFamily returns an attribute KeyValue conforming to the
// "host.cpu.family" semantic conventions. It represents the numeric value
// specifying the family or generation of the CPU.
func HostCPUFamily(val int) attribute.KeyValue {
	return HostCPUFamilyKey.Int(val)
}

// HostCPUModelID returns an attribute KeyValue conforming to the
// "host.cpu.model.id" semantic conventions. It represents the model
// identifier. It provides more granular information about the CPU,
// distinguishing it from other CPUs within the same family.
func HostCPUModelID(val int) attribute.KeyValue {
	return HostCPUModelIDKey.Int(val)
}

// HostCPUModelName returns an attribute KeyValue conforming to the
// "host.cpu.model.name" semantic conventions. It represents the model
// designation of the processor.
func HostCPUModelName(val string) attribute.KeyValue {
	return HostCPUModelNameKey.String(val)
}

// HostCPUStepping returns an attribute KeyValue conforming to the
// "host.cpu.stepping" semantic conventions. It represents the stepping or core
// revisions.
func HostCPUStepping(val int) attribute.KeyValue {
	return HostCPUSteppingKey.Int(val)
}

// HostCPUVendorID returns an attribute KeyValue conforming to the
// "host.cpu.vendor.id" semantic conventions. It represents the processor
// manufacturer identifier. A maximum 12-character string.
func HostCPUVendorID(val string) attribute.KeyValue {
	return HostCPUVendorIDKey.String(val)
}

// A Kubernetes Cluster.
const (
	// K8SClusterNameKey is the attribute Key conforming to the
	// "k8s.cluster.name" semantic conventions. It represents the name of the
	// cluster.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'opentelemetry-cluster'
	K8SClusterNameKey = attribute.Key("k8s.cluster.name")

	// K8SClusterUIDKey is the attribute Key conforming to the
	// "k8s.cluster.uid" semantic conventions. It represents a pseudo-ID for
	// the cluster, set to the UID of the `kube-system` namespace.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '218fc5a9-a5f1-4b54-aa05-46717d0ab26d'
	// Note: K8S doesn't have support for obtaining a cluster ID. If this is
	// ever
	// added, we will recommend collecting the `k8s.cluster.uid` through the
	// official APIs. In the meantime, we are able to use the `uid` of the
	// `kube-system` namespace as a proxy for cluster ID. Read on for the
	// rationale.
	//
	// Every object created in a K8S cluster is assigned a distinct UID. The
	// `kube-system` namespace is used by Kubernetes itself and will exist
	// for the lifetime of the cluster. Using the `uid` of the `kube-system`
	// namespace is a reasonable proxy for the K8S ClusterID as it will only
	// change if the cluster is rebuilt. Furthermore, Kubernetes UIDs are
	// UUIDs as standardized by
	// [ISO/IEC 9834-8 and ITU-T
	// X.667](https://www.itu.int/ITU-T/studygroups/com17/oid.html).
	// Which states:
	//
	// > If generated according to one of the mechanisms defined in Rec.
	//   ITU-T X.667 | ISO/IEC 9834-8, a UUID is either guaranteed to be
	//   different from all other UUIDs generated before 3603 A.D., or is
	//   extremely likely to be different (depending on the mechanism chosen).
	//
	// Therefore, UIDs between clusters should be extremely unlikely to
	// conflict.
	K8SClusterUIDKey = attribute.Key("k8s.cluster.uid")
)

// K8SClusterName returns an attribute KeyValue conforming to the
// "k8s.cluster.name" semantic conventions. It represents the name of the
// cluster.
func K8SClusterName(val string) attribute.KeyValue {
	return K8SClusterNameKey.String(val)
}

// K8SClusterUID returns an attribute KeyValue conforming to the
// "k8s.cluster.uid" semantic conventions. It represents a pseudo-ID for the
// cluster, set to the UID of the `kube-system` namespace.
func K8SClusterUID(val string) attribute.KeyValue {
	return K8SClusterUIDKey.String(val)
}

// A Kubernetes Node object.
const (
	// K8SNodeNameKey is the attribute Key conforming to the "k8s.node.name"
	// semantic conventions. It represents the name of the Node.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'node-1'
	K8SNodeNameKey = attribute.Key("k8s.node.name")

	// K8SNodeUIDKey is the attribute Key conforming to the "k8s.node.uid"
	// semantic conventions. It represents the UID of the Node.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '1eb3a0c6-0477-4080-a9cb-0cb7db65c6a2'
	K8SNodeUIDKey = attribute.Key("k8s.node.uid")
)

// K8SNodeName returns an attribute KeyValue conforming to the
// "k8s.node.name" semantic conventions. It represents the name of the Node.
func K8SNodeName(val string) attribute.KeyValue {
	return K8SNodeNameKey.String(val)
}

// K8SNodeUID returns an attribute KeyValue conforming to the "k8s.node.uid"
// semantic conventions. It represents the UID of the Node.
func K8SNodeUID(val string) attribute.KeyValue {
	return K8SNodeUIDKey.String(val)
}

// A Kubernetes Namespace.
const (
	// K8SNamespaceNameKey is the attribute Key conforming to the
	// "k8s.namespace.name" semantic conventions. It represents the name of the
	// namespace that the pod is running in.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'default'
	K8SNamespaceNameKey = attribute.Key("k8s.namespace.name")
)

// K8SNamespaceName returns an attribute KeyValue conforming to the
// "k8s.namespace.name" semantic conventions. It represents the name of the
// namespace that the pod is running in.
func K8SNamespaceName(val string) attribute.KeyValue {
	return K8SNamespaceNameKey.String(val)
}

// A Kubernetes Pod object.
const (
	// K8SPodNameKey is the attribute Key conforming to the "k8s.pod.name"
	// semantic conventions. It represents the name of the Pod.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'opentelemetry-pod-autoconf'
	K8SPodNameKey = attribute.Key("k8s.pod.name")

	// K8SPodUIDKey is the attribute Key conforming to the "k8s.pod.uid"
	// semantic conventions. It represents the UID of the Pod.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '275ecb36-5aa8-4c2a-9c47-d8bb681b9aff'
	K8SPodUIDKey = attribute.Key("k8s.pod.uid")
)

// K8SPodName returns an attribute KeyValue conforming to the "k8s.pod.name"
// semantic conventions. It represents the name of the Pod.
func K8SPodName(val string) attribute.KeyValue {
	return K8SPodNameKey.String(val)
}

// K8SPodUID returns an attribute KeyValue conforming to the "k8s.pod.uid"
// semantic conventions. It represents the UID of the Pod.
func K8SPodUID(val string) attribute.KeyValue {
	return K8SPodUIDKey.String(val)
}

// A container in a
// [PodTemplate](https://kubernetes.io/docs/concepts/workloads/pods/#pod-templates).
const (
	// K8SContainerNameKey is the attribute Key conforming to the
	// "k8s.container.name" semantic conventions. It represents the name of the
	// Container from Pod specification, must be unique within a Pod. Container
	// runtime usually uses different globally unique name (`container.name`).
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'redis'
	K8SContainerNameKey = attribute.Key("k8s.container.name")

	// K8SContainerRestartCountKey is the attribute Key conforming to the
	// "k8s.container.restart_count" semantic conventions. It represents the
	// number of times the container was restarted. This attribute can be used
	// to identify a particular container (running or stopped) within a
	// container spec.
	//
	// Type: int
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 0, 2
	K8SContainerRestartCountKey = attribute.Key("k8s.container.restart_count")
)

// K8SContainerName returns an attribute KeyValue conforming to the
// "k8s.container.name" semantic conventions. It represents the name of the
// Container from Pod specification, must be unique within a Pod. Container
// runtime usually uses different globally unique name (`container.name`).
func K8SContainerName(val string) attribute.KeyValue {
	return K8SContainerNameKey.String(val)
}

// K8SContainerRestartCount returns an attribute KeyValue conforming to the
// "k8s.container.restart_count" semantic conventions. It represents the number
// of times the container was restarted. This attribute can be used to identify
// a particular container (running or stopped) within a container spec.
func K8SContainerRestartCount(val int) attribute.KeyValue {
	return K8SContainerRestartCountKey.Int(val)
}

// A Kubernetes ReplicaSet object.
const (
	// K8SReplicaSetNameKey is the attribute Key conforming to the
	// "k8s.replicaset.name" semantic conventions. It represents the name of
	// the ReplicaSet.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'opentelemetry'
	K8SReplicaSetNameKey = attribute.Key("k8s.replicaset.name")

	// K8SReplicaSetUIDKey is the attribute Key conforming to the
	// "k8s.replicaset.uid" semantic conventions. It represents the UID of the
	// ReplicaSet.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '275ecb36-5aa8-4c2a-9c47-d8bb681b9aff'
	K8SReplicaSetUIDKey = attribute.Key("k8s.replicaset.uid")
)

// K8SReplicaSetName returns an attribute KeyValue conforming to the
// "k8s.replicaset.name" semantic conventions. It represents the name of the
// ReplicaSet.
func K8SReplicaSetName(val string) attribute.KeyValue {
	return K8SReplicaSetNameKey.String(val)
}

// K8SReplicaSetUID returns an attribute KeyValue conforming to the
// "k8s.replicaset.uid" semantic conventions. It represents the UID of the
// ReplicaSet.
func K8SReplicaSetUID(val string) attribute.KeyValue {
	return K8SReplicaSetUIDKey.String(val)
}

// A Kubernetes Deployment object.
const (
	// K8SDeploymentNameKey is the attribute Key conforming to the
	// "k8s.deployment.name" semantic conventions. It represents the name of
	// the Deployment.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'opentelemetry'
	K8SDeploymentNameKey = attribute.Key("k8s.deployment.name")

	// K8SDeploymentUIDKey is the attribute Key conforming to the
	// "k8s.deployment.uid" semantic conventions. It represents the UID of the
	// Deployment.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '275ecb36-5aa8-4c2a-9c47-d8bb681b9aff'
	K8SDeploymentUIDKey = attribute.Key("k8s.deployment.uid")
)

// K8SDeploymentName returns an attribute KeyValue conforming to the
// "k8s.deployment.name" semantic conventions. It represents the name of the
// Deployment.
func K8SDeploymentName(val string) attribute.KeyValue {
	return K8SDeploymentNameKey.String(val)
}

// K8SDeploymentUID returns an attribute KeyValue conforming to the
// "k8s.deployment.uid" semantic conventions. It represents the UID of the
// Deployment.
func K8SDeploymentUID(val string) attribute.KeyValue {
	return K8SDeploymentUIDKey.String(val)
}

// A Kubernetes StatefulSet object.
const (
	// K8SStatefulSetNameKey is the attribute Key conforming to the
	// "k8s.statefulset.name" semantic conventions. It represents the name of
	// the StatefulSet.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'opentelemetry'
	K8SStatefulSetNameKey = attribute.Key("k8s.statefulset.name")

	// K8SStatefulSetUIDKey is the attribute Key conforming to the
	// "k8s.statefulset.uid" semantic conventions. It represents the UID of the
	// StatefulSet.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '275ecb36-5aa8-4c2a-9c47-d8bb681b9aff'
	K8SStatefulSetUIDKey = attribute.Key("k8s.statefulset.uid")
)

// K8SStatefulSetName returns an attribute KeyValue conforming to the
// "k8s.statefulset.name" semantic conventions. It represents the name of the
// StatefulSet.
func K8SStatefulSetName(val string) attribute.KeyValue {
	return K8SStatefulSetNameKey.String(val)
}

// K8SStatefulSetUID returns an attribute KeyValue conforming to the
// "k8s.statefulset.uid" semantic conventions. It represents the UID of the
// StatefulSet.
func K8SStatefulSetUID(val string) attribute.KeyValue {
	return K8SStatefulSetUIDKey.String(val)
}

// A Kubernetes DaemonSet object.
const (
	// K8SDaemonSetNameKey is the attribute Key conforming to the
	// "k8s.daemonset.name" semantic conventions. It represents the name of the
	// DaemonSet.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'opentelemetry'
	K8SDaemonSetNameKey = attribute.Key("k8s.daemonset.name")

	// K8SDaemonSetUIDKey is the attribute Key conforming to the
	// "k8s.daemonset.uid" semantic conventions. It represents the UID of the
	// DaemonSet.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '275ecb36-5aa8-4c2a-9c47-d8bb681b9aff'
	K8SDaemonSetUIDKey = attribute.Key("k8s.daemonset.uid")
)

// K8SDaemonSetName returns an attribute KeyValue conforming to the
// "k8s.daemonset.name" semantic conventions. It represents the name of the
// DaemonSet.
func K8SDaemonSetName(val string) attribute.KeyValue {
	return K8SDaemonSetNameKey.String(val)
}

// K8SDaemonSetUID returns an attribute KeyValue conforming to the
// "k8s.daemonset.uid" semantic conventions. It represents the UID of the
// DaemonSet.
func K8SDaemonSetUID(val string) attribute.KeyValue {
	return K8SDaemonSetUIDKey.String(val)
}

// A Kubernetes Job object.
const (
	// K8SJobNameKey is the attribute Key conforming to the "k8s.job.name"
	// semantic conventions. It represents the name of the Job.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'opentelemetry'
	K8SJobNameKey = attribute.Key("k8s.job.name")

	// K8SJobUIDKey is the attribute Key conforming to the "k8s.job.uid"
	// semantic conventions. It represents the UID of the Job.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '275ecb36-5aa8-4c2a-9c47-d8bb681b9aff'
	K8SJobUIDKey = attribute.Key("k8s.job.uid")
)

// K8SJobName returns an attribute KeyValue conforming to the "k8s.job.name"
// semantic conventions. It represents the name of the Job.
func K8SJobName(val string) attribute.KeyValue {
	return K8SJobNameKey.String(val)
}

// K8SJobUID returns an attribute KeyValue conforming to the "k8s.job.uid"
// semantic conventions. It represents the UID of the Job.
func K8SJobUID(val string) attribute.KeyValue {
	return K8SJobUIDKey.String(val)
}

// A Kubernetes CronJob object.
const (
	// K8SCronJobNameKey is the attribute Key conforming to the
	// "k8s.cronjob.name" semantic conventions. It represents the name of the
	// CronJob.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'opentelemetry'
	K8SCronJobNameKey = attribute.Key("k8s.cronjob.name")

	// K8SCronJobUIDKey is the attribute Key conforming to the
	// "k8s.cronjob.uid" semantic conventions. It represents the UID of the
	// CronJob.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '275ecb36-5aa8-4c2a-9c47-d8bb681b9aff'
	K8SCronJobUIDKey = attribute.Key("k8s.cronjob.uid")
)

// K8SCronJobName returns an attribute KeyValue conforming to the
// "k8s.cronjob.name" semantic conventions. It represents the name of the
// CronJob.
func K8SCronJobName(val string) attribute.KeyValue {
	return K8SCronJobNameKey.String(val)
}

// K8SCronJobUID returns an attribute KeyValue conforming to the
// "k8s.cronjob.uid" semantic conventions. It represents the UID of the
// CronJob.
func K8SCronJobUID(val string) attribute.KeyValue {
	return K8SCronJobUIDKey.String(val)
}

// The operating system (OS) on which the process represented by this resource
// is running.
const (
	// OSBuildIDKey is the attribute Key conforming to the "os.build_id"
	// semantic conventions. It represents the unique identifier for a
	// particular build or compilation of the operating system.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'TQ3C.230805.001.B2', '20E247', '22621'
	OSBuildIDKey = attribute.Key("os.build_id")

	// OSDescriptionKey is the attribute Key conforming to the "os.description"
	// semantic conventions. It represents the human readable (not intended to
	// be parsed) OS version information, like e.g. reported by `ver` or
	// `lsb_release -a` commands.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'Microsoft Windows [Version 10.0.18363.778]', 'Ubuntu 18.04.1
	// LTS'
	OSDescriptionKey = attribute.Key("os.description")

	// OSNameKey is the attribute Key conforming to the "os.name" semantic
	// conventions. It represents the human readable operating system name.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'iOS', 'Android', 'Ubuntu'
	OSNameKey = attribute.Key("os.name")

	// OSTypeKey is the attribute Key conforming to the "os.type" semantic
	// conventions. It represents the operating system type.
	//
	// Type: Enum
	// RequirementLevel: Required
	// Stability: experimental
	OSTypeKey = attribute.Key("os.type")

	// OSVersionKey is the attribute Key conforming to the "os.version"
	// semantic conventions. It represents the version string of the operating
	// system as defined in [Version
	// Attributes](/docs/resource/README.md#version-attributes).
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '14.2.1', '18.04.1'
	OSVersionKey = attribute.Key("os.version")
)

var (
	// Microsoft Windows
	OSTypeWindows = OSTypeKey.String("windows")
	// Linux
	OSTypeLinux = OSTypeKey.String("linux")
	// Apple Darwin
	OSTypeDarwin = OSTypeKey.String("darwin")
	// FreeBSD
	OSTypeFreeBSD = OSTypeKey.String("freebsd")
	// NetBSD
	OSTypeNetBSD = OSTypeKey.String("netbsd")
	// OpenBSD
	OSTypeOpenBSD = OSTypeKey.String("openbsd")
	// DragonFly BSD
	OSTypeDragonflyBSD = OSTypeKey.String("dragonflybsd")
	// HP-UX (Hewlett Packard Unix)
	OSTypeHPUX = OSTypeKey.String("hpux")
	// AIX (Advanced Interactive eXecutive)
	OSTypeAIX = OSTypeKey.String("aix")
	// SunOS, Oracle Solaris
	OSTypeSolaris = OSTypeKey.String("solaris")
	// IBM z/OS
	OSTypeZOS = OSTypeKey.String("z_os")
)

// OSBuildID returns an attribute KeyValue conforming to the "os.build_id"
// semantic conventions. It represents the unique identifier for a particular
// build or compilation of the operating system.
func OSBuildID(val string) attribute.KeyValue {
	return OSBuildIDKey.String(val)
}

// OSDescription returns an attribute KeyValue conforming to the
// "os.description" semantic conventions. It represents the human readable (not
// intended to be parsed) OS version information, like e.g. reported by `ver`
// or `lsb_release -a` commands.
func OSDescription(val string) attribute.KeyValue {
	return OSDescriptionKey.String(val)
}

// OSName returns an attribute KeyValue conforming to the "os.name" semantic
// conventions. It represents the human readable operating system name.
func OSName(val string) attribute.KeyValue {
	return OSNameKey.String(val)
}

// OSVersion returns an attribute KeyValue conforming to the "os.version"
// semantic conventions. It represents the version string of the operating
// system as defined in [Version
// Attributes](/docs/resource/README.md#version-attributes).
func OSVersion(val string) attribute.KeyValue {
	return OSVersionKey.String(val)
}

// An operating system process.
const (
	// ProcessCommandKey is the attribute Key conforming to the
	// "process.command" semantic conventions. It represents the command used
	// to launch the process (i.e. the command name). On Linux based systems,
	// can be set to the zeroth string in `proc/[pid]/cmdline`. On Windows, can
	// be set to the first parameter extracted from `GetCommandLineW`.
	//
	// Type: string
	// RequirementLevel: ConditionallyRequired (See alternative attributes
	// below.)
	// Stability: experimental
	// Examples: 'cmd/otelcol'
	ProcessCommandKey = attribute.Key("process.command")

	// ProcessCommandArgsKey is the attribute Key conforming to the
	// "process.command_args" semantic conventions. It represents the all the
	// command arguments (including the command/executable itself) as received
	// by the process. On Linux-based systems (and some other Unixoid systems
	// supporting procfs), can be set according to the list of null-delimited
	// strings extracted from `proc/[pid]/cmdline`. For libc-based executables,
	// this would be the full argv vector passed to `main`.
	//
	// Type: string[]
	// RequirementLevel: ConditionallyRequired (See alternative attributes
	// below.)
	// Stability: experimental
	// Examples: 'cmd/otecol', '--config=config.yaml'
	ProcessCommandArgsKey = attribute.Key("process.command_args")

	// ProcessCommandLineKey is the attribute Key conforming to the
	// "process.command_line" semantic conventions. It represents the full
	// command used to launch the process as a single string representing the
	// full command. On Windows, can be set to the result of `GetCommandLineW`.
	// Do not set this if you have to assemble it just for monitoring; use
	// `process.command_args` instead.
	//
	// Type: string
	// RequirementLevel: ConditionallyRequired (See alternative attributes
	// below.)
	// Stability: experimental
	// Examples: 'C:\\cmd\\otecol --config="my directory\\config.yaml"'
	ProcessCommandLineKey = attribute.Key("process.command_line")

	// ProcessExecutableNameKey is the attribute Key conforming to the
	// "process.executable.name" semantic conventions. It represents the name
	// of the process executable. On Linux based systems, can be set to the
	// `Name` in `proc/[pid]/status`. On Windows, can be set to the base name
	// of `GetProcessImageFileNameW`.
	//
	// Type: string
	// RequirementLevel: ConditionallyRequired (See alternative attributes
	// below.)
	// Stability: experimental
	// Examples: 'otelcol'
	ProcessExecutableNameKey = attribute.Key("process.executable.name")

	// ProcessExecutablePathKey is the attribute Key conforming to the
	// "process.executable.path" semantic conventions. It represents the full
	// path to the process executable. On Linux based systems, can be set to
	// the target of `proc/[pid]/exe`. On Windows, can be set to the result of
	// `GetProcessImageFileNameW`.
	//
	// Type: string
	// RequirementLevel: ConditionallyRequired (See alternative attributes
	// below.)
	// Stability: experimental
	// Examples: '/usr/bin/cmd/otelcol'
	ProcessExecutablePathKey = attribute.Key("process.executable.path")

	// ProcessOwnerKey is the attribute Key conforming to the "process.owner"
	// semantic conventions. It represents the username of the user that owns
	// the process.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'root'
	ProcessOwnerKey = attribute.Key("process.owner")

	// ProcessParentPIDKey is the attribute Key conforming to the
	// "process.parent_pid" semantic conventions. It represents the parent
	// Process identifier (PID).
	//
	// Type: int
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 111
	ProcessParentPIDKey = attribute.Key("process.parent_pid")

	// ProcessPIDKey is the attribute Key conforming to the "process.pid"
	// semantic conventions. It represents the process identifier (PID).
	//
	// Type: int
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 1234
	ProcessPIDKey = attribute.Key("process.pid")
)

// ProcessCommand returns an attribute KeyValue conforming to the
// "process.command" semantic conventions. It represents the command used to
// launch the process (i.e. the command name). On Linux based systems, can be
// set to the zeroth string in `proc/[pid]/cmdline`. On Windows, can be set to
// the first parameter extracted from `GetCommandLineW`.
func ProcessCommand(val string) attribute.KeyValue {
	return ProcessCommandKey.String(val)
}

// ProcessCommandArgs returns an attribute KeyValue conforming to the
// "process.command_args" semantic conventions. It represents the all the
// command arguments (including the command/executable itself) as received by
// the process. On Linux-based systems (and some other Unixoid systems
// supporting procfs), can be set according to the list of null-delimited
// strings extracted from `proc/[pid]/cmdline`. For libc-based executables,
// this would be the full argv vector passed to `main`.
func ProcessCommandArgs(val ...string) attribute.KeyValue {
	return ProcessCommandArgsKey.StringSlice(val)
}

// ProcessCommandLine returns an attribute KeyValue conforming to the
// "process.command_line" semantic conventions. It represents the full command
// used to launch the process as a single string representing the full command.
// On Windows, can be set to the result of `GetCommandLineW`. Do not set this
// if you have to assemble it just for monitoring; use `process.command_args`
// instead.
func ProcessCommandLine(val string) attribute.KeyValue {
	return ProcessCommandLineKey.String(val)
}

// ProcessExecutableName returns an attribute KeyValue conforming to the
// "process.executable.name" semantic conventions. It represents the name of
// the process executable. On Linux based systems, can be set to the `Name` in
// `proc/[pid]/status`. On Windows, can be set to the base name of
// `GetProcessImageFileNameW`.
func ProcessExecutableName(val string) attribute.KeyValue {
	return ProcessExecutableNameKey.String(val)
}

// ProcessExecutablePath returns an attribute KeyValue conforming to the
// "process.executable.path" semantic conventions. It represents the full path
// to the process executable. On Linux based systems, can be set to the target
// of `proc/[pid]/exe`. On Windows, can be set to the result of
// `GetProcessImageFileNameW`.
func ProcessExecutablePath(val string) attribute.KeyValue {
	return ProcessExecutablePathKey.String(val)
}

// ProcessOwner returns an attribute KeyValue conforming to the
// "process.owner" semantic conventions. It represents the username of the user
// that owns the process.
func ProcessOwner(val string) attribute.KeyValue {
	return ProcessOwnerKey.String(val)
}

// ProcessParentPID returns an attribute KeyValue conforming to the
// "process.parent_pid" semantic conventions. It represents the parent Process
// identifier (PID).
func ProcessParentPID(val int) attribute.KeyValue {
	return ProcessParentPIDKey.Int(val)
}

// ProcessPID returns an attribute KeyValue conforming to the "process.pid"
// semantic conventions. It represents the process identifier (PID).
func ProcessPID(val int) attribute.KeyValue {
	return ProcessPIDKey.Int(val)
}

// The single (language) runtime instance which is monitored.
const (
	// ProcessRuntimeDescriptionKey is the attribute Key conforming to the
	// "process.runtime.description" semantic conventions. It represents an
	// additional description about the runtime of the process, for example a
	// specific vendor customization of the runtime environment.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'Eclipse OpenJ9 Eclipse OpenJ9 VM openj9-0.21.0'
	ProcessRuntimeDescriptionKey = attribute.Key("process.runtime.description")

	// ProcessRuntimeNameKey is the attribute Key conforming to the
	// "process.runtime.name" semantic conventions. It represents the name of
	// the runtime of this process. For compiled native binaries, this SHOULD
	// be the name of the compiler.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'OpenJDK Runtime Environment'
	ProcessRuntimeNameKey = attribute.Key("process.runtime.name")

	// ProcessRuntimeVersionKey is the attribute Key conforming to the
	// "process.runtime.version" semantic conventions. It represents the
	// version of the runtime of this process, as returned by the runtime
	// without modification.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '14.0.2'
	ProcessRuntimeVersionKey = attribute.Key("process.runtime.version")
)

// ProcessRuntimeDescription returns an attribute KeyValue conforming to the
// "process.runtime.description" semantic conventions. It represents an
// additional description about the runtime of the process, for example a
// specific vendor customization of the runtime environment.
func ProcessRuntimeDescription(val string) attribute.KeyValue {
	return ProcessRuntimeDescriptionKey.String(val)
}

// ProcessRuntimeName returns an attribute KeyValue conforming to the
// "process.runtime.name" semantic conventions. It represents the name of the
// runtime of this process. For compiled native binaries, this SHOULD be the
// name of the compiler.
func ProcessRuntimeName(val string) attribute.KeyValue {
	return ProcessRuntimeNameKey.String(val)
}

// ProcessRuntimeVersion returns an attribute KeyValue conforming to the
// "process.runtime.version" semantic conventions. It represents the version of
// the runtime of this process, as returned by the runtime without
// modification.
func ProcessRuntimeVersion(val string) attribute.KeyValue {
	return ProcessRuntimeVersionKey.String(val)
}

// A service instance.
const (
	// ServiceNameKey is the attribute Key conforming to the "service.name"
	// semantic conventions. It represents the logical name of the service.
	//
	// Type: string
	// RequirementLevel: Required
	// Stability: experimental
	// Examples: 'shoppingcart'
	// Note: MUST be the same for all instances of horizontally scaled
	// services. If the value was not specified, SDKs MUST fallback to
	// `unknown_service:` concatenated with
	// [`process.executable.name`](process.md#process), e.g.
	// `unknown_service:bash`. If `process.executable.name` is not available,
	// the value MUST be set to `unknown_service`.
	ServiceNameKey = attribute.Key("service.name")

	// ServiceVersionKey is the attribute Key conforming to the
	// "service.version" semantic conventions. It represents the version string
	// of the service API or implementation. The format is not defined by these
	// conventions.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '2.0.0', 'a01dbef8a'
	ServiceVersionKey = attribute.Key("service.version")
)

// ServiceName returns an attribute KeyValue conforming to the
// "service.name" semantic conventions. It represents the logical name of the
// service.
func ServiceName(val string) attribute.KeyValue {
	return ServiceNameKey.String(val)
}

// ServiceVersion returns an attribute KeyValue conforming to the
// "service.version" semantic conventions. It represents the version string of
// the service API or implementation. The format is not defined by these
// conventions.
func ServiceVersion(val string) attribute.KeyValue {
	return ServiceVersionKey.String(val)
}

// A service instance.
const (
	// ServiceInstanceIDKey is the attribute Key conforming to the
	// "service.instance.id" semantic conventions. It represents the string ID
	// of the service instance.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'my-k8s-pod-deployment-1',
	// '627cc493-f310-47de-96bd-71410b7dec09'
	// Note: MUST be unique for each instance of the same
	// `service.namespace,service.name` pair (in other words
	// `service.namespace,service.name,service.instance.id` triplet MUST be
	// globally unique). The ID helps to distinguish instances of the same
	// service that exist at the same time (e.g. instances of a horizontally
	// scaled service). It is preferable for the ID to be persistent and stay
	// the same for the lifetime of the service instance, however it is
	// acceptable that the ID is ephemeral and changes during important
	// lifetime events for the service (e.g. service restarts). If the service
	// has no inherent unique ID that can be used as the value of this
	// attribute it is recommended to generate a random Version 1 or Version 4
	// RFC 4122 UUID (services aiming for reproducible UUIDs may also use
	// Version 5, see RFC 4122 for more recommendations).
	ServiceInstanceIDKey = attribute.Key("service.instance.id")

	// ServiceNamespaceKey is the attribute Key conforming to the
	// "service.namespace" semantic conventions. It represents a namespace for
	// `service.name`.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'Shop'
	// Note: A string value having a meaning that helps to distinguish a group
	// of services, for example the team name that owns a group of services.
	// `service.name` is expected to be unique within the same namespace. If
	// `service.namespace` is not specified in the Resource then `service.name`
	// is expected to be unique for all services that have no explicit
	// namespace defined (so the empty/unspecified namespace is simply one more
	// valid namespace). Zero-length namespace string is assumed equal to
	// unspecified namespace.
	ServiceNamespaceKey = attribute.Key("service.namespace")
)

// ServiceInstanceID returns an attribute KeyValue conforming to the
// "service.instance.id" semantic conventions. It represents the string ID of
// the service instance.
func ServiceInstanceID(val string) attribute.KeyValue {
	return ServiceInstanceIDKey.String(val)
}

// ServiceNamespace returns an attribute KeyValue conforming to the
// "service.namespace" semantic conventions. It represents a namespace for
// `service.name`.
func ServiceNamespace(val string) attribute.KeyValue {
	return ServiceNamespaceKey.String(val)
}

// The telemetry SDK used to capture data recorded by the instrumentation
// libraries.
const (
	// TelemetrySDKLanguageKey is the attribute Key conforming to the
	// "telemetry.sdk.language" semantic conventions. It represents the
	// language of the telemetry SDK.
	//
	// Type: Enum
	// RequirementLevel: Required
	// Stability: experimental
	TelemetrySDKLanguageKey = attribute.Key("telemetry.sdk.language")

	// TelemetrySDKNameKey is the attribute Key conforming to the
	// "telemetry.sdk.name" semantic conventions. It represents the name of the
	// telemetry SDK as defined above.
	//
	// Type: string
	// RequirementLevel: Required
	// Stability: experimental
	// Examples: 'opentelemetry'
	// Note: The OpenTelemetry SDK MUST set the `telemetry.sdk.name` attribute
	// to `opentelemetry`.
	// If another SDK, like a fork or a vendor-provided implementation, is
	// used, this SDK MUST set the
	// `telemetry.sdk.name` attribute to the fully-qualified class or module
	// name of this SDK's main entry point
	// or another suitable identifier depending on the language.
	// The identifier `opentelemetry` is reserved and MUST NOT be used in this
	// case.
	// All custom identifiers SHOULD be stable across different versions of an
	// implementation.
	TelemetrySDKNameKey = attribute.Key("telemetry.sdk.name")

	// TelemetrySDKVersionKey is the attribute Key conforming to the
	// "telemetry.sdk.version" semantic conventions. It represents the version
	// string of the telemetry SDK.
	//
	// Type: string
	// RequirementLevel: Required
	// Stability: experimental
	// Examples: '1.2.3'
	TelemetrySDKVersionKey = attribute.Key("telemetry.sdk.version")
)

var (
	// cpp
	TelemetrySDKLanguageCPP = TelemetrySDKLanguageKey.String("cpp")
	// dotnet
	TelemetrySDKLanguageDotnet = TelemetrySDKLanguageKey.String("dotnet")
	// erlang
	TelemetrySDKLanguageErlang = TelemetrySDKLanguageKey.String("erlang")
	// go
	TelemetrySDKLanguageGo = TelemetrySDKLanguageKey.String("go")
	// java
	TelemetrySDKLanguageJava = TelemetrySDKLanguageKey.String("java")
	// nodejs
	TelemetrySDKLanguageNodejs = TelemetrySDKLanguageKey.String("nodejs")
	// php
	TelemetrySDKLanguagePHP = TelemetrySDKLanguageKey.String("php")
	// python
	TelemetrySDKLanguagePython = TelemetrySDKLanguageKey.String("python")
	// ruby
	TelemetrySDKLanguageRuby = TelemetrySDKLanguageKey.String("ruby")
	// rust
	TelemetrySDKLanguageRust = TelemetrySDKLanguageKey.String("rust")
	// swift
	TelemetrySDKLanguageSwift = TelemetrySDKLanguageKey.String("swift")
	// webjs
	TelemetrySDKLanguageWebjs = TelemetrySDKLanguageKey.String("webjs")
)

// TelemetrySDKName returns an attribute KeyValue conforming to the
// "telemetry.sdk.name" semantic conventions. It represents the name of the
// telemetry SDK as defined above.
func TelemetrySDKName(val string) attribute.KeyValue {
	return TelemetrySDKNameKey.String(val)
}

// TelemetrySDKVersion returns an attribute KeyValue conforming to the
// "telemetry.sdk.version" semantic conventions. It represents the version
// string of the telemetry SDK.
func TelemetrySDKVersion(val string) attribute.KeyValue {
	return TelemetrySDKVersionKey.String(val)
}

// The telemetry SDK used to capture data recorded by the instrumentation
// libraries.
const (
	// TelemetryDistroNameKey is the attribute Key conforming to the
	// "telemetry.distro.name" semantic conventions. It represents the name of
	// the auto instrumentation agent or distribution, if used.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'parts-unlimited-java'
	// Note: Official auto instrumentation agents and distributions SHOULD set
	// the `telemetry.distro.name` attribute to
	// a string starting with `opentelemetry-`, e.g.
	// `opentelemetry-java-instrumentation`.
	TelemetryDistroNameKey = attribute.Key("telemetry.distro.name")

	// TelemetryDistroVersionKey is the attribute Key conforming to the
	// "telemetry.distro.version" semantic conventions. It represents the
	// version string of the auto instrumentation agent or distribution, if
	// used.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '1.2.3'
	TelemetryDistroVersionKey = attribute.Key("telemetry.distro.version")
)

// TelemetryDistroName returns an attribute KeyValue conforming to the
// "telemetry.distro.name" semantic conventions. It represents the name of the
// auto instrumentation agent or distribution, if used.
func TelemetryDistroName(val string) attribute.KeyValue {
	return TelemetryDistroNameKey.String(val)
}

// TelemetryDistroVersion returns an attribute KeyValue conforming to the
// "telemetry.distro.version" semantic conventions. It represents the version
// string of the auto instrumentation agent or distribution, if used.
func TelemetryDistroVersion(val string) attribute.KeyValue {
	return TelemetryDistroVersionKey.String(val)
}

// Resource describing the packaged software running the application code. Web
// engines are typically executed using process.runtime.
const (
	// WebEngineDescriptionKey is the attribute Key conforming to the
	// "webengine.description" semantic conventions. It represents the
	// additional description of the web engine (e.g. detailed version and
	// edition information).
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'WildFly Full 21.0.0.Final (WildFly Core 13.0.1.Final) -
	// 2.2.2.Final'
	WebEngineDescriptionKey = attribute.Key("webengine.description")

	// WebEngineNameKey is the attribute Key conforming to the "webengine.name"
	// semantic conventions. It represents the name of the web engine.
	//
	// Type: string
	// RequirementLevel: Required
	// Stability: experimental
	// Examples: 'WildFly'
	WebEngineNameKey = attribute.Key("webengine.name")

	// WebEngineVersionKey is the attribute Key conforming to the
	// "webengine.version" semantic conventions. It represents the version of
	// the web engine.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '21.0.0'
	WebEngineVersionKey = attribute.Key("webengine.version")
)

// WebEngineDescription returns an attribute KeyValue conforming to the
// "webengine.description" semantic conventions. It represents the additional
// description of the web engine (e.g. detailed version and edition
// information).
func WebEngineDescription(val string) attribute.KeyValue {
	return WebEngineDescriptionKey.String(val)
}

// WebEngineName returns an attribute KeyValue conforming to the
// "webengine.name" semantic conventions. It represents the name of the web
// engine.
func WebEngineName(val string) attribute.KeyValue {
	return WebEngineNameKey.String(val)
}

// WebEngineVersion returns an attribute KeyValue conforming to the
// "webengine.version" semantic conventions. It represents the version of the
// web engine.
func WebEngineVersion(val string) attribute.KeyValue {
	return WebEngineVersionKey.String(val)
}

// Attributes used by non-OTLP exporters to represent OpenTelemetry Scope's
// concepts.
const (
	// OTelScopeNameKey is the attribute Key conforming to the
	// "otel.scope.name" semantic conventions. It represents the name of the
	// instrumentation scope - (`InstrumentationScope.Name` in OTLP).
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'io.opentelemetry.contrib.mongodb'
	OTelScopeNameKey = attribute.Key("otel.scope.name")

	// OTelScopeVersionKey is the attribute Key conforming to the
	// "otel.scope.version" semantic conventions. It represents the version of
	// the instrumentation scope - (`InstrumentationScope.Version` in OTLP).
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '1.0.0'
	OTelScopeVersionKey = attribute.Key("otel.scope.version")
)

// OTelScopeName returns an attribute KeyValue conforming to the
// "otel.scope.name" semantic conventions. It represents the name of the
// instrumentation scope - (`InstrumentationScope.Name` in OTLP).
func OTelScopeName(val string) attribute.KeyValue {
	return OTelScopeNameKey.String(val)
}

// OTelScopeVersion returns an attribute KeyValue conforming to the
// "otel.scope.version" semantic conventions. It represents the version of the
// instrumentation scope - (`InstrumentationScope.Version` in OTLP).
func OTelScopeVersion(val string) attribute.KeyValue {
	return OTelScopeVersionKey.String(val)
}

// Span attributes used by non-OTLP exporters to represent OpenTelemetry
// Scope's concepts.
const (
	// OTelLibraryNameKey is the attribute Key conforming to the
	// "otel.library.name" semantic conventions.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: deprecated
	// Examples: 'io.opentelemetry.contrib.mongodb'
	// Deprecated: use the `otel.scope.name` attribute.
	OTelLibraryNameKey = attribute.Key("otel.library.name")

	// OTelLibraryVersionKey is the attribute Key conforming to the
	// "otel.library.version" semantic conventions.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: deprecated
	// Examples: '1.0.0'
	// Deprecated: use the `otel.scope.version` attribute.
	OTelLibraryVersionKey = attribute.Key("otel.library.version")
)

// OTelLibraryName returns an attribute KeyValue conforming to the
// "otel.library.name" semantic conventions.
//
// Deprecated: use the `otel.scope.name` attribute.
func OTelLibraryName(val string) attribute.KeyValue {
	return OTelLibraryNameKey.String(val)
}

// OTelLibraryVersion returns an attribute KeyValue conforming to the
// "otel.library.version" semantic conventions.
//
// Deprecated: use the `otel.scope.version` attribute.
func OTelLibraryVersion(val string) attribute.KeyValue {
	return OTelLibraryVersionKey.String(val)
}
