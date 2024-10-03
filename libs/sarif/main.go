package sarif

import (
	"encoding/json"
	"fmt"
)

type Sarif struct {
	Schema  string `json:"$schema,omitempty"`
	Version string `json:"version,omitempty"`
	Runs    []Runs `json:"runs,omitempty"`
}
type Runs struct {
	Tool struct {
		Driver struct {
			Name            string `json:"name,omitempty"`
			Organization    string `json:"organization,omitempty"`
			SemanticVersion string `json:"semanticVersion,omitempty"`
			Notifications   []struct {
				ID               string `json:"id,omitempty"`
				Name             string `json:"name,omitempty"`
				ShortDescription struct {
					Text string `json:"text,omitempty"`
				} `json:"shortDescription,omitempty"`
				FullDescription struct {
					Text string `json:"text,omitempty"`
				} `json:"fullDescription,omitempty"`
				DefaultConfiguration struct {
					Enabled bool `json:"enabled,omitempty"`
				} `json:"defaultConfiguration,omitempty"`
				Properties struct {
					Description string `json:"description,omitempty"`
					ID          string `json:"id,omitempty"`
					Kind        string `json:"kind,omitempty"`
					Name        string `json:"name,omitempty"`
				} `json:"properties,omitempty,omitempty"`
				Properties0 struct {
					Tags        []string `json:"tags,omitempty"`
					Description string   `json:"description,omitempty"`
					ID          string   `json:"id,omitempty"`
					Kind        string   `json:"kind,omitempty"`
					Name        string   `json:"name,omitempty"`
				} `json:"properties,omitempty"`
			} `json:"notifications,omitempty"`
			Rules []struct {
				ID               string `json:"id,omitempty"`
				Name             string `json:"name,omitempty"`
				ShortDescription struct {
					Text string `json:"text,omitempty"`
				} `json:"shortDescription,omitempty"`
				FullDescription struct {
					Text string `json:"text,omitempty"`
				} `json:"fullDescription,omitempty"`
				DefaultConfiguration struct {
					Enabled bool   `json:"enabled,omitempty"`
					Level   string `json:"level,omitempty"`
				} `json:"defaultConfiguration,omitempty"`
				Help struct {
					Text     string `json:"text,omitempty"`
					Markdown string `json:"markdown,omitempty"`
				} `json:"help,omitempty"`
				Properties struct {
					Tags             []string `json:"tags,omitempty"`
					Description      string   `json:"description,omitempty"`
					ID               string   `json:"id,omitempty"`
					Kind             string   `json:"kind,omitempty"`
					Name             string   `json:"name,omitempty"`
					Precision        string   `json:"precision,omitempty"`
					ProblemSeverity  string   `json:"problem.severity,omitempty"`
					SecuritySeverity string   `json:"security-severity,omitempty"`
				} `json:"properties,omitempty"`
			} `json:"rules,omitempty"`
		} `json:"driver"`
		Extensions []Extension `json:"extensions,omitempty"`
	} `json:"tool,omitempty"`
	Invocations []struct {
		ToolExecutionNotifications []struct {
			Message struct {
				Text string `json:"text,omitempty"`
			} `json:"message,omitempty"`
			Descriptor struct {
				ID    string `json:"id,omitempty"`
				Index int    `json:"index,omitempty"`
			} `json:"descriptor,omitempty"`
			Properties struct {
				FormattedMessage struct {
					Text string `json:"text,omitempty"`
				} `json:"formattedMessage,omitempty"`
			} `json:"properties,omitempty"`
			Locations []struct {
				PhysicalLocation struct {
					ArtifactLocation struct {
						URI       string `json:"uri,omitempty"`
						URIBaseID string `json:"uriBaseId,omitempty"`
						Index     int    `json:"index,omitempty"`
					} `json:"artifactLocation,omitempty"`
				} `json:"physicalLocation,omitempty"`
			} `json:"locations,omitempty,omitempty"`
			Level string `json:"level,omitempty,omitempty"`
		} `json:"toolExecutionNotifications,omitempty"`
		ExecutionSuccessful bool `json:"executionSuccessful,omitempty"`
	} `json:"invocations,omitempty"`
	Artifacts []struct {
		Location struct {
			URI       string `json:"uri,omitempty"`
			URIBaseID string `json:"uriBaseId,omitempty"`
			Index     int    `json:"index,omitempty"`
		} `json:"location,omitempty"`
	} `json:"artifacts,omitempty"`
	Results           []Result `json:"results,omitempty"`
	AutomationDetails struct {
		ID string `json:"id,omitempty"`
	} `json:"automationDetails,omitempty"`
	ColumnKind string `json:"columnKind,omitempty"`
	Properties struct {
		MetricResults []struct {
			Rule struct {
				ID    string `json:"id,omitempty"`
				Index int    `json:"index,omitempty"`
			} `json:"rule,omitempty"`
			RuleID    string `json:"ruleId,omitempty"`
			RuleIndex int    `json:"ruleIndex,omitempty"`
			Value     int    `json:"value,omitempty"`
			Message   struct {
				Text string `json:"text,omitempty"`
			} `json:"message,omitempty"`
			Baseline int `json:"baseline,omitempty"`
		} `json:"metricResults"`
		SemmleFormatSpecifier string `json:"semmle.formatSpecifier"`
	} `json:"properties,omitempty"`
}

type Extension struct {
	Name            string          `json:"name,omitempty"`
	SemanticVersion string          `json:"semanticVersion,omitempty"`
	Rules           []ExtensionRule `json:"rules,omitempty"`
	Locations       []struct {
		URI         string `json:"uri,omitempty"`
		Description struct {
			Text string `json:"text,omitempty"`
		} `json:"description,omitempty"`
	} `json:"locations,omitempty"`
}

type ExtensionRule struct {
	ID         string                  `json:"id,omitempty"`
	Name       string                  `json:"name,omitempty"`
	Properties ExtensionRuleProperties `json:"properties,omitempty"`
}

type ExtensionRuleProperties struct {
	Tags            []string `json:"tags,omitempty"`
	Description     string   `json:"description,omitempty"`
	ID              string   `json:"id,omitempty"`
	Kind            string   `json:"kind,omitempty"`
	Name            string   `json:"name,omitempty"`
	Precision       string   `json:"precision,omitempty"`
	ProblemSeverity string   `json:"problem.severity,omitempty"`
}

type Fixes struct {
	Description FixesDescription `json:"description,omitempty"`
}

type FixesDescription struct {
	Text string `json:"text,omitempty"`
}

type Result struct {
	RuleID    string `json:"ruleId,omitempty"`
	RuleIndex int    `json:"ruleIndex,omitempty"`
	Rule      struct {
		ID    string `json:"id,omitempty"`
		Index int    `json:"index,omitempty"`
	} `json:"rule,omitempty"`
	Message struct {
		Text string `json:"text,omitempty"`
	} `json:"message,omitempty"`
	Locations []struct {
		PhysicalLocation struct {
			ArtifactLocation struct {
				URI       string `json:"uri,omitempty"`
				URIBaseID string `json:"uriBaseId,omitempty"`
				Index     int    `json:"index,omitempty"`
			} `json:"artifactLocation,omitempty"`
			Region struct {
				StartLine   int `json:"startLine,omitempty"`
				StartColumn int `json:"startColumn,omitempty"`
				EndColumn   int `json:"endColumn,omitempty"`
			} `json:"region,omitempty"`
		} `json:"physicalLocation,omitempty"`
	} `json:"locations,omitempty"`
	PartialFingerprints struct {
		PrimaryLocationLineHash               string `json:"primaryLocationLineHash,omitempty"`
		PrimaryLocationStartColumnFingerprint string `json:"primaryLocationStartColumnFingerprint,omitempty"`
	} `json:"partialFingerprints,omitempty"`
	RelatedLocations []struct {
		ID               int `json:"id,omitempty"`
		PhysicalLocation struct {
			ArtifactLocation struct {
				URI       string `json:"uri,omitempty"`
				URIBaseID string `json:"uriBaseId,omitempty"`
				Index     int    `json:"index,omitempty"`
			} `json:"artifactLocation,omitempty"`
			Region struct {
				StartLine   int `json:"startLine,omitempty"`
				StartColumn int `json:"startColumn,omitempty"`
				EndColumn   int `json:"endColumn,omitempty"`
			} `json:"region,omitempty"`
		} `json:"physicalLocation,omitempty"`
		Message struct {
			Text string `json:"text,omitempty"`
		} `json:"message,omitempty"`
	} `json:"relatedLocations,omitempty,omitempty"`
	CodeFlows []struct {
		ThreadFlows []struct {
			Locations []struct {
				Location struct {
					PhysicalLocation struct {
						ArtifactLocation struct {
							URI       string `json:"uri,omitempty"`
							URIBaseID string `json:"uriBaseId,omitempty"`
							Index     int    `json:"index,omitempty"`
						} `json:"artifactLocation,omitempty"`
						Region struct {
							StartLine   int `json:"startLine,omitempty"`
							StartColumn int `json:"startColumn,omitempty"`
							EndColumn   int `json:"endColumn,omitempty"`
						} `json:"region,omitempty"`
					} `json:"physicalLocation,omitempty"`
					Message struct {
						Text string `json:"text,omitempty"`
					} `json:"message,omitempty"`
				} `json:"location,omitempty"`
			} `json:"locations,omitempty"`
		} `json:"threadFlows,omitempty"`
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
