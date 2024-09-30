package folder

import (
	"errors"
	"strings"

	"github.com/gofrs/uuid"
)

func GetAllFolders() []Folder {
	return GetSampleData()
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	folders := f.folders

	res := []Folder{}
	for _, f := range folders {
		if f.OrgId == orgID {
			res = append(res, f)
		}
	}

	return res

}

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) ([]Folder, error) {

	// Check if orgID is valid, (non nil)
	if orgID == uuid.Nil {
		return []Folder{}, errors.New("invalid orgID, cannot be nil UUID")
	}

	// Get folders
	folders := f.folders

	// Find folder name, check if it is a folder that exists and that it's in the organization.
	foundFolder := containsFolder(folders, name)
	if foundFolder == nil {
		return []Folder{}, errors.New("Folder does not exist")
	} else if foundFolder.OrgId != orgID {
		return []Folder{}, errors.New("Folder does not exist in the specified organization")
	}

	// Get all folders that match the orgID and contain the name as part of the filepath
	res := []Folder{}
	for _, folder := range folders {
		if folder.OrgId == orgID && strings.Contains(folder.Paths, (name + ".")) {
			res = append(res, folder)
		}
	}

	return res, nil
}

// Helper function to check if the list of folders contains a folder with a given name, returns that folder if it exists
func containsFolder(folders []Folder, name string) (*Folder) {
	for _, folder := range folders {
		if folder.Name == name {
			return &folder
		}
	}
	return nil
}