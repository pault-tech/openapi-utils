
import (
	"fmt"
	"github.com/pb33f/libopenapi"
	"io/ioutil"
)

func main() {

	// load an OpenAPI 3 specification from bytes
	petstore, _ := ioutil.ReadFile("test_specs/petstorev3.json")

	// create a new document from specification bytes
	document, err := libopenapi.NewDocument(petstore)

	// if anything went wrong, an error is thrown
	if err != nil {
		panic(fmt.Sprintf("cannot create new document: %e", err))

	}

	// because we know this is a v3 spec,
	// we can build a ready to go model from it.
	v3Model, errors := document.BuildV3Model()

	// if anything went wrong when building the v3 model,
	// a slice of errors will be returned
	if len(errors) > 0 {
		for i := range errors {
			fmt.Printf("error: %e\n", errors[i])

		}
		panic(fmt.Sprintf("cannot create v3 model from "+
			"document: %d errors reported",
			len(errors)))

	}

	// get a count of the number of paths and schemas.
	paths := len(v3Model.Model.Paths.PathItems)
	schemas := len(v3Model.Model.Components.Schemas)

	// print the number of paths and schemas in the document
	fmt.Printf("There are %d paths and %d schemas "+
		"in the document", paths, schemas)

}
