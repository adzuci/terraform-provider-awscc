package reflect

import (
	"context"
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/internal/diagnostics"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// Pointer builds a new zero value of the concrete type that `target`
// references, populates it with BuildValue, and takes a pointer to it.
//
// It is meant to be called through Into, not directly.
func Pointer(ctx context.Context, typ attr.Type, val tftypes.Value, target reflect.Value, opts Options, path *tftypes.AttributePath) (reflect.Value, []*tfprotov6.Diagnostic) {
	var diags []*tfprotov6.Diagnostic

	if target.Kind() != reflect.Ptr {
		err := fmt.Errorf("cannot dereference pointer, not a pointer, is a %s (%s)", target.Type(), target.Kind())
		return target, append(diags, &tfprotov6.Diagnostic{
			Severity:  tfprotov6.DiagnosticSeverityError,
			Summary:   "Value Conversion Error",
			Detail:    "An unexpected error was encountered trying to convert to pointer value. This is always an error in the provider. Please report the following to the provider developer:\n\n" + err.Error(),
			Attribute: path,
		})
	}
	// we may have gotten a nil pointer, so we need to create our own that
	// we can set
	pointer := reflect.New(target.Type().Elem())
	// build out whatever the pointer is pointing to
	pointed, pointedDiags := BuildValue(ctx, typ, val, pointer.Elem(), opts, path)
	diags = append(diags, pointedDiags...)

	if diagnostics.DiagsHasErrors(diags) {
		return target, diags
	}
	// to be able to set the pointer to our new pointer, we need to create
	// a pointer to the pointer
	pointerPointer := reflect.New(pointer.Type())
	// we set the pointer we created on the pointer to the pointer
	pointerPointer.Elem().Set(pointer)
	// then it's settable, so we can now set the concrete value we created
	// on the pointer
	pointerPointer.Elem().Elem().Set(pointed)
	// return the pointer we created
	return pointerPointer.Elem(), diags
}

// create a zero value of concrete type underlying any number of pointers, then
// wrap it in that number of pointers again. The end result is to wind up with
// the same exact type, except now you can be sure it's pointing to actual data
// and will not give you a nil pointer dereference panic unexpectedly.
func pointerSafeZeroValue(ctx context.Context, target reflect.Value) reflect.Value {
	pointer := target.Type()
	var pointers int
	for pointer.Kind() == reflect.Ptr {
		pointer = pointer.Elem()
		pointers++
	}
	receiver := reflect.Zero(pointer)
	for i := 0; i < pointers; i++ {
		newReceiver := reflect.New(receiver.Type())
		newReceiver.Elem().Set(receiver)
		receiver = newReceiver
	}
	return receiver
}

// FromPointer turns a pointer into an attr.Value using `typ`. If the pointer
// is nil, the attr.Value will use its null representation. If it is not nil,
// it will recurse into FromValue to find the attr.Value of the type the value
// the pointer is referencing.
//
// It is meant to be called through OutOf, not directly.
func FromPointer(ctx context.Context, typ attr.Type, value reflect.Value, path *tftypes.AttributePath) (attr.Value, []*tfprotov6.Diagnostic) {
	var diags []*tfprotov6.Diagnostic

	if value.Kind() != reflect.Ptr {
		err := fmt.Errorf("cannot use type %s as a pointer", value.Type())
		return nil, append(diags, &tfprotov6.Diagnostic{
			Severity:  tfprotov6.DiagnosticSeverityError,
			Summary:   "Value Conversion Error",
			Detail:    "An unexpected error was encountered trying to convert from pointer value. This is always an error in the provider. Please report the following to the provider developer:\n\n" + err.Error(),
			Attribute: path,
		})
	}
	if value.IsNil() {
		tfVal := tftypes.NewValue(typ.TerraformType(ctx), nil)

		if typeWithValidate, ok := typ.(attr.TypeWithValidate); ok {
			diags = append(diags, typeWithValidate.Validate(ctx, tfVal)...)

			if diagnostics.DiagsHasErrors(diags) {
				return nil, diags
			}
		}

		attrVal, err := typ.ValueFromTerraform(ctx, tfVal)

		if err != nil {
			return nil, append(diags, &tfprotov6.Diagnostic{
				Severity:  tfprotov6.DiagnosticSeverityError,
				Summary:   "Value Conversion Error",
				Detail:    "An unexpected error was encountered trying to convert from pointer value. This is always an error in the provider. Please report the following to the provider developer:\n\n" + err.Error(),
				Attribute: path,
			})
		}

		return attrVal, diags
	}

	attrVal, attrValDiags := FromValue(ctx, typ, value.Elem().Interface(), path)
	diags = append(diags, attrValDiags...)

	return attrVal, diags
}
