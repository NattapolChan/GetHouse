// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NattapolChan/ent/predicate"
	"github.com/NattapolChan/ent/propertylisting"
	"github.com/NattapolChan/ent/user"
	"github.com/google/uuid"
)

// PropertyListingUpdate is the builder for updating PropertyListing entities.
type PropertyListingUpdate struct {
	config
	hooks    []Hook
	mutation *PropertyListingMutation
}

// Where appends a list predicates to the PropertyListingUpdate builder.
func (plu *PropertyListingUpdate) Where(ps ...predicate.PropertyListing) *PropertyListingUpdate {
	plu.mutation.Where(ps...)
	return plu
}

// SetName sets the "name" field.
func (plu *PropertyListingUpdate) SetName(s string) *PropertyListingUpdate {
	plu.mutation.SetName(s)
	return plu
}

// SetHouseType sets the "houseType" field.
func (plu *PropertyListingUpdate) SetHouseType(s string) *PropertyListingUpdate {
	plu.mutation.SetHouseType(s)
	return plu
}

// SetArea sets the "area" field.
func (plu *PropertyListingUpdate) SetArea(i int) *PropertyListingUpdate {
	plu.mutation.ResetArea()
	plu.mutation.SetArea(i)
	return plu
}

// SetNillableArea sets the "area" field if the given value is not nil.
func (plu *PropertyListingUpdate) SetNillableArea(i *int) *PropertyListingUpdate {
	if i != nil {
		plu.SetArea(*i)
	}
	return plu
}

// AddArea adds i to the "area" field.
func (plu *PropertyListingUpdate) AddArea(i int) *PropertyListingUpdate {
	plu.mutation.AddArea(i)
	return plu
}

// SetNumOfBedroom sets the "num_of_bedroom" field.
func (plu *PropertyListingUpdate) SetNumOfBedroom(i int) *PropertyListingUpdate {
	plu.mutation.ResetNumOfBedroom()
	plu.mutation.SetNumOfBedroom(i)
	return plu
}

// AddNumOfBedroom adds i to the "num_of_bedroom" field.
func (plu *PropertyListingUpdate) AddNumOfBedroom(i int) *PropertyListingUpdate {
	plu.mutation.AddNumOfBedroom(i)
	return plu
}

// SetLeaseDate sets the "lease_date" field.
func (plu *PropertyListingUpdate) SetLeaseDate(s string) *PropertyListingUpdate {
	plu.mutation.SetLeaseDate(s)
	return plu
}

// SetNumberOfData sets the "number_of_data" field.
func (plu *PropertyListingUpdate) SetNumberOfData(i int) *PropertyListingUpdate {
	plu.mutation.ResetNumberOfData()
	plu.mutation.SetNumberOfData(i)
	return plu
}

// AddNumberOfData adds i to the "number_of_data" field.
func (plu *PropertyListingUpdate) AddNumberOfData(i int) *PropertyListingUpdate {
	plu.mutation.AddNumberOfData(i)
	return plu
}

// SetRentalPrice sets the "rental_price" field.
func (plu *PropertyListingUpdate) SetRentalPrice(i int) *PropertyListingUpdate {
	plu.mutation.ResetRentalPrice()
	plu.mutation.SetRentalPrice(i)
	return plu
}

// AddRentalPrice adds i to the "rental_price" field.
func (plu *PropertyListingUpdate) AddRentalPrice(i int) *PropertyListingUpdate {
	plu.mutation.AddRentalPrice(i)
	return plu
}

// SetY sets the "y" field.
func (plu *PropertyListingUpdate) SetY(f float64) *PropertyListingUpdate {
	plu.mutation.ResetY()
	plu.mutation.SetY(f)
	return plu
}

// AddY adds f to the "y" field.
func (plu *PropertyListingUpdate) AddY(f float64) *PropertyListingUpdate {
	plu.mutation.AddY(f)
	return plu
}

// SetX sets the "x" field.
func (plu *PropertyListingUpdate) SetX(f float64) *PropertyListingUpdate {
	plu.mutation.ResetX()
	plu.mutation.SetX(f)
	return plu
}

// AddX adds f to the "x" field.
func (plu *PropertyListingUpdate) AddX(f float64) *PropertyListingUpdate {
	plu.mutation.AddX(f)
	return plu
}

// SetStreet sets the "street" field.
func (plu *PropertyListingUpdate) SetStreet(s string) *PropertyListingUpdate {
	plu.mutation.SetStreet(s)
	return plu
}

// SetNillableStreet sets the "street" field if the given value is not nil.
func (plu *PropertyListingUpdate) SetNillableStreet(s *string) *PropertyListingUpdate {
	if s != nil {
		plu.SetStreet(*s)
	}
	return plu
}

// ClearStreet clears the value of the "street" field.
func (plu *PropertyListingUpdate) ClearStreet() *PropertyListingUpdate {
	plu.mutation.ClearStreet()
	return plu
}

// SetDistrict sets the "district" field.
func (plu *PropertyListingUpdate) SetDistrict(s string) *PropertyListingUpdate {
	plu.mutation.SetDistrict(s)
	return plu
}

// SetUUID sets the "uuid" field.
func (plu *PropertyListingUpdate) SetUUID(u uuid.UUID) *PropertyListingUpdate {
	plu.mutation.SetUUID(u)
	return plu
}

// SetNillableUUID sets the "uuid" field if the given value is not nil.
func (plu *PropertyListingUpdate) SetNillableUUID(u *uuid.UUID) *PropertyListingUpdate {
	if u != nil {
		plu.SetUUID(*u)
	}
	return plu
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (plu *PropertyListingUpdate) AddUserIDs(ids ...string) *PropertyListingUpdate {
	plu.mutation.AddUserIDs(ids...)
	return plu
}

// AddUsers adds the "users" edges to the User entity.
func (plu *PropertyListingUpdate) AddUsers(u ...*User) *PropertyListingUpdate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return plu.AddUserIDs(ids...)
}

// Mutation returns the PropertyListingMutation object of the builder.
func (plu *PropertyListingUpdate) Mutation() *PropertyListingMutation {
	return plu.mutation
}

// ClearUsers clears all "users" edges to the User entity.
func (plu *PropertyListingUpdate) ClearUsers() *PropertyListingUpdate {
	plu.mutation.ClearUsers()
	return plu
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (plu *PropertyListingUpdate) RemoveUserIDs(ids ...string) *PropertyListingUpdate {
	plu.mutation.RemoveUserIDs(ids...)
	return plu
}

// RemoveUsers removes "users" edges to User entities.
func (plu *PropertyListingUpdate) RemoveUsers(u ...*User) *PropertyListingUpdate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return plu.RemoveUserIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (plu *PropertyListingUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, plu.sqlSave, plu.mutation, plu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (plu *PropertyListingUpdate) SaveX(ctx context.Context) int {
	affected, err := plu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (plu *PropertyListingUpdate) Exec(ctx context.Context) error {
	_, err := plu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (plu *PropertyListingUpdate) ExecX(ctx context.Context) {
	if err := plu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (plu *PropertyListingUpdate) check() error {
	if v, ok := plu.mutation.LeaseDate(); ok {
		if err := propertylisting.LeaseDateValidator(v); err != nil {
			return &ValidationError{Name: "lease_date", err: fmt.Errorf(`ent: validator failed for field "PropertyListing.lease_date": %w`, err)}
		}
	}
	return nil
}

func (plu *PropertyListingUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := plu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(propertylisting.Table, propertylisting.Columns, sqlgraph.NewFieldSpec(propertylisting.FieldID, field.TypeInt))
	if ps := plu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := plu.mutation.Name(); ok {
		_spec.SetField(propertylisting.FieldName, field.TypeString, value)
	}
	if value, ok := plu.mutation.HouseType(); ok {
		_spec.SetField(propertylisting.FieldHouseType, field.TypeString, value)
	}
	if value, ok := plu.mutation.Area(); ok {
		_spec.SetField(propertylisting.FieldArea, field.TypeInt, value)
	}
	if value, ok := plu.mutation.AddedArea(); ok {
		_spec.AddField(propertylisting.FieldArea, field.TypeInt, value)
	}
	if value, ok := plu.mutation.NumOfBedroom(); ok {
		_spec.SetField(propertylisting.FieldNumOfBedroom, field.TypeInt, value)
	}
	if value, ok := plu.mutation.AddedNumOfBedroom(); ok {
		_spec.AddField(propertylisting.FieldNumOfBedroom, field.TypeInt, value)
	}
	if value, ok := plu.mutation.LeaseDate(); ok {
		_spec.SetField(propertylisting.FieldLeaseDate, field.TypeString, value)
	}
	if value, ok := plu.mutation.NumberOfData(); ok {
		_spec.SetField(propertylisting.FieldNumberOfData, field.TypeInt, value)
	}
	if value, ok := plu.mutation.AddedNumberOfData(); ok {
		_spec.AddField(propertylisting.FieldNumberOfData, field.TypeInt, value)
	}
	if value, ok := plu.mutation.RentalPrice(); ok {
		_spec.SetField(propertylisting.FieldRentalPrice, field.TypeInt, value)
	}
	if value, ok := plu.mutation.AddedRentalPrice(); ok {
		_spec.AddField(propertylisting.FieldRentalPrice, field.TypeInt, value)
	}
	if value, ok := plu.mutation.Y(); ok {
		_spec.SetField(propertylisting.FieldY, field.TypeFloat64, value)
	}
	if value, ok := plu.mutation.AddedY(); ok {
		_spec.AddField(propertylisting.FieldY, field.TypeFloat64, value)
	}
	if value, ok := plu.mutation.X(); ok {
		_spec.SetField(propertylisting.FieldX, field.TypeFloat64, value)
	}
	if value, ok := plu.mutation.AddedX(); ok {
		_spec.AddField(propertylisting.FieldX, field.TypeFloat64, value)
	}
	if value, ok := plu.mutation.Street(); ok {
		_spec.SetField(propertylisting.FieldStreet, field.TypeString, value)
	}
	if plu.mutation.StreetCleared() {
		_spec.ClearField(propertylisting.FieldStreet, field.TypeString)
	}
	if value, ok := plu.mutation.District(); ok {
		_spec.SetField(propertylisting.FieldDistrict, field.TypeString, value)
	}
	if value, ok := plu.mutation.UUID(); ok {
		_spec.SetField(propertylisting.FieldUUID, field.TypeUUID, value)
	}
	if plu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   propertylisting.UsersTable,
			Columns: propertylisting.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := plu.mutation.RemovedUsersIDs(); len(nodes) > 0 && !plu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   propertylisting.UsersTable,
			Columns: propertylisting.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := plu.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   propertylisting.UsersTable,
			Columns: propertylisting.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, plu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{propertylisting.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	plu.mutation.done = true
	return n, nil
}

// PropertyListingUpdateOne is the builder for updating a single PropertyListing entity.
type PropertyListingUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PropertyListingMutation
}

// SetName sets the "name" field.
func (pluo *PropertyListingUpdateOne) SetName(s string) *PropertyListingUpdateOne {
	pluo.mutation.SetName(s)
	return pluo
}

// SetHouseType sets the "houseType" field.
func (pluo *PropertyListingUpdateOne) SetHouseType(s string) *PropertyListingUpdateOne {
	pluo.mutation.SetHouseType(s)
	return pluo
}

// SetArea sets the "area" field.
func (pluo *PropertyListingUpdateOne) SetArea(i int) *PropertyListingUpdateOne {
	pluo.mutation.ResetArea()
	pluo.mutation.SetArea(i)
	return pluo
}

// SetNillableArea sets the "area" field if the given value is not nil.
func (pluo *PropertyListingUpdateOne) SetNillableArea(i *int) *PropertyListingUpdateOne {
	if i != nil {
		pluo.SetArea(*i)
	}
	return pluo
}

// AddArea adds i to the "area" field.
func (pluo *PropertyListingUpdateOne) AddArea(i int) *PropertyListingUpdateOne {
	pluo.mutation.AddArea(i)
	return pluo
}

// SetNumOfBedroom sets the "num_of_bedroom" field.
func (pluo *PropertyListingUpdateOne) SetNumOfBedroom(i int) *PropertyListingUpdateOne {
	pluo.mutation.ResetNumOfBedroom()
	pluo.mutation.SetNumOfBedroom(i)
	return pluo
}

// AddNumOfBedroom adds i to the "num_of_bedroom" field.
func (pluo *PropertyListingUpdateOne) AddNumOfBedroom(i int) *PropertyListingUpdateOne {
	pluo.mutation.AddNumOfBedroom(i)
	return pluo
}

// SetLeaseDate sets the "lease_date" field.
func (pluo *PropertyListingUpdateOne) SetLeaseDate(s string) *PropertyListingUpdateOne {
	pluo.mutation.SetLeaseDate(s)
	return pluo
}

// SetNumberOfData sets the "number_of_data" field.
func (pluo *PropertyListingUpdateOne) SetNumberOfData(i int) *PropertyListingUpdateOne {
	pluo.mutation.ResetNumberOfData()
	pluo.mutation.SetNumberOfData(i)
	return pluo
}

// AddNumberOfData adds i to the "number_of_data" field.
func (pluo *PropertyListingUpdateOne) AddNumberOfData(i int) *PropertyListingUpdateOne {
	pluo.mutation.AddNumberOfData(i)
	return pluo
}

// SetRentalPrice sets the "rental_price" field.
func (pluo *PropertyListingUpdateOne) SetRentalPrice(i int) *PropertyListingUpdateOne {
	pluo.mutation.ResetRentalPrice()
	pluo.mutation.SetRentalPrice(i)
	return pluo
}

// AddRentalPrice adds i to the "rental_price" field.
func (pluo *PropertyListingUpdateOne) AddRentalPrice(i int) *PropertyListingUpdateOne {
	pluo.mutation.AddRentalPrice(i)
	return pluo
}

// SetY sets the "y" field.
func (pluo *PropertyListingUpdateOne) SetY(f float64) *PropertyListingUpdateOne {
	pluo.mutation.ResetY()
	pluo.mutation.SetY(f)
	return pluo
}

// AddY adds f to the "y" field.
func (pluo *PropertyListingUpdateOne) AddY(f float64) *PropertyListingUpdateOne {
	pluo.mutation.AddY(f)
	return pluo
}

// SetX sets the "x" field.
func (pluo *PropertyListingUpdateOne) SetX(f float64) *PropertyListingUpdateOne {
	pluo.mutation.ResetX()
	pluo.mutation.SetX(f)
	return pluo
}

// AddX adds f to the "x" field.
func (pluo *PropertyListingUpdateOne) AddX(f float64) *PropertyListingUpdateOne {
	pluo.mutation.AddX(f)
	return pluo
}

// SetStreet sets the "street" field.
func (pluo *PropertyListingUpdateOne) SetStreet(s string) *PropertyListingUpdateOne {
	pluo.mutation.SetStreet(s)
	return pluo
}

// SetNillableStreet sets the "street" field if the given value is not nil.
func (pluo *PropertyListingUpdateOne) SetNillableStreet(s *string) *PropertyListingUpdateOne {
	if s != nil {
		pluo.SetStreet(*s)
	}
	return pluo
}

// ClearStreet clears the value of the "street" field.
func (pluo *PropertyListingUpdateOne) ClearStreet() *PropertyListingUpdateOne {
	pluo.mutation.ClearStreet()
	return pluo
}

// SetDistrict sets the "district" field.
func (pluo *PropertyListingUpdateOne) SetDistrict(s string) *PropertyListingUpdateOne {
	pluo.mutation.SetDistrict(s)
	return pluo
}

// SetUUID sets the "uuid" field.
func (pluo *PropertyListingUpdateOne) SetUUID(u uuid.UUID) *PropertyListingUpdateOne {
	pluo.mutation.SetUUID(u)
	return pluo
}

// SetNillableUUID sets the "uuid" field if the given value is not nil.
func (pluo *PropertyListingUpdateOne) SetNillableUUID(u *uuid.UUID) *PropertyListingUpdateOne {
	if u != nil {
		pluo.SetUUID(*u)
	}
	return pluo
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (pluo *PropertyListingUpdateOne) AddUserIDs(ids ...string) *PropertyListingUpdateOne {
	pluo.mutation.AddUserIDs(ids...)
	return pluo
}

// AddUsers adds the "users" edges to the User entity.
func (pluo *PropertyListingUpdateOne) AddUsers(u ...*User) *PropertyListingUpdateOne {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return pluo.AddUserIDs(ids...)
}

// Mutation returns the PropertyListingMutation object of the builder.
func (pluo *PropertyListingUpdateOne) Mutation() *PropertyListingMutation {
	return pluo.mutation
}

// ClearUsers clears all "users" edges to the User entity.
func (pluo *PropertyListingUpdateOne) ClearUsers() *PropertyListingUpdateOne {
	pluo.mutation.ClearUsers()
	return pluo
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (pluo *PropertyListingUpdateOne) RemoveUserIDs(ids ...string) *PropertyListingUpdateOne {
	pluo.mutation.RemoveUserIDs(ids...)
	return pluo
}

// RemoveUsers removes "users" edges to User entities.
func (pluo *PropertyListingUpdateOne) RemoveUsers(u ...*User) *PropertyListingUpdateOne {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return pluo.RemoveUserIDs(ids...)
}

// Where appends a list predicates to the PropertyListingUpdate builder.
func (pluo *PropertyListingUpdateOne) Where(ps ...predicate.PropertyListing) *PropertyListingUpdateOne {
	pluo.mutation.Where(ps...)
	return pluo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pluo *PropertyListingUpdateOne) Select(field string, fields ...string) *PropertyListingUpdateOne {
	pluo.fields = append([]string{field}, fields...)
	return pluo
}

// Save executes the query and returns the updated PropertyListing entity.
func (pluo *PropertyListingUpdateOne) Save(ctx context.Context) (*PropertyListing, error) {
	return withHooks(ctx, pluo.sqlSave, pluo.mutation, pluo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pluo *PropertyListingUpdateOne) SaveX(ctx context.Context) *PropertyListing {
	node, err := pluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pluo *PropertyListingUpdateOne) Exec(ctx context.Context) error {
	_, err := pluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pluo *PropertyListingUpdateOne) ExecX(ctx context.Context) {
	if err := pluo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pluo *PropertyListingUpdateOne) check() error {
	if v, ok := pluo.mutation.LeaseDate(); ok {
		if err := propertylisting.LeaseDateValidator(v); err != nil {
			return &ValidationError{Name: "lease_date", err: fmt.Errorf(`ent: validator failed for field "PropertyListing.lease_date": %w`, err)}
		}
	}
	return nil
}

func (pluo *PropertyListingUpdateOne) sqlSave(ctx context.Context) (_node *PropertyListing, err error) {
	if err := pluo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(propertylisting.Table, propertylisting.Columns, sqlgraph.NewFieldSpec(propertylisting.FieldID, field.TypeInt))
	id, ok := pluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PropertyListing.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, propertylisting.FieldID)
		for _, f := range fields {
			if !propertylisting.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != propertylisting.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pluo.mutation.Name(); ok {
		_spec.SetField(propertylisting.FieldName, field.TypeString, value)
	}
	if value, ok := pluo.mutation.HouseType(); ok {
		_spec.SetField(propertylisting.FieldHouseType, field.TypeString, value)
	}
	if value, ok := pluo.mutation.Area(); ok {
		_spec.SetField(propertylisting.FieldArea, field.TypeInt, value)
	}
	if value, ok := pluo.mutation.AddedArea(); ok {
		_spec.AddField(propertylisting.FieldArea, field.TypeInt, value)
	}
	if value, ok := pluo.mutation.NumOfBedroom(); ok {
		_spec.SetField(propertylisting.FieldNumOfBedroom, field.TypeInt, value)
	}
	if value, ok := pluo.mutation.AddedNumOfBedroom(); ok {
		_spec.AddField(propertylisting.FieldNumOfBedroom, field.TypeInt, value)
	}
	if value, ok := pluo.mutation.LeaseDate(); ok {
		_spec.SetField(propertylisting.FieldLeaseDate, field.TypeString, value)
	}
	if value, ok := pluo.mutation.NumberOfData(); ok {
		_spec.SetField(propertylisting.FieldNumberOfData, field.TypeInt, value)
	}
	if value, ok := pluo.mutation.AddedNumberOfData(); ok {
		_spec.AddField(propertylisting.FieldNumberOfData, field.TypeInt, value)
	}
	if value, ok := pluo.mutation.RentalPrice(); ok {
		_spec.SetField(propertylisting.FieldRentalPrice, field.TypeInt, value)
	}
	if value, ok := pluo.mutation.AddedRentalPrice(); ok {
		_spec.AddField(propertylisting.FieldRentalPrice, field.TypeInt, value)
	}
	if value, ok := pluo.mutation.Y(); ok {
		_spec.SetField(propertylisting.FieldY, field.TypeFloat64, value)
	}
	if value, ok := pluo.mutation.AddedY(); ok {
		_spec.AddField(propertylisting.FieldY, field.TypeFloat64, value)
	}
	if value, ok := pluo.mutation.X(); ok {
		_spec.SetField(propertylisting.FieldX, field.TypeFloat64, value)
	}
	if value, ok := pluo.mutation.AddedX(); ok {
		_spec.AddField(propertylisting.FieldX, field.TypeFloat64, value)
	}
	if value, ok := pluo.mutation.Street(); ok {
		_spec.SetField(propertylisting.FieldStreet, field.TypeString, value)
	}
	if pluo.mutation.StreetCleared() {
		_spec.ClearField(propertylisting.FieldStreet, field.TypeString)
	}
	if value, ok := pluo.mutation.District(); ok {
		_spec.SetField(propertylisting.FieldDistrict, field.TypeString, value)
	}
	if value, ok := pluo.mutation.UUID(); ok {
		_spec.SetField(propertylisting.FieldUUID, field.TypeUUID, value)
	}
	if pluo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   propertylisting.UsersTable,
			Columns: propertylisting.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pluo.mutation.RemovedUsersIDs(); len(nodes) > 0 && !pluo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   propertylisting.UsersTable,
			Columns: propertylisting.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pluo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   propertylisting.UsersTable,
			Columns: propertylisting.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &PropertyListing{config: pluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{propertylisting.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	pluo.mutation.done = true
	return _node, nil
}
