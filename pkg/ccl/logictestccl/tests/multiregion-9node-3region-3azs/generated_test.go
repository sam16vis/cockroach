// Copyright 2022 The Cockroach Authors.
//
// Licensed as a CockroachDB Enterprise file under the Cockroach Community
// License (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
//     https://github.com/cockroachdb/cockroach/blob/master/licenses/CCL.txt

// Code generated by generate-logictest, DO NOT EDIT.

package testmultiregion_9node_3region_3azs

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/cockroachdb/cockroach/pkg/build/bazel"
	"github.com/cockroachdb/cockroach/pkg/ccl"
	"github.com/cockroachdb/cockroach/pkg/security/securityassets"
	"github.com/cockroachdb/cockroach/pkg/security/securitytest"
	"github.com/cockroachdb/cockroach/pkg/server"
	"github.com/cockroachdb/cockroach/pkg/sql/logictest"
	"github.com/cockroachdb/cockroach/pkg/testutils/serverutils"
	"github.com/cockroachdb/cockroach/pkg/testutils/skip"
	"github.com/cockroachdb/cockroach/pkg/testutils/testcluster"
	"github.com/cockroachdb/cockroach/pkg/util/leaktest"
	"github.com/cockroachdb/cockroach/pkg/util/randutil"
)

const configIdx = 13

var cclLogicTestDir string

func init() {
	if bazel.BuiltWithBazel() {
		var err error
		cclLogicTestDir, err = bazel.Runfile("pkg/ccl/logictestccl/testdata/logic_test")
		if err != nil {
			panic(err)
		}
	} else {
		cclLogicTestDir = "../../../../ccl/logictestccl/testdata/logic_test"
	}
}

func TestMain(m *testing.M) {
	defer ccl.TestingEnableEnterprise()()
	securityassets.SetLoader(securitytest.EmbeddedAssets)
	randutil.SeedForTests()
	serverutils.InitTestServerFactory(server.TestServerFactory)
	serverutils.InitTestClusterFactory(testcluster.TestClusterFactory)
	os.Exit(m.Run())
}

func runCCLLogicTest(t *testing.T, file string) {
	skip.UnderDeadlock(t, "times out and/or hangs")
	logictest.RunLogicTest(t, logictest.TestServerArgs{}, configIdx, filepath.Join(cclLogicTestDir, file))
}

// TestLogic_tmp runs any tests that are prefixed with "_", in which a dedicated
// test is not generated for. This allows developers to create and run temporary
// test files that are not checked into the repository, without repeatedly
// regenerating and reverting changes to this file, generated_test.go.
//
// TODO(mgartner): Add file filtering so that individual files can be run,
// instead of all files with the "_" prefix.
func TestLogic_tmp(t *testing.T) {
	defer leaktest.AfterTest(t)()
	var glob string
	glob = filepath.Join(cclLogicTestDir, "_*")
	logictest.RunLogicTests(t, logictest.TestServerArgs{}, configIdx, glob)
}

func TestCCLLogic_alter_table_locality(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "alter_table_locality")
}

func TestCCLLogic_global_placement_restricted(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "global_placement_restricted")
}

func TestCCLLogic_multi_region(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "multi_region")
}

func TestCCLLogic_multi_region_backup(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "multi_region_backup")
}

func TestCCLLogic_multi_region_default_primary_region(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "multi_region_default_primary_region")
}

func TestCCLLogic_multi_region_drop_region(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "multi_region_drop_region")
}

func TestCCLLogic_multi_region_foreign_key_lookup_join(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "multi_region_foreign_key_lookup_join")
}

func TestCCLLogic_multi_region_import_export(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "multi_region_import_export")
}

func TestCCLLogic_multi_region_locality_optimized_search_query_behavior(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "multi_region_locality_optimized_search_query_behavior")
}

func TestCCLLogic_multi_region_privileges(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "multi_region_privileges")
}

func TestCCLLogic_multi_region_query_behavior(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "multi_region_query_behavior")
}

func TestCCLLogic_multi_region_remote_access_error(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "multi_region_remote_access_error")
}

func TestCCLLogic_multi_region_show(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "multi_region_show")
}

func TestCCLLogic_multi_region_zone_config_extensions(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "multi_region_zone_config_extensions")
}

func TestCCLLogic_multi_region_zone_configs(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "multi_region_zone_configs")
}

func TestCCLLogic_placement(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "placement")
}

func TestCCLLogic_regional_by_row(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "regional_by_row")
}

func TestCCLLogic_regional_by_row_auto_rehoming(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "regional_by_row_auto_rehoming")
}

func TestCCLLogic_regional_by_row_hash_sharded_index(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "regional_by_row_hash_sharded_index")
}

func TestCCLLogic_regional_by_row_hash_sharded_index_query_plan(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "regional_by_row_hash_sharded_index_query_plan")
}

func TestCCLLogic_regional_by_row_placement_restricted(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "regional_by_row_placement_restricted")
}

func TestCCLLogic_regional_by_row_query_behavior(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "regional_by_row_query_behavior")
}

func TestCCLLogic_regional_by_row_rename_column(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "regional_by_row_rename_column")
}

func TestCCLLogic_regional_by_table_placement_restricted(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "regional_by_table_placement_restricted")
}
