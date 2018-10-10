// Copyright 2018 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package opt

// RuleName enumerates the names of all the optimizer rules. Manual rule names
// are defined in this file and rule names generated by Optgen are defined in
// rule_name.og.go.
type RuleName uint32

// Enumeration of all manual rule names.
const (
	InvalidRuleName RuleName = iota

	// ------------------------------------------------------------
	// Manual Rule Names
	// ------------------------------------------------------------

	SimplifyRootOrdering
	PruneRootCols
	SimplifyZeroCardinalityGroup

	// NumManualRules tracks the number of manually-defined rules.
	NumManualRuleNames
)

// IsNormalize returns true if r is a normalization rule.
func (r RuleName) IsNormalize() bool {
	return r < startExploreRule
}

// IsExplore returns true if r is an exploration rule.
func (r RuleName) IsExplore() bool {
	return r > startExploreRule
}

// Make linter happy.
var _ = InvalidRuleName
var _ = NumManualRuleNames
var _ = RuleName.IsNormalize
var _ = RuleName.IsExplore
