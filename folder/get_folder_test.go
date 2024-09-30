package folder_test

import (
	"errors"
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

// feel free to change how the unit test is structured
func Test_folder_GetAllFolders(t *testing.T) {
	sampleData := folder.GetSampleData()

	t.Parallel()
	tests := [...]struct {
		testName    string
		want    []folder.Folder
	}{
		{
			testName:  "Get all folders",
			want: sampleData, 
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			get := folder.GetAllFolders()
			assert.Equal(t, tt.want, get)
		})
	}
}

func Test_folder_GetFoldersByOrgID(t *testing.T) {
	sampleData := folder.GetSampleData()
	defaultOrgId := uuid.FromStringOrNil(folder.DefaultOrgID)

	t.Parallel()
	tests := [...]struct {
		testName    string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
	}{
		{
			testName:  "Folders with a matching orgID",
			orgID: defaultOrgId,
			folders: sampleData,
			want: sampleData[79:220], // items 79-220 are all the same folders that have the default orgID
		},
		{
			testName:  "No Folders match the given orgID",
			orgID: uuid.FromStringOrNil("123e4567-e89b-12d3-a456-426655440000"), // random unrelated uuid
			folders: sampleData,
			want: []folder.Folder{}, // so no folders returned
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get := f.GetFoldersByOrgID(tt.orgID)
			assert.Equal(t, tt.want, get)
		})
	}
}

func Test_folder_GetAllChildFolders(t *testing.T) {

	sampleData := folder.GetSampleData()
	defaultOrgId := uuid.FromStringOrNil(folder.DefaultOrgID)

	t.Parallel()
	tests := [...]struct {
		testName    string
		orgID   uuid.UUID
		givenFolder string
		folders []folder.Folder
		wantFolders    []folder.Folder
		wantError error
	}{
		{
			testName:  "Error: invalid orgID",
			orgID: uuid.FromStringOrNil("this will be invalid!"),
			givenFolder: "noble-vixen",
			folders: sampleData,
			wantFolders: []folder.Folder{},
			wantError: errors.New("invalid orgID, cannot be nil UUID"), 
		},
		{
			testName:  "Error: folder does not exist",
			orgID: defaultOrgId,
			givenFolder: "not-in-sample-data",
			folders: sampleData,
			wantFolders: []folder.Folder{},
			wantError: errors.New("Folder does not exist"), 
		},
		{
			testName:  "Error: folder exists, under a diff orgID",
			orgID: uuid.FromStringOrNil("123e4567-e89b-12d3-a456-426655440000"),
			givenFolder: "noble-vixen",
			folders: sampleData,
			wantFolders: []folder.Folder{},
			wantError: errors.New("Folder does not exist in the specified organization"), 
		},
		{
			testName:  "Root folder with multiple kids",
			orgID: defaultOrgId,
			givenFolder: "noble-vixen",
			folders: sampleData,
			wantFolders: sampleData[80:139], // Items 80-139 are all the items that have noble-vixen as a parent
			wantError: nil, 
		},
		{
			testName:  "Penultimate folder, one child",
			orgID: defaultOrgId,
			givenFolder: "super-cobweb",
			folders: sampleData,
			wantFolders: []folder.Folder{
				{
					Name: "perfect-vanisher",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "noble-vixen.fast-watchmen.growing-menace.super-cobweb.perfect-vanisher",
				},
			}, 
			wantError: nil, 
		},
		{
			testName:  "Final folder with no children",
			orgID: defaultOrgId,
			givenFolder: "amazing-bubbles",
			folders: sampleData,
			wantFolders: []folder.Folder{}, // No output, folder doesn't see itself
			wantError: nil, 
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			getFolders, getError := f.GetAllChildFolders(tt.orgID, tt.givenFolder)
			assert.Equal(t, tt.wantFolders, getFolders)
			assert.Equal(t, tt.wantError, getError)
		})
	}
}
