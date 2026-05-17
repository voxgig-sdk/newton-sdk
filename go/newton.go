package voxgignewtonsdk

import (
	"github.com/voxgig-sdk/newton-sdk/go/core"
	"github.com/voxgig-sdk/newton-sdk/go/entity"
	"github.com/voxgig-sdk/newton-sdk/go/feature"
	_ "github.com/voxgig-sdk/newton-sdk/go/utility"
)

// Type aliases preserve external API.
type NewtonSDK = core.NewtonSDK
type Context = core.Context
type Utility = core.Utility
type Feature = core.Feature
type Entity = core.Entity
type NewtonEntity = core.NewtonEntity
type FetcherFunc = core.FetcherFunc
type Spec = core.Spec
type Result = core.Result
type Response = core.Response
type Operation = core.Operation
type Control = core.Control
type NewtonError = core.NewtonError

// BaseFeature from feature package.
type BaseFeature = feature.BaseFeature

func init() {
	core.NewBaseFeatureFunc = func() core.Feature {
		return feature.NewBaseFeature()
	}
	core.NewTestFeatureFunc = func() core.Feature {
		return feature.NewTestFeature()
	}
	core.NewAbsEntityFunc = func(client *core.NewtonSDK, entopts map[string]any) core.NewtonEntity {
		return entity.NewAbsEntity(client, entopts)
	}
	core.NewArccoEntityFunc = func(client *core.NewtonSDK, entopts map[string]any) core.NewtonEntity {
		return entity.NewArccoEntity(client, entopts)
	}
	core.NewArcsinEntityFunc = func(client *core.NewtonSDK, entopts map[string]any) core.NewtonEntity {
		return entity.NewArcsinEntity(client, entopts)
	}
	core.NewArctanEntityFunc = func(client *core.NewtonSDK, entopts map[string]any) core.NewtonEntity {
		return entity.NewArctanEntity(client, entopts)
	}
	core.NewAreaEntityFunc = func(client *core.NewtonSDK, entopts map[string]any) core.NewtonEntity {
		return entity.NewAreaEntity(client, entopts)
	}
	core.NewCosEntityFunc = func(client *core.NewtonSDK, entopts map[string]any) core.NewtonEntity {
		return entity.NewCosEntity(client, entopts)
	}
	core.NewDeriveEntityFunc = func(client *core.NewtonSDK, entopts map[string]any) core.NewtonEntity {
		return entity.NewDeriveEntity(client, entopts)
	}
	core.NewFactorEntityFunc = func(client *core.NewtonSDK, entopts map[string]any) core.NewtonEntity {
		return entity.NewFactorEntity(client, entopts)
	}
	core.NewIntegrateEntityFunc = func(client *core.NewtonSDK, entopts map[string]any) core.NewtonEntity {
		return entity.NewIntegrateEntity(client, entopts)
	}
	core.NewLogEntityFunc = func(client *core.NewtonSDK, entopts map[string]any) core.NewtonEntity {
		return entity.NewLogEntity(client, entopts)
	}
	core.NewSimplifyEntityFunc = func(client *core.NewtonSDK, entopts map[string]any) core.NewtonEntity {
		return entity.NewSimplifyEntity(client, entopts)
	}
	core.NewSinEntityFunc = func(client *core.NewtonSDK, entopts map[string]any) core.NewtonEntity {
		return entity.NewSinEntity(client, entopts)
	}
	core.NewTanEntityFunc = func(client *core.NewtonSDK, entopts map[string]any) core.NewtonEntity {
		return entity.NewTanEntity(client, entopts)
	}
	core.NewTangentEntityFunc = func(client *core.NewtonSDK, entopts map[string]any) core.NewtonEntity {
		return entity.NewTangentEntity(client, entopts)
	}
	core.NewZeroEntityFunc = func(client *core.NewtonSDK, entopts map[string]any) core.NewtonEntity {
		return entity.NewZeroEntity(client, entopts)
	}
}

// Constructor re-exports.
var NewNewtonSDK = core.NewNewtonSDK
var TestSDK = core.TestSDK
var NewContext = core.NewContext
var NewSpec = core.NewSpec
var NewResult = core.NewResult
var NewResponse = core.NewResponse
var NewOperation = core.NewOperation
var MakeConfig = core.MakeConfig
var NewBaseFeature = feature.NewBaseFeature
var NewTestFeature = feature.NewTestFeature
