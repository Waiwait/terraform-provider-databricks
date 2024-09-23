package tfschema

import (
	"fmt"

	"github.com/databricks/terraform-provider-databricks/common"
	dataschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// ListNestedBlockBuilder represents a list of complex (non-primitive) types.
type ListNestedBlockBuilder struct {
	NestedObject       NestedBlockObject
	Optional           bool
	Required           bool
	Sensitive          bool
	Computed           bool
	DeprecationMessage string
	Validators         []validator.List
	PlanModifiers      []planmodifier.List
}

func (a ListNestedBlockBuilder) BuildDataSourceAttribute() dataschema.Attribute {
	panic(fmt.Errorf("BuildDataSourceBlock should never be called for ListNestedAttributeBuilder. %s", common.TerraformBugErrorMessage))
}

func (a ListNestedBlockBuilder) BuildResourceAttribute() schema.Attribute {
	panic(fmt.Errorf("BuildResourceBlock should never be called for ListNestedAttributeBuilder. %s", common.TerraformBugErrorMessage))
}

func (a ListNestedBlockBuilder) BuildDataSourceBlock() dataschema.Block {
	return dataschema.ListNestedBlock{
		NestedObject:       a.NestedObject.BuildDataSourceAttribute(),
		DeprecationMessage: a.DeprecationMessage,
		Validators:         a.Validators,
	}
}

func (a ListNestedBlockBuilder) BuildResourceBlock() schema.Block {
	return schema.ListNestedBlock{
		NestedObject:       a.NestedObject.BuildResourceAttribute(),
		DeprecationMessage: a.DeprecationMessage,
		Validators:         a.Validators,
		PlanModifiers:      a.PlanModifiers,
	}
}

func (a ListNestedBlockBuilder) SetOptional() AttributeBuilder {
	if a.Optional && !a.Required {
		panic("attribute is already optional")
	}
	a.Optional = true
	a.Required = false
	return a
}

func (a ListNestedBlockBuilder) SetRequired() AttributeBuilder {
	if !a.Optional && a.Required {
		panic("attribute is already required")
	}
	a.Optional = false
	a.Required = true
	return a
}

func (a ListNestedBlockBuilder) SetSensitive() AttributeBuilder {
	if a.Sensitive {
		panic("attribute is already sensitive")
	}
	a.Sensitive = true
	return a
}

func (a ListNestedBlockBuilder) SetComputed() AttributeBuilder {
	if a.Computed {
		panic("attribute is already computed")
	}
	a.Computed = true
	return a
}

func (a ListNestedBlockBuilder) SetReadOnly() AttributeBuilder {
	if a.Computed && !a.Optional && !a.Required {
		panic("attribute is already read only")
	}
	a.Computed = true
	a.Optional = false
	a.Required = false
	return a
}

func (a ListNestedBlockBuilder) SetDeprecated(msg string) AttributeBuilder {
	a.DeprecationMessage = msg
	return a
}

func (a ListNestedBlockBuilder) AddValidator(v validator.List) AttributeBuilder {
	a.Validators = append(a.Validators, v)
	return a
}

func (a ListNestedBlockBuilder) AddPlanModifier(v planmodifier.List) AttributeBuilder {
	a.PlanModifiers = append(a.PlanModifiers, v)
	return a
}
