package lib

import (
	"context"
	"fmt"
	"sort"

	"github.com/google/go-github/v49/github"
	semver "github.com/hashicorp/go-version"
)

// GetSemver : returns version that will be installed based on server constaint provided
func GetSemver(ctx context.Context, tgconstraint *string, ghClient *github.Client) (string, error) {

	tglist := GetAppList(ctx, ghClient)
	fmt.Printf("Reading required version from constraint: %s\n", *tgconstraint)
	tgversion, err := SemVerParser(tgconstraint, tglist)
	return tgversion, err
}

// ValidateSemVer : Goes through the list of terragrunt version, return a valid tf version for contraint provided
func SemVerParser(tfconstraint *string, tflist []string) (string, error) {
	tgversion := ""
	constraints, err := semver.NewConstraint(*tfconstraint) //NewConstraint returns a Constraints instance that a Version instance can be checked against
	if err != nil {
		return "", fmt.Errorf("error parsing constraint: %s", err)
	}
	versions := make([]*semver.Version, len(tflist))
	//put tgversion into semver object
	for i, tfvals := range tflist {
		version, err := semver.NewVersion(tfvals) //NewVersion parses a given version and returns an instance of Version or an error if unable to parse the version.
		if err != nil {
			return "", fmt.Errorf("error parsing constraint: %s", err)
		}
		versions[i] = version
	}

	sort.Sort(sort.Reverse(semver.Collection(versions)))

	for _, element := range versions {
		if constraints.Check(element) { // Validate a version against a constraint
			tgversion = element.String()
			fmt.Printf("Matched version: %s\n", tgversion)
			if ValidVersionFormat(tgversion) { //check if version format is correct
				return tgversion, nil
			}
		}
	}

	PrintInvalidTFVersion()
	return "", fmt.Errorf("error parsing constraint: %s", *tfconstraint)
}

// Print invalid TF version
func PrintInvalidTFVersion() {
	fmt.Println("Version does not exist or invalid terraform version format.\n Format should be #.#.# or #.#.#-@# where # are numbers and @ are word characters.\n For example, 0.11.7 and 0.11.9-beta1 are valid versions")
}
