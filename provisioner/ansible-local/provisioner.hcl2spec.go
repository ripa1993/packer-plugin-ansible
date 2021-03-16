// Code generated by "packer-sdc mapstructure-to-hcl2"; DO NOT EDIT.

package ansiblelocal

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName       *string           `mapstructure:"packer_build_name" cty:"packer_build_name" hcl:"packer_build_name"`
	PackerBuilderType     *string           `mapstructure:"packer_builder_type" cty:"packer_builder_type" hcl:"packer_builder_type"`
	PackerCoreVersion     *string           `mapstructure:"packer_core_version" cty:"packer_core_version" hcl:"packer_core_version"`
	PackerDebug           *bool             `mapstructure:"packer_debug" cty:"packer_debug" hcl:"packer_debug"`
	PackerForce           *bool             `mapstructure:"packer_force" cty:"packer_force" hcl:"packer_force"`
	PackerOnError         *string           `mapstructure:"packer_on_error" cty:"packer_on_error" hcl:"packer_on_error"`
	PackerUserVars        map[string]string `mapstructure:"packer_user_variables" cty:"packer_user_variables" hcl:"packer_user_variables"`
	PackerSensitiveVars   []string          `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables" hcl:"packer_sensitive_variables"`
	Command               *string           `cty:"command" hcl:"command"`
	ExtraArguments        []string          `mapstructure:"extra_arguments" cty:"extra_arguments" hcl:"extra_arguments"`
	GroupVars             *string           `mapstructure:"group_vars" cty:"group_vars" hcl:"group_vars"`
	HostVars              *string           `mapstructure:"host_vars" cty:"host_vars" hcl:"host_vars"`
	PlaybookDir           *string           `mapstructure:"playbook_dir" cty:"playbook_dir" hcl:"playbook_dir"`
	PlaybookFile          *string           `mapstructure:"playbook_file" cty:"playbook_file" hcl:"playbook_file"`
	PlaybookFiles         []string          `mapstructure:"playbook_files" cty:"playbook_files" hcl:"playbook_files"`
	PlaybookPaths         []string          `mapstructure:"playbook_paths" cty:"playbook_paths" hcl:"playbook_paths"`
	RolePaths             []string          `mapstructure:"role_paths" cty:"role_paths" hcl:"role_paths"`
	CollectionPaths       []string          `mapstructure:"collection_paths" cty:"collection_paths" hcl:"collection_paths"`
	StagingDir            *string           `mapstructure:"staging_directory" cty:"staging_directory" hcl:"staging_directory"`
	CleanStagingDir       *bool             `mapstructure:"clean_staging_directory" cty:"clean_staging_directory" hcl:"clean_staging_directory"`
	InventoryFile         *string           `mapstructure:"inventory_file" cty:"inventory_file" hcl:"inventory_file"`
	InventoryGroups       []string          `mapstructure:"inventory_groups" cty:"inventory_groups" hcl:"inventory_groups"`
	GalaxyFile            *string           `mapstructure:"galaxy_file" cty:"galaxy_file" hcl:"galaxy_file"`
	GalaxyCommand         *string           `mapstructure:"galaxy_command" cty:"galaxy_command" hcl:"galaxy_command"`
	GalaxyForceInstall    *bool             `mapstructure:"galaxy_force_install" cty:"galaxy_force_install" hcl:"galaxy_force_install"`
	GalaxyRolesPath       *string           `mapstructure:"galaxy_roles_path" cty:"galaxy_roles_path" hcl:"galaxy_roles_path"`
	GalaxyCollectionsPath *string           `mapstructure:"galaxy_collections_path" cty:"galaxy_collections_path" hcl:"galaxy_collections_path"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatConfig)
}

// HCL2Spec returns the hcl spec of a Config.
// This spec is used by HCL to read the fields of Config.
// The decoded values from this spec will then be applied to a FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"packer_build_name":          &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":        &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_core_version":        &hcldec.AttrSpec{Name: "packer_core_version", Type: cty.String, Required: false},
		"packer_debug":               &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":               &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":            &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":      &hcldec.AttrSpec{Name: "packer_user_variables", Type: cty.Map(cty.String), Required: false},
		"packer_sensitive_variables": &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"command":                    &hcldec.AttrSpec{Name: "command", Type: cty.String, Required: false},
		"extra_arguments":            &hcldec.AttrSpec{Name: "extra_arguments", Type: cty.List(cty.String), Required: false},
		"group_vars":                 &hcldec.AttrSpec{Name: "group_vars", Type: cty.String, Required: false},
		"host_vars":                  &hcldec.AttrSpec{Name: "host_vars", Type: cty.String, Required: false},
		"playbook_dir":               &hcldec.AttrSpec{Name: "playbook_dir", Type: cty.String, Required: false},
		"playbook_file":              &hcldec.AttrSpec{Name: "playbook_file", Type: cty.String, Required: false},
		"playbook_files":             &hcldec.AttrSpec{Name: "playbook_files", Type: cty.List(cty.String), Required: false},
		"playbook_paths":             &hcldec.AttrSpec{Name: "playbook_paths", Type: cty.List(cty.String), Required: false},
		"role_paths":                 &hcldec.AttrSpec{Name: "role_paths", Type: cty.List(cty.String), Required: false},
		"collection_paths":           &hcldec.AttrSpec{Name: "collection_paths", Type: cty.List(cty.String), Required: false},
		"staging_directory":          &hcldec.AttrSpec{Name: "staging_directory", Type: cty.String, Required: false},
		"clean_staging_directory":    &hcldec.AttrSpec{Name: "clean_staging_directory", Type: cty.Bool, Required: false},
		"inventory_file":             &hcldec.AttrSpec{Name: "inventory_file", Type: cty.String, Required: false},
		"inventory_groups":           &hcldec.AttrSpec{Name: "inventory_groups", Type: cty.List(cty.String), Required: false},
		"galaxy_file":                &hcldec.AttrSpec{Name: "galaxy_file", Type: cty.String, Required: false},
		"galaxy_command":             &hcldec.AttrSpec{Name: "galaxy_command", Type: cty.String, Required: false},
		"galaxy_force_install":       &hcldec.AttrSpec{Name: "galaxy_force_install", Type: cty.Bool, Required: false},
		"galaxy_roles_path":          &hcldec.AttrSpec{Name: "galaxy_roles_path", Type: cty.String, Required: false},
		"galaxy_collections_path":    &hcldec.AttrSpec{Name: "galaxy_collections_path", Type: cty.String, Required: false},
	}
	return s
}
