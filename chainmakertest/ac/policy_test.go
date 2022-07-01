package ac

import (
	"bufio"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"testing"

	"chainmaker.org/chainmaker/common/v2/json"
	"github.com/stretchr/testify/assert"
)

func TestPlicyDiff(t *testing.T) {
	//load policy.json
	policyJson, _ := ioutil.ReadFile("./policy.json")

	var resourcePolicy ResoucePolicy
	err := json.Unmarshal(policyJson, &resourcePolicy)
	assert.NoError(t, err)

	//load policy.txt
	var policyMap = make(map[string]struct{})
	file, err := os.Open("./policy.txt")
	assert.NoError(t, err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		policyMap[line] = struct{}{}
	}

	var notFound ResoucePolicy
	for _, policy := range resourcePolicy {
		if _, ok := policyMap[policy.ResourceName]; !ok {
			if !strings.Contains(policy.ResourceName, "GET") &&
				!strings.Contains(policy.ResourceName, "QUERY") &&
				!strings.Contains(policy.ResourceName, "test") &&
				!strings.Contains(policy.ResourceName, "TEST") {
				notFound = append(notFound, policy)
			}
		}
	}

	sort.Slice(notFound, func(i, j int) bool {
		if notFound[i].ResourceName < notFound[j].ResourceName {
			return true
		}
		return false
	})

	jsonBytes, err := json.MarshalIndent(notFound, "", " ")
	assert.NoError(t, err)

	//fmt.Println(string(jsonBytes))

	err = ioutil.WriteFile("./policy.diff", jsonBytes, fs.ModePerm)
	assert.NoError(t, err)

	fmt.Println(len(notFound))
}
