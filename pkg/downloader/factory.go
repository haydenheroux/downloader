package downloader

func CreateDownloader(name string, format string) Downloader {
	switch name {
	case "mock":
		return MockDownloader{
			Format: format,
		}
	case "yt-dlp":
		fallthrough
	case "ytdlp":
		return YoutubeDLCompatibleDownloader{
			Executable: "yt-dlp",
			Format:     format,
		}
	case "ytdl":
		fallthrough
	default:
		return YoutubeDLCompatibleDownloader{
			Executable: "youtube_dl",
			Format:     format,
		}
	}
}
