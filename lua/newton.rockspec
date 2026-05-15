package = "voxgig-sdk-newton"
version = "0.0-1"
source = {
  url = "git://github.com/voxgig-sdk/newton-sdk.git"
}
description = {
  summary = "Newton SDK for Lua",
  license = "MIT"
}
dependencies = {
  "lua >= 5.3",
  "dkjson >= 2.5",
  "dkjson >= 2.5",
}
build = {
  type = "builtin",
  modules = {
    ["newton_sdk"] = "newton_sdk.lua",
    ["config"] = "config.lua",
    ["features"] = "features.lua",
  }
}
