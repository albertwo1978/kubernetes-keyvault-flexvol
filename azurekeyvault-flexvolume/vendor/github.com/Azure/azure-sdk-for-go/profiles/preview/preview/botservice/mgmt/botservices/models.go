// +build go1.9

// Copyright 2018 Microsoft Corporation
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

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package botservice

import original "github.com/Azure/azure-sdk-for-go/services/preview/botservice/mgmt/2017-12-01/botservices"

type BotsClient = original.BotsClient
type ChannelsClient = original.ChannelsClient

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient
type ChannelName = original.ChannelName

const (
	ChannelNameDirectLineChannel ChannelName = original.ChannelNameDirectLineChannel
	ChannelNameEmailChannel      ChannelName = original.ChannelNameEmailChannel
	ChannelNameFacebookChannel   ChannelName = original.ChannelNameFacebookChannel
	ChannelNameKikChannel        ChannelName = original.ChannelNameKikChannel
	ChannelNameMsTeamsChannel    ChannelName = original.ChannelNameMsTeamsChannel
	ChannelNameSkypeChannel      ChannelName = original.ChannelNameSkypeChannel
	ChannelNameSlackChannel      ChannelName = original.ChannelNameSlackChannel
	ChannelNameSmsChannel        ChannelName = original.ChannelNameSmsChannel
	ChannelNameTelegramChannel   ChannelName = original.ChannelNameTelegramChannel
	ChannelNameWebChatChannel    ChannelName = original.ChannelNameWebChatChannel
)

type ChannelNameBasicChannel = original.ChannelNameBasicChannel

const (
	ChannelNameChannel            ChannelNameBasicChannel = original.ChannelNameChannel
	ChannelNameDirectLineChannel1 ChannelNameBasicChannel = original.ChannelNameDirectLineChannel1
	ChannelNameEmailChannel1      ChannelNameBasicChannel = original.ChannelNameEmailChannel1
	ChannelNameFacebookChannel1   ChannelNameBasicChannel = original.ChannelNameFacebookChannel1
	ChannelNameKikChannel1        ChannelNameBasicChannel = original.ChannelNameKikChannel1
	ChannelNameMsTeamsChannel1    ChannelNameBasicChannel = original.ChannelNameMsTeamsChannel1
	ChannelNameSkypeChannel1      ChannelNameBasicChannel = original.ChannelNameSkypeChannel1
	ChannelNameSlackChannel1      ChannelNameBasicChannel = original.ChannelNameSlackChannel1
	ChannelNameSmsChannel1        ChannelNameBasicChannel = original.ChannelNameSmsChannel1
	ChannelNameTelegramChannel1   ChannelNameBasicChannel = original.ChannelNameTelegramChannel1
	ChannelNameWebChatChannel1    ChannelNameBasicChannel = original.ChannelNameWebChatChannel1
)

type Kind = original.Kind

const (
	KindBot      Kind = original.KindBot
	KindDesigner Kind = original.KindDesigner
	KindFunction Kind = original.KindFunction
	KindSdk      Kind = original.KindSdk
)

type SkuName = original.SkuName

const (
	F0 SkuName = original.F0
	S1 SkuName = original.S1
)

type SkuTier = original.SkuTier

const (
	Free     SkuTier = original.Free
	Standard SkuTier = original.Standard
)

type Bot = original.Bot
type BotChannel = original.BotChannel
type BotProperties = original.BotProperties
type BotResponseList = original.BotResponseList
type BotResponseListIterator = original.BotResponseListIterator
type BotResponseListPage = original.BotResponseListPage
type BasicChannel = original.BasicChannel
type Channel = original.Channel
type ChannelResponseList = original.ChannelResponseList
type ChannelResponseListIterator = original.ChannelResponseListIterator
type ChannelResponseListPage = original.ChannelResponseListPage
type CheckNameAvailabilityRequestBody = original.CheckNameAvailabilityRequestBody
type CheckNameAvailabilityResponseBody = original.CheckNameAvailabilityResponseBody
type DirectLineChannel = original.DirectLineChannel
type DirectLineChannelProperties = original.DirectLineChannelProperties
type DirectLineSite = original.DirectLineSite
type EmailChannel = original.EmailChannel
type EmailChannelProperties = original.EmailChannelProperties
type Error = original.Error
type ErrorBody = original.ErrorBody
type FacebookChannel = original.FacebookChannel
type FacebookChannelProperties = original.FacebookChannelProperties
type FacebookPage = original.FacebookPage
type KikChannel = original.KikChannel
type KikChannelProperties = original.KikChannelProperties
type MsTeamsChannel = original.MsTeamsChannel
type MsTeamsChannelProperties = original.MsTeamsChannelProperties
type OperationDisplayInfo = original.OperationDisplayInfo
type OperationEntity = original.OperationEntity
type OperationEntityListResult = original.OperationEntityListResult
type OperationEntityListResultIterator = original.OperationEntityListResultIterator
type OperationEntityListResultPage = original.OperationEntityListResultPage
type Resource = original.Resource
type Sku = original.Sku
type SkypeChannel = original.SkypeChannel
type SkypeChannelProperties = original.SkypeChannelProperties
type SlackChannel = original.SlackChannel
type SlackChannelProperties = original.SlackChannelProperties
type SmsChannel = original.SmsChannel
type SmsChannelProperties = original.SmsChannelProperties
type TelegramChannel = original.TelegramChannel
type TelegramChannelProperties = original.TelegramChannelProperties
type WebChatChannel = original.WebChatChannel
type WebChatChannelProperties = original.WebChatChannelProperties
type WebChatSite = original.WebChatSite
type OperationsClient = original.OperationsClient

func NewBotsClient(subscriptionID string) BotsClient {
	return original.NewBotsClient(subscriptionID)
}
func NewBotsClientWithBaseURI(baseURI string, subscriptionID string) BotsClient {
	return original.NewBotsClientWithBaseURI(baseURI, subscriptionID)
}
func NewChannelsClient(subscriptionID string) ChannelsClient {
	return original.NewChannelsClient(subscriptionID)
}
func NewChannelsClientWithBaseURI(baseURI string, subscriptionID string) ChannelsClient {
	return original.NewChannelsClientWithBaseURI(baseURI, subscriptionID)
}
func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleChannelNameValues() []ChannelName {
	return original.PossibleChannelNameValues()
}
func PossibleChannelNameBasicChannelValues() []ChannelNameBasicChannel {
	return original.PossibleChannelNameBasicChannelValues()
}
func PossibleKindValues() []Kind {
	return original.PossibleKindValues()
}
func PossibleSkuNameValues() []SkuName {
	return original.PossibleSkuNameValues()
}
func PossibleSkuTierValues() []SkuTier {
	return original.PossibleSkuTierValues()
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
