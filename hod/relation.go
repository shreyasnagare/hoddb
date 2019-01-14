package hod

import (
	"fmt"
	"github.com/RoaringBitmap/roaring"
	//"github.com/gtfierro/hod/storage"
)

type relation struct {
	rows []*relationRow

	multiindex map[string]map[EntityKey]*roaring.Bitmap

	// map variable name to position in row
	vars map[string]int
	keys []string
}

func newRelation(vars []string) *relation {
	rel := &relation{
		keys:       vars,
		vars:       make(map[string]int),
		multiindex: make(map[string]map[EntityKey]*roaring.Bitmap),
	}
	for idx, varname := range vars {
		rel.vars[varname] = idx
		rel.multiindex[varname] = make(map[EntityKey]*roaring.Bitmap)
	}
	return rel
}

func (rel *relation) done() {
	for _, row := range rel.rows {
		row.release()
	}
}

func (rel *relation) add1Value(key1 string, values entityset) {
	key1pos, found := rel.vars[key1]
	if !found {
		rel.vars[key1] = len(rel.vars)
		key1pos = rel.vars[key1]
		rel.multiindex[key1] = make(map[EntityKey]*roaring.Bitmap)
	}

	// For each value (for this variable), we want to check
	// if the bitmap is non-zero. If it is, then this value already
	// exists inside the relation. Otherwise, we can add it ourselves
	if len(rel.rows) == 0 {
		rel.rows = make([]*relationRow, 0, len(values))
	}
	for value := range values {
		bitmap := rel.multiindex[key1][value]

		// if this is non-nil, then the value exists already
		if bitmap != nil {
			return
		}

		row := newRelationRow()
		row.addValue(key1pos, value)
		rel.rows = append(rel.rows, row)
		// add the row to the multiindex
		rel.multiindex[key1][value] = roaring.New()
		rel.multiindex[key1][value].AddInt(len(rel.rows) - 1)
	}
}

func (rel *relation) add2Values(key1, key2 string, values [][]EntityKey) {
	key1pos, found := rel.vars[key1]
	if !found {
		rel.vars[key1] = len(rel.vars) + 1
		key1pos = rel.vars[key1]
		rel.multiindex[key1] = make(map[EntityKey]*roaring.Bitmap)
	}
	key2pos, found := rel.vars[key2]
	if !found {
		rel.vars[key2] = len(rel.vars) + 1
		key2pos = rel.vars[key2]
		rel.multiindex[key2] = make(map[EntityKey]*roaring.Bitmap)
	}

	if len(rel.rows) == 0 {
		rel.rows = make([]*relationRow, 0, len(values))
	}
	for _, valuepair := range values {
		bitmap1 := rel.multiindex[key1][valuepair[0]]
		bitmap2 := rel.multiindex[key2][valuepair[1]]

		// if the bitmaps are all non-nil, and the intersection is non-nil, then the value pair exists already
		if bitmap1 != nil && bitmap2 != nil && !roaring.And(bitmap1, bitmap2).IsEmpty() {
			continue
		}

		row := newRelationRow()
		row.addValue(key1pos, valuepair[0])
		row.addValue(key2pos, valuepair[1])
		rel.rows = append(rel.rows, row)

		if bitmap1 == nil {
			rel.multiindex[key1][valuepair[0]] = roaring.New()
		}
		if bitmap2 == nil {
			rel.multiindex[key2][valuepair[1]] = roaring.New()
		}

		// add the row to the multiindex
		rel.multiindex[key1][valuepair[0]].AddInt(len(rel.rows) - 1)
		rel.multiindex[key2][valuepair[1]].AddInt(len(rel.rows) - 1)
	}
}

func (rel *relation) add3Values(key1, key2, key3 string, values [][]EntityKey) {
	key1pos, found := rel.vars[key1]
	if !found {
		rel.vars[key1] = len(rel.vars) + 1
		key1pos = rel.vars[key1]
		rel.multiindex[key1] = make(map[EntityKey]*roaring.Bitmap)
	}
	key2pos, found := rel.vars[key2]
	if !found {
		rel.vars[key2] = len(rel.vars) + 1
		key2pos = rel.vars[key2]
		rel.multiindex[key2] = make(map[EntityKey]*roaring.Bitmap)
	}
	key3pos, found := rel.vars[key3]
	if !found {
		rel.vars[key3] = len(rel.vars) + 1
		key3pos = rel.vars[key3]
		rel.multiindex[key3] = make(map[EntityKey]*roaring.Bitmap)
	}

	if len(rel.rows) == 0 {
		rel.rows = make([]*relationRow, 0, len(values))
	}
	for _, valuepair := range values {
		bitmap1 := rel.multiindex[key1][valuepair[0]]
		bitmap2 := rel.multiindex[key2][valuepair[1]]
		bitmap3 := rel.multiindex[key3][valuepair[2]]

		// if the bitmaps are all non-nil, and the intersection is non-nil, then the value pair exists already
		if bitmap1 != nil && bitmap2 != nil && bitmap3 != nil && !roaring.FastAnd(bitmap1, bitmap2, bitmap3).IsEmpty() {
			continue
		}

		row := newRelationRow()
		row.addValue(key1pos, valuepair[0])
		row.addValue(key2pos, valuepair[1])
		row.addValue(key3pos, valuepair[2])
		rel.rows = append(rel.rows, row)

		if bitmap1 == nil {
			rel.multiindex[key1][valuepair[0]] = roaring.New()
		}
		if bitmap2 == nil {
			rel.multiindex[key2][valuepair[1]] = roaring.New()
		}
		if bitmap3 == nil {
			rel.multiindex[key3][valuepair[2]] = roaring.New()
		}

		// add the row to the multiindex
		rel.multiindex[key1][valuepair[0]].AddInt(len(rel.rows) - 1)
		rel.multiindex[key2][valuepair[1]].AddInt(len(rel.rows) - 1)
		rel.multiindex[key3][valuepair[2]].AddInt(len(rel.rows) - 1)
	}

}

func (rel *relation) join(other *relation, on []string, cursor *Cursor) {
	// get the variable positions for the join variables for
	// each of the relations (these may be different)

	var joinedRows = make([]*relationRow, 0, len(rel.rows))
innerRows:
	for _, innerRow := range rel.rows {
		// find all the rows in [other] that share the values
		var otherBitmaps []*roaring.Bitmap
		for _, joinVarName := range on {
			myVarPos := rel.vars[joinVarName]
			innerRowValue := innerRow.valueAt(myVarPos)
			if otherBitmap := other.multiindex[joinVarName][innerRowValue]; otherBitmap != nil {
				otherBitmaps = append(otherBitmaps, otherBitmap)
			} else {
				innerRow.release()
				continue innerRows // skip this row
			}
		}
		otherRowsBitmap := roaring.FastAnd(otherBitmaps...)
		if otherRowsBitmap.IsEmpty() {
			innerRow.release()
			continue innerRows // skip this row because there are no values to join
		}
		iter := otherRowsBitmap.Iterator()
		for iter.HasNext() {
			row := other.rows[iter.Next()]
			newRow := innerRow.copy()
			for otherVarname, otherIdx := range other.vars {
				newRow.addValue(rel.vars[otherVarname], row.valueAt(otherIdx))
			}

			joinedRows = append(joinedRows, newRow)
		}
		innerRow.release() // now done with this row
	}
	rel.rows = joinedRows
}

func (rel *relation) dumpRows(prefix string, cursor *Cursor) {
	for _, row := range rel.rows {
		//rel.dumpRow(row, ctx)
		//cursor.dumpRow(prefix, row)
		fmt.Println("}", row)
	}
}