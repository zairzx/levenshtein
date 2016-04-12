// Copyright 2016 ALRUX Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package levenshtein

// Params represents a set of parameter values for the various formulas involved
// in the calculation of the Levenshtein string metrics.
type Params struct {
	insCost        int
	subCost        int
	delCost        int
	maxCost        int
	minScore       float64
	bonusPrefix    int
	bonusScale     float64
	bonusThreshold float64
}

var (
	defaultParams = NewParams()
)

// NewParams creates a new set of parameters and initializes it with the default values.
func NewParams() *Params {
	return &Params{
		insCost:        1,
		subCost:        1,
		delCost:        1,
		maxCost:        0,
		minScore:       0,
		bonusPrefix:    4,
		bonusScale:     .1,
		bonusThreshold: .7,
	}
}

// Clone returns a pointer to a copy of the receiver parameter set, or of a new
// default parameter set if the receiver is nil.
func (this *Params) Clone() *Params {
	if this == nil {
		return NewParams()
	}
	return &Params{
		insCost:        this.insCost,
		subCost:        this.subCost,
		delCost:        this.delCost,
		maxCost:        this.maxCost,
		minScore:       this.minScore,
		bonusPrefix:    this.bonusPrefix,
		bonusScale:     this.bonusScale,
		bonusThreshold: this.bonusThreshold,
	}
}

// InsCost overrides the default value of 1 for the cost of insertion.
// The new value must be zero or positive.
func (this *Params) InsCost(v int) *Params {
	if v >= 0 {
		this.insCost = v
	}
	return this
}

// SubCost overrides the default value of 1 for the cost of substitution.
// The new value must be zero or positive.
func (this *Params) SubCost(v int) *Params {
	if v >= 0 {
		this.subCost = v
	}
	return this
}

// DelCost overrides the default value of 1 for the cost of deletion.
// The new value must be zero or positive.
func (this *Params) DelCost(v int) *Params {
	if v >= 0 {
		this.delCost = v
	}
	return this
}

// MaxCost overrides the default value of 0 (meaning unlimited) for the maximum cost.
// The calculation of Distance() stops when the result is guaranteed to exceed
// this maximum, returning a lower-bound rather than exact value.
// The new value must be zero or positive.
func (this *Params) MaxCost(v int) *Params {
	if v >= 0 {
		this.maxCost = v
	}
	return this
}

// MinScore overrides the default value of 0 for the minimum similarity score.
// Scores below this threshold are returned as 0 by Similarity() and Match().
// The new value must be zero or positive. Note that a minimum greater than 1
// can never be satisfied, resulting in a score of 0 for any pair of strings.
func (this *Params) MinScore(v float64) *Params {
	if v >= 0 {
		this.minScore = v
	}
	return this
}

// BonusPrefix overrides the default value for the maximum length of
// common prefix to be considered for bonus by Match().
// The new value must be zero or positive.
func (this *Params) BonusPrefix(v int) *Params {
	if v >= 0 {
		this.bonusPrefix = v
	}
	return this
}

// BonusScale overrides the default value for the scaling factor used by Match()
// in calculating the bonus.
// The new value must be zero or positive. To guarantee that the similarity score
// remains in the interval 0..1, this scaling factor is not allowed to exceed
// 1 / BonusPrefix.
func (this *Params) BonusScale(v float64) *Params {
	if v >= 0 {
		this.bonusScale = v
	}

	// the bonus cannot exceed (1-sim), or the score may become greater than 1.
	if float64(this.bonusPrefix)*this.bonusScale > 1 {
		this.bonusScale = 1 / float64(this.bonusPrefix)
	}

	return this
}

// BonusThreshold overrides the default value for the minimum similarity score
// for which Match() can assign a bonus.
// The new value must be zero or positive. Note that a threshold greater than 1
// effectively makes Match() become the equivalent of Similarity().
func (this *Params) BonusThreshold(v float64) *Params {
	if v >= 0 {
		this.bonusThreshold = v
	}
	return this
}
