package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
)

type PluginInfo struct {
	Name        string
	Description string
	Author      string
	Version     string
}

// Get available plugins
func getPlugins(dir *string) (*[]PluginInfo, error) {
	var plugins []PluginInfo

	files, err := os.ReadDir(path.Clean(fmt.Sprint(*dir, "/plugins")))
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if !strings.HasSuffix(file.Name(), ".js") {
			continue
		}

		// Check if plugin has info file
		infoFilePath := path.Clean(fmt.Sprintf("%s/plugins/%s.json", *dir, strings.TrimSuffix(file.Name(), ".js")))
		if _, err := os.Stat(infoFilePath); errors.Is(err, os.ErrNotExist) {
			continue
		}

		// Get the plugin's info
		infoFile, err := os.Open(infoFilePath)
		if err != nil {
			return nil, err
		}

		infoFileContent, err := io.ReadAll(infoFile)
		if err != nil {
			return nil, err
		}

		var info PluginInfo
		if err := json.Unmarshal(infoFileContent, &info); err != nil {
			return nil, err
		}

		plugins = append(plugins, info)
	}

	return &plugins, nil
}

// Load a file
func loadFile(name string) (*string, error) {
	file, err := os.Open(path.Clean(fmt.Sprint("./files/", name)))
	if err != nil {
		return nil, err
	}

	fileContent, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	file.Close()

	fileString := string(fileContent)
	return &fileString, nil
}

// Write to a file
func writeFile(name string, data string) error {
	out, err := os.OpenFile(path.Clean(name), os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0650)
	if err != nil {
		return err
	}

	_, err = out.WriteString(data)
	if err != nil {
		return err
	}

	out.Close()

	return nil
}
