package main

// go build -ldflags -H=windowsgui
import (
	"log"

	"github.com/dianchi/PowerNukkitX-Panel/ui"
	"github.com/jchv/go-webview2"
)

func main() {
	w := webview2.NewWithOptions(webview2.WebViewOptions{
		Debug:     true,
		AutoFocus: true,
		WindowOptions: webview2.WindowOptions{
			Title:  "Minimal webview example",
			Width:  800,
			Height: 600,
			IconId: 2, // icon resource id
			Center: true,
		},
	})
	if w == nil {
		log.Fatalln("Failed to load webview.")
	}
	defer w.Destroy()
	w.SetSize(800, 600, webview2.HintFixed)
	html := ui.FileReader("./ui/html/index.html")
	w.Navigate("data:text/html," + html)
	w.Run()
}
