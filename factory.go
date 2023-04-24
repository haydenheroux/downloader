package main

import "music_dl/downloader"

func createDownloader(name string) downloader.Downloader {
	switch name {
	case "mock":
		return downloader.MockDownloader{
			Format:          outputFormat,
			OutputDirectory: outputDirectory,
		}
	case "ytdlp":
		return downloader.YoutubeDLCompatibleDownloader{
			Executable:      "yt-dlp",
			Format:          outputFormat,
			OutputDirectory: outputDirectory,
		}
	case "ytdl":
		fallthrough
	default:
		return downloader.YoutubeDLCompatibleDownloader{
			Executable:      "youtube_dl",
			Format:          outputFormat,
			OutputDirectory: outputDirectory,
		}
	}
}
