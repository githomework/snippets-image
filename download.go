// download image to folder
// return path and whether image is in the portrait orientation.
func getImage(sku string) (string, bool) {
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	resp, err := client.Get("https://server/" + sku + ".jpg")
	if err != nil {
		log.Println(err)
		return global.FolderAndSlash + "images/blank.jpg", false
	}
	defer resp.Body.Close()

	var buf []byte
	buffer := bytes.NewBuffer(buf)
	buffer.ReadFrom(resp.Body)
	img, _ := os.Create(global.FolderAndSlash + "images/" + sku + ".jpg")
	defer img.Close()
	img.Write(buffer.Bytes())
	info, _, err := image.DecodeConfig(bytes.NewReader(buffer.Bytes()))
	if err != nil {
		return strings.ReplaceAll(global.FolderAndSlash+"images/blank.jpg", "\\", "/"), false
	}
	if info.Height == 0 {
		return strings.ReplaceAll(global.FolderAndSlash+"images/blank.jpg", "\\", "/"), false
	}

	if info.Height > info.Width {
		return strings.ReplaceAll(global.FolderAndSlash+"images/"+sku+".jpg", "\\", "/"), true
	} else {
		return strings.ReplaceAll(global.FolderAndSlash+"images/"+sku+".jpg", "\\", "/"), false
	}
}
