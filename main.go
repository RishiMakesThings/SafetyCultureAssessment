package main

import (
	"fmt"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
)

func main() {
	orgID := uuid.FromStringOrNil(folder.DefaultOrgID)

	res := folder.GetAllFolders()

	// example usage
	folderDriver := folder.NewDriver(res)
	childFolders, error := folderDriver.GetAllChildFolders(orgID, "magnetic-sinister-six")

	folder.PrettyPrint(res[78:221])
	fmt.Printf("\n Folders for orgID: %s\n", orgID)

	if error != nil {
		fmt.Print(error, "\n")
	} else {
		folder.PrettyPrint(childFolders)
	}
}
