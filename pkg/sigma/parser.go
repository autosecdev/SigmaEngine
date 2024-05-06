package sigma

import (
	"github.com/autosecdev/SigmaEngine/config"
	"github.com/autosecdev/SigmaEngine/utils"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"strings"
)

// GetRuleList returns a list of paths to Sigma rules
func GetRuleList(config config.SigmaConfig) (ruleList []Rule, err error) {

	err = filepath.Walk(config.RuleDir, func(
		path string,
		info os.FileInfo,
		err error,
	) error {
		if !info.IsDir() && strings.HasSuffix(path, "yml") {
			rule := Rule{}
			err = yaml.Unmarshal([]byte(utils.ReadFile(path)), &rule)
			if err != nil {
				return err
			}

			ruleList = append(ruleList, rule)
		}

		return err
	})

	return ruleList, nil
}
