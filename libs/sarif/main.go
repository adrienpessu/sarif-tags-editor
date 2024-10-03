package sarif

import (
	"encoding/json"
	"fmt"
)

type Sarif struct {
	Schema  string `json:"$schema"`
	Version string `json:"version"`
	Runs    []Runs `json:"runs"`
}
type Runs struct {
	Tool struct {
		Driver struct {
			Name            string `json:"name"`
			Organization    string `json:"organization"`
			SemanticVersion string `json:"semanticVersion"`
			Notifications   []struct {
				ID               string `json:"id"`
				Name             string `json:"name"`
				ShortDescription struct {
					Text string `json:"text"`
				} `json:"shortDescription"`
				FullDescription struct {
					Text string `json:"text"`
				} `json:"fullDescription"`
				DefaultConfiguration struct {
					Enabled bool `json:"enabled"`
				} `json:"defaultConfiguration"`
				Properties struct {
					Description string `json:"description"`
					ID          string `json:"id"`
					Kind        string `json:"kind"`
					Name        string `json:"name"`
				} `json:"properties,omitempty"`
				Properties0 struct {
					Tags        []string `json:"tags"`
					Description string   `json:"description"`
					ID          string   `json:"id"`
					Kind        string   `json:"kind"`
					Name        string   `json:"name"`
				} `json:"properties,omitempty"`
			} `json:"notifications"`
			Rules []struct {
				ID               string `json:"id"`
				Name             string `json:"name"`
				ShortDescription struct {
					Text string `json:"text"`
				} `json:"shortDescription"`
				FullDescription struct {
					Text string `json:"text"`
				} `json:"fullDescription"`
				DefaultConfiguration struct {
					Enabled bool   `json:"enabled"`
					Level   string `json:"level"`
				} `json:"defaultConfiguration,omitempty"`
				Help struct {
					Text     string `json:"text"`
					Markdown string `json:"markdown"`
				} `json:"help,omitempty"`
				Properties struct {
					Tags             []string `json:"tags"`
					Description      string   `json:"description"`
					ID               string   `json:"id"`
					Kind             string   `json:"kind"`
					Name             string   `json:"name"`
					Precision        string   `json:"precision"`
					ProblemSeverity  string   `json:"problem.severity"`
					SecuritySeverity string   `json:"security-severity"`
				} `json:"properties,omitempty"`
			} `json:"rules"`
		} `json:"driver"`
		Extensions []Extension `json:"extensions"`
	} `json:"tool"`
	Invocations []struct {
		ToolExecutionNotifications []struct {
			Message struct {
				Text string `json:"text"`
			} `json:"message"`
			Descriptor struct {
				ID    string `json:"id"`
				Index int    `json:"index"`
			} `json:"descriptor"`
			Properties struct {
				FormattedMessage struct {
					Text string `json:"text"`
				} `json:"formattedMessage"`
			} `json:"properties"`
			Locations []struct {
				PhysicalLocation struct {
					ArtifactLocation struct {
						URI       string `json:"uri"`
						URIBaseID string `json:"uriBaseId"`
						Index     int    `json:"index"`
					} `json:"artifactLocation"`
				} `json:"physicalLocation"`
			} `json:"locations,omitempty"`
			Level string `json:"level,omitempty"`
		} `json:"toolExecutionNotifications"`
		ExecutionSuccessful bool `json:"executionSuccessful"`
	} `json:"invocations"`
	Artifacts []struct {
		Location struct {
			URI       string `json:"uri"`
			URIBaseID string `json:"uriBaseId"`
			Index     int    `json:"index"`
		} `json:"location"`
	} `json:"artifacts"`
	Results           []Result `json:"results"`
	AutomationDetails struct {
		ID string `json:"id"`
	} `json:"automationDetails"`
	ColumnKind string `json:"columnKind"`
	Properties struct {
		MetricResults []struct {
			Rule struct {
				ID    string `json:"id"`
				Index int    `json:"index"`
			} `json:"rule"`
			RuleID    string `json:"ruleId"`
			RuleIndex int    `json:"ruleIndex"`
			Value     int    `json:"value"`
			Message   struct {
				Text string `json:"text"`
			} `json:"message,omitempty"`
			Baseline int `json:"baseline,omitempty"`
		} `json:"metricResults"`
		SemmleFormatSpecifier string `json:"semmle.formatSpecifier"`
	} `json:"properties"`
}

type Extension struct {
	Name            string          `json:"name"`
	SemanticVersion string          `json:"semanticVersion"`
	Rules           []ExtensionRule `json:"rules"`
	Locations       []struct {
		URI         string `json:"uri"`
		Description struct {
			Text string `json:"text"`
		} `json:"description"`
	} `json:"locations"`
}

type ExtensionRule struct {
	ID         string                  `json:"id"`
	Name       string                  `json:"name"`
	Properties ExtensionRuleProperties `json:"properties"`
}

type ExtensionRuleProperties struct {
	Tags            []string `json:"tags"`
	Description     string   `json:"description"`
	ID              string   `json:"id"`
	Kind            string   `json:"kind"`
	Name            string   `json:"name"`
	Precision       string   `json:"precision"`
	ProblemSeverity string   `json:"problem.severity"`
}

type Fixes struct {
	Description FixesDescription `json:"description"`
}

type FixesDescription struct {
	Text string `json:"text"`
}

type Result struct {
	RuleID    string `json:"ruleId"`
	RuleIndex int    `json:"ruleIndex"`
	Rule      struct {
		ID    string `json:"id"`
		Index int    `json:"index"`
	} `json:"rule"`
	Message struct {
		Text string `json:"text"`
	} `json:"message"`
	Locations []struct {
		PhysicalLocation struct {
			ArtifactLocation struct {
				URI       string `json:"uri"`
				URIBaseID string `json:"uriBaseId"`
				Index     int    `json:"index"`
			} `json:"artifactLocation"`
			Region struct {
				StartLine   int `json:"startLine"`
				StartColumn int `json:"startColumn"`
				EndColumn   int `json:"endColumn"`
			} `json:"region"`
		} `json:"physicalLocation"`
	} `json:"locations"`
	PartialFingerprints struct {
		PrimaryLocationLineHash               string `json:"primaryLocationLineHash"`
		PrimaryLocationStartColumnFingerprint string `json:"primaryLocationStartColumnFingerprint"`
	} `json:"partialFingerprints"`
	RelatedLocations []struct {
		ID               int `json:"id"`
		PhysicalLocation struct {
			ArtifactLocation struct {
				URI       string `json:"uri"`
				URIBaseID string `json:"uriBaseId"`
				Index     int    `json:"index"`
			} `json:"artifactLocation"`
			Region struct {
				StartLine   int `json:"startLine"`
				StartColumn int `json:"startColumn"`
				EndColumn   int `json:"endColumn"`
			} `json:"region"`
		} `json:"physicalLocation"`
		Message struct {
			Text string `json:"text"`
		} `json:"message"`
	} `json:"relatedLocations,omitempty"`
	CodeFlows []struct {
		ThreadFlows []struct {
			Locations []struct {
				Location struct {
					PhysicalLocation struct {
						ArtifactLocation struct {
							URI       string `json:"uri"`
							URIBaseID string `json:"uriBaseId"`
							Index     int    `json:"index"`
						} `json:"artifactLocation"`
						Region struct {
							StartLine   int `json:"startLine"`
							StartColumn int `json:"startColumn"`
							EndColumn   int `json:"endColumn"`
						} `json:"region"`
					} `json:"physicalLocation"`
					Message struct {
						Text string `json:"text"`
					} `json:"message"`
				} `json:"location"`
			} `json:"locations"`
		} `json:"threadFlows"`
	} `json:"codeFlows,omitempty"`
	Fixes []Fixes `json:"fixes,omitempty"`
}

func SearchAndAssignOrReplaceTags(sarifInput string, searchedTags []string, addTags []string, removeTags []string) string {
	var structuredSarif Sarif
	err2 := json.Unmarshal([]byte(sarifInput), &structuredSarif)
	if err2 != nil {
		fmt.Println(err2)
	}

	var runs = make([]Runs, 0)

	for _, run := range structuredSarif.Runs {
		extensionsToAdd := []Extension{}
		for _, extensions := range run.Tool.Extensions {
			rulesToAdd := []ExtensionRule{}
			for _, rule := range extensions.Rules {

				// if tags contains the searched tags
				tags := rule.Properties.Tags
				newTags := []string{}
				if contains(tags, searchedTags) {
					println("tags: ", tags)

					// remove the removeTags
					for _, tag := range tags {
						if !contains(removeTags, []string{tag}) {
							newTags = append(newTags, tag)
						}
					}
					// add the addTags
					newTags = append(newTags, addTags...)
					rule.Properties.Tags = newTags
				}
				rulesToAdd = append(rulesToAdd, rule)
			}
			extensions.Rules = rulesToAdd
			extensionsToAdd = append(extensionsToAdd, extensions)
		}
		run.Tool.Extensions = extensionsToAdd
		runs = append(runs, run)
	}
	structuredSarif.Runs = runs

	json, err := json.Marshal(structuredSarif)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	sarif := string(json)
	return sarif
}

// Define the contains function
func contains(tags []string, searchedTags []string) bool {
	for _, tag := range tags {
		for _, searchedTag := range searchedTags {
			if tag == searchedTag {
				return true
			}
		}
	}
	return false
}
