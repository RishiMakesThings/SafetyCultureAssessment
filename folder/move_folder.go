package folder

import (
	"errors"
	"strings"
)

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	
	// Get folders list
	folders := f.folders

	// Get pointers to the source and destination folders
	sourceFolder := containsFolder(folders, name)
	dstFolder := containsFolder(folders, dst)

	// Error checking
	if sourceFolder == nil {
		return []Folder{}, errors.New("Error: Source folder does not exist")
	} else if dstFolder == nil {
		return []Folder{}, errors.New("Error: Destination folder does not exist")
	} else if sourceFolder.Name == dstFolder.Name {
		return []Folder{}, errors.New("Error: Cannot move a folder to itself")
	} else if sourceFolder.OrgId != dstFolder.OrgId {
		return []Folder{}, errors.New("Error: Cannot move a folder to a different organization")
	} else if strings.Contains(dstFolder.Paths, name) {
		return []Folder{}, errors.New("Error: Cannot move a folder to a child of itself")
	}

	// Iterate through folders updating paths
	res := []Folder{}
	newPath := dstFolder.Paths + "." + sourceFolder.Name
	for _, folder := range folders {
		// Update source folder path
		if folder.OrgId == sourceFolder.OrgId && folder.Name == sourceFolder.Name {
			res = append(res, Folder{
				Name: sourceFolder.Name,
				OrgId: sourceFolder.OrgId,
				Paths: newPath,
			})
		// Update paths of children of source folder
		} else if folder.OrgId == sourceFolder.OrgId && strings.Contains(folder.Paths, (name + ".")) {
			res = append(res, Folder{
				Name: folder.Name,
				OrgId: folder.OrgId,
				Paths: newPath + "." + strings.TrimPrefix(folder.Paths, sourceFolder.Paths + "."),
			})
		} else {
			res = append(res, folder)
		}
	}

	return res, nil
}
