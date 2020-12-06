// Code generated by execgen; DO NOT EDIT.
// Copyright 2018 The Cockroach Authors.
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package colexecagg

import (
	"strings"
	"unsafe"

	"github.com/cockroachdb/cockroach/pkg/col/coldata"
	"github.com/cockroachdb/cockroach/pkg/sql/colexecbase/colexecerror"
	"github.com/cockroachdb/cockroach/pkg/sql/colmem"
	"github.com/cockroachdb/cockroach/pkg/sql/sem/tree"
	"github.com/cockroachdb/cockroach/pkg/sql/types"
	"github.com/cockroachdb/errors"
)

func newSumIntOrderedAggAlloc(
	allocator *colmem.Allocator, t *types.T, allocSize int64,
) (aggregateFuncAlloc, error) {
	allocBase := aggAllocBase{allocator: allocator, allocSize: allocSize}
	switch t.Family() {
	case types.IntFamily:
		switch t.Width() {
		case 16:
			return &sumIntInt16OrderedAggAlloc{aggAllocBase: allocBase}, nil
		case 32:
			return &sumIntInt32OrderedAggAlloc{aggAllocBase: allocBase}, nil
		default:
			return &sumIntInt64OrderedAggAlloc{aggAllocBase: allocBase}, nil
		}
	default:
		return nil, errors.Errorf("unsupported sum %s agg type %s", strings.ToLower("Int"), t.Name())
	}
}

type sumIntInt16OrderedAgg struct {
	orderedAggregateFuncBase
	// curAgg holds the running total, so we can index into the slice once per
	// group, instead of on each iteration.
	curAgg int64
	// col points to the output vector we are updating.
	col []int64
	// foundNonNullForCurrentGroup tracks if we have seen any non-null values
	// for the group that is currently being aggregated.
	foundNonNullForCurrentGroup bool
}

var _ AggregateFunc = &sumIntInt16OrderedAgg{}

func (a *sumIntInt16OrderedAgg) SetOutput(vec coldata.Vec) {
	a.orderedAggregateFuncBase.SetOutput(vec)
	a.col = vec.Int64()
}

func (a *sumIntInt16OrderedAgg) Compute(
	vecs []coldata.Vec, inputIdxs []uint32, inputLen int, sel []int,
) {
	var oldCurAggSize uintptr
	vec := vecs[inputIdxs[0]]
	col, nulls := vec.Int16(), vec.Nulls()
	a.allocator.PerformOperation([]coldata.Vec{a.vec}, func() {
		// Capture col to force bounds check to work. See
		// https://github.com/golang/go/issues/39756
		col := col
		groups := a.groups
		if sel == nil {
			_ = groups[inputLen-1]
			col = col[:inputLen]
			if nulls.MaybeHasNulls() {
				for i := range col {

					if groups[i] {
						// If we encounter a new group, and we haven't found any non-nulls for the
						// current group, the output for this group should be null.
						if !a.foundNonNullForCurrentGroup {
							a.nulls.SetNull(a.curIdx)
						} else {
							a.col[a.curIdx] = a.curAgg
						}
						a.curIdx++
						a.curAgg = zeroInt64Value

						a.foundNonNullForCurrentGroup = false
					}

					var isNull bool
					isNull = nulls.NullAt(i)
					if !isNull {

						{
							result := int64(a.curAgg) + int64(col[i])
							if (result < int64(a.curAgg)) != (int64(col[i]) < 0) {
								colexecerror.ExpectedError(tree.ErrIntOutOfRange)
							}
							a.curAgg = result
						}

						a.foundNonNullForCurrentGroup = true
					}
				}
			} else {
				for i := range col {

					if groups[i] {
						// If we encounter a new group, and we haven't found any non-nulls for the
						// current group, the output for this group should be null.
						if !a.foundNonNullForCurrentGroup {
							a.nulls.SetNull(a.curIdx)
						} else {
							a.col[a.curIdx] = a.curAgg
						}
						a.curIdx++
						a.curAgg = zeroInt64Value

					}

					var isNull bool
					isNull = false
					if !isNull {

						{
							result := int64(a.curAgg) + int64(col[i])
							if (result < int64(a.curAgg)) != (int64(col[i]) < 0) {
								colexecerror.ExpectedError(tree.ErrIntOutOfRange)
							}
							a.curAgg = result
						}

						a.foundNonNullForCurrentGroup = true
					}
				}
			}
		} else {
			sel = sel[:inputLen]
			if nulls.MaybeHasNulls() {
				for _, i := range sel {

					if groups[i] {
						// If we encounter a new group, and we haven't found any non-nulls for the
						// current group, the output for this group should be null.
						if !a.foundNonNullForCurrentGroup {
							a.nulls.SetNull(a.curIdx)
						} else {
							a.col[a.curIdx] = a.curAgg
						}
						a.curIdx++
						a.curAgg = zeroInt64Value

						a.foundNonNullForCurrentGroup = false
					}

					var isNull bool
					isNull = nulls.NullAt(i)
					if !isNull {

						{
							result := int64(a.curAgg) + int64(col[i])
							if (result < int64(a.curAgg)) != (int64(col[i]) < 0) {
								colexecerror.ExpectedError(tree.ErrIntOutOfRange)
							}
							a.curAgg = result
						}

						a.foundNonNullForCurrentGroup = true
					}
				}
			} else {
				for _, i := range sel {

					if groups[i] {
						// If we encounter a new group, and we haven't found any non-nulls for the
						// current group, the output for this group should be null.
						if !a.foundNonNullForCurrentGroup {
							a.nulls.SetNull(a.curIdx)
						} else {
							a.col[a.curIdx] = a.curAgg
						}
						a.curIdx++
						a.curAgg = zeroInt64Value

					}

					var isNull bool
					isNull = false
					if !isNull {

						{
							result := int64(a.curAgg) + int64(col[i])
							if (result < int64(a.curAgg)) != (int64(col[i]) < 0) {
								colexecerror.ExpectedError(tree.ErrIntOutOfRange)
							}
							a.curAgg = result
						}

						a.foundNonNullForCurrentGroup = true
					}
				}
			}
		}
	},
	)
	var newCurAggSize uintptr
	if newCurAggSize != oldCurAggSize {
		a.allocator.AdjustMemoryUsage(int64(newCurAggSize - oldCurAggSize))
	}
}

func (a *sumIntInt16OrderedAgg) Flush(outputIdx int) {
	// The aggregation is finished. Flush the last value. If we haven't found
	// any non-nulls for this group so far, the output for this group should be
	// null.
	// Go around "argument overwritten before first use" linter error.
	_ = outputIdx
	outputIdx = a.curIdx
	a.curIdx++
	if !a.foundNonNullForCurrentGroup {
		a.nulls.SetNull(outputIdx)
	} else {
		a.col[outputIdx] = a.curAgg
	}
}

type sumIntInt16OrderedAggAlloc struct {
	aggAllocBase
	aggFuncs []sumIntInt16OrderedAgg
}

var _ aggregateFuncAlloc = &sumIntInt16OrderedAggAlloc{}

const sizeOfSumIntInt16OrderedAgg = int64(unsafe.Sizeof(sumIntInt16OrderedAgg{}))
const sumIntInt16OrderedAggSliceOverhead = int64(unsafe.Sizeof([]sumIntInt16OrderedAgg{}))

func (a *sumIntInt16OrderedAggAlloc) newAggFunc() AggregateFunc {
	if len(a.aggFuncs) == 0 {
		a.allocator.AdjustMemoryUsage(sumIntInt16OrderedAggSliceOverhead + sizeOfSumIntInt16OrderedAgg*a.allocSize)
		a.aggFuncs = make([]sumIntInt16OrderedAgg, a.allocSize)
	}
	f := &a.aggFuncs[0]
	f.allocator = a.allocator
	a.aggFuncs = a.aggFuncs[1:]
	return f
}

type sumIntInt32OrderedAgg struct {
	orderedAggregateFuncBase
	// curAgg holds the running total, so we can index into the slice once per
	// group, instead of on each iteration.
	curAgg int64
	// col points to the output vector we are updating.
	col []int64
	// foundNonNullForCurrentGroup tracks if we have seen any non-null values
	// for the group that is currently being aggregated.
	foundNonNullForCurrentGroup bool
}

var _ AggregateFunc = &sumIntInt32OrderedAgg{}

func (a *sumIntInt32OrderedAgg) SetOutput(vec coldata.Vec) {
	a.orderedAggregateFuncBase.SetOutput(vec)
	a.col = vec.Int64()
}

func (a *sumIntInt32OrderedAgg) Compute(
	vecs []coldata.Vec, inputIdxs []uint32, inputLen int, sel []int,
) {
	var oldCurAggSize uintptr
	vec := vecs[inputIdxs[0]]
	col, nulls := vec.Int32(), vec.Nulls()
	a.allocator.PerformOperation([]coldata.Vec{a.vec}, func() {
		// Capture col to force bounds check to work. See
		// https://github.com/golang/go/issues/39756
		col := col
		groups := a.groups
		if sel == nil {
			_ = groups[inputLen-1]
			col = col[:inputLen]
			if nulls.MaybeHasNulls() {
				for i := range col {

					if groups[i] {
						// If we encounter a new group, and we haven't found any non-nulls for the
						// current group, the output for this group should be null.
						if !a.foundNonNullForCurrentGroup {
							a.nulls.SetNull(a.curIdx)
						} else {
							a.col[a.curIdx] = a.curAgg
						}
						a.curIdx++
						a.curAgg = zeroInt64Value

						a.foundNonNullForCurrentGroup = false
					}

					var isNull bool
					isNull = nulls.NullAt(i)
					if !isNull {

						{
							result := int64(a.curAgg) + int64(col[i])
							if (result < int64(a.curAgg)) != (int64(col[i]) < 0) {
								colexecerror.ExpectedError(tree.ErrIntOutOfRange)
							}
							a.curAgg = result
						}

						a.foundNonNullForCurrentGroup = true
					}
				}
			} else {
				for i := range col {

					if groups[i] {
						// If we encounter a new group, and we haven't found any non-nulls for the
						// current group, the output for this group should be null.
						if !a.foundNonNullForCurrentGroup {
							a.nulls.SetNull(a.curIdx)
						} else {
							a.col[a.curIdx] = a.curAgg
						}
						a.curIdx++
						a.curAgg = zeroInt64Value

					}

					var isNull bool
					isNull = false
					if !isNull {

						{
							result := int64(a.curAgg) + int64(col[i])
							if (result < int64(a.curAgg)) != (int64(col[i]) < 0) {
								colexecerror.ExpectedError(tree.ErrIntOutOfRange)
							}
							a.curAgg = result
						}

						a.foundNonNullForCurrentGroup = true
					}
				}
			}
		} else {
			sel = sel[:inputLen]
			if nulls.MaybeHasNulls() {
				for _, i := range sel {

					if groups[i] {
						// If we encounter a new group, and we haven't found any non-nulls for the
						// current group, the output for this group should be null.
						if !a.foundNonNullForCurrentGroup {
							a.nulls.SetNull(a.curIdx)
						} else {
							a.col[a.curIdx] = a.curAgg
						}
						a.curIdx++
						a.curAgg = zeroInt64Value

						a.foundNonNullForCurrentGroup = false
					}

					var isNull bool
					isNull = nulls.NullAt(i)
					if !isNull {

						{
							result := int64(a.curAgg) + int64(col[i])
							if (result < int64(a.curAgg)) != (int64(col[i]) < 0) {
								colexecerror.ExpectedError(tree.ErrIntOutOfRange)
							}
							a.curAgg = result
						}

						a.foundNonNullForCurrentGroup = true
					}
				}
			} else {
				for _, i := range sel {

					if groups[i] {
						// If we encounter a new group, and we haven't found any non-nulls for the
						// current group, the output for this group should be null.
						if !a.foundNonNullForCurrentGroup {
							a.nulls.SetNull(a.curIdx)
						} else {
							a.col[a.curIdx] = a.curAgg
						}
						a.curIdx++
						a.curAgg = zeroInt64Value

					}

					var isNull bool
					isNull = false
					if !isNull {

						{
							result := int64(a.curAgg) + int64(col[i])
							if (result < int64(a.curAgg)) != (int64(col[i]) < 0) {
								colexecerror.ExpectedError(tree.ErrIntOutOfRange)
							}
							a.curAgg = result
						}

						a.foundNonNullForCurrentGroup = true
					}
				}
			}
		}
	},
	)
	var newCurAggSize uintptr
	if newCurAggSize != oldCurAggSize {
		a.allocator.AdjustMemoryUsage(int64(newCurAggSize - oldCurAggSize))
	}
}

func (a *sumIntInt32OrderedAgg) Flush(outputIdx int) {
	// The aggregation is finished. Flush the last value. If we haven't found
	// any non-nulls for this group so far, the output for this group should be
	// null.
	// Go around "argument overwritten before first use" linter error.
	_ = outputIdx
	outputIdx = a.curIdx
	a.curIdx++
	if !a.foundNonNullForCurrentGroup {
		a.nulls.SetNull(outputIdx)
	} else {
		a.col[outputIdx] = a.curAgg
	}
}

type sumIntInt32OrderedAggAlloc struct {
	aggAllocBase
	aggFuncs []sumIntInt32OrderedAgg
}

var _ aggregateFuncAlloc = &sumIntInt32OrderedAggAlloc{}

const sizeOfSumIntInt32OrderedAgg = int64(unsafe.Sizeof(sumIntInt32OrderedAgg{}))
const sumIntInt32OrderedAggSliceOverhead = int64(unsafe.Sizeof([]sumIntInt32OrderedAgg{}))

func (a *sumIntInt32OrderedAggAlloc) newAggFunc() AggregateFunc {
	if len(a.aggFuncs) == 0 {
		a.allocator.AdjustMemoryUsage(sumIntInt32OrderedAggSliceOverhead + sizeOfSumIntInt32OrderedAgg*a.allocSize)
		a.aggFuncs = make([]sumIntInt32OrderedAgg, a.allocSize)
	}
	f := &a.aggFuncs[0]
	f.allocator = a.allocator
	a.aggFuncs = a.aggFuncs[1:]
	return f
}

type sumIntInt64OrderedAgg struct {
	orderedAggregateFuncBase
	// curAgg holds the running total, so we can index into the slice once per
	// group, instead of on each iteration.
	curAgg int64
	// col points to the output vector we are updating.
	col []int64
	// foundNonNullForCurrentGroup tracks if we have seen any non-null values
	// for the group that is currently being aggregated.
	foundNonNullForCurrentGroup bool
}

var _ AggregateFunc = &sumIntInt64OrderedAgg{}

func (a *sumIntInt64OrderedAgg) SetOutput(vec coldata.Vec) {
	a.orderedAggregateFuncBase.SetOutput(vec)
	a.col = vec.Int64()
}

func (a *sumIntInt64OrderedAgg) Compute(
	vecs []coldata.Vec, inputIdxs []uint32, inputLen int, sel []int,
) {
	var oldCurAggSize uintptr
	vec := vecs[inputIdxs[0]]
	col, nulls := vec.Int64(), vec.Nulls()
	a.allocator.PerformOperation([]coldata.Vec{a.vec}, func() {
		// Capture col to force bounds check to work. See
		// https://github.com/golang/go/issues/39756
		col := col
		groups := a.groups
		if sel == nil {
			_ = groups[inputLen-1]
			col = col[:inputLen]
			if nulls.MaybeHasNulls() {
				for i := range col {

					if groups[i] {
						// If we encounter a new group, and we haven't found any non-nulls for the
						// current group, the output for this group should be null.
						if !a.foundNonNullForCurrentGroup {
							a.nulls.SetNull(a.curIdx)
						} else {
							a.col[a.curIdx] = a.curAgg
						}
						a.curIdx++
						a.curAgg = zeroInt64Value

						a.foundNonNullForCurrentGroup = false
					}

					var isNull bool
					isNull = nulls.NullAt(i)
					if !isNull {

						{
							result := int64(a.curAgg) + int64(col[i])
							if (result < int64(a.curAgg)) != (int64(col[i]) < 0) {
								colexecerror.ExpectedError(tree.ErrIntOutOfRange)
							}
							a.curAgg = result
						}

						a.foundNonNullForCurrentGroup = true
					}
				}
			} else {
				for i := range col {

					if groups[i] {
						// If we encounter a new group, and we haven't found any non-nulls for the
						// current group, the output for this group should be null.
						if !a.foundNonNullForCurrentGroup {
							a.nulls.SetNull(a.curIdx)
						} else {
							a.col[a.curIdx] = a.curAgg
						}
						a.curIdx++
						a.curAgg = zeroInt64Value

					}

					var isNull bool
					isNull = false
					if !isNull {

						{
							result := int64(a.curAgg) + int64(col[i])
							if (result < int64(a.curAgg)) != (int64(col[i]) < 0) {
								colexecerror.ExpectedError(tree.ErrIntOutOfRange)
							}
							a.curAgg = result
						}

						a.foundNonNullForCurrentGroup = true
					}
				}
			}
		} else {
			sel = sel[:inputLen]
			if nulls.MaybeHasNulls() {
				for _, i := range sel {

					if groups[i] {
						// If we encounter a new group, and we haven't found any non-nulls for the
						// current group, the output for this group should be null.
						if !a.foundNonNullForCurrentGroup {
							a.nulls.SetNull(a.curIdx)
						} else {
							a.col[a.curIdx] = a.curAgg
						}
						a.curIdx++
						a.curAgg = zeroInt64Value

						a.foundNonNullForCurrentGroup = false
					}

					var isNull bool
					isNull = nulls.NullAt(i)
					if !isNull {

						{
							result := int64(a.curAgg) + int64(col[i])
							if (result < int64(a.curAgg)) != (int64(col[i]) < 0) {
								colexecerror.ExpectedError(tree.ErrIntOutOfRange)
							}
							a.curAgg = result
						}

						a.foundNonNullForCurrentGroup = true
					}
				}
			} else {
				for _, i := range sel {

					if groups[i] {
						// If we encounter a new group, and we haven't found any non-nulls for the
						// current group, the output for this group should be null.
						if !a.foundNonNullForCurrentGroup {
							a.nulls.SetNull(a.curIdx)
						} else {
							a.col[a.curIdx] = a.curAgg
						}
						a.curIdx++
						a.curAgg = zeroInt64Value

					}

					var isNull bool
					isNull = false
					if !isNull {

						{
							result := int64(a.curAgg) + int64(col[i])
							if (result < int64(a.curAgg)) != (int64(col[i]) < 0) {
								colexecerror.ExpectedError(tree.ErrIntOutOfRange)
							}
							a.curAgg = result
						}

						a.foundNonNullForCurrentGroup = true
					}
				}
			}
		}
	},
	)
	var newCurAggSize uintptr
	if newCurAggSize != oldCurAggSize {
		a.allocator.AdjustMemoryUsage(int64(newCurAggSize - oldCurAggSize))
	}
}

func (a *sumIntInt64OrderedAgg) Flush(outputIdx int) {
	// The aggregation is finished. Flush the last value. If we haven't found
	// any non-nulls for this group so far, the output for this group should be
	// null.
	// Go around "argument overwritten before first use" linter error.
	_ = outputIdx
	outputIdx = a.curIdx
	a.curIdx++
	if !a.foundNonNullForCurrentGroup {
		a.nulls.SetNull(outputIdx)
	} else {
		a.col[outputIdx] = a.curAgg
	}
}

type sumIntInt64OrderedAggAlloc struct {
	aggAllocBase
	aggFuncs []sumIntInt64OrderedAgg
}

var _ aggregateFuncAlloc = &sumIntInt64OrderedAggAlloc{}

const sizeOfSumIntInt64OrderedAgg = int64(unsafe.Sizeof(sumIntInt64OrderedAgg{}))
const sumIntInt64OrderedAggSliceOverhead = int64(unsafe.Sizeof([]sumIntInt64OrderedAgg{}))

func (a *sumIntInt64OrderedAggAlloc) newAggFunc() AggregateFunc {
	if len(a.aggFuncs) == 0 {
		a.allocator.AdjustMemoryUsage(sumIntInt64OrderedAggSliceOverhead + sizeOfSumIntInt64OrderedAgg*a.allocSize)
		a.aggFuncs = make([]sumIntInt64OrderedAgg, a.allocSize)
	}
	f := &a.aggFuncs[0]
	f.allocator = a.allocator
	a.aggFuncs = a.aggFuncs[1:]
	return f
}
