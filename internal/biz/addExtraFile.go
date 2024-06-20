package biz

import (
	"blog-processor/global"
	"blog-processor/internal/model"
	"blog-processor/pkg/utils/file_handler"
)

func AddSearchFile() {
	if !global.SearchSetting.Require {
		return
	}
	buildAndWriteSearchFile()
}

func AddArchivesFile() {
	if !global.ArchivesSetting.Require {
		return
	}
	buildAndWriteArchivesFile()
}

func buildAndWriteSearchFile() {
	var header model.BlogHeader
	header.Title = global.SearchSetting.Title
	header.Layout = global.SearchSetting.Layout
	header.Summary = global.SearchSetting.Summary
	header.PlaceHolder = global.SearchSetting.Placeholder

	path := global.BasicSetting.BlogDir + "/" + "search.md"
	err := file_handler.WriteProcessedMarkdown(path, &header, "")
	if err != nil {
		return
	}
}

func buildAndWriteArchivesFile() {
	var header model.BlogHeader
	header.Title = global.ArchivesSetting.Title
	header.Layout = global.ArchivesSetting.Layout
	header.Summary = global.ArchivesSetting.Summary
	header.Url = global.ArchivesSetting.Url

	path := global.BasicSetting.BlogDir + "/" + "archives.md"
	err := file_handler.WriteProcessedMarkdown(path, &header, "")
	if err != nil {
		return
	}
}
