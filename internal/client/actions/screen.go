package actions

import (
  "net"
  "log"
  "image"
	"image/png"
	"github.com/kbinani/screenshot"
)

func Screen(conn net.Conn, args []string) {

  if buf, err := takeScreenshot(); err != nil {
  	log.Println("Error taking screenshot:", err.Error())
  	conn.Write([]byte{0})
  } else {
  	log.Println("Sending screenshot...")
  	png.Encode(conn, buf)
  }

}

func takeScreenshot() (image.Image, error) {
	// Take a screenshot and return the bytes
	bounds := screenshot.GetDisplayBounds(0)
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		return nil, err
	}
	return img, nil
}
