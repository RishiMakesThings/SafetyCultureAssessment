package folder_test

import (
	"errors"
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_folder_MoveFolder(t *testing.T) {
	sampleData := folder.GetSampleData()

	t.Parallel()
	tests := [...]struct {
		testName    string
		sourceFolder string
		destinationFolder string
		folders	[]folder.Folder
		wantFolders    []folder.Folder
		wantError	error
	}{
		{
			testName: "Error: Source folder doesnt exist",
			sourceFolder: "not-in-sample-data", // folder that doesn't exist in sample data
			destinationFolder: "hip-stingray", // folder that does exist in sample data
			folders: sampleData,
			wantFolders: []folder.Folder{},
			wantError: errors.New("Error: Source folder does not exist"),
		},
		{
			testName: "Error: Destination folder doesnt exist",
			sourceFolder: "magnetic-sinister-six", // folder that does exist in sample data
			destinationFolder: "not-in-sample-data", // folder that doesn't exist in sample data
			folders: sampleData,
			wantFolders: []folder.Folder{},
			wantError: errors.New("Error: Destination folder does not exist"),
		},
		{
			testName: "Error: Moving into folder to itself",
			sourceFolder: "magnetic-sinister-six", 
			destinationFolder: "magnetic-sinister-six",
			folders: sampleData,
			wantFolders: []folder.Folder{},
			wantError: errors.New("Error: Cannot move a folder to itself"),
		},
		{
			testName: "Error: Moving into a folder from a different organization",
			sourceFolder: "topical-micromax", // OrgId of "38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"
			destinationFolder: "magnetic-sinister-six", // OrgId of "c1556e17-b7c0-45a3-a6ae-9546248fb17a"
			folders: sampleData,
			wantFolders: []folder.Folder{},
			wantError: errors.New("Error: Cannot move a folder to a different organization"),
		},
		{
			testName: "Error: Moving a folder to a child of itself",
			sourceFolder: "magnetic-sinister-six", 
			destinationFolder: "stirred-rainbow", // stirred rainbow is a child of magnetic-sinister-six
			folders: sampleData,
			wantFolders: []folder.Folder{},
			wantError: errors.New("Error: Cannot move a folder to a child of itself"),
		},
		{
			testName: "Succesful move",
			sourceFolder: "magnetic-sinister-six", 
			destinationFolder: "hip-stingray",
			folders: sampleData[81:90],
			wantFolders: []folder.Folder{ 
				// note that hipstingray itself does not change but magnetic-sinister-six and all of it's kids do
				{
					Name: "magnetic-sinister-six",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "noble-vixen.nearby-secret.hip-stingray.magnetic-sinister-six",
				},
				{
					Name: "stirred-rainbow",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "noble-vixen.nearby-secret.hip-stingray.magnetic-sinister-six.stirred-rainbow",
				},
				{
					Name: "smashing-abyss",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "noble-vixen.nearby-secret.hip-stingray.magnetic-sinister-six.stirred-rainbow.smashing-abyss",
				},
				{
					Name: "strong-spoiler",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "noble-vixen.nearby-secret.hip-stingray.magnetic-sinister-six.stirred-rainbow.strong-spoiler",
				},
				{
					Name: "warm-thunderball",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "noble-vixen.nearby-secret.hip-stingray.magnetic-sinister-six.stirred-rainbow.warm-thunderball",
				},
				{
					Name: "healthy-hiroim",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "noble-vixen.nearby-secret.hip-stingray.magnetic-sinister-six.healthy-hiroim",
				},
				{
					Name: "outgoing-network",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "noble-vixen.nearby-secret.hip-stingray.magnetic-sinister-six.healthy-hiroim.outgoing-network",
				},
				{
					Name: "social-wasp",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "noble-vixen.nearby-secret.hip-stingray.magnetic-sinister-six.healthy-hiroim.social-wasp",
				},
				{
					Name: "hip-stingray",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "noble-vixen.nearby-secret.hip-stingray",
				},			
			},
			wantError: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			getFolders, getError := f.MoveFolder(tt.sourceFolder, tt.destinationFolder)
			assert.Equal(t, tt.wantFolders, getFolders)
			assert.Equal(t, tt.wantError, getError)
		})
	}
}