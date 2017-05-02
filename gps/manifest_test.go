// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gps

import "testing"

// Test that prep manifest sanitizes manifests appropriately
func TestPrepManifest(t *testing.T) {
	m := SimpleManifest{
		Deps: ProjectConstraints{
			ProjectRoot("foo"): ProjectProperties{},
			ProjectRoot("bar"): ProjectProperties{
				Source: "whatever",
			},
		},
		TestDeps: ProjectConstraints{
			ProjectRoot("baz"): ProjectProperties{},
			ProjectRoot("qux"): ProjectProperties{
				Source: "whatever",
			},
		},
	}

	prepped := prepManifest(m)
	d := prepped.DependencyConstraints()
	td := prepped.TestDependencyConstraints()
	if len(d) != 1 {
		t.Error("prepManifest did not eliminate empty ProjectProperties from deps map")
	}
	if len(td) != 1 {
		t.Error("prepManifest did not eliminate empty ProjectProperties from test deps map")
	}

	if d[ProjectRoot("bar")].Constraint != any {
		t.Error("prepManifest did not normalize nil constraint to anyConstraint in deps map")
	}
	if td[ProjectRoot("qux")].Constraint != any {
		t.Error("prepManifest did not normalize nil constraint to anyConstraint in test deps map")
	}
}
