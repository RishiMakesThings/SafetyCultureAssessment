package main

import (
	"fmt"

	"github.com/georgechieng-sc/interns-2022/folder"
)

func main() {
	//orgID := uuid.FromStringOrNil(folder.DefaultOrgID)

	res := folder.GetAllFolders()

	// example usage
	//folderDriver := folder.NewDriver(res[78:90])
	//childFolders, error := folderDriver.GetAllChildFolders(orgID, "magnetic-sinister-six")

	fmt.Printf("\n Folders for Child Magnetic Sinister Six:\n")
	folder.PrettyPrint(res[78:90])

	fmt.Printf("\n After Move: \n")

	childDriver := folder.NewDriver(res[78:90])
	movedFolders, error := childDriver.MoveFolder("magnetic-sinister-six", "hip-stingray")
	folder.PrettyPrint(movedFolders)
	fmt.Print(error, "\n")
}
