package docsrc

import (
	"github.com/qor5/docs/docsrc/content"
	advanced_functions "github.com/qor5/docs/docsrc/content/advanced-functions"
	"github.com/qor5/docs/docsrc/content/basics"
	digging_deeper "github.com/qor5/docs/docsrc/content/digging-deeper"
	getting_started "github.com/qor5/docs/docsrc/content/getting-started"
	"github.com/qor5/docs/docsrc/utils"
	"github.com/theplant/docgo"
)

var DocTree = []interface{}{
	content.Home,
	&docgo.DocsGroup{
		Title: "Getting Started",
		Docs: []*docgo.DocBuilder{
			getting_started.OneMinuteQuickStart,
		},
	},
	&docgo.DocsGroup{
		Title: "Building Admin",
		Docs: []*docgo.DocBuilder{
			// listing
			basics.Listing,
			basics.Filter,
			// editing
			basics.EditingCustomizations,
			// brand
			basics.Brand,
			// menu
			basics.ManageMenu,
			advanced_functions.DetailPageForComplexObject,
			// permission
			basics.Permissions,
			basics.Role,
			// other basics
			basics.NotificationCenter,
			basics.ShortCut,
			basics.ConfirmDialog,
			basics.Slug,
			basics.SEO,
			basics.Activity,
			basics.Worker,
		},
	},

	&docgo.DocsGroup{
		Title: "Web Application",
		Docs: []*docgo.DocBuilder{
			advanced_functions.PageFuncAndEventFunc,
			advanced_functions.TheGoHTMLBuilder,
			advanced_functions.ItsTheWholeHouse,
			advanced_functions.LazyPortalsAndReload,
			advanced_functions.LayoutFunctionAndPageInjector,
			advanced_functions.SwitchPagesWithPushState,
			advanced_functions.ReloadPageWithAFlash,
			advanced_functions.PartialRefreshWithPortal,
			advanced_functions.ManipulatePageURLInEventFunc,
			advanced_functions.SummaryOfEventResponse,
			advanced_functions.WebScope,
			advanced_functions.EventHandling,
			basics.FormHandling,
		},
	},

	&docgo.DocsGroup{
		Title: "UI Components",
		Docs: []*docgo.DocBuilder{
			// TODO: move BasicInputs to ATasteOfUsingVuetifyInGo
			basics.BasicInputs,
			advanced_functions.ATasteOfUsingVuetifyInGo,
			// vuetifyx
			basics.LinkageSelect,
			basics.AutoComplete,
			// build ui component
			digging_deeper.CompositeNewComponentWithGo,
			digging_deeper.IntegrateAHeavyVueComponent,
		},
	},

	&docgo.DocsGroup{
		Title: "Appendix",
		Docs: []*docgo.DocBuilder{
			docgo.Doc(utils.ExamplesDoc()).
				Title("All Demo Examples").
				Slug("appendix/all-demo-examples"),
		},
	},
}
