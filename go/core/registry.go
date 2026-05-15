package core

var UtilityRegistrar func(u *Utility)

var NewBaseFeatureFunc func() Feature

var NewTestFeatureFunc func() Feature

var NewAbsEntityFunc func(client *NewtonSDK, entopts map[string]any) NewtonEntity

var NewArccoEntityFunc func(client *NewtonSDK, entopts map[string]any) NewtonEntity

var NewArcsinEntityFunc func(client *NewtonSDK, entopts map[string]any) NewtonEntity

var NewArctanEntityFunc func(client *NewtonSDK, entopts map[string]any) NewtonEntity

var NewAreaEntityFunc func(client *NewtonSDK, entopts map[string]any) NewtonEntity

var NewCosEntityFunc func(client *NewtonSDK, entopts map[string]any) NewtonEntity

var NewDeriveEntityFunc func(client *NewtonSDK, entopts map[string]any) NewtonEntity

var NewFactorEntityFunc func(client *NewtonSDK, entopts map[string]any) NewtonEntity

var NewIntegrateEntityFunc func(client *NewtonSDK, entopts map[string]any) NewtonEntity

var NewLogEntityFunc func(client *NewtonSDK, entopts map[string]any) NewtonEntity

var NewSimplifyEntityFunc func(client *NewtonSDK, entopts map[string]any) NewtonEntity

var NewSinEntityFunc func(client *NewtonSDK, entopts map[string]any) NewtonEntity

var NewTanEntityFunc func(client *NewtonSDK, entopts map[string]any) NewtonEntity

var NewTangentEntityFunc func(client *NewtonSDK, entopts map[string]any) NewtonEntity

var NewZeroEntityFunc func(client *NewtonSDK, entopts map[string]any) NewtonEntity

