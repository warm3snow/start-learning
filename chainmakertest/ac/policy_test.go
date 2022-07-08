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

func TestPolicyFormat(t *testing.T) {

	//load policy.json
	policyJson, _ := ioutil.ReadFile("./pk_dpos.list")

	var resourcePolicy ResoucePolicy
	err := json.Unmarshal(policyJson, &resourcePolicy)
	assert.NoError(t, err)

	sort.Slice(resourcePolicy, func(i, j int) bool {
		if resourcePolicy[i].ResourceName < resourcePolicy[j].ResourceName {
			return true
		}
		return false
	})

	//in order to catalogue
	type policyDetail struct {
		contractName string
		methodName   string
		resourceName string
		policy       Policy
		desc         string
		other        string
	}
	var policyDetails []policyDetail
	for _, p := range resourcePolicy {
		contractName, methodName, resourceName := getNames(p.ResourceName)
		policyDetails = append(policyDetails, policyDetail{
			contractName: contractName,
			methodName:   methodName,
			resourceName: resourceName,
			policy:       p.Policy,
			desc:         "",
			other:        "",
		})
	}
	sort.Slice(policyDetails, func(i, j int) bool {
		if policyDetails[i].contractName < policyDetails[j].contractName {
			return true
		} else if policyDetails[i].contractName == policyDetails[j].contractName {
			if policyDetails[i].methodName < policyDetails[j].methodName {
				return true
			} else if policyDetails[i].methodName == policyDetails[j].methodName {
				return true
			}
		}
		return false
	})

	fmt.Printf("| 合约名 | 方法名 | 资源名 | 功能 | 默认权限 | 备注 | \n")
	fmt.Printf("|----------------|----------------|--------------------------------|----------------|----------------|----------------|\n")
	for _, policy := range policyDetails {
		fmt.Printf("| %s | %s | %s | %s | %s | %s |\n", policy.contractName, policy.methodName, policy.resourceName,
			policy.desc, policy.policy, policy.other)
	}
}

func getNames(name string) (contractName, methodName, resourceName string) {
	first_index := 0
	for i := 0; i < len(name); i++ {
		if name[i] == '-' {
			first_index = i
			break
		}
	}
	if first_index == 0 || first_index == len(name) {
		return "", "", name
	}

	return name[:first_index], name[first_index+1:], name
}
