package panorama_api

import "github.com/untrustedmodders/go-plugify"

var _ = plugify.ApiVersion

// Generated from panorama_api

// OnCsScriptReadyCallback - CS Script has come up and is ready to use.
type OnCsScriptReadyCallback func()


// OnHudClickedCallback - A hud button click. Called when a button in a CustomHudLayout is clicked.
type OnHudClickedCallback func(playerSlot int32, name string, buttonId string)


